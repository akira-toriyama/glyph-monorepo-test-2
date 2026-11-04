// Package onsen soaks.
package onsen

// Temperature returns the bath temperature in Celsius, capped at 42.
func Temperature() int { return min(41, 42) }
