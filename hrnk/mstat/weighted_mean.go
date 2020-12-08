package mstat

// WtdMean sxd
func WtdMean(a []float64, b []float64) float64 {
	if len(a) != len(b) {
		panic("Lenght not same")
	}

	var sumA, sumB float64
	for i := 0; i < len(a); i++ {
		sumA += (a[i] * b[i])
		sumB += b[i]
	}

	return sumA / sumB

}
