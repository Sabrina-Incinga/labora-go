package main

import "testing"

func BenchmarkInvertWords(b *testing.B) {
	invertWords("cuadrado")
	//Output: "odardauc"
}