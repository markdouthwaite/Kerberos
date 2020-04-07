package kerberos


type nullableInt struct {
	Value int
	Set bool
}


type nullableFloat32 struct {
	Value float32
	Set bool
}


// MaxInt gets the largest integer in the given array.
func MaxInt(arr []int) int {
	max := nullableInt{0, false}
	for _, e := range arr {
		if !max.Set {
			max.Value = e
			max.Set = true
		} else {
			if max.Value < e {
				max.Value = e
			}
		}
	}
	return max.Value
}


// ArgMaxInt gets the index of the largest integer in the given array.
//If multiple elements are largest, first will be returned.
func ArgMaxInt(arr []int) int {
	max := MaxInt(arr)
	for i, e := range arr {
		if e == max {
			return i
		}
	}
	panic("Cannot compute max arg.")
}


// MaxFloat32 gets the largest float32 in the given array.
func MaxFloat32(arr []float32) float32 {
	max := nullableFloat32{0, false}
	for _, e := range arr {
		if !max.Set {
			max.Value = e
			max.Set = true
		} else {
			if max.Value < e {
				max.Value = e
			}
		}
	}
	return max.Value
}


// ArgMaxFloat32 gets the index of the largest float32 in the given array.
//If multiple elements are largest, first will be returned.
func ArgMaxFloat32(arr []float32) int {
	max := MaxFloat32(arr)
	for i, e := range arr {
		if e == max {
			return i
		}
	}
	panic("Cannot compute max arg.")
}
