package main

import (
	"errors"
	"fmt"
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

func a() {
	defer fmt.Println("a")
	b()
}

// на этаж выше
func b() {
	defer fmt.Println("b")
	/*defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Перехвачено на этаж выше! Ошибка: %v\n", r)
		}
	}()*/
	c()
}

// recover в defer и в теле функции
func c() {
	defer fmt.Println("c")
	/*defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Перехвачено сразу! Ошибка: %v\n", r)
		}
	}()*/
	if r := recover(); r != nil {
		fmt.Printf("В теле функции! Ошибка: %v\n", r)
	}
	//panic("test")
	d()
}

type ValidationBug struct{ Field string }

func d() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Паника перехвачена!")
			switch v := r.(type) {
			case ValidationBug:
				fmt.Printf("Поле Field: %s\n", v.Field)
			case string:
				fmt.Printf("Строковая паника: %s\n", v)
			case int:
				fmt.Printf("Числовая паника: %d\n", v)
			default:
				fmt.Printf("Неизвестный тип паники: %v (тип: %T)\n", r, r)
			}
		}
	}()
	//panic(ValidationBug{Field: "panic d"})
	e()
}

func e() {
	defer func() {
		if err := recover(); err != nil {
			switch v := err.(type) {
			case ValidationBug:
				fmt.Println("Перехвачена нужная паника:", v.Field)

			case string:
				panic(err)
			case int:
				panic(err)
			default:
				panic(err)
			}
		}
	}()

	panic(ValidationBug{Field: "panic d"})
}

func f() (result int) {
	defer func() { result *= 2 }()
	return 5
}
func g() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("в функции g: %w", err)
		}
	}()
	return errors.New("дно")
}

func SafeCall(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("паника: %v", r)
		}
	}()
	f()
	return nil
}

func main() {

	//17 defer B
	//неверное утверждение типа
	/*var i interface{} = "привет"
	n := i.(int)
	fmt.Println(n)*/
	//деление на 0
	/*a := 10 / 0
	fmt.Println(a)*/
	//запись в nil-мапу
	/*m := make(map[string]int)
	m = nil
	m["a"] = 1*/
	//разыменование nil
	/*var p *int
	_ = *p*/
	//выход за границу среза
	/*slice := []int{10, 20, 30}
	_ = slice[3]*/

	//14 defer B
	/*err := SafeCall(func() {
		var slice []int
		_ = slice[1]
	})

	err := SafeCall(func() {
		s := 1
		fmt.Println(s)
	})

	if err != nil {
		fmt.Println("Перехвачена паника:", err)
	}*/

	//9-13 defer B
	//defer fmt.Println("main")
	//a()

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
