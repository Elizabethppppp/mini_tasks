package methordStruct

type Animal struct{ Name string }

func (a Animal) Describe() string { return "это " + a.Name }

func (d Dog) Describe() string {
	return d.Animal.Describe() + ", breed: " + d.Breed
}

type Dog struct {
	Animal
	Breed string
}
