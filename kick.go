package kick

import (
	"io"
	"math"
	"math/rand"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

// Waveform types
const (
	WaveSine = iota
	WaveTriangle
)

// Noise types
const (
	NoiseNone = iota
	NoiseWhite
	NoisePink
	NoiseBrown
)

// GenerateKickWithEffects generates a kick drum sound and writes it to a provided io.WriteSeeker.
// It also allows mixing in white, pink, or brown noise.
func GenerateKickWithEffects(
	startFreq, endFreq float64, sampleRate int, duration float64,
	waveformType int, attack, decay, sustain, release, drive, filterCutoff, sweep, pitchDecay float64,
	noiseType int, noiseAmount float64,
	output io.WriteSeeker,
) error {
	// Generate the kick drum samples
	samples := generateSamples(startFreq, endFreq, sampleRate, duration, waveformType, attack, decay, sustain, release, drive, filterCutoff, sweep, pitchDecay)

	// Mix in noise if specified
	if noiseType != NoiseNone {
		mixNoise(samples, noiseType, noiseAmount, sampleRate)
	}

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
			sample = 2*math.Abs(2*((t*frequency)-math.Floor((t*frequency)+0.5))) - 1
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

// mixNoise mixes the specified noise into the generated kick drum samples.
func mixNoise(samples []int, noiseType int, noiseAmount float64, sampleRate int) {
	for i := range samples {
		noiseSample := generateNoiseSample(noiseType, i, sampleRate)
		samples[i] += int(noiseSample * noiseAmount * 32767.0)
		// Clip to 16-bit range
		if samples[i] > 32767 {
			samples[i] = 32767
		} else if samples[i] < -32767 {
			samples[i] = -32767
		}
	}
}

// generateNoiseSample generates a noise sample based on the noise type.
func generateNoiseSample(noiseType int, i int, sampleRate int) float64 {
	switch noiseType {
	case NoiseWhite:
		return rand.Float64()*2 - 1 // White noise
	case NoisePink:
		return generatePinkNoise(i) // Pink noise
	case NoiseBrown:
		return generateBrownNoise(i) // Brown noise
	default:
		return 0.0
	}
}

// Pink noise generation using a simple filtering approach
var pinkNoiseAccumulator float64

func generatePinkNoise(i int) float64 {
	// A simple, naive pink noise generation (can be improved with filtering)
	pinkNoiseAccumulator += rand.Float64()*2 - 1
	return pinkNoiseAccumulator / float64(i+1)
}

// Brown noise generation using a simple filtering approach
var brownNoiseAccumulator float64

func generateBrownNoise(i int) float64 {
	// A simple, naive brown noise generation (can be improved with filtering)
	brownNoiseAccumulator += (rand.Float64()*2 - 1) * 0.1
	if brownNoiseAccumulator > 1 {
		brownNoiseAccumulator = 1
	} else if brownNoiseAccumulator < -1 {
		brownNoiseAccumulator = -1
	}
	return brownNoiseAccumulator
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
