package iter_test

import (
	"testing"

	"iter"
	"maps"
	"slices"

	it "github.com/takanoriyanagitani/go-avro-head/util/iter"
)

func TestIter(t *testing.T) {
	t.Parallel()

	t.Run("Take2", func(t *testing.T) {
		t.Parallel()

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			var original iter.Seq2[int, string] = slices.All([]string{})
			var taken iter.Seq2[int, string] = it.Take2(original, 0)
			var m map[int]string = maps.Collect(taken)
			var msz int = len(m)
			if 0 != msz {
				t.Fatalf("must be empty: %v\n", msz)
			}
		})

		t.Run("take 0", func(t *testing.T) {
			t.Parallel()

			var original iter.Seq2[int, string] = slices.All([]string{
				"helo",
				"wrld",
			})
			var taken iter.Seq2[int, string] = it.Take2(original, 0)
			var m map[int]string = maps.Collect(taken)
			var msz int = len(m)
			if 0 != msz {
				t.Fatalf("must be empty: %v\n", msz)
			}
		})

		t.Run("take 1", func(t *testing.T) {
			t.Parallel()

			var original iter.Seq2[int, string] = slices.All([]string{
				"helo",
				"wrld",
			})
			var taken iter.Seq2[int, string] = it.Take2(original, 1)
			var m map[int]string = maps.Collect(taken)
			var msz int = len(m)
			if 1 != msz {
				t.Fatalf("must not be empty: %v\n", msz)
			}

			var first string = m[0]
			if "helo" != first {
				t.Fatalf("unexpected value: %s\n", first)
			}
		})
	})
}
