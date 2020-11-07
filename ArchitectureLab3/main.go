package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var correctString bool  = true

func ToBool(a int) bool {
	if a == 1 {return true }
	return false
}

func correctAtom(s string) bool {
	arrAtom := []rune{'1', '0', 'x', 'y', 'z', 'w'}
	f1 := false
	for _, value := range arrAtom {
		f1 = f1  || strings.ContainsRune(s, value)
	}
	return f1
}
func Atom (s string, x, y, z, w int) bool {
	switch s {
	case "x": return ToBool(x)
	case "y": return ToBool(y)
	case "z": return ToBool(z)
	case "w": return ToBool(w)
	case "1": return true
	default:
		return false
	}
}

func Not(x bool) bool {
	return !x
}
func And(x, y bool) bool {
	return x && y
}
func Or(x, y bool) bool {
	return x || y
}
func Equal(x, y bool) bool {
	return x == y
}
func Implecation(x, y bool) bool {
	return (!x || y)
}
func Sum(x, y bool) bool {
	return x != y
}
func Sheff(x, y bool ) bool {
	return !(x && y)
}
func Pirs(x, y bool) bool {
	return !(x || y)
}

func Solution (s string, x, y, z, w int) bool {
	if len(s) == 1 {
		if(correctAtom(s)) {
			return Atom(s, x,y,z,w)
		} else {
			correctString = false
			return correctString
		}
	}
	if(strings.ContainsRune(s, '/') || strings.ContainsRune(s, '|')) {
		var countBracket int = 0 // count '('
		for i := len(s) -1; i >=0; i-- {
			if countBracket == 0 {
				switch s[i] {
				case '/': return Sheff(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w))
				case '|': return Pirs(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w))
				}
			}
			if s[i] == ')' { countBracket++ }
			if  s[i] == '(' { countBracket-- }
		}
	}
	if strings.ContainsRune(s, '=') {
		var countBracket int = 0 // count '('
		for i := len(s) - 1; i >= 0; i-- {
			if countBracket == 0 && s[i] == '=' { return Equal(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w)) }
		    if s[i] == ')' { countBracket++ }
			if s[i] == '(' { countBracket-- }
		}
	}
	if strings.ContainsRune(s, '-') {
		var countBracket int = 0 // count '('
		for i := len(s) - 1; i >= 0; i-- {
			if countBracket == 0 && s[i] == '-' { return Implecation(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w)) }
			if s[i] == ')' { countBracket++ }
			if s[i] == '(' { countBracket-- }
		}
	}
	if(strings.ContainsRune(s, 'V') || strings.ContainsRune(s, '+')) {
		var countBracket int = 0 // count '('
		for i := len(s) -1; i >=0; i-- {
			if countBracket == 0 {
				switch s[i] {
				case '+': return Sum(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w))
				case 'V': return Or(Solution(s[0:i], x,y,z,w), Solution(s[i+1:], x,y,z,w))
				}
			}
			if s[i] == ')' { countBracket++ }
			if  s[i] == '(' { countBracket-- }
		}
	}
	if strings.ContainsRune(s, '&') {
		var countBracket int = 0
		for i := len(s) - 1; i >= 0; i-- {

			if countBracket == 0 && s[i] == '&' {

				s1 := s[:i]
				s2 := s[i+1:]
                fmt.Println(s1, s2)
				return And(Solution(s1, x,y,z,w), Solution(s2, x,y,z,w)) }
			if s[i] == ')' { countBracket++ }
			if s[i] == '(' { countBracket-- }
		}
	}
	if strings.HasPrefix(s, "!") { return Not(Solution(s[1:], x,y,z,w))}
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		return Solution(s[1:len(s)-1], x,y,z,w)
	} else  {correctString = false}
	return false

}
func main() {
	// задание входного файла
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

    // задание выходного файла
	file2, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
    defer file2.Close()

	s := bufio.NewScanner(file)
	if s == nil  {
		_, _ = file2.WriteString("Исходный файл пуст")
		return
	}

	strArray := make([]string, 0, 0)  // create slice string str
    for s.Scan() {
		strArray = append(strArray, s.Text())
	}
	for _, str := range strArray {
		_, _ = file2.Write([]byte("F = " + str + "\n"))
		_, _ = file2.Write([]byte("x y z w  F" + "\n"))
		x, y, z, w := 0, 0, 0, 0
		for x = 0; x < 2 && correctString; x++ {
			for y = 0; y < 2 && correctString; y++ {
				for z = 0; z < 2 && correctString; z++ {
					for w = 0; w < 2; w++ {
						res := "0"
						if Solution(str, x, y, z, w)  {
							res = "1"
						}
						if !correctString { break; }
						xStr := strconv.Itoa(x)
						yStr := strconv.Itoa(y)
						zStr := strconv.Itoa(z)
						wStr := strconv.Itoa(w)

						file2.Write([]byte(xStr + " " + yStr + " " + zStr + " " + wStr + "  " + res + "\n"))
					}
				}
			}
		}
		if !correctString {
			file2.Write([]byte("Incorrect form of function"))
		}
		file2.Write([]byte("\n-----------------------------\n"))
	}

}