package kick

import (
	"io"
)

// NewExperimental with more extreme parameters for a truly experimental sound
func NewExperimental(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(80.0, 20.0, sampleRate, duration, bitDepth, output) // More extreme pitch decay
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSawtooth // Sawtooth for a sharper, edgier sound
	cfg.Attack = 0.001              // Very quick attack
	cfg.Decay = 0.7                 // Longer decay
	cfg.Release = 0.4               // Extended release
	cfg.Drive = 0.8                 // Strong drive for distortion
	cfg.FilterCutoff = 3000         // Low cutoff for a dark, experimental tone
	cfg.Sweep = 1.2                 // Extreme sweep for dramatic pitch variation
	cfg.PitchDecay = 0.8            // Heavily exaggerated pitch decay
	return cfg, nil
}

// NewLinnDrum emulates the LinnDrum bass drum, known for its punchy and iconic sound
func NewLinnDrum(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(60.0, 40.0, sampleRate, duration, bitDepth, output)
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
	cfg.PitchDecay = 0.4 // Gentle pitch decay for depth
	return cfg, nil
}

// NewDeepHouse creates a kick drum perfect for Deep House music
func NewDeepHouse(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(45.0, 25.0, sampleRate, duration, bitDepth, output)
	if err != nil {
		return nil, err
	}
	cfg.WaveformType = WaveSine
	cfg.Attack = 0.005      // Soft attack for a smooth start
	cfg.Decay = 0.9         // Long decay for a deep, lingering sound
	cfg.Sustain = 0.3       // Medium sustain for warmth
	cfg.Release = 0.7       // Extended release for a smooth tail
	cfg.Drive = 0.6         // Moderate drive for warmth and body
	cfg.FilterCutoff = 3500 // Low cutoff for deep, smooth bass
	cfg.Sweep = 0.8         // Gentle sweep to keep it subtle
	cfg.PitchDecay = 0.6    // Slight pitch decay for that deep house feel
	return cfg, nil
}

func New606(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(65.0, 45.0, sampleRate, duration, bitDepth, output)
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
	return cfg, nil
}

func New707(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(60.0, 40.0, sampleRate, duration, bitDepth, output)
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
	return cfg, nil
}

func New808(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(55.0, 30.0, sampleRate, duration, bitDepth, output)
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
	return cfg, nil
}

func New909(sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	cfg, err := NewConfig(70.0, 50.0, sampleRate, duration, bitDepth, output)
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
	return cfg, nil
}
