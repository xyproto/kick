package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xyproto/kick"
)

const version = "1.6.0"

func main() {
	// Command-line flags for sound customization
	kick808 := flag.Bool("808", false, "Generate a kick.wav like an 808 kick drum")
	kick909 := flag.Bool("909", false, "Generate a kick.wav like a 909 kick drum")
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

	// Determine characteristics for 808 or 909 kick
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
		*drive, *filterCutoff, *sweep, *pitchDecay, outFile,
	)
	if err != nil {
		fmt.Println("Failed to generate kick:", err)
		return
	}

	fmt.Println("Kick drum sound generated and written to", *outputFile)
}
