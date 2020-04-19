package main

func main() {
	// Simple for loop
	var i int
	for i < 10 {
		println("i =", i)
		i++
		if i == 4 {
			continue
		}
		if i == 7 {
			break
		}
		println("(end of iteration)")
	}
	println()

	// Pre-condition initializer post-condition
	for j := 0; j < 8; j++ {
		println(j)
	}
	println()

	// Infinite loop
	k := 0
	for {
		if k == 12 {
			break
		}
		println(k)
		k++
	}
	println()

	// Iterate through a slice
	sl1 := []int{10, 20, 40, 80}
	for i := 0; i < len(sl1); i++ {
		println(i, ":", sl1[i])
	}
	println()

	// Iterate through a slice with range
	for i, v := range sl1 {
		println(i, ":", v)
	}
	println()

	// Range on maps
	m := map[string]int{"one": 100, "two": 200, "three": 300}
	for k, v := range m {
		println(k, ":", v)
	}
	println()

	// Use the write-only underscore var to ignore keys
	for _, v := range m {
		println(v)
	}
	println()

	// Switch statements
	// No implicit fallthrough
	s := "quine"
	switch s {
	case "carnap":
		println("Carnap")
	case "smullyan":
		println("Smullyan")
	case "quine":
		println("Quine")
		fallthrough
	case "gardner":
		println("Gardner")
	default:
		println("None of the above")
	}

	// Panic
	panic("This is a panic")
}
