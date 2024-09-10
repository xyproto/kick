
# Kick

This is a Go package for generating kick drum audio samples.

Note that the project is a bit experimental, a work in progress and that the generated samples aren't quite right, yet.

The repository includes:

- A Go package for generating kick drum samples with a wide variety of parameters.
- A command-line utility for generating kick drum samples (`cmd/kick`).
- A GUI that allows the user to randomly mutate and refine kick drum samples in a genetic algorithm style (`cmd/mutator`).

## Requirements

* `mpv` or `ffmpeg` for playing the generated samples, when using `cmd/kick` or `cmd/mutator`.

## Features

- **Customizable Kick Drum Generation**: Use a wide range of parameters such as waveform type, attack, decay, sustain, release, drive, and more to generate kick drum sounds tailored to your needs.
- **Preconfigured Kick Styles**: Generate kick drums inspired by famous drum machines like the Roland 808, 909, 707, and others. Plus, create Deep House-style and experimental kicks.
- **Waveform Types**: Choose from Sine, Triangle, Sawtooth, Square, and various noise types (White, Pink, Brown).
- **Noise Integration**: Add noise to your kick drum samples for extra texture and uniqueness.
- **Multi-Oscillator Support**: Layer multiple oscillators to create complex and rich kick sounds.
- **Command-line Utility**: Generate `.wav` files via the command-line using predefined or custom parameters.
- **Graphical User Interface**: Experiment with randomly generated kicks using a genetic algorithm approach.

## Installation

```bash
go get github.com/xyproto/kick
```

To use the CLI, UI, or Mutator, navigate to the respective directory and build the Go program:

```bash
cd cmd/kick
go build
```

Or for the GUI that uses Fyne:

```bash
cd cmd/ui
go build
```

Or for the Mutator GUI that uses gio:

```bash
cd cmd/mutator
go build
```

## Usage

### Command-line Utility

Generate a kick drum inspired by the Roland TR-808:

```bash
kick --808 -o kick808.wav
```

You can customize various parameters like the waveform, attack, decay, and more:

```bash
kick --waveform 0 --attack 0.005 --decay 0.3 --release 0.2 --drive 0.4 --o custom_kick.wav
```

Available drum machine styles:

- `--606` for 606-style kicks.
- `--707` for 707-style kicks.
- `--808` for 808-style kicks.
- `--909` for 909-style kicks.
- `--deephouse` for deep house kicks.
- `--linn` for LinnDrum-style kicks.
- `--experimental` for unique and experimental sounds.

For more help:

```bash
./kick --help
```

### GUI for Specifying Parameters (`cmd/ui`)

The GUI allows you to adjust kick drum parameters like length, waveform, attack, decay, and more via a simple form.

To run the GUI:

```bash
cd cmd/ui
go run main.go
```

### GUI with Genetic Algorithm for Randomized Kick Samples (`cmd/mutator`)

The Mutator GUI generates 16 randomized kick drums at once. You can mutate and refine these kick drums in a genetic algorithm style by selecting "mutate" or save your favorite as a `.wav` file.

To run the Mutator:

```bash
cd cmd/mutator
go run main.go
```

## Audio Samples

Generated samples can be found in the `samples/` directory. They include both `.wav` and `.mp4` formats to showcase the sounds in GitHub's README.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to submit issues and pull requests. Contributions are welcome!
