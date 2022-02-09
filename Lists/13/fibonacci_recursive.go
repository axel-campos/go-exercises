package main

var fibonaaci_map map[int]uint64 = make(map[int]uint64)

func main() {
	println("Fibonacci sequence:")
	for i := 0; i < 100; i++ {
		print(fibonacci_sequence(i), " ")
	}
}

func fibonacci_sequence(n int) uint64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n_sequence, exist := fibonaaci_map[n]
	if exist {
		return n_sequence
	} else {
		fibonaaci_map[n] = fibonacci_sequence(n-1) + fibonacci_sequence(n-2)
		return fibonaaci_map[n]
	}
}
