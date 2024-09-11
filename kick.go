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
	StartFreq        float64
	EndFreq          float64
	SampleRate       int
	Duration         float64
	WaveformType     int
	Attack           float64
	Decay            float64
	Sustain          float64
	Release          float64
	Drive            float64
	FilterCutoff     float64
	FilterResonance  float64
	Sweep            float64
	PitchDecay       float64
	NoiseType        int
	NoiseAmount      float64
	Output           io.WriteSeeker
	NumOscillators   int
	OscillatorLevels []float64
	SaturatorAmount  float64
	FilterBands      []float64
	BitDepth         int
}

var pinkNoiseAccumulator float64
var brownNoiseAccumulator float64

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
func PlayWav(filePath string) {
	cmd := exec.Command("mpv", filePath)
	fmt.Printf("Running command: %v\n", cmd.Args)
	err := cmd.Start()
	if err != nil {
		// Fallback to ffmpeg if mpv is not available
		cmd = exec.Command("ffmpeg", "-i", filePath, "-f", "null", "-")
		fmt.Printf("Running command: %v\n", cmd.Args)
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error playing sound with both mpv and ffmpeg:", err)
			return
		}
	}
	cmd.Wait()
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

// GenerateKickTemp generates a kick and writes it to a temporary file
func (cfg *Config) GenerateKickTemp() (string, error) {
	file, err := os.CreateTemp("", "kick_*.wav")
	if err != nil {
		return "", err
	}
	defer file.Close()

	cfg.Output = file

	if err := cfg.GenerateKick(); err != nil {
		return "", err
	}

	return file.Name(), nil
}

// Play generates a kick and plays it, removing the temporary file afterward
func (cfg *Config) Play() error {
	filePath, err := cfg.GenerateKickTemp()
	if err != nil {
		return err
	}
	defer os.Remove(filePath)
	PlayWav(filePath)
	return nil
}

func (cfg *Config) GenerateKick() error {
	samples := cfg.generateMultiOscillatorSamples()
	applySaturator(samples, cfg.SaturatorAmount)
	applyMultiBandFiltering(samples, cfg.FilterBands, cfg.SampleRate)

	if cfg.NoiseType != NoiseNone {
		mixNoise(samples, cfg)
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
			decayFactor := math.Pow(cfg.EndFreq/cfg.StartFreq, (t/cfg.Duration)*cfg.Sweep)
			frequency := cfg.StartFreq * decayFactor * pitchMod[i]

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

		samples[i] = int(totalSample * float64(int(1)<<(cfg.BitDepth-1)))
	}

	return samples
}

func applySaturator(samples []int, amount float64) {
	for i := range samples {
		sample := float64(samples[i]) / float64(1<<(16-1))
		sample = math.Tanh(sample * (1.0 + amount))
		samples[i] = int(sample * float64(1<<15))
	}
}

func applyMultiBandFiltering(samples []int, bands []float64, sampleRate int) {
	numSamples := len(samples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		frequency := 440.0 * math.Pow(2.0, t)

		if frequency < bands[0] {
			samples[i] = int(float64(samples[i]) * 0.9)
		} else if frequency < bands[1] {
			samples[i] = int(float64(samples[i]) * 0.8)
		} else if frequency < bands[2] {
			samples[i] = int(float64(samples[i]) * 0.7)
		} else {
			samples[i] = int(float64(samples[i]) * 0.6)
		}
	}
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

func applyDrive(sample, driveAmount float64) float64 {
	if driveAmount > 0 {
		return sample * (1 + driveAmount) / (1 + driveAmount*math.Abs(sample))
	}
	return sample
}

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
