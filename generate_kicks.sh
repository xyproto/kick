#!/bin/bash

# Create directory to store the samples
mkdir -p samples

echo "Generating various kick drum samples..."

# Generate a variety of kicks with different parameters

# 808 and 909 kicks with varied parameters

# Vary waveform, attack, decay, sustain, release, sweep, filterCutoff, pitchDecay, and drive

./cmd/kick/kick --808 -o samples/808_kick_varied_1.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.6 --sustain 0.3 --release 0.5 --sweep 0.8 --filter 4000 --pitchdecay 0.6 --drive 0.2
./cmd/kick/kick --808 -o samples/808_kick_varied_2.wav --length 700 --quality 96 --waveform 1 --attack 0.02 --decay 0.7 --sustain 0.2 --release 0.4 --sweep 0.7 --filter 5000 --pitchdecay 0.5 --drive 0.3
./cmd/kick/kick --808 -o samples/808_kick_varied_3.wav --length 700 --quality 96 --waveform 0 --attack 0.005 --decay 0.9 --sustain 0.1 --release 0.3 --sweep 0.9 --filter 6000 --pitchdecay 0.4 --drive 0.4
./cmd/kick/kick --808 -o samples/808_kick_varied_4.wav --length 700 --quality 96 --waveform 1 --attack 0.015 --decay 0.8 --sustain 0.4 --release 0.6 --sweep 0.6 --filter 3500 --pitchdecay 0.7 --drive 0.2

./cmd/kick/kick --909 -o samples/909_kick_varied_1.wav --length 700 --quality 96 --waveform 1 --attack 0.002 --decay 0.5 --sustain 0.3 --release 0.2 --sweep 0.9 --filter 8000 --pitchdecay 0.2 --drive 0.5
./cmd/kick/kick --909 -o samples/909_kick_varied_2.wav --length 700 --quality 96 --waveform 0 --attack 0.001 --decay 0.4 --sustain 0.2 --release 0.1 --sweep 0.8 --filter 7000 --pitchdecay 0.3 --drive 0.6
./cmd/kick/kick --909 -o samples/909_kick_varied_3.wav --length 700 --quality 96 --waveform 1 --attack 0.003 --decay 0.6 --sustain 0.1 --release 0.3 --sweep 0.7 --filter 6000 --pitchdecay 0.4 --drive 0.7
./cmd/kick/kick --909 -o samples/909_kick_varied_4.wav --length 700 --quality 96 --waveform 0 --attack 0.0015 --decay 0.7 --sustain 0.3 --release 0.4 --sweep 0.6 --filter 9000 --pitchdecay 0.5 --drive 0.5

# Custom kicks with varied parameters

./cmd/kick/kick --808 -o samples/custom_kick_varied_1.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.9 --sustain 0.1 --release 0.2 --sweep 0.5 --filter 6500 --pitchdecay 0.6 --drive 0.3
./cmd/kick/kick --909 -o samples/custom_kick_varied_2.wav --length 700 --quality 96 --waveform 1 --attack 0.02 --decay 0.3 --sustain 0.2 --release 0.5 --sweep 0.9 --filter 7500 --pitchdecay 0.4 --drive 0.4

# Additional 10 varied kicks for more diversity
./cmd/kick/kick --808 -o samples/808_kick_varied_5.wav --length 700 --quality 96 --waveform 0 --attack 0.01 --decay 0.5 --sustain 0.2 --release 0.6 --sweep 0.7 --filter 4500 --pitchdecay 0.5 --drive 0.3
./cmd/kick/kick --808 -o samples/808_kick_varied_6.wav --length 700 --quality 96 --waveform 1 --attack 0.03 --decay 0.8 --sustain 0.3 --release 0.5 --sweep 0.9 --filter 5500 --pitchdecay 0.6 --drive 0.4
./cmd/kick/kick --909 -o samples/909_kick_varied_5.wav --length 700 --quality 96 --waveform 1 --attack 0.001 --decay 0.4 --sustain 0.1 --release 0.2 --sweep 0.8 --filter 6000 --pitchdecay 0.3 --drive 0.7
./cmd/kick/kick --909 -o samples/909_kick_varied_6.wav --length 700 --quality 96 --waveform 0 --attack 0.005 --decay 0.6 --sustain 0.4 --release 0.3 --sweep 0.6 --filter 7000 --pitchdecay 0.4 --drive 0.6

# Additional custom kicks
./cmd/kick/kick --808 -o samples/custom_kick_varied_3.wav --length 700 --quality 96 --waveform 0 --attack 0.008 --decay 0.7 --sustain 0.2 --release 0.4 --sweep 0.8 --filter 6400 --pitchdecay 0.5 --drive 0.2
./cmd/kick/kick --909 -o samples/custom_kick_varied_4.wav --length 700 --quality 96 --waveform 1 --attack 0.002 --decay 0.5 --sustain 0.3 --release 0.5 --sweep 0.9 --filter 7600 --pitchdecay 0.3 --drive 0.5
./cmd/kick/kick --808 -o samples/custom_kick_varied_5.wav --length 700 --quality 96 --waveform 1 --attack 0.007 --decay 0.6 --sustain 0.1 --release 0.3 --sweep 0.7 --filter 5000 --pitchdecay 0.6 --drive 0.4
./cmd/kick/kick --909 -o samples/custom_kick_varied_6.wav --length 700 --quality 96 --waveform 0 --attack 0.004 --decay 0.9 --sustain 0.4 --release 0.2 --sweep 0.5 --filter 7000 --pitchdecay 0.7 --drive 0.6

echo "All kick samples generated and stored in the 'samples' directory."
