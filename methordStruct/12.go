package methordStruct

import "strings"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToF() Fahrenheit {

	return Fahrenheit((c * 9 / 5) + 32)
}

type Path []string

func (p Path) String() string {
	return strings.Join(p, "/")
}
