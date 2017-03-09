package easing

// No easing, no acceleration.
func Linear(t float64) float64 {
	return t
}

// Accelerating from zero velocity.
func EaseInQuad(t float64) float64 {
	return t * t
}

// Decelerating to zero velocity.
func EaseOutQuad(t float64) float64 {
	return t * (2 - t)
}

// Acceleration until halfway, then deceleration.
func EaseInOutQuad(t float64) float64 {
	if t < .5 {
		return 2 * t * t
	}

	return -1 + (4-2*t)*t
}

// Accelerating from zero velocity.
func EaseInCubic(t float64) float64 {
	return t * t * t
}

// Decelerating to zero velocity.
func EaseOutCubic(t float64) float64 {
	t -= 1
	return t*t*t + 1
}

// Acceleration until halfway, then deceleration.
func EaseInOutCubic(t float64) float64 {
	if t < .5 {
		return 4 * t * t * t
	}

	return (t-1)*(2*t-2)*(2*t-2) + 1
}

// Accelerating from zero velocity.
func EaseInQuart(t float64) float64 {
	return t * t * t * t
}

// Decelerating to zero velocity.
func EaseOutQuart(t float64) float64 {
	t -= 1
	return 1 - t*t*t*t
}

// Acceleration until halfway, then deceleration.
func EaseInOutQuart(t float64) float64 {
	if t < .5 {
		return 8 * t * t * t * t
	}

	t -= 1
	return 1 - 8*t*t*t*t
}

// Accelerating from zero velocity.
func EaseInQuint(t float64) float64 {
	return t * t * t * t * t

}

// Decelerating to zero velocity.
func EaseOutQuint(t float64) float64 {
	t -= 1
	return 1 + t*t*t*t*t
}

// Acceleration until halfway, then deceleration.
func EaseInOutQuint(t float64) float64 {
	if t < .5 {
		return 16 * t * t * t * t * t
	}

	t -= 1
	return 1 + 16*t*t*t*t*t
}
