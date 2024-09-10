package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xyproto/kick"
)

const version = "1.7.0"

func main() {
	// Command-line flags for sound customization
	kick808 := flag.Bool("808", false, "Generate a kick.wav like an 808 kick drum")
	kick909 := flag.Bool("909", false, "Generate a kick.wav like a 909 kick drum")
	kick707 := flag.Bool("707", false, "Generate a kick.wav like a 707 kick drum")
	kickKORG := flag.Bool("korg", false, "Generate a kick.wav with KORG-style characteristics")
	kickSOMA := flag.Bool("soma", false, "Generate a kick.wav with SOMA-style characteristics")
	noiseType := flag.String("noise", "none", "Type of noise to mix in (none, white, pink, brown)")
	noiseAmount := flag.Float64("noiseamount", 0.0, "Amount of noise to mix in (0.0 to 1.0)")
	length := flag.Float64("length", 1000, "Length of the kick drum sample in milliseconds")
	quality := flag.Int("quality", 44, "Sample rate in kHz (44, 48, 96, or 192)")
	attack := flag.Float64("attack", 0.005, "Attack time in seconds")
	decay := flag.Float64("decay", 0.3, "Decay time in seconds")
	sustain := flag.Float64("sustain", 0.2, "Sustain level (0.0 to 1.0)")
	release := flag.Float64("release", 0.2, "Release time in seconds")
	waveform := flag.Int("waveform", kick.WaveSine, "Waveform type (0: Sine, 1: Triangle)")
	drive := flag.Float64("drive", 0.0, "Amount of distortion/drive (0.0 to 1.0)")
	filterCutoff := flag.Float64("filter", 10000.0, "Low-pass filter cutoff frequency (Hz)")
	sweep := flag.Float64("sweep", 1.0, "Pitch sweep rate (0.1 to 1.0)")
	pitchDecay := flag.Float64("pitchdecay", 0.1, "Pitch envelope decay (shorter for punch, longer for boom)")
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
		fmt.Println("Invalid sample rate. Choose 44, 48, 96, or 192 kHz.")
		os.Exit(1)
	}

	// Determine characteristics for different kick drum styles
	if *kick808 {
		*waveform = kick.WaveSine
		*attack = 0.01
		*decay = 0.8
		*release = 0.6
		*drive = 0.2
		*filterCutoff = 5000
		*sweep = 0.9
		*pitchDecay = 0.5
		fmt.Println("Generating 808 kick with deep sub-bass and smooth characteristics.")
	} else if *kick909 {
		*waveform = kick.WaveTriangle
		*attack = 0.002
		*decay = 0.2
		*release = 0.15
		*drive = 0.4
		*filterCutoff = 8000
		*sweep = 0.7
		*pitchDecay = 0.2
		fmt.Println("Generating 909 kick with punchy, mid-range presence and quick decay.")
	} else if *kick707 {
		*waveform = kick.WaveTriangle
		*attack = 0.005
		*decay = 0.3
		*release = 0.2
		*drive = 0.3
		*filterCutoff = 7000
		*sweep = 0.6
		*pitchDecay = 0.3
		fmt.Println("Generating 707 kick with a classic, shorter punchy sound.")
	} else if *kickKORG {
		*waveform = kick.WaveSine
		*attack = 0.008
		*decay = 0.5
		*release = 0.4
		*drive = 0.6
		*filterCutoff = 6000
		*sweep = 0.8
		*pitchDecay = 0.4
		fmt.Println("Generating KORG-style kick with distinct analog punch.")
	} else if *kickSOMA {
		*waveform = kick.WaveTriangle
		*attack = 0.004
		*decay = 0.6
		*release = 0.3
		*drive = 0.5
		*filterCutoff = 5500
		*sweep = 0.7
		*pitchDecay = 0.4
		fmt.Println("Generating SOMA-style kick with experimental texture.")
	}

	// Determine noise type
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

	// Open the output file
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outFile.Close()

	// Generate the kick drum sound and write to the output file
	err = kick.GenerateKickWithEffects(
		150.0, 40.0, sampleRate, *length/1000.0,
		*waveform, *attack, *decay, *sustain, *release,
		*drive, *filterCutoff, *sweep, *pitchDecay, noise, *noiseAmount, outFile,
	)
	if err != nil {
		fmt.Println("Failed to generate kick:", err)
		return
	}

	fmt.Println("Kick drum sound generated and written to", *outputFile)
}
