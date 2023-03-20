package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup main")
	res := m.Run()
	fmt.Println("tear-down main")

	os.Exit(res)
}

func TestMultiple(t *testing.T) {
	// setup
	// insert test data within db
	t.Run("group A", func(t *testing.T) {
		t.Log("A")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 22, 22, 484

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("group B", func(t *testing.T) {
		t.Log("B")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 22, 22, 484

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("group C", func(t *testing.T) {
		t.Log("C")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 22, 22, 484

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})
	})

	t.Run("group D", func(t *testing.T) {
		t.Log("D")
		t.Run("simple", func(t *testing.T) {
			t.Parallel()
			t.Log("simple")
			var x, y, result = 2, 2, 4

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}

			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel()
			t.Log("medium")
			var x, y, result = 22, 22, 484

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})

		t.Run("negative", func(t *testing.T) {
			t.Parallel()
			t.Log("negative")
			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("expected result: %d != %d real result", result, realResult)
			}
		})

	})
	// tearDown
	// delete test data within db
}
