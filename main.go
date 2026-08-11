package main

type Address struct{ City, Street string }
type Person struct {
	Name string
	Addr Address
}

type Group struct {
	Name    string
	Members []string
}

func main() {



	//any

	/*var x any

	x = 1
	fmt.Printf("int: %T", x)
	fmt.Println()
	x = "hello"
	fmt.Printf("string: %T", x)
	fmt.Println()
	x = []int{1, 2, 3}
	fmt.Printf("[]int: %T", x)
	fmt.Println()
	x = Circle{4}
	fmt.Printf("my struct Circle: %T", x)
	fmt.Println()
	x = nil
	fmt.Printf("nil: %T", x)
	fmt.Println()


	val,ok:=x.(int)
	if ok{
		fmt.Println(val+1)
	}*/

	//срез анонимных структур
	/*people := []struct {
		Name string
		Age  int
	}{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	for _, p := range people {
		fmt.Println(p.Name, p.Age)
	}*/

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
