package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"

	g "github.com/AllenDang/giu"
	"github.com/xyproto/kick"
)

const numPads = 16
const buttonSize = 100

// Kick drum settings for the selected pad
var selectedConfig *kick.Config
var activePadIndex int = 0 // Always start with the first pad selected

// Dropdown selection index for the waveform
var waveformSelectedIndex int32

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

// Function to mutate all pads based on a base config
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
	}
}

// Deep copy function for Config struct
func copyConfig(cfg *kick.Config) *kick.Config {
	newCfg := *cfg
	newCfg.OscillatorLevels = append([]float64(nil), cfg.OscillatorLevels...) // Deep copy
	return &newCfg
}

// Function to create a widget for a Config struct, with the color based on the config
func createPadWidget(cfg *kick.Config, padLabel string, padIndex int) g.Widget {
	return g.Style().SetColor(g.StyleColorButton, cfg.Color()).To(
		g.Column(
			g.Button(padLabel).Size(buttonSize, buttonSize).OnClick(func() {
				selectedConfig = cfg      // Set the selected pad
				activePadIndex = padIndex // Set the active pad index
			}),
			g.Button("Mutate").OnClick(func() {
				mutateAllPads(cfg)
			}),
			g.Button("Save").OnClick(func() {
				generateAndHandleKick(cfg, true)
			}),
		),
	)
}

// Pads state
var pads [numPads]*kick.Config

// Function to create sliders and dropdown for viewing and editing the selected pad's settings
func createSlidersForSelectedPad() g.Widget {
	if selectedConfig == nil {
		return g.Label("No pad selected")
	}

	// Convert float64 to float32 for sliders
	attack := float32(selectedConfig.Attack)
	decay := float32(selectedConfig.Decay)
	sustain := float32(selectedConfig.Sustain)
	release := float32(selectedConfig.Release)
	drive := float32(selectedConfig.Drive)
	filterCutoff := float32(selectedConfig.FilterCutoff)
	sweep := float32(selectedConfig.Sweep)
	pitchDecay := float32(selectedConfig.PitchDecay)

	// List of available waveforms
	waveforms := []string{"Sine", "Triangle", "Sawtooth", "Square", "Noise White", "Noise Pink", "Noise Brown"}
	waveformSelectedIndex = int32(selectedConfig.WaveformType)

	return g.Column(
		g.Label(fmt.Sprintf("Adjust Settings for Selected Pad: Pad %d", activePadIndex+1)),
		g.Row(
			g.Label("Waveform"),
			g.Combo("Waveform", waveforms[waveformSelectedIndex], waveforms, &waveformSelectedIndex).Size(150).OnChange(func() {
				selectedConfig.WaveformType = int(waveformSelectedIndex)
			}),
		),
		g.Row(
			g.Label("Attack"),
			g.SliderFloat(&attack, 0.0, 1.0).Size(150).OnChange(func() { selectedConfig.Attack = float64(attack) }),
		),
		g.Row(
			g.Label("Decay"),
			g.SliderFloat(&decay, 0.1, 1.0).Size(150).OnChange(func() { selectedConfig.Decay = float64(decay) }),
		),
		g.Row(
			g.Label("Sustain"),
			g.SliderFloat(&sustain, 0.0, 1.0).Size(150).OnChange(func() { selectedConfig.Sustain = float64(sustain) }),
		),
		g.Row(
			g.Label("Release"),
			g.SliderFloat(&release, 0.1, 1.0).Size(150).OnChange(func() { selectedConfig.Release = float64(release) }),
		),
		g.Row(
			g.Label("Drive"),
			g.SliderFloat(&drive, 0.0, 1.0).Size(150).OnChange(func() { selectedConfig.Drive = float64(drive) }),
		),
		g.Row(
			g.Label("Filter Cutoff"),
			g.SliderFloat(&filterCutoff, 1000, 8000).Size(150).OnChange(func() { selectedConfig.FilterCutoff = float64(filterCutoff) }),
		),
		g.Row(
			g.Label("Sweep"),
			g.SliderFloat(&sweep, 0.1, 2.0).Size(150).OnChange(func() { selectedConfig.Sweep = float64(sweep) }),
		),
		g.Row(
			g.Label("Pitch Decay"),
			g.SliderFloat(&pitchDecay, 0.1, 1.5).Size(150).OnChange(func() { selectedConfig.PitchDecay = float64(pitchDecay) }),
		),
		// Buttons under the sliders: Play, Mutate, Save
		g.Button("Play").OnClick(func() {
			filePath := generateAndHandleKick(selectedConfig, false)
			playKick(filePath)
		}),
		g.Button("Mutate All Pads").OnClick(func() {
			mutateAllPads(selectedConfig)
		}),
		g.Button("Save").OnClick(func() {
			generateAndHandleKick(selectedConfig, true)
		}),
	)
}

// Function to create the UI layout
func loop() {
	// Display the 16 pads in a 4x4 grid
	padGrid := []g.Widget{}
	padIndex := 0 // Declare padIndex before the loop
	for row := 0; row < 4; row++ {
		rowWidgets := []g.Widget{}
		for col := 0; col < 4; col++ {
			rowWidgets = append(rowWidgets, createPadWidget(pads[padIndex], fmt.Sprintf("Pad %d", padIndex+1), padIndex))
			padIndex++ // Increment padIndex
		}
		padGrid = append(padGrid, g.Row(rowWidgets...))
	}

	// Build the layout with the grid on the left and the sliders on the right
	g.SingleWindow().Layout(
		g.Row(
			g.Column(padGrid...),
			createSlidersForSelectedPad(),
		),
	)
}

func main() {
	// Initialize random settings for the 16 pads using kick.NewRandom()
	for i := 0; i < numPads; i++ {
		pads[i] = kick.NewRandom() // Generate random settings using the kick package
	}

	// Set the first pad as selected
	selectedConfig = pads[0]

	// Adjust the window size to fit the grid, buttons, and sliders better
	wnd := g.NewMasterWindow("Kick Drum Generator", 760, 640, g.MasterWindowFlagsNotResizable) // 40px taller, 40px narrower
	wnd.Run(loop)
}
