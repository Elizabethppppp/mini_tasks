package main

import (
	"io"
)

type Address struct{ City, Street string }
type Person struct {
	Name string
	Addr Address
}

type Group struct {
	Name    string
	Members []string
}

type CountingWriter struct{ N int }

func (w *CountingWriter) Write(p []byte) (int, error) {
	w.N += len(p)
	return len(p), nil
}

type UpperReader struct{ src io.Reader }

func (u *UpperReader) Read(p []byte) (int, error) {
	n, err := u.src.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= 'a' - 'A'
		}
	}
	return n, nil
}

func main() {

	//11 myErrors B
	/*base := &myErrors.OutOfRange{
		Value: 150,
		Min:   100,
		Max:   0,
	}

	layer1 := fmt.Errorf("проверка 1: %w", base)
	layer2 := fmt.Errorf("проверка 2: %w", layer1)

	err := layer2

	var result *myErrors.OutOfRange
	if errors.As(err, &result) {
		fmt.Printf("Нашли ошибку\n")
		fmt.Printf("Допустимо от %d до %d, получено %d\n", result.Min, result.Max, result.Value)
	} else {
		fmt.Println("Не найдено")
	}*/

	/*cw := CountingWriter{}
	fmt.Fprintf(&cw, "привет, %s!", "мир")
	fmt.Printf("Записано байт: %d\n", cw.N)

	data, _ := io.ReadAll(&UpperReader{strings.NewReader("go go go")})
	fmt.Println(string(data))*/

	//18 task methordStruct B
	/*//m := map[string]methordStruct.Counter{"a": {}}
	//m["a"].Inc()                  // (а) нет, не адресуемое
	//methordStruct.Counter{}.Inc() // (б) нет, не адресуемое
	c := methordStruct.Counter{}
	c.Inc() // (в) всё хорошо, так как переменная */

	///16 task methordStruct B
	/*dog := methordStruct.Dog{
		Animal: methordStruct.Animal{Name: "Lola"},
		Breed:  "shpiz",
	}

	fmt.Println(dog.Animal.Name)
	fmt.Println(dog.Name)
	fmt.Println(dog.Breed)
	fmt.Println(dog.Describe())
	fmt.Println(dog.Animal.Describe())*/

	//14 task methordStruct B
	/*p := methordStruct.Point{X: 3, Y: 4}
	fmt.Println(p)

	f := p.Dist
	fmt.Printf("type %T", f)
	fmt.Println()
	g := methordStruct.Point.Dist
	fmt.Printf("type %T", g)
	fmt.Println()*/

	/*fmt.Println(p.Dist())
	fmt.Println(methordStruct.Dist2(p))

	p.MoveWrong(5, 6)
	fmt.Println(p)

	p.Move(5, 6)
	fmt.Println(p)*/

	//18 B map module
	/*m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	m[2] = 3
	m[3] = 4
	m[4] = 5
	m[5] = 6
	m[6] = 7

	fmt.Println(m)

	for k := range m {
		if k == 3 {
			delete(m, k)
		}
	}

	fmt.Println(m)

	for k := range m {
		if k == 5 {
			m[7] = 8
			m[8] = 9
		}
	}

	fmt.Println(m)*/

	//channels
	/*ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	for val := range ch {

		fmt.Println(val)
	}*/

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
