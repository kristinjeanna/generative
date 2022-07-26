package misc

import "fmt"

func ExampleLastNIndexes() {
	s, _ := LastNIndexes(97, 5)
	fmt.Println(s)

	s, _ = LastNIndexes(2, 5)
	fmt.Println(s)
	// Output: [93 94 95 96 97]
	// [0 1 2]
}

func ExampleLastNIndexesWrap() {
	s, _ := LastNIndexesWrap(97, 5, 100)
	fmt.Println(s)

	s, _ = LastNIndexesWrap(2, 5, 10)
	fmt.Println(s)
	// Output: [93 94 95 96 97]
	// [8 9 0 1 2]
}
