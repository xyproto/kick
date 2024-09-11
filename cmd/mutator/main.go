package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"

	g "github.com/AllenDang/giu"
	"github.com/go-audio/wav"
	"github.com/xyproto/kick"
)

const (
	buttonSize    = 100
	numPads       = 16
	maxGenerations = 1000
)

// Kick drum settings for the selected pad
var (
	activePadIndex   int
	activeConfig     *kick.Config
	pads             [numPads]*kick.Config
	loadedWaveform   []int  // Loaded .wav file waveform data
	trainingOngoing  bool   // Indicates whether the genetic algorithm is running
	wavFilePath      string // Path to the .wav file
)

// Dropdown selection index for the waveform
var waveformSelectedIndex int32

// Load a .wav file and store the waveform for comparison
func loadWavFile() error {
	// Open and decode the .wav file
	file, err := os.Open(wavFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	buffer, err := decoder.FullPCMBuffer()
	if err != nil {
		return err
	}

	loadedWaveform = buffer.Data // Store the waveform data for comparison
	fmt.Printf("Loaded .wav file: %s\n", wavFilePath)
	return nil
}

// Compare two waveforms using Mean Squared Error (MSE)
func compareWaveforms(waveform1, waveform2 []int) float64 {
	if len(waveform1) != len(waveform2) {
		return math.Inf(1) // Return infinity if the waveforms are of different lengths
	}

	mse := 0.0
	for i := 0; i < len(waveform1); i++ {
		diff := float64(waveform1[i] - waveform2[i])
		mse += diff * diff
	}
	return mse / float64(len(waveform1))
}

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

// Function to optimize the settings using a genetic algorithm
func optimizeSettings() {
	if len(loadedWaveform) == 0 {
		fmt.Println("No .wav file loaded. Please load a .wav file first.")
		return
	}

	population := make([]*kick.Config, 100) // Initial population
	for i := 0; i < len(population); i++ {
		population[i] = kick.NewRandom()
	}

	bestConfig := activeConfig
	bestFitness := math.Inf(1)

	for generation := 0; generation < maxGenerations && trainingOngoing; generation++ {
		// Evaluate fitness of each individual
		for _, individual := range population {
			// Generate the current individual's kick waveform
			filePath, err := individual.GenerateKickTemp()
			if err != nil {
				fmt.Println("Error generating kick:", err)
				continue
			}
			defer os.Remove(filePath) // Clean up temp file

			// Load and compare the generated waveform with the target waveform
			generatedWaveform, err := loadWavWaveform(filePath)
			if err != nil {
				fmt.Println("Error loading generated .wav file:", err)
				continue
			}
			fitness := compareWaveforms(generatedWaveform, loadedWaveform)

			// Update best config if the fitness is better
			if fitness < bestFitness {
				bestFitness = fitness
				bestConfig = kick.CopyConfig(individual)
				if bestFitness < 1e-3 { // Consider this a global optimum
					fmt.Printf("Global optimum found at generation %d!\n", generation)
					trainingOngoing = false
					break
				}
			}
		}

		// Perform mutation and crossover for next generation
		for i := 0; i < len(population); i++ {
			mutateConfig(population[i])
		}

		fmt.Printf("Generation %d: Best fitness = %f\n", generation, bestFitness)
	}

	// Set the best configuration as the active one
	activeConfig = bestConfig
}

// Helper function to mutate a configuration (for the genetic algorithm)
func mutateConfig(cfg *kick.Config) {
	mutationRate := 0.1
	if rand.Float64() < mutationRate {
		cfg.Attack = cfg.Attack * (0.9 + rand.Float64()*0.2) // Mutate Attack, for example
	}
}

// Helper function to load a waveform from a .wav file
func loadWavWaveform(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	buffer, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	return buffer.Data, nil
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

	// Build the layout with the grid on the left, sliders and buttons on the right, and text input for the .wav file path below
	g.SingleWindow().Layout(
		g.Row(
			g.Column(padGrid...),
			createSlidersForSelectedPad(),
		),
		g.Row(
			g.InputText(&wavFilePath).Label("Enter path to .wav file"),
			g.Button("Load WAV").OnClick(func() {
				err := loadWavFile()
				if err != nil {
					fmt.Println("Failed to load .wav file:", err)
				}
			}),
			g.Button("Start Training").OnClick(func() {
				if !trainingOngoing {
					trainingOngoing = true
					go optimizeSettings()
				}
			}),
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
