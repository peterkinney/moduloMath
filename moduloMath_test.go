package moduloMath

import (
	"testing"
)

func TestSum (tt *testing.T) {
	for mm := uint64(1); mm <= 10; mm++ {
		for ii := uint64(0); ii < 20; ii++ {
			for jj := uint64(0); jj < 20; jj++ {
				want, got := (ii + jj) % mm, Sum(mm, ii, jj)
				if want != got {
					tt.Errorf("Sum(%d,%d,%d) Failed: want %d, got %d",mm,ii,jj,want,got)
				}
			}
		}
	}
}

func TestSumRange(tt *testing.T){
	for mm := uint64(1); mm <= 10; mm++ {
		for start := uint64(0); start < 20; start++ {
			for len := uint64(0); len < 20; len++ {
				want, got := (((start + start + len - 1) * len) / 2) % mm, SumRange(mm, start, len)
				if want != got {
					tt.Errorf("SumRange(%d,%d,%d) Failed: want %d, got %d",mm,start,len,want,got)
				}
			}
		}
	}
}