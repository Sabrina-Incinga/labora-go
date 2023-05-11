package main

import "testing"

func BenchmarkInvertWords(b *testing.B) {
	for i:=0;i<b.N;i++{
		invertWords("cuadrado")
	}
}