package iter

import (
	"iter"
)

func Take2[K, V any](
	original iter.Seq2[K, V],
	max uint64,
) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var cnt uint64 = 0
		for k, v := range original {
			if max <= cnt {
				return
			}
			cnt += 1

			if !yield(k, v) {
				return
			}
		}
	}
}
