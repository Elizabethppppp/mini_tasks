package main

import "fmt"

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

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	for val := range ch {

		fmt.Println(val)
	}

	//gorutine with receive
	/*group := sync.WaitGroup{}
	group.Add(1)
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		defer group.Done()
		val := <-ch
		fmt.Println(val)
	}()
	group.Wait()*/

	//gorutine send
	/*ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)*/

	//with bufer
	/*ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)*/

	//gorutine 2
	/*result := make([]int, 5)
	group := sync.WaitGroup{}
	group.Add(5)
	go func(index int) {
		defer group.Done()
		result[index] = 0

	}(0)
	go func(index int) {
		defer group.Done()
		result[index] = 1

	}(1)
	go func(index int) {
		defer group.Done()
		result[index] = 2

	}(2)
	go func(index int) {
		defer group.Done()
		result[index] = 3

	}(3)
	go func(index int) {
		defer group.Done()
		result[index] = 4

	}(4)

	group.Wait()
	for k, v := range result {
		fmt.Printf("index: %d, value: %d ", k, v)
	}*/

	//gorutine

	/*group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		defer group.Done()
		fmt.Println("привет")
	}()
	group.Wait()*/

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
