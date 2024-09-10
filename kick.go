package kick

import (
	"io"
	"math"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// Waveform types
const (
	WaveSine = iota
	WaveTriangle
)

// GenerateKickWithEffects generates a kick drum sound and writes it to a provided io.WriteSeeker.
func GenerateKickWithEffects(
	startFreq, endFreq float64, sampleRate int, duration float64,
	waveformType int, attack, decay, sustain, release, drive, filterCutoff, sweep, pitchDecay float64,
	output io.WriteSeeker,
) error {
	// Generate the kick drum samples
	samples := generateSamples(startFreq, endFreq, sampleRate, duration, waveformType, attack, decay, sustain, release, drive, filterCutoff, sweep, pitchDecay)

	// Prepare the audio buffer
	buffer := &audio.IntBuffer{
		Data:           samples,
		Format:         &audio.Format{SampleRate: sampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}

	// Write the WAV-encoded audio data to the provided io.WriteSeeker
	encoder := wav.NewEncoder(output, sampleRate, 16, 1, 1)
	if err := encoder.Write(buffer); err != nil {
		return err
	}
	return encoder.Close()
}

// generateSamples generates the kick drum samples with distortion, filtering, pitch sweep, and pitch envelope.
func generateSamples(
	startFreq, endFreq float64, sampleRate int, duration float64,
	waveformType int, attack, decay, sustain, release, drive, filterCutoff, sweep, pitchDecay float64,
) []int {
	numSamples := int(float64(sampleRate) * duration)
	samples := make([]int, numSamples)

	for i := 0; i < numSamples; i++ {
		// Time in seconds
		t := float64(i) / float64(sampleRate)

		// Exponential pitch sweep with pitch envelope decay for punch
		decayFactor := math.Pow(endFreq/startFreq, (t/duration)*sweep)
		frequency := startFreq * decayFactor * math.Pow(2, -t/pitchDecay)

		// Generate waveform
		var sample float64
		if waveformType == WaveSine {
			sample = math.Sin(2 * math.Pi * frequency * t)
		} else if waveformType == WaveTriangle {
			sample = 2*math.Abs(2*((t*frequency)-math.Floor((t*frequency)+0.5)))-1
		}

		// Apply drive (simple soft clipping)
		sample = applyDrive(sample, drive)

		// Apply envelope (ADSR)
		envelopeValue := applyEnvelope(t, attack, decay, sustain, release, duration)
		sample *= envelopeValue

		// Apply low-pass filter
		sample = applyLowPassFilter(sample, frequency, filterCutoff, sampleRate)

		// Scale to 16-bit integer range [-32767, 32767]
		samples[i] = int(sample * 32767.0)
	}

	return samples
}

// applyDrive applies a simple distortion/drive effect by soft-clipping the signal.
func applyDrive(sample, driveAmount float64) float64 {
	if driveAmount > 0 {
		return sample * (1 + driveAmount) / (1 + driveAmount*math.Abs(sample))
	}
	return sample
}

// applyEnvelope applies an Attack-Decay-Sustain-Release (ADSR) envelope to the signal.
func applyEnvelope(t, attack, decay, sustain, release, duration float64) float64 {
	if t < attack {
		return t / attack
	}
	if t < attack+decay {
		return 1.0 - ((t - attack) / decay * (1.0 - sustain))
	}
	if t < duration-release {
		return sustain
	}
	if t < duration {
		return sustain * (1.0 - (t-(duration-release))/release)
	}
	return 0.0
}

// applyLowPassFilter applies a simple low-pass filter to the signal.
func applyLowPassFilter(sample, freq, cutoff float64, sampleRate int) float64 {
	if freq > cutoff {
		// Simple low-pass filter: attenuate frequencies above cutoff
		sample *= cutoff / freq
	}
	return sample
}
