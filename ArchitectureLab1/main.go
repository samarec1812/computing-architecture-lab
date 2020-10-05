package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)
var(
    mu sync.Mutex
    wg sync.WaitGroup
)

type Numbers struct {
	s string    // строковое представление числа в файле
	integerPartNumber int64  // целая часть числа в файле
	fractionalPartNumber int64 // дробная часть числа в файле
	lenAccuracy int // точность числа в 10 c/c
	lenAccuracy2 int // точность в троичной с/c
	sign        int // знак числа
}


func (n *Numbers) splitNumber () (string, string) {
	flag := strings.Contains(n.s, ".")
    if !flag {
    	return n.s, ""
	}
	index := strings.Index(n.s, ".")
	if n.s[0] == '-' {
		(*n).sign = 1
		s1 := n.s[1:index]
		s2 := n.s[index + 1:]
		return s1, s2
	}
	(*n).sign = 0
	s1 := n.s[:index]
	s2 := n.s[index + 1:]
	return s1, s2
}



func (n *Numbers) correctNumbers() bool  {
	  s1, s2 := n.splitNumber()
	  var res2 int64 = 0

	  res1, err := strconv.ParseInt(s1, 10, 64)
	  if err != nil {
		return false
	  }
	 if s2 == "" {
		res2 = 0
	 } else {
		 res2, err = strconv.ParseInt(s2, 10, 64)
		 if err != nil {
			 return false
		 }
	 }
	(*n).integerPartNumber = res1
	(*n).fractionalPartNumber = res2
	(*n).lenAccuracy = len(s2)

	return true
}

func (n *Numbers)PerformIntegerNumbers() string{
	var ostatki int64
	delimoe := n.integerPartNumber
    if delimoe == 0 {
    	return "0"
	}
	ostatki, i := 0, 0
	for delimoe > 0 {
		ostatki += delimoe % 3 * int64(math.Pow10(i))
		delimoe /=  3
		i++

	}
	str := strconv.Itoa(int(ostatki))
	if n.sign == 1 {
		return "-" + str
	}
	return str
}

func Accuracy(length int) int {
	i := 0
	for math.Pow10(length) >= math.Pow(3, float64(i)) {
		i++
	}
	return i
}

func (n *Numbers) PerformFractionalNumber() string  {

	mnog := n.fractionalPartNumber
	if mnog == 0 {
		return " "
	}
	str := "0."
	length := Accuracy(n.lenAccuracy) // точность в троичной с/c
	for i := 0; i < length; i++ {
		str += strconv.Itoa(int(mnog * 3)/int(math.Pow10(n.lenAccuracy)))
		mnog = (mnog *3) % int64(math.Pow10(n.lenAccuracy))
	}
	(*n).lenAccuracy2 = length
    return str
}
func Equal(s string)(bool, error) {
	number, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println()
		return false, err
	}
	fmt.Println(number)
	return true, nil
}

func Equall(s string, accuracy, accuracy2 int ) string {
	mu.Lock()
	defer mu.Unlock()
	var number10System float64 = 0
	strAccur := strconv.Itoa(accuracy2)
	pow := 0

	for i:=0; i < len(s); i++ {
		if s[len(s)-1-i] == '.' {
			continue
		}
		pow++
		num, _ := strconv.ParseFloat(string(s[len(s)-1-i]), 64)
		number10System += num* math.Pow(3.0, float64(-accuracy + pow-1))
		fmt.Println(num, " :: ",  number10System, " :: ", -accuracy + pow-1)



	}
	//return fmt.Sprintf("%." + strAccur + "f", number10System)
	  fmt.Printf("%." + strAccur + "f\n", number10System)
      return ""
}
func main() {
    file, err := os.Open("input.txt")

    if err != nil {
    	panic(err)
	}
	defer file.Close()
    numbers := make([]Numbers, 0, 0)  // create slice Struct Numbers
    numbersAnotherSystem := make(map[string][]int)
	file2, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
    s := bufio.NewScanner(file)
    if s == nil  {
        _, _ = file2.WriteString("Исходный файл пуст")
        return
	}
    for s.Scan() {
    	numbers = append(numbers, Numbers{s.Text(), 0, 0, 0, 0, 0})
	}

	defer file2.Close()
	numbers2 := make([]Numbers, len(numbers)) // create slice Numbers2

	for index, value := range numbers {
		if !value.correctNumbers() {
			numbers2[index] = Numbers{"", 0,0,0,  0,  0}
			continue
		}
		numbers2[index] = Numbers{value.s, value.integerPartNumber, value.fractionalPartNumber, value.lenAccuracy, value.lenAccuracy2, value.sign}
	    fmt.Println(value.integerPartNumber, value.fractionalPartNumber)


	}
	fmt.Println("\n")

    for index, value := range numbers2 {
    	if value.s == "" {
    	   _, err = file2.Write([]byte("Число под номером " + strconv.Itoa(index + 1) +  " некорректно\n"))
    	   if err != nil {
			   fmt.Println(err)
		   }
    	   continue
		}

    	 finallyNumber := value.PerformIntegerNumbers() + (value.PerformFractionalNumber())[1:]
    	 numbersAnotherSystem[finallyNumber] = []int {value.lenAccuracy2, value.lenAccuracy}
    	 _, err = file2.Write([]byte("Число с номером " + strconv.Itoa(index + 1) + " из 10 с/c: " + value.s +
		 " в 3-ой с/с: " + finallyNumber + "\n"))
	     if err != nil {
		     fmt.Println(err)
	        }
		}

       for key, value := range numbersAnotherSystem {
			     Equall(key, value[0], value[1])
	   }

}
