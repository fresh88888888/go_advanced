package main

import "fmt"

//go:noinline
func ebpf_demo(a1 int, a2 bool, a3 float32) (r1 int64, r2 int32, r3 string) {
	fmt.Printf("ebpf_demo:: a1=%d a2=%t a3=%.2f\n", a1, a2, a3)
	return 100, 200, "test for ebpf"
}

func main() {
	r1, r2, r3 := ebpf_demo(100, true, 68.34)
	fmt.Printf("main:: r1=%d r2=%d r3=%s\n", r1, r2, r3)

	return
}
