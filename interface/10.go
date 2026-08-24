package _interface

import (
	"fmt"
	"unicode/utf8"
)

func Describe(x any) string {
	switch x.(type) {
	case int:
		return fmt.Sprintf("целое %d", x)
	case string:
		return fmt.Sprintf("строка %s длины %d", x, utf8.RuneCountInString(x.(string)))
	case []any:
		return fmt.Sprintf("срез из %d элементов", len(x.([]any)))
	case bool:
		return fmt.Sprintf("булево %t", x)
	default:
		return fmt.Sprintf("что-то типа %T", x)
	}
}
