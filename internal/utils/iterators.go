package utils

import "iter"

func NormalTransform[T1, T2 any](list []T1, transform func(T1) T2) []T2 {
	transformed := make([]T2, len(list))

	for i, t := range list {
		transformed[i] = transform(t)
	}
	return transformed
}

func IteratorTransform[T1, T2 any](list []T1, transform func(T1) T2) iter.Seq2[int, T2] {
	return func(yield func(int, T2) bool) {
		for i, t := range list {
			if !yield(i, transform(t)) {
				return
			}
		}
	}
}
