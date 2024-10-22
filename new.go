package kick

import (
	"io"
	"math/rand"
)

// Generate random kick settings
func NewRandom() *Settings {
	cfg, _ := NewSettings(55.0, 30.0, 96000, 1.0, 16, nil)
	cfg.Attack = rand.Float64() * 0.02
	cfg.Decay = 0.2 + rand.Float64()*0.8
	cfg.Sustain = rand.Float64() * 0.5
	cfg.Release = 0.2 + rand.Float64()*0.5
	cfg.Drive = rand.Float64()
	cfg.FilterCutoff = 2000 + rand.Float64()*6000
	cfg.Sweep = rand.Float64() * 1.5
	cfg.PitchDecay = rand.Float64() * 1.5
	cfg.FadeDuration = rand.Float64() * 0.1
	if rand.Float64() < 0.1 {
		cfg.SmoothFrequencyTransitions = false
	} else {
		cfg.SmoothFrequencyTransitions = true
	}
	if rand.Float64() < 0.1 {
		cfg.WaveformType = rand.Intn(7)
	} else {
		cfg.WaveformType = rand.Intn(2)
	}
	return cfg
}

// NewExperimental with more extreme parameters for a truly experimental sound
func NewExperimental(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(80.0, 20.0, sampleRate, duration, bitDepth, output) // More extreme pitch decay
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSawtooth       // Sawtooth for a sharper, edgier sound
	cfg.Attack = 0.001                    // Very quick attack
	cfg.Decay = 0.7                       // Longer decay
	cfg.Release = 0.4                     // Extended release
	cfg.Drive = 0.8                       // Strong drive for distortion
	cfg.FilterCutoff = 3000               // Low cutoff for a dark, experimental tone
	cfg.Sweep = 1.2                       // Extreme sweep for dramatic pitch variation
	cfg.PitchDecay = 0.8                  // Heavily exaggerated pitch decay
	cfg.FadeDuration = 0.01               // 10ms fade in/out
	cfg.SmoothFrequencyTransitions = true // Disable smooth frequency transitions

	return cfg, nil
}

// NewLinnDrum emulates the LinnDrum bass drum, known for its punchy and iconic sound
func NewLinnDrum(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(60.0, 40.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSine
	cfg.Attack = 0.01 // Smooth attack
	cfg.Decay = 0.5   // Moderate decay for punchiness
	cfg.Sustain = 0.1 // Low sustain for a tight sound
	cfg.Release = 0.3
	cfg.Drive = 0.4
	cfg.FilterCutoff = 5000 // Balanced cutoff for clarity
	cfg.Sweep = 0.6
	cfg.PitchDecay = 0.4                  // Gentle pitch decay for depth
	cfg.FadeDuration = 0.02               // 20ms fade in/out for a punchy yet smooth sound
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}

// NewDeepHouse creates a kick drum perfect for Deep House music
func NewDeepHouse(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(45.0, 25.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSine
	cfg.Attack = 0.005                    // Soft attack for a smooth start
	cfg.Decay = 0.9                       // Long decay for a deep, lingering sound
	cfg.Sustain = 0.3                     // Medium sustain for warmth
	cfg.Release = 0.7                     // Extended release for a smooth tail
	cfg.Drive = 0.6                       // Moderate drive for warmth and body
	cfg.FilterCutoff = 3500               // Low cutoff for deep, smooth bass
	cfg.Sweep = 0.8                       // Gentle sweep to keep it subtle
	cfg.PitchDecay = 0.6                  // Slight pitch decay for that deep house feel
	cfg.FadeDuration = 0.03               // 30ms fade in/out for a more gradual sound
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}

func New606(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(65.0, 45.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSine
	cfg.Attack = 0.01
	cfg.Decay = 0.3
	cfg.Sustain = 0.1
	cfg.Release = 0.2
	cfg.Drive = 0.4
	cfg.FilterCutoff = 5000
	cfg.Sweep = 0.7
	cfg.PitchDecay = 0.5
	cfg.FadeDuration = 0.015              // 15ms fade in/out for a balanced sound
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}

func New707(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(60.0, 40.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveTriangle
	cfg.Attack = 0.005
	cfg.Decay = 0.3
	cfg.Sustain = 0.2
	cfg.Release = 0.2
	cfg.Drive = 0.3
	cfg.FilterCutoff = 5000
	cfg.Sweep = 0.6
	cfg.PitchDecay = 0.3
	cfg.FadeDuration = 0.01               // 10ms fade in/out
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}

func New808(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(55.0, 30.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSine
	cfg.Attack = 0.01
	cfg.Decay = 0.8
	cfg.Sustain = 0.2
	cfg.Release = 0.6
	cfg.Drive = 0.2
	cfg.FilterCutoff = 4000
	cfg.Sweep = 0.9
	cfg.PitchDecay = 0.5
	cfg.FadeDuration = 0.02               // 20ms fade in/out for a fuller sound
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}

func New909(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Settings, error) {
	cfg, err := NewSettings(70.0, 50.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveTriangle
	cfg.Attack = 0.002
	cfg.Decay = 0.2
	cfg.Sustain = 0.1
	cfg.Release = 0.3
	cfg.Drive = 0.4
	cfg.FilterCutoff = 8000
	cfg.Sweep = 0.7
	cfg.PitchDecay = 0.2
	cfg.FadeDuration = 0.015              // 15ms fade in/out to balance between smoothness and clarity
	cfg.SmoothFrequencyTransitions = true // Enable smooth frequency transitions

	return cfg, nil
}
