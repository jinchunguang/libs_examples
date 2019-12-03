package calc

import (
	"math"
	"testing"
)

// TestXxx 函数式放在一个文件尾部名 _test.go 中
func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %v; want 1", got)
	}
}



