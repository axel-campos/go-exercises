package main

func main() {
	println("\t1\t2\t3\t4\t5\t6\t7\t8\t9\t10\t11\t12")
	for i := 1; i <= 12; i++ {
		print(i)
		for j := 1; j <= 12; j++ {
			print("\t", i*j)
		}
		println()
	}
}
