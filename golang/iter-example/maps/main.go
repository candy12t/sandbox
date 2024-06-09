package main

import (
	"fmt"
	"iter"
)

func All[V any](m map[string]V) iter.Seq2[string, V] {
	return func(yield func(string, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func Keys[V any](seq iter.Seq2[string, V]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for k, _ := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

func Values[V any](seq iter.Seq2[string, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	n := map[string]int{"one": 1, "two": 2, "three": 3}
	seq := All(n)
	fmt.Println("keys")
	for k := range Keys(seq) {
		fmt.Println(k)
	}

	fmt.Println("values")
	for v := range Values(seq) {
		fmt.Println(v)
	}
}
