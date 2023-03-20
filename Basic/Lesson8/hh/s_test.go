package hh

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup hh")
	res := m.Run()
	fmt.Println("tear-down hh")

	os.Exit(res)
}

func TestAdd(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		t.Log("simple")
		var x, y, result = 2, 2, 4

		realResult := Add(x, y)

		if realResult != result {
			t.Errorf("expected result: %d != %d real result", result, realResult)
		}
		t.Run("1", func(t *testing.T) {
			r := Add(1, 1)
			if r != 2 {
				t.Errorf("failed")
			}
		})
	})
}