package kick

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"image/color"
	"io"
	"math"
	"math/rand"
	"os"
	"os/exec"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

const (
	WaveSine = iota
	WaveTriangle
	WaveSawtooth
	WaveSquare
	WaveNoiseWhite
	WaveNoisePink
	WaveNoiseBrown
)

const (
	NoiseNone = iota
	NoiseWhite
	NoisePink
	NoiseBrown
)

type Config struct {
	StartFreq                  float64
	EndFreq                    float64
	SampleRate                 int
	Duration                   float64
	WaveformType               int
	Attack                     float64
	Decay                      float64
	Sustain                    float64
	Release                    float64
	Drive                      float64
	FilterCutoff               float64
	FilterResonance            float64
	Sweep                      float64
	PitchDecay                 float64
	NoiseType                  int
	NoiseAmount                float64
	Output                     io.WriteSeeker
	NumOscillators             int
	OscillatorLevels           []float64
	SaturatorAmount            float64
	FilterBands                []float64
	BitDepth                   int
	FadeDuration               float64
	SmoothFrequencyTransitions bool
}

var (
	brownNoiseAccumulator float64
	pinkNoiseAccumulator  float64
)

func NewConfig(startFreq, endFreq float64, sampleRate int, duration float64, bitDepth int, output io.WriteSeeker) (*Config, error) {
	if sampleRate <= 0 || duration <= 0 {
		return nil, errors.New("invalid sample rate or duration")
	}

	return &Config{
		StartFreq:        startFreq,
		EndFreq:          endFreq,
		SampleRate:       sampleRate,
		Duration:         duration,
		WaveformType:     WaveSine,
		Attack:           0.005,
		Decay:            0.3,
		Sustain:          0.2,
		Release:          0.3,
		Drive:            0.2,
		FilterCutoff:     5000,
		FilterResonance:  0.2,
		Sweep:            0.7,
		PitchDecay:       0.4,
		NoiseType:        NoiseNone,
		NoiseAmount:      0.0,
		Output:           output,
		NumOscillators:   1,
		OscillatorLevels: []float64{1.0},
		SaturatorAmount:  0.3,
		FilterBands:      []float64{200.0, 1000.0, 3000.0},
		BitDepth:         bitDepth,
	}, nil
}

// CopyConfig creates a deep copy of a Config struct
func CopyConfig(cfg *Config) *Config {
	newCfg := *cfg
	newCfg.OscillatorLevels = append([]float64(nil), cfg.OscillatorLevels...) // Deep copy the slice
	return &newCfg
}

// PlayWav plays a WAV file using mpv or ffmpeg
func PlayWav(filePath string) error {
	cmd := exec.Command("mpv", filePath)
	fmt.Printf("Running command: %v\n", cmd.Args)
	err := cmd.Start()
	if err != nil {
		// Fallback to ffmpeg if mpv is not available
		cmd = exec.Command("ffmpeg", "-i", filePath, "-f", "null", "-")
		fmt.Printf("Running command: %v\n", cmd.Args)
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("Error playing sound with both mpv and ffmpeg: %v", err)
		}
	}
	cmd.Wait()
	return nil
}

// PlayWaveform writes the waveform to a temporary .wav file and plays it using mpv or ffmpeg
func PlayWaveform(wave []int, sampleRate int) error {
	// Create a temporary file
	file, err := os.CreateTemp("", "waveform_*.wav")
	if err != nil {
		return fmt.Errorf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	// Write the waveform as a .wav file
	buffer := &audio.IntBuffer{
		Data:           wave,
		Format:         &audio.Format{SampleRate: sampleRate, NumChannels: 1},
		SourceBitDepth: 16,
	}

	encoder := wav.NewEncoder(file, sampleRate, 16, 1, 1)
	if err := encoder.Write(buffer); err != nil {
		return fmt.Errorf("Error writing waveform to WAV file: %v", err)
	}
	encoder.Close()

	// Play the WAV file
	return PlayWav(file.Name())
}

// SaveTo saves the generated kick to a specified directory, avoiding filename collisions
func (cfg *Config) SaveTo(directory string) (string, error) {
	n := 1
	var fileName string
	for {
		fileName = fmt.Sprintf("%s/kick%d.wav", directory, n)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			break
		}
		n++
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	cfg.Output = file

	if err := cfg.GenerateKick(); err != nil {
		return "", err
	}

	return fileName, nil
}

// Play generates a kick and plays it directly using mpv or ffmpeg
func (cfg *Config) Play() error {
	// Generate the kick waveform in memory
	waveform, err := cfg.GenerateKickInMemory()
	if err != nil {
		return err
	}

	// Play the generated waveform directly
	err = PlayWaveform(waveform, cfg.SampleRate)
	if err != nil {
		return fmt.Errorf("Error playing generated waveform: %v", err)
	}
	return nil
}

func (cfg *Config) GenerateKick() error {
	samples := cfg.generateMultiOscillatorSamples()

	applySaturator(samples, cfg.SaturatorAmount)

	applyMultiBandFiltering(samples, cfg.FilterBands, cfg.SampleRate)

	if cfg.NoiseType != NoiseNone {
		mixNoise(samples, cfg)
	}

	// Apply fade in/out if FadeDuration is set
	if cfg.FadeDuration > 0 {
		applyFadeInOut(samples, cfg.SampleRate, cfg.FadeDuration)
	}

	buffer := &audio.IntBuffer{
		Data:           samples,
		Format:         &audio.Format{SampleRate: cfg.SampleRate, NumChannels: 1},
		SourceBitDepth: cfg.BitDepth,
	}

	encoder := wav.NewEncoder(cfg.Output, cfg.SampleRate, cfg.BitDepth, 1, 1)
	if err := encoder.Write(buffer); err != nil {
		return err
	}

	return encoder.Close()
}

func (cfg *Config) generateMultiOscillatorSamples() []int {
	numSamples := int(float64(cfg.SampleRate) * cfg.Duration)
	samples := make([]int, numSamples)

	pitchMod := generatePitchModulation(cfg.StartFreq, cfg.EndFreq, cfg.SampleRate, cfg.Duration)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(cfg.SampleRate)
		var totalSample float64

		for oscIndex := 0; oscIndex < cfg.NumOscillators; oscIndex++ {
			var frequency float64
			if cfg.SmoothFrequencyTransitions {
				// Smoother frequency transition
				decayFactor := math.Pow(cfg.EndFreq/cfg.StartFreq, (t/cfg.Duration)*cfg.Sweep)
				frequency = cfg.StartFreq * decayFactor * pitchMod[i]
			} else {
				// Abrupt frequency transition
				if t < cfg.Duration/2 {
					frequency = cfg.StartFreq
				} else {
					frequency = cfg.EndFreq
				}
			}

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
			case WaveNoiseWhite:
				sample = rand.Float64()*2 - 1
			case WaveNoisePink:
				sample = generatePinkNoise(i)
			case WaveNoiseBrown:
				sample = generateBrownNoise(i)
			}

			sample = applyDrive(sample, cfg.Drive)
			envelopeValue := applyEnvelope(t, cfg.Attack, cfg.Decay, cfg.Sustain, cfg.Release, cfg.Duration)
			sample *= envelopeValue

			sample *= cfg.OscillatorLevels[oscIndex]
			totalSample += sample
		}

		samples[i] = int(totalSample * float64(int(1)<<cfg.BitDepth-1))
	}

	return samples
}

func generatePitchModulation(startFreq, endFreq float64, sampleRate int, duration float64) []float64 {
	numSamples := int(float64(sampleRate) * duration)
	pitchMod := make([]float64, numSamples)

	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		pitchMod[i] = 1.0 + 0.1*math.Exp(-3.0*t)
	}

	return pitchMod
}

func generatePinkNoise(i int) float64 {
	pinkNoiseAccumulator += rand.Float64()*2 - 1
	return pinkNoiseAccumulator / float64(i+1)
}

func generateBrownNoise(i int) float64 {
	brownNoiseAccumulator += (rand.Float64()*2 - 1) * 0.1
	if brownNoiseAccumulator > 1 {
		brownNoiseAccumulator = 1
	} else if brownNoiseAccumulator < -1 {
		brownNoiseAccumulator = -1
	}
	return brownNoiseAccumulator
}

func mixNoise(samples []int, cfg *Config) {
	for i := range samples {
		var noiseSample float64
		switch cfg.NoiseType {
		case NoiseWhite:
			noiseSample = rand.Float64()*2 - 1
		case NoisePink:
			noiseSample = generatePinkNoise(i)
		case NoiseBrown:
			noiseSample = generateBrownNoise(i)
		}

		samples[i] += int(noiseSample * cfg.NoiseAmount * float64(int(1)<<(cfg.BitDepth-1)))

		if samples[i] > 32767 {
			samples[i] = 32767
		} else if samples[i] < -32767 {
			samples[i] = -32767
		}
	}
}

// Color returns a color that very approximately represents the current kick config
func (cfg *Config) Color() color.RGBA {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%d%f%f%f%f%f%f%f%f", cfg.WaveformType, cfg.Attack, cfg.Decay, cfg.Sustain, cfg.Release, cfg.Drive, cfg.FilterCutoff, cfg.Sweep, cfg.PitchDecay)))
	hashBytes := hasher.Sum(nil)

	// Convert the first few bytes of the hash into an RGB color
	r := hashBytes[0]
	g := hashBytes[1]
	b := hashBytes[2]
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

// GenerateKickInMemory generates the kick waveform and returns it as a slice of integers.
func (cfg *Config) GenerateKickInMemory() ([]int, error) {
	samples := cfg.generateMultiOscillatorSamples()
	applySaturator(samples, cfg.SaturatorAmount)
	applyMultiBandFiltering(samples, cfg.FilterBands, cfg.SampleRate)

	if cfg.NoiseType != NoiseNone {
		mixNoise(samples, cfg)
	}

	return samples, nil
}
