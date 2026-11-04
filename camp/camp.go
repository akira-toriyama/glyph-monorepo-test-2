// Package camp pitches tents.
package camp

// Pitch returns the tent count.
func Pitch() int { return PitchN(1) }

// PitchN returns n tents, at least one.
func PitchN(n int) int { // n below one pitches one
	if n < 1 {
		return 1 // a camp always has one tent
	}
	return n
}
