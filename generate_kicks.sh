#!/bin/bash

# Create directory to store the samples
mkdir -p samples

echo "Generating various kick drum samples..."

# Generate classic 707, 808, 909, KORG, and SOMA kicks with default parameters

# 707 classic
./cmd/kick/kick --707 -o samples/707_classic.wav --length 700 --quality 96 --attack 0.002 --decay 0.3 --sustain 0.2 --release 0.3 --sweep 0.6 --filter 4000 --pitchdecay 0.4 --drive 0.3

# 808 classic
./cmd/kick/kick --808 -o samples/808_classic.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.8 --sustain 0.2 --release 0.5 --sweep 0.9 --filter 5000 --pitchdecay 0.5 --drive 0.2

# 909 classic
./cmd/kick/kick --909 -o samples/909_classic.wav --length 700 --quality 96 --waveform 1 --attack 0.002 --decay 0.5 --sustain 0.1 --release 0.4 --sweep 0.7 --filter 8000 --pitchdecay 0.3 --drive 0.4

# KORG-style
./cmd/kick/kick --korg -o samples/korg_classic.wav --length 700 --quality 96 --waveform 0 --attack 0.008 --decay 0.6 --sustain 0.3 --release 0.4 --sweep 0.8 --filter 6000 --pitchdecay 0.4 --drive 0.5

# SOMA-style
./cmd/kick/kick --soma -o samples/soma_classic.wav --length 700 --quality 96 --waveform 1 --attack 0.005 --decay 0.4 --sustain 0.2 --release 0.3 --sweep 0.7 --filter 5500 --pitchdecay 0.3 --drive 0.6

# Custom kicks with varied parameters
./cmd/kick/kick --808 -o samples/custom_kick_1.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.9 --sustain 0.1 --release 0.2 --sweep 0.5 --filter 6500 --pitchdecay 0.6 --drive 0.3
./cmd/kick/kick --909 -o samples/custom_kick_2.wav --length 700 --quality 96 --waveform 1 --attack 0.02 --decay 0.3 --sustain 0.2 --release 0.5 --sweep 0.9 --filter 7500 --pitchdecay 0.4 --drive 0.4

# Additional 10 varied kicks for more diversity
./cmd/kick/kick --808 -o samples/808_varied_1.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.5 --sustain 0.2 --release 0.6 --sweep 0.7 --filter 4500 --pitchdecay 0.5 --drive 0.3
./cmd/kick/kick --808 -o samples/808_varied_2.wav --length 700 --quality 96 --waveform 1 --attack 0.03 --decay 0.8 --sustain 0.3 --release 0.5 --sweep 0.9 --filter 5500 --pitchdecay 0.6 --drive 0.4
./cmd/kick/kick --909 -o samples/909_varied_1.wav --length 700 --quality 96 --waveform 1 --attack 0.001 --decay 0.4 --sustain 0.1 --release 0.2 --sweep 0.8 --filter 6000 --pitchdecay 0.3 --drive 0.7
./cmd/kick/kick --909 -o samples/909_varied_2.wav --length 700 --quality 96 --waveform 0 --attack 0.005 --decay 0.6 --sustain 0.4 --release 0.3 --sweep 0.6 --filter 7000 --pitchdecay 0.4 --drive 0.6

# Generate weird and wacky samples
./cmd/kick/kick --808 -o samples/weird_kick_1.wav --length 1000 --quality 96 --waveform 0 --attack 0.1 --decay 0.2 --sustain 0.1 --release 0.8 --sweep 1.0 --filter 3000 --pitchdecay 0.7 --drive 0.9 --noise pink --noiseamount 0.5
./cmd/kick/kick --909 -o samples/weird_kick_2.wav --length 1200 --quality 96 --waveform 1 --attack 0.005 --decay 1.0 --sustain 0.1 --release 0.6 --sweep 0.3 --filter 4000 --pitchdecay 0.8 --drive 0.8 --noise white --noiseamount 0.7
./cmd/kick/kick --korg -o samples/weird_kick_3.wav --length 500 --quality 96 --waveform 0 --attack 0.002 --decay 0.5 --sustain 0.2 --release 0.1 --sweep 0.9 --filter 1000 --pitchdecay 0.3 --drive 1.0 --noise brown --noiseamount 0.9
./cmd/kick/kick --soma -o samples/weird_kick_4.wav --length 900 --quality 96 --waveform 1 --attack 0.005 --decay 0.3 --sustain 0.5 --release 0.3 --sweep 1.2 --filter 2500 --pitchdecay 0.5 --drive 0.6 --noise pink --noiseamount 0.4
./cmd/kick/kick --909 -o samples/weird_kick_5.wav --length 1100 --quality 96 --waveform 1 --attack 0.02 --decay 0.7 --sustain 0.4 --release 0.2 --sweep 1.5 --filter 7000 --pitchdecay 1.0 --drive 0.5 --noise white --noiseamount 0.3

# Convert all .wav files to .mp4 using ffmpeg
for wav_file in samples/*.wav; do
  mp4_file="${wav_file%.wav}.mp4"
  echo "Converting $wav_file to $mp4_file..."
  ffmpeg -i "$wav_file" -c:a aac "$mp4_file" -y
done

echo "All kick samples generated, converted to .mp4, and stored in the 'samples' directory."
