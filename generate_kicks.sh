#!/bin/bash

# Create directory to store the samples
mkdir -p samples

echo "Generating a variety of kick drum samples..."

# Generate classic kicks
./cmd/kick/kick --808 -o samples/808_classic.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.8 --sustain 0.2 --release 0.5 --sweep 0.9 --filter 4000 --pitchdecay 0.5 --drive 0.2 --bitdepth 16
./cmd/kick/kick --909 -o samples/909_classic.wav --length 700 --quality 96 --waveform 1 --attack 0.002 --decay 0.5 --sustain 0.1 --release 0.4 --sweep 0.7 --filter 8000 --pitchdecay 0.3 --drive 0.4 --bitdepth 16
./cmd/kick/kick --707 -o samples/707_classic.wav --length 700 --quality 96 --waveform 1 --attack 0.005 --decay 0.3 --sustain 0.2 --release 0.3 --sweep 0.6 --filter 5000 --pitchdecay 0.4 --drive 0.3 --bitdepth 16
./cmd/kick/kick --606 -o samples/606_classic.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.3 --sustain 0.1 --release 0.2 --sweep 0.7 --filter 5000 --pitchdecay 0.5 --drive 0.4 --bitdepth 16

# Generate modern and experimental kicks
./cmd/kick/kick --deephouse -o samples/deephouse_classic.wav --length 700 --quality 96 --waveform 0 --attack 0.005 --decay 0.9 --sustain 0.3 --release 0.7 --sweep 0.8 --filter 3500 --pitchdecay 0.6 --drive 0.6 --bitdepth 16
./cmd/kick/kick --experimental -o samples/experimental_classic.wav --length 700 --quality 96 --waveform 2 --attack 0.001 --decay 0.7 --sustain 0.1 --release 0.4 --sweep 1.2 --filter 3000 --pitchdecay 0.8 --drive 0.8 --bitdepth 16

echo "Generating varied kick drum samples with different parameters..."

# 10 varied samples demonstrating different parameter combinations

./cmd/kick/kick --808 -o samples/808_varied_1.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.6 --sustain 0.3 --release 0.6 --sweep 0.7 --filter 4500 --pitchdecay 0.5 --drive 0.3 --bitdepth 16
./cmd/kick/kick --909 -o samples/909_varied_1.wav --length 700 --quality 96 --waveform 1 --attack 0.001 --decay 0.4 --sustain 0.1 --release 0.2 --sweep 0.8 --filter 6000 --pitchdecay 0.3 --drive 0.6 --bitdepth 16
./cmd/kick/kick --606 -o samples/606_varied_1.wav --length 700 --quality 96 --waveform 0 --attack 0.02 --decay 0.4 --sustain 0.2 --release 0.3 --sweep 0.8 --filter 5000 --pitchdecay 0.6 --drive 0.5 --bitdepth 16
./cmd/kick/kick --deephouse -o samples/deephouse_varied_1.wav --length 900 --quality 96 --waveform 0 --attack 0.008 --decay 0.7 --sustain 0.3 --release 0.8 --sweep 0.9 --filter 3200 --pitchdecay 0.7 --drive 0.7 --bitdepth 16
./cmd/kick/kick --experimental -o samples/experimental_varied_1.wav --length 900 --quality 96 --waveform 2 --attack 0.006 --decay 0.7 --sustain 0.2 --release 0.4 --sweep 1.0 --filter 3000 --pitchdecay 0.8 --drive 0.7 --bitdepth 16
./cmd/kick/kick --707 -o samples/707_varied_1.wav --length 800 --quality 96 --waveform 1 --attack 0.004 --decay 0.6 --sustain 0.2 --release 0.4 --sweep 0.7 --filter 5000 --pitchdecay 0.3 --drive 0.5 --bitdepth 16
./cmd/kick/kick --deephouse -o samples/deephouse_varied_2.wav --length 900 --quality 96 --waveform 0 --attack 0.01 --decay 0.9 --sustain 0.4 --release 0.9 --sweep 0.9 --filter 3500 --pitchdecay 0.6 --drive 0.8 --bitdepth 16

# Video game-style kicks
./cmd/kick/kick --experimental -o samples/videogame_kick_1.wav --length 500 --quality 96 --waveform 2 --attack 0.001 --decay 0.2 --sustain 0.0 --release 0.1 --sweep 1.5 --filter 2000 --pitchdecay 1.0 --drive 0.7 --bitdepth 16
./cmd/kick/kick --experimental -o samples/videogame_kick_2.wav --length 500 --quality 96 --waveform 3 --attack 0.001 --decay 0.3 --sustain 0.0 --release 0.2 --sweep 2.0 --filter 1500 --pitchdecay 1.5 --drive 0.9 --bitdepth 16

# Convert all .wav files to .mp4 using ffmpeg
for wav_file in samples/*.wav; do
  mp4_file="${wav_file%.wav}.mp4"
  echo "Converting $wav_file to $mp4_file..."
  ffmpeg -i "$wav_file" -c:a aac "$mp4_file" -y
done

echo "All kick samples generated, converted to .mp4, and stored in the 'samples' directory."
