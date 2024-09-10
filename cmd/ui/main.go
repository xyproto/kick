package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/xyproto/kick"
)

// Function to play the generated kick using mpv or ffmpeg
func playKick(filePath string) {
	// Try using mpv first
	cmd := exec.Command("mpv", filePath)
	err := cmd.Start()
	if err != nil {
		// Fallback to ffmpeg if mpv is not available
		cmd = exec.Command("ffmpeg", "-i", filePath, "-f", "null", "-")
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error playing sound with both mpv and ffmpeg:", err)
			return
		}
	}
	cmd.Wait()
	fmt.Println("Finished playing sound")
}

func main() {
	// Create the app and a window
	a := app.New()
	w := a.NewWindow("Kick Drum Generator")

	// Default parameters
	var sampleRate int = 96000
	var length float64 = 1000
	var bitDepth int = 16
	var waveformType int = kick.WaveSine
	var attack float64 = 0.005
	var decay float64 = 0.5
	var sustain float64 = 0.3
	var release float64 = 0.6
	var sweep float64 = 0.9
	var drive float64 = 0.4
	var filterCutoff float64 = 4000
	var pitchDecay float64 = 0.5

	// Create input fields for parameters
	lengthEntry := widget.NewEntry()
	lengthEntry.SetText(strconv.FormatFloat(length, 'f', 2, 64))

	attackEntry := widget.NewEntry()
	attackEntry.SetText(strconv.FormatFloat(attack, 'f', 3, 64))

	decayEntry := widget.NewEntry()
	decayEntry.SetText(strconv.FormatFloat(decay, 'f', 3, 64))

	sustainEntry := widget.NewEntry()
	sustainEntry.SetText(strconv.FormatFloat(sustain, 'f', 2, 64))

	releaseEntry := widget.NewEntry()
	releaseEntry.SetText(strconv.FormatFloat(release, 'f', 3, 64))

	driveEntry := widget.NewEntry()
	driveEntry.SetText(strconv.FormatFloat(drive, 'f', 2, 64))

	filterCutoffEntry := widget.NewEntry()
	filterCutoffEntry.SetText(strconv.FormatFloat(filterCutoff, 'f', 2, 64))

	sweepEntry := widget.NewEntry()
	sweepEntry.SetText(strconv.FormatFloat(sweep, 'f', 2, 64))

	pitchDecayEntry := widget.NewEntry()
	pitchDecayEntry.SetText(strconv.FormatFloat(pitchDecay, 'f', 2, 64))

	waveformSelect := widget.NewSelect([]string{"Sine", "Triangle", "Sawtooth", "Square"}, func(value string) {
		switch value {
		case "Sine":
			waveformType = kick.WaveSine
		case "Triangle":
			waveformType = kick.WaveTriangle
		case "Sawtooth":
			waveformType = kick.WaveSawtooth
		case "Square":
			waveformType = kick.WaveSquare
		}
	})
	waveformSelect.SetSelected("Sine")

	// Function to generate the kick sound
	generateKick := func() {
		length, _ = strconv.ParseFloat(lengthEntry.Text, 64)
		attack, _ = strconv.ParseFloat(attackEntry.Text, 64)
		decay, _ = strconv.ParseFloat(decayEntry.Text, 64)
		sustain, _ = strconv.ParseFloat(sustainEntry.Text, 64)
		release, _ = strconv.ParseFloat(releaseEntry.Text, 64)
		drive, _ = strconv.ParseFloat(driveEntry.Text, 64)
		filterCutoff, _ = strconv.ParseFloat(filterCutoffEntry.Text, 64)
		sweep, _ = strconv.ParseFloat(sweepEntry.Text, 64)
		pitchDecay, _ = strconv.ParseFloat(pitchDecayEntry.Text, 64)

		// Create a temporary file for the .wav output
		tmpFile, err := ioutil.TempFile("", "kick_*.wav")
		if err != nil {
			fmt.Println("Failed to create temporary file:", err)
			return
		}
		defer tmpFile.Close()

		cfg, err := kick.NewConfig(55.0, 30.0, sampleRate, length/1000.0, bitDepth, tmpFile)
		if err != nil {
			fmt.Println("Error creating config:", err)
			return
		}
		cfg.WaveformType = waveformType
		cfg.Attack = attack
		cfg.Decay = decay
		cfg.Sustain = sustain
		cfg.Release = release
		cfg.Drive = drive
		cfg.FilterCutoff = filterCutoff
		cfg.Sweep = sweep
		cfg.PitchDecay = pitchDecay

		// Generate the kick sound
		if err := cfg.GenerateKick(); err != nil {
			fmt.Println("Failed to generate kick:", err)
			return
		}

		fmt.Println("Kick drum generated:", tmpFile.Name())
		playKick(tmpFile.Name())
	}

	// Button to generate the kick sound
	generateButton := widget.NewButton("Generate Kick", generateKick)

	// Layout for the UI with labels
	w.SetContent(container.NewVBox(
		widget.NewLabel("Kick Drum Generator"),
		container.New(layout.NewFormLayout(),
			widget.NewLabel("Length (ms):"), lengthEntry,
			widget.NewLabel("Waveform:"), waveformSelect,
			widget.NewLabel("Attack (seconds):"), attackEntry,
			widget.NewLabel("Decay (seconds):"), decayEntry,
			widget.NewLabel("Sustain (level):"), sustainEntry,
			widget.NewLabel("Release (seconds):"), releaseEntry,
			widget.NewLabel("Drive (0.0 - 1.0):"), driveEntry,
			widget.NewLabel("Filter Cutoff (Hz):"), filterCutoffEntry,
			widget.NewLabel("Sweep (rate):"), sweepEntry,
			widget.NewLabel("Pitch Decay:"), pitchDecayEntry,
		),
		generateButton,
	))

	// Show the window
	w.Resize(fyne.NewSize(400, 500))
	w.ShowAndRun()
}
