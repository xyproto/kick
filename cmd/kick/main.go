package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xyproto/kick"
)

const version = "1.8.1"

func main() {
	// Command-line flags for sound customization
	kick808 := flag.Bool("808", false, "Generate a kick.wav like an 808 kick drum")
	kick909 := flag.Bool("909", false, "Generate a kick.wav like a 909 kick drum")
	kick707 := flag.Bool("707", false, "Generate a kick.wav like a 707 kick drum")
	kick606 := flag.Bool("606", false, "Generate a kick.wav like a 606 kick drum")
	kickLinnDrum := flag.Bool("linn", false, "Generate a kick.wav like a LinnDrum kick drum")
	kickDeepHouse := flag.Bool("deephouse", false, "Generate a deep house kick drum")
	kickExperimental := flag.Bool("experimental", false, "Generate a kick.wav with experimental-style characteristics")
	noiseType := flag.String("noise", "none", "Type of noise to mix in (none, white, pink, brown)")
	noiseAmount := flag.Float64("noiseamount", 0.0, "Amount of noise to mix in (0.0 to 1.0)")
	length := flag.Float64("length", 1000, "Length of the kick drum sample in milliseconds")
	quality := flag.Int("quality", 96, "Sample rate in kHz (48 or 96)")
	bitDepth := flag.Int("bitdepth", 16, "Bit depth of the audio (16 or 24)")
	waveform := flag.Int("waveform", kick.WaveSine, "Waveform type (0: Sine, 1: Triangle, 2: Sawtooth, 3: Square)")
	attack := flag.Float64("attack", 0.003, "Attack time in seconds")
	decay := flag.Float64("decay", 0.3, "Decay time in seconds")
	sustain := flag.Float64("sustain", 0.1, "Sustain level (0.0 to 1.0)")
	release := flag.Float64("release", 0.15, "Release time in seconds")
	sweep := flag.Float64("sweep", 0.8, "Pitch sweep rate")
	filterCutoff := flag.Float64("filter", 5000.0, "Low-pass filter cutoff frequency (Hz)")
	pitchDecay := flag.Float64("pitchdecay", 0.2, "Pitch envelope decay time")
	drive := flag.Float64("drive", 0.1, "Amount of distortion/drive")
	numOscillators := flag.Int("numoscillators", 1, "Number of oscillators for layering")
	oscillatorLevels := flag.String("oscillatorlevels", "1.0", "Comma-separated levels for each oscillator")
	saturatorAmount := flag.Float64("saturator", 0.3, "Amount of saturation to apply")
	filterBands := flag.String("filterbands", "200,1000,3000", "Comma-separated multi-band filter cutoff frequencies")
	outputFile := flag.String("o", "kick.wav", "Output file path")
	showVersion := flag.Bool("version", false, "Show the current version")
	showHelp := flag.Bool("help", false, "Display this help")

	flag.Parse()

	// Display help or version
	if *showHelp {
		flag.Usage()
		return
	}
	if *showVersion {
		fmt.Println("kick version", version)
		return
	}

	// Set sample rate based on the quality flag
	var sampleRate int
	switch *quality {
	case 44:
		sampleRate = 44100
	case 48:
		sampleRate = 48000
	case 96:
		sampleRate = 96000
	case 192:
		sampleRate = 192000
	default:
		fmt.Println("Invalid sample rate. Choose 48 or 96 kHz.")
		os.Exit(1)
	}

	// Open the output file
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outFile.Close()

	// Use the appropriate constructor based on the selected kick style
	var cfg *kick.Settings
	switch {
	case *kick808:
		cfg, err = kick.New808(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating 808 kick with deep sub-bass and smooth characteristics.")
	case *kick909:
		cfg, err = kick.New909(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating 909 kick with punchy, mid-range presence and quick decay.")
	case *kick707:
		cfg, err = kick.New707(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating 707 kick with a classic, shorter punchy sound.")
	case *kick606:
		cfg, err = kick.New606(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating 606 kick with a punchy, shorter sound.")
	case *kickLinnDrum:
		cfg, err = kick.NewLinnDrum(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating LinnDrum kick with an iconic, punchy sound.")
	case *kickDeepHouse:
		cfg, err = kick.NewDeepHouse(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating Deep House kick with smooth, warm bass.")
	case *kickExperimental:
		cfg, err = kick.NewExperimental(sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating experimental-style kick with unique texture.")
	default:
		cfg, err = kick.NewSettings(150.0, 40.0, sampleRate, *length/1000.0, *bitDepth, outFile)
		fmt.Println("Generating default kick with user-defined characteristics.")
	}

	if err != nil {
		fmt.Println("Error creating config:", err)
		os.Exit(1)
	}

	// Set additional parameters from command-line flags
	cfg.WaveformType = *waveform
	cfg.Attack = *attack
	cfg.Decay = *decay
	cfg.Sustain = *sustain
	cfg.Release = *release
	cfg.Sweep = *sweep
	cfg.FilterCutoff = *filterCutoff
	cfg.PitchDecay = *pitchDecay
	cfg.Drive = *drive
	cfg.NumOscillators = *numOscillators
	cfg.OscillatorLevels = parseCommaSeparatedFloats(*oscillatorLevels)
	cfg.SaturatorAmount = *saturatorAmount
	cfg.FilterBands = parseCommaSeparatedFloats(*filterBands)
	cfg.FadeDuration = 0.01
	cfg.SmoothFrequencyTransitions = true

	// Set noise type
	var noise int
	switch *noiseType {
	case "white":
		noise = kick.NoiseWhite
	case "pink":
		noise = kick.NoisePink
	case "brown":
		noise = kick.NoiseBrown
	case "none":
		noise = kick.NoiseNone
	default:
		fmt.Println("Invalid noise type. Choose from: none, white, pink, brown.")
		os.Exit(1)
	}
	cfg.NoiseType = noise
	cfg.NoiseAmount = *noiseAmount

	// Generate the kick drum sound
	if err := cfg.GenerateKick(); err != nil {
		fmt.Println("Failed to generate kick:", err)
		return
	}

	fmt.Println("Kick drum sound generated and written to", *outputFile)
}

// parseCommaSeparatedFloats parses a comma-separated string into a slice of float64s
func parseCommaSeparatedFloats(input string) []float64 {
	var result []float64
	for _, s := range strings.Split(input, ",") {
		value, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			fmt.Println("Error parsing value:", s)
			continue
		}
		result = append(result, value)
	}
	return result
}
