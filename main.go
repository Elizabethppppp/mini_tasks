package main

import "fmt"

type Address struct{ City, Street string }
type Person struct {
	Name string
	Addr Address
}

func main() {

	p1 := Person{
		Name: "Alice",
		Addr: Address{
			City:   "Moscow",
			Street: "Tverskaya",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.Addr.City)
	p1.Addr.Street = "Nevski"
	fmt.Println(p1)

	//анонимные функции

	/*func() {
		fmt.Println("hello world")
	}()

	msg := func(msg string) {
		fmt.Println(msg)
	}

	msg("hello world")

	nums := []int{1, 7, 3, 4, 5}

	slices.SortFunc(nums, func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	})
	fmt.Println(nums)*/

}
