package main

import (
	"fmt"
	"math/rand"

	g "github.com/AllenDang/giu"
	"github.com/xyproto/kick"
)

const (
	buttonSize = 100
	numPads    = 16
)

// Kick drum settings for the selected pad
var (
	activePadIndex int
	activeConfig   *kick.Config
	pads           [numPads]*kick.Config
)

// Dropdown selection index for the waveform
var waveformSelectedIndex int32

// Function to mutate all pads based on a base config
func mutateAllPads(base *kick.Config) {
	mutationFactor := 0.4
	mutate := func(value float64) float64 {
		return value + (rand.Float64()-0.5)*mutationFactor*value
	}

	for i := 0; i < numPads; i++ {
		pads[i] = kick.CopyConfig(base)
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

// Function to create a widget for a Config struct, with the color based on the config
func createPadWidget(cfg *kick.Config, padLabel string, padIndex int) g.Widget {
	return g.Style().SetColor(g.StyleColorButton, cfg.Color()).To(
		g.Column(
			g.Button(padLabel).Size(buttonSize, buttonSize).OnClick(func() {
				activeConfig = cfg
				activePadIndex = padIndex
				err := cfg.Play()
				if err != nil {
					fmt.Println("Failed to play kick:", err)
				}
			}),
			g.Button("Mutate").OnClick(func() {
				mutateAllPads(cfg)
			}),
			g.Button("Save").OnClick(func() {
				_, err := cfg.SaveTo(".")
				if err != nil {
					fmt.Println("Failed to save kick:", err)
				}
			}),
		),
	)
}

// Function to create sliders and dropdown for viewing and editing the selected pad's settings
func createSlidersForSelectedPad() g.Widget {
	if activeConfig == nil {
		return g.Label("No pad selected")
	}

	// Convert float64 to float32 for sliders
	attack := float32(activeConfig.Attack)
	decay := float32(activeConfig.Decay)
	sustain := float32(activeConfig.Sustain)
	release := float32(activeConfig.Release)
	drive := float32(activeConfig.Drive)
	filterCutoff := float32(activeConfig.FilterCutoff)
	sweep := float32(activeConfig.Sweep)
	pitchDecay := float32(activeConfig.PitchDecay)

	// List of available waveforms
	waveforms := []string{"Sine", "Triangle", "Sawtooth", "Square", "Noise White", "Noise Pink", "Noise Brown"}
	waveformSelectedIndex = int32(activeConfig.WaveformType)

	return g.Column(
		g.Label(fmt.Sprintf("Adjust Settings for Selected Pad: Pad %d", activePadIndex+1)),
		g.Row(
			g.Label("Waveform"),
			g.Combo("Waveform", waveforms[waveformSelectedIndex], waveforms, &waveformSelectedIndex).Size(150).OnChange(func() {
				activeConfig.WaveformType = int(waveformSelectedIndex)
			}),
		),
		g.Row(
			g.Label("Attack"),
			g.SliderFloat(&attack, 0.0, 1.0).Size(150).OnChange(func() { activeConfig.Attack = float64(attack) }),
		),
		g.Row(
			g.Label("Decay"),
			g.SliderFloat(&decay, 0.1, 1.0).Size(150).OnChange(func() { activeConfig.Decay = float64(decay) }),
		),
		g.Row(
			g.Label("Sustain"),
			g.SliderFloat(&sustain, 0.0, 1.0).Size(150).OnChange(func() { activeConfig.Sustain = float64(sustain) }),
		),
		g.Row(
			g.Label("Release"),
			g.SliderFloat(&release, 0.1, 1.0).Size(150).OnChange(func() { activeConfig.Release = float64(release) }),
		),
		g.Row(
			g.Label("Drive"),
			g.SliderFloat(&drive, 0.0, 1.0).Size(150).OnChange(func() { activeConfig.Drive = float64(drive) }),
		),
		g.Row(
			g.Label("Filter Cutoff"),
			g.SliderFloat(&filterCutoff, 1000, 8000).Size(150).OnChange(func() { activeConfig.FilterCutoff = float64(filterCutoff) }),
		),
		g.Row(
			g.Label("Sweep"),
			g.SliderFloat(&sweep, 0.1, 2.0).Size(150).OnChange(func() { activeConfig.Sweep = float64(sweep) }),
		),
		g.Row(
			g.Label("Pitch Decay"),
			g.SliderFloat(&pitchDecay, 0.1, 1.5).Size(150).OnChange(func() { activeConfig.PitchDecay = float64(pitchDecay) }),
		),
		// Buttons under the sliders: Play, Mutate all, Save
		g.Button("Play").OnClick(func() {
			err := activeConfig.Play()
			if err != nil {
				fmt.Println("Failed to play kick:", err)
			}
		}),
		g.Button("Mutate all").OnClick(func() {
			mutateAllPads(activeConfig)
		}),
		g.Button("Save").OnClick(func() {
			_, err := activeConfig.SaveTo(".")
			if err != nil {
				fmt.Println("Failed to save kick:", err)
			}
		}),
	)
}

// Function to create the UI layout
func loop() {
	// Display the 16 pads in a 4x4 grid
	padGrid := []g.Widget{}
	padIndex := 0
	for row := 0; row < 4; row++ {
		rowWidgets := []g.Widget{}
		for col := 0; col < 4; col++ {
			rowWidgets = append(rowWidgets, createPadWidget(pads[padIndex], fmt.Sprintf("Pad %d", padIndex+1), padIndex))
			padIndex++
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
		pads[i] = kick.NewRandom()
	}

	// Set the first pad as selected
	activePadIndex = 0
	activeConfig = pads[activePadIndex]

	// Adjust the window size to fit the grid, buttons, and sliders better
	wnd := g.NewMasterWindow("Kick Drum Generator", 760, 640, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
