package main

import "testing"

func BenchmarkQueueSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	// b.
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runQueueSlice()
		}
	})

	// run()
}

func BenchmarkQueueLL(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runLL()
		}
	})
}
