package moduloMath

func Sum (mod, aa, bb uint64) uint64 {
    aa, bb = aa % mod, bb % mod
	return (aa + bb) % mod
}

func SumRange (mod, start, length uint64) uint64 {
    a, b := start % (mod * 2), length % (mod * 2)
    return (((a + (a + b - 1)) * b) / 2) % mod
}