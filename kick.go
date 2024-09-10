package kick

import (
	"errors"
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
	WaveSawtooth
	WaveSquare
)

// Noise types
const (
	NoiseNone = iota
	NoiseWhite
	NoisePink
	NoiseBrown
)

// Config holds all parameters for kick drum generation
type Config struct {
	StartFreq       float64
	EndFreq         float64
	SampleRate      int
	Duration        float64
	WaveformType    int
	Attack          float64
	Decay           float64
	Sustain         float64
	Release         float64
	Drive           float64
	FilterCutoff    float64
	FilterResonance float64
	Sweep           float64
	PitchDecay      float64
	NoiseType       int
	NoiseAmount     float64
	NoiseAttack     float64
	NoiseDecay      float64
	Output          io.WriteSeeker
}

// NewConfig returns a new Config object with default values
func NewConfig(startFreq, endFreq float64, sampleRate int, duration float64, output io.WriteSeeker) (*Config, error) {
	if sampleRate <= 0 || duration <= 0 {
		return nil, errors.New("invalid sample rate or duration")
	}

	return &Config{
		StartFreq:       startFreq,
		EndFreq:         endFreq,
		SampleRate:      sampleRate,
		Duration:        duration,
		WaveformType:    WaveSine,
		Attack:          0.005,
		Decay:           0.3,
		Sustain:         0.2,
		Release:         0.2,
		Drive:           0.0,
		FilterCutoff:    10000.0,
		FilterResonance: 0.0,
		Sweep:           1.0,
		PitchDecay:      0.1,
		NoiseType:       NoiseNone,
		NoiseAmount:     0.0,
		Output:          output,
	}, nil
}

// GenerateKick generates a kick drum sound using the given configuration
func (cfg *Config) GenerateKick() error {
	// Generate the base samples for the kick drum
	samples := cfg.generateSamples()

	// Add resonator effects (simulating the effect of a drum cavity)
	applyResonatorEffect(samples, cfg.SampleRate, cfg.Duration)

	// Mix in noise with its own envelope if specified
	if cfg.NoiseType != NoiseNone {
		mixNoiseWithEnvelope(samples, cfg)
	}

	// Prepare the audio buffer
	buffer := &audio.IntBuffer{
		Data:           samples,
		Format:         &audio.Format{SampleRate: cfg.SampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}

	// Write the WAV-encoded audio data to the provided io.WriteSeeker
	encoder := wav.NewEncoder(cfg.Output, cfg.SampleRate, 16, 1, 1)
	if err := encoder.Write(buffer); err != nil {
		return err
	}
	return encoder.Close()
}

// generateSamples generates the base kick drum samples with configurable parameters
func (cfg *Config) generateSamples() []int {
	numSamples := int(float64(cfg.SampleRate) * cfg.Duration)
	samples := make([]int, numSamples)

	// Simulate membrane tension effects (pitch modulation over time)
	pitchMod := generatePitchModulation(cfg.StartFreq, cfg.EndFreq, cfg.SampleRate, cfg.Duration)

	for i := 0; i < numSamples; i++ {
		// Time in seconds
		t := float64(i) / float64(cfg.SampleRate)

		// Exponential pitch sweep with pitch envelope decay for punch
		decayFactor := math.Pow(cfg.EndFreq/cfg.StartFreq, (t/cfg.Duration)*cfg.Sweep)
		frequency := cfg.StartFreq * decayFactor * pitchMod[i]

		// Generate waveform
		var sample float64
		switch cfg.WaveformType {
		case WaveSine:
			sample = math.Sin(2 * math.Pi * frequency * t)
		case WaveTriangle:
			sample = 2*math.Abs(2*((t*frequency)-math.Floor((t*frequency)+0.5))) - 1
		case WaveSawtooth:
			sample = 2 * (t*frequency - math.Floor(0.5+t*frequency))
		case WaveSquare:
			sample = math.Copysign(1.0, math.Sin(2*math.Pi*frequency*t))
		}

		// Apply drive (soft clipping)
		sample = applyDrive(sample, cfg.Drive)

		// Apply envelope (ADSR)
		envelopeValue := applyEnvelope(t, cfg.Attack, cfg.Decay, cfg.Sustain, cfg.Release, cfg.Duration)
		sample *= envelopeValue

		// Apply resonant low-pass filter
		sample = applyResonantLowPassFilter(sample, frequency, cfg.FilterCutoff, cfg.FilterResonance, cfg.SampleRate)

		// Scale to 16-bit integer range [-32767, 32767]
		samples[i] = int(sample * 32767.0)
	}

	return samples
}

// applyResonatorEffect simulates resonating modes by adding harmonic overtones to the base sound
func applyResonatorEffect(samples []int, sampleRate int, duration float64) {
	numSamples := len(samples)
	resonance := 0.03 // A small amount of resonance

	for i := 0; i < numSamples; i++ {
		// Add harmonic overtones
		t := float64(i) / float64(sampleRate)
		harmonicFreq := 2.0 * math.Pi * 80.0 * t // Second mode at 80Hz
		thirdMode := 2.0 * math.Pi * 120.0 * t   // Third mode at 120Hz

		samples[i] += int(resonance * math.Sin(harmonicFreq) * 32767)
		samples[i] += int(resonance * math.Sin(thirdMode) * 32767)

		// Clip to 16-bit range
		if samples[i] > 32767 {
			samples[i] = 32767
		} else if samples[i] < -32767 {
			samples[i] = -32767
		}
	}
}

// generatePitchModulation simulates the effect of membrane tension shifting the pitch when struck
func generatePitchModulation(startFreq, endFreq float64, sampleRate int, duration float64) []float64 {
	numSamples := int(float64(sampleRate) * duration)
	pitchMod := make([]float64, numSamples)

	// Modulate the pitch upwards initially, then decay it
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		// Exponential pitch decay
		pitchMod[i] = 1.0 + 0.1*math.Exp(-3.0*t) // 10% initial pitch rise, decaying
	}

	return pitchMod
}

// mixNoiseWithEnvelope mixes noise with its own envelope into the generated kick drum samples
func mixNoiseWithEnvelope(samples []int, cfg *Config) {
	numSamples := len(samples)
	noiseEnvelope := generateNoiseEnvelope(cfg.NoiseAttack, cfg.NoiseDecay, numSamples, cfg.SampleRate)

	for i := range samples {
		noiseSample := generateNoiseSample(cfg.NoiseType, i, cfg.SampleRate)
		samples[i] += int(noiseSample * noiseEnvelope[i] * cfg.NoiseAmount * 32767.0)

		// Clip to 16-bit range
		if samples[i] > 32767 {
			samples[i] = 32767
		} else if samples[i] < -32767 {
			samples[i] = -32767
		}
	}
}

// generateNoiseEnvelope generates an envelope for noise based on attack and decay parameters
func generateNoiseEnvelope(attack, decay float64, numSamples, sampleRate int) []float64 {
	envelope := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		if t < attack {
			envelope[i] = t / attack
		} else {
			envelope[i] = math.Exp(-5.0 * (t - attack) / decay)
		}
	}

	return envelope
}

// generateNoiseSample generates a noise sample based on the noise type
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
	// Simple, naive pink noise generation
	pinkNoiseAccumulator += rand.Float64()*2 - 1
	return pinkNoiseAccumulator / float64(i+1)
}

// Brown noise generation using a simple filtering approach
var brownNoiseAccumulator float64

func generateBrownNoise(i int) float64 {
	// Simple, naive brown noise generation
	brownNoiseAccumulator += (rand.Float64()*2 - 1) * 0.1
	if brownNoiseAccumulator > 1 {
		brownNoiseAccumulator = 1
	} else if brownNoiseAccumulator < -1 {
		brownNoiseAccumulator = -1
	}
	return brownNoiseAccumulator
}

// applyDrive applies a simple distortion/drive effect by soft-clipping the signal
func applyDrive(sample, driveAmount float64) float64 {
	if driveAmount > 0 {
		return sample * (1 + driveAmount) / (1 + driveAmount*math.Abs(sample))
	}
	return sample
}

// applyEnvelope applies an Attack-Decay-Sustain-Release (ADSR) envelope to the signal
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

// applyResonantLowPassFilter applies a resonant low-pass filter to the signal
func applyResonantLowPassFilter(sample, freq, cutoff, resonance float64, sampleRate int) float64 {
	// Simple resonant low-pass filter (for demonstration purposes)
	if freq > cutoff {
		// Resonance effect (if any) and attenuation
		sample *= math.Exp(-resonance * (freq - cutoff) / float64(sampleRate))
		sample *= cutoff / freq
	}
	return sample
}
