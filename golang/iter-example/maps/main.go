package main

import (
	"fmt"
	"iter"
)

func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				return
			}
		}
	}
}

func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range m {
			if !yield(k) {
				return
			}
		}
	}
}

func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range m {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	ns := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("All")
	for k, v := range All(ns) {
		fmt.Println(k, v)
	}
	fmt.Println("Keys")
	for key := range Keys(ns) {
		fmt.Println(key)
	}
	fmt.Println("Values")
	for value := range Values(ns) {
		fmt.Println(value)
	}
}
