package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"os/exec"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/xyproto/kick"
)

const numPads = 16
const buttonSize = 100

// Function to play or save the generated kick
func generateAndHandleKick(cfg *kick.Config, save bool) string {
	var fileName string
	if save {
		n := 1
		for {
			fileName = fmt.Sprintf("kick%d.wav", n)
			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				break
			}
			n++
		}
		tmpFile, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return ""
		}
		defer tmpFile.Close()
		cfg.Output = tmpFile
		fmt.Printf("Saved kick drum to %s\n", fileName)
	} else {
		tmpFile, err := os.CreateTemp("", "kick_*.wav")
		if err != nil {
			fmt.Println("Failed to create temporary file:", err)
			return ""
		}
		defer tmpFile.Close()
		cfg.Output = tmpFile
		fileName = tmpFile.Name()
	}

	if err := cfg.GenerateKick(); err != nil {
		fmt.Println("Failed to generate kick:", err)
		return ""
	}

	return fileName
}

// Function to play the generated kick using mpv or ffmpeg
func playKick(filePath string) {
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

// Generate random kick.Config settings
func randomKickSettings() *kick.Config {
	cfg, _ := kick.NewConfig(55.0, 30.0, 96000, 1.0, 16, nil)
	cfg.Attack = rand.Float64() * 0.02
	cfg.Decay = 0.2 + rand.Float64()*0.8
	cfg.Sustain = rand.Float64() * 0.5
	cfg.Release = 0.2 + rand.Float64()*0.5
	cfg.Drive = rand.Float64()
	cfg.FilterCutoff = 2000 + rand.Float64()*6000
	cfg.Sweep = rand.Float64() * 1.5
	cfg.PitchDecay = rand.Float64() * 1.5

	if rand.Float64() < 0.1 {
		cfg.WaveformType = rand.Intn(7)
	} else {
		cfg.WaveformType = rand.Intn(1)
	}

	return cfg
}

// Function to get a waveform abbreviation
func waveformAbbreviation(waveformType int) string {
	switch waveformType {
	case kick.WaveSine:
		return "sin"
	case kick.WaveTriangle:
		return "tri"
	case kick.WaveSawtooth:
		return "saw"
	case kick.WaveSquare:
		return "sqr"
	case kick.WaveNoiseWhite:
		return "nwh"
	case kick.WaveNoisePink:
		return "npk"
	case kick.WaveNoiseBrown:
		return "nbr"
	default:
		return "unk"
	}
}

// Mutate all the settings to create significant variations based on a base config
func mutateAllPads(base *kick.Config) {
	mutationFactor := 0.4
	mutate := func(value float64) float64 {
		return value + (rand.Float64()-0.5)*mutationFactor*value
	}

	for i := 0; i < numPads; i++ {
		pads[i] = copyConfig(base)
		pads[i].Attack = mutate(base.Attack)
		pads[i].Decay = mutate(base.Decay)
		pads[i].Sustain = mutate(base.Sustain)
		pads[i].Release = mutate(base.Release)
		pads[i].Drive = mutate(base.Drive)
		pads[i].FilterCutoff = mutate(base.FilterCutoff)
		pads[i].Sweep = mutate(base.Sweep)
		pads[i].PitchDecay = mutate(base.PitchDecay)

		// In 90% of cases, keep the waveform as Sine or Triangle; in 10%, mutate to any waveform (0 to 7)
		if rand.Float64() < 0.1 {
			pads[i].WaveformType = rand.Intn(7) // Randomly choose a new waveform
		} else {
			pads[i].WaveformType = rand.Intn(2) // Keep as Sine or Triangle
		}

		// Output a summary of the settings being used for each pad
		fmt.Printf("Pad %d: Waveform=%s, Attack=%.3f, Decay=%.3f, Sustain=%.3f, Release=%.3f, Drive=%.3f, FilterCutoff=%.3f, Sweep=%.3f, PitchDecay=%.3f\n",
			i+1, waveformAbbreviation(pads[i].WaveformType), pads[i].Attack, pads[i].Decay, pads[i].Sustain, pads[i].Release, pads[i].Drive, pads[i].FilterCutoff, pads[i].Sweep, pads[i].PitchDecay)
	}
}

// Deep copy function for Config struct
func copyConfig(cfg *kick.Config) *kick.Config {
	newCfg := *cfg
	newCfg.OscillatorLevels = append([]float64(nil), cfg.OscillatorLevels...) // Deep copy
	return &newCfg
}

// Generate the color for the button based on the kick settings
func generateColor(cfg *kick.Config) color.RGBA {
	r := uint8(cfg.Drive * 255)
	g := uint8(cfg.Sweep / 1.5 * 255)
	b := uint8(cfg.PitchDecay / 1.5 * 255)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

// Function to create a widget for a Config struct, consisting of a Play Pad button, mutate button, and save button
func createPadWidget(cfg *kick.Config, padLabel string) g.Widget {
	return g.Column(
		g.Style().SetColor(g.StyleColorButton, generateColor(cfg)).To(
			g.Button(padLabel).Size(buttonSize, buttonSize).OnClick(func() {
				filePath := generateAndHandleKick(cfg, false)
				playKick(filePath)
			}),
		),
		g.Button("mutate").OnClick(func() {
			// Mutate all pads to be variations of the selected pad
			mutateAllPads(cfg)
		}),
		g.Button("save").OnClick(func() {
			// Save the kick drum as kickN.wav
			generateAndHandleKick(cfg, true)
		}),
	)
}

// Pads state
var pads [numPads]*kick.Config

// Function to create the UI layout
func loop() {
	layout := []g.Widget{
		g.Label("Kick Drum Generator"),
	}

	// Display the 16 pads in a 4x4 grid
	for row := 0; row < 4; row++ {
		rowWidgets := []g.Widget{}
		for col := 0; col < 4; col++ {
			padIndex := row*4 + col
			rowWidgets = append(rowWidgets, createPadWidget(pads[padIndex], fmt.Sprintf("Pad %d", padIndex+1)))
		}
		layout = append(layout, g.Row(rowWidgets...))
	}

	// Build the layout
	g.SingleWindow().Layout(layout...)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize random settings for the 16 pads
	for i := 0; i < numPads; i++ {
		pads[i] = randomKickSettings()
	}

	// Adjust the window size to fit the grid and buttons better
	wnd := g.NewMasterWindow("Kick Drum Generator", 440, 650, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
