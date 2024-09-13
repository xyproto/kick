package kick

import "math"

func applySaturator(samples []int, amount float64) {
	for i := range samples {
		sample := float64(samples[i]) / float64(1<<(16-1))
		sample = math.Tanh(sample * (1.0 + amount))
		samples[i] = int(sample * float64(1<<15))
	}
}

func applyMultiBandFiltering(samples []int, bands []float64, sampleRate int) {
	numSamples := len(samples)
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(sampleRate)
		frequency := 440.0 * math.Pow(2.0, t)

		if frequency < bands[0] {
			samples[i] = int(float64(samples[i]) * 0.9)
		} else if frequency < bands[1] {
			samples[i] = int(float64(samples[i]) * 0.8)
		} else if frequency < bands[2] {
			samples[i] = int(float64(samples[i]) * 0.7)
		} else {
			samples[i] = int(float64(samples[i]) * 0.6)
		}
	}
}

func applyDrive(sample, driveAmount float64) float64 {
	if driveAmount > 0 {
		return sample * (1 + driveAmount) / (1 + driveAmount*math.Abs(sample))
	}
	return sample
}

func applyEnvelope(t, attack, decay, sustain, release, duration float64) float64 {
	if t < attack {
		return t / attack
	}
	if t < attack+decay {
		return 1.0 - ((t - attack) / decay * (1.0 - sustain))
	}
	if t < duration-release {
		return sustain
	}
	if t < duration {
		return sustain * (1.0 - (t-(duration-release))/release)
	}
	return 0.0
}

func applyFadeInOut(samples []int, sampleRate int, fadeDuration float64) {
	fadeSamples := int(fadeDuration * float64(sampleRate))
	if fadeSamples > len(samples)/2 {
		fadeSamples = len(samples) / 2
	}

	// Apply fade-in
	for i := 0; i < fadeSamples; i++ {
		fadeFactor := float64(i) / float64(fadeSamples)
		samples[i] = int(float64(samples[i]) * fadeFactor)
	}

	// Apply fade-out
	for i := len(samples) - fadeSamples; i < len(samples); i++ {
		fadeFactor := float64(len(samples)-i) / float64(fadeSamples)
		samples[i] = int(float64(samples[i]) * fadeFactor)
	}
}
