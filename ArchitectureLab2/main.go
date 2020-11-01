package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*var(
	mu sync.Mutex
	wg sync.WaitGroup
)*/

type Numbers struct {
	s string    // строковое представление числа в файле
	integerPartNumber int64  // целая часть числа в файле
	fractionalPartNumber int64 // дробная часть числа в файле
	lenAccuracy int // точность числа в 10 c/c
	lenAccuracy2 int // точность в четверичной с/c
	sign        int // знак числа
}
type FinallyAccuracy int
type FinallySign int

type Numbers2 struct {
	integerNumber string // целая часть в 4-ой с/c
	fractionalNumber string // дробная часть в 4-ой с/c
	sign int
}

type FinallyNumber struct{
	finallyNumber string
	sign int
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
		ostatki += delimoe % 4 * int64(math.Pow10(i))
		delimoe /=  4
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
	for math.Pow10(length) >= math.Pow(4, float64(i)) {
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
		str += strconv.Itoa(int(mnog * 4)/int(math.Pow10(n.lenAccuracy)))
		mnog = (mnog *4) % int64(math.Pow10(n.lenAccuracy))
	}
	(*n).lenAccuracy2 = length
	return str
}

func SumIntegerNumber(n string, m string, addit int) string {
	s := ""
    finallyLen := 0
    if n[0] == '-' { n = n[1:]}
    if m[0] == '-' { m = m[1:]}
   //  fmt.Println(len(n), len(m))
    if len(n) == len(m) { finallyLen = len(m)}
	if len(n) != len(m) {
		if len(n) < len(m) {
			for i := 0; i <= len(m) - len(n) + 1; i++ {
				n = "0" + n
			}
			finallyLen = len(m)
		} else {
			for i := 0; i < len(n) - len(m) + 1; i++ {
				m = "0" + m
			}
			finallyLen = len(n)
		}
	}

	//if n[0] == '-' {finallyLen--}
	fmt.Println(n, m, addit)
	masAddit := make([]int, finallyLen+1)
    if addit == 1 { masAddit[0] = 1 }

	for i:=0; i < finallyLen; i++ {

		number1, _ := strconv.Atoi(string(n[finallyLen - i - 1]))
		number2, _ := strconv.Atoi(string(m[finallyLen - i - 1]))

		if number1 + number2 + masAddit[i] >= 4  {
			masAddit[i + 1] = 1

			s1 := strconv.Itoa((number1 + number2 + masAddit[i])% 4) + s
			s = s1
			fmt.Println(s)
		} else {
			s1 := strconv.Itoa((number1+number2+masAddit[i])%4) + s
			s = s1
			fmt.Println(s)
		}
	}
	if masAddit[finallyLen] == 1 {
		s1 := "1" + s
		s = s1
	}
	return s
}
func SumFractionalNumber(n string, m string) (string, int) {
	s := ""

    if n == " " && m == " " { return "", 0}

	if n != " " { n = n[2:] }
	if m != " " { m = m[2:] }
    if n == " " && len(m) >= 1 {
    	return m, 0
	} else if m == " " && len(n) >= 1 {
		return n, 0
	}
    finallyLen := 0

	if len(n) != len(m)  {

		if len(n) < len(m) {
			for i := 0; i < len(m) - len(n) + 1; i++ {
				n += "0"
			}
			finallyLen = len(m)
		} else {
			for i := 0; i < len(n) - len(m) + 1; i++ {
				m += "0"
			}
			finallyLen = len(n)
		}

	}

        masAddit := make([]int, finallyLen + 1)

		for i:=0; i < finallyLen; i++ {
			number1, _ := strconv.Atoi(string(n[finallyLen - i - 1]))
			number2, _ := strconv.Atoi(string(m[finallyLen - i - 1]))
			if number1 + number2 + masAddit[i] >= 4  {
				     masAddit[i + 1] = 1
                     s1 := strconv.Itoa((number1 + number2 + masAddit[i])% 4) + s
                     s = s1
                     fmt.Println(s)
			} else {
				s1 := strconv.Itoa((number1+number2+masAddit[i])%4) + s
				s = s1
				fmt.Println(s)
			}
		}
    p := 0
    if masAddit[finallyLen] == 1 {
    	p = 1
	}
	return s, p
}
func InSumIntegerNumber(n string, m string, addit int) string {
	s := ""
    if n[0] == '-' { n = n[1:]}
	if m[0] == '-' { m = m[1:]  }

	fmt.Println(n, m)
	if len(n) != len(m) {
		if len(n) < len(m) {
			for i := 0; i <= len(m) - len(n) + 1; i++ {
				n = "0" + n
			}
		} else {
			for i := 0; i <= len(n) - len(m) + 1; i++ {
				m = "0" + m
			}
		}
	}
	fmt.Println(n, m)
	fmt.Println(n, m, addit)
	masAddit := make([]int, len(n)+1)
	if addit == -1 { masAddit[0] = -1 }
	for i:=0; i < len(n); i++ {

		number1, _ := strconv.Atoi(string(n[len(n) - i - 1]))
		number2, _ := strconv.Atoi(string(m[len(n) - i - 1]))

		if number1 - number2 + masAddit[i] < 0 {
			masAddit[i + 1] = -1

			s1 := strconv.Itoa((number1 - number2 + 4+ masAddit[i])% 4) + s
			s = s1

		} else {
			s1 := strconv.Itoa((number1-number2+masAddit[i])%4) + s
			s = s1

		}
	}
	if masAddit[len(n)] == -1 {
		s1 := "-1" + s
		s = s1
	}
	return s
}


func InSumFractionalNumber(n string, m string) (string, int) {
	s := ""

	if n == " " && m == " " { return "", 0}

	if n != " " { n = n[2:] }
	if m != " " { m = m[2:] }
	if n == " " && len(m) >= 1 {
		return m, 0
	} else if m == " " && len(n) >= 1 {
		return n, 0
	}

	if len(n) != len(m)  {

		if len(n) < len(m) {
			for i := 0; i < len(m) - len(n) + 1; i++ {
				n += "0"
			}
		} else {
			for i := 0; i < len(n) - len(m) + 1; i++ {
				m += "0"
			}
		}

	}
	maxI := int(math.Max(float64(len (n)), float64(len(m))))
	masAddit := make([]int, maxI + 1)
    fmt.Println(n, m)
	for i:=0; i < maxI; i++ {
		number1, _ := strconv.Atoi(string(n[len(n) - i - 1]))
		number2, _ := strconv.Atoi(string(m[len(n) - i - 1]))
		if number1 - number2 + masAddit[i] < 0  {
			masAddit[i + 1] = -1
			s1 := strconv.Itoa((number1 - number2+4+ masAddit[i])%4) + s
			s = s1

		} else {
			s1 := strconv.Itoa((number1 - number2 +masAddit[i])%4) + s
			s = s1

		}
	}
	p := 0
	if masAddit[len(n)] == -1 {
		p = -1
	}
	return s, p
}

func Equall(s string, accuracy, accuracy2 int ) string {
	//mu.Lock()
	//defer mu.Unlock()
	var number10System float64 = 0
	strAccur := strconv.Itoa(accuracy2)
	pow := 0

	for i:=0; i < len(s); i++ {
		if s[len(s)-1-i] == '.' {
			continue
		}
		pow++
		num, _ := strconv.ParseFloat(string(s[len(s)-1-i]), 64)
		number10System += num* math.Pow(4.0, float64(-accuracy + pow-1))
		fmt.Println(num, " :: ",  number10System, " :: ", -accuracy + pow-1)



	}
	//return fmt.Sprintf("%." + strAccur + "f", number10System)
	fmt.Printf("%." + strAccur + "f\n", number10System)
	return ""
}
// определяем знак суммы
func  Sign(t1, t2 Numbers2) string {
	if t1.sign == 0 && t2.sign == 0 {
		return ""
	} else if t1.sign == 1 && t2.sign == 1 {
		return "-"
	} else if t1.sign != t2.sign {
		if t1.integerNumber < t2.integerNumber && t1.sign == 1 {
			return ""
		} else if t1.integerNumber < t2.integerNumber && t1.sign == 0 {
			return "-"
		} else if  t1.integerNumber == t2.integerNumber {
			if t1.fractionalNumber < t2.fractionalNumber && t1.sign == 1 {
				return ""
			} else {
				return "-"
			}
		}
	}
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

	file3, err := os.Create("output1.txt")
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
			" в 4-ой с/с: " + finallyNumber + "\n"))
		if err != nil {
			fmt.Println(err)
		}
	}

	for key, value := range numbersAnotherSystem {
		Equall(key, value[0], value[1])
	}

	fmt.Println("Сложение чисел в 4-ой с.с")
	for i := 0; i < len(numbers2) - 1; i += 2 {
		   if (numbers2[i].sign == 0 && numbers2[i+1].sign == 0) || (numbers2[i].sign == 1 && numbers2[i+1].sign == 1)  {
                first4Number := Numbers2{numbers2[i].PerformIntegerNumbers(), numbers2[i].PerformFractionalNumber(), numbers2[i].sign}
                second4Number := Numbers2{numbers2[i+1].PerformIntegerNumbers(), numbers2[i+1].PerformFractionalNumber(), numbers2[i+1].sign}
                fmt.Println(first4Number, second4Number)
                fractional, addPoint := SumFractionalNumber(first4Number.fractionalNumber, second4Number.fractionalNumber)
                fmt.Println(fractional, addPoint)
			   integer := SumIntegerNumber(first4Number.integerNumber, second4Number.integerNumber, addPoint)
			   fmt.Println(integer)
			   sign := Sign(first4Number, second4Number)
			   first10Number, err := strconv.ParseFloat(numbers2[i].s, 64)
			   if err != nil {
				   fmt.Println("ошибка")
			   }
			   second10Number, err := strconv.ParseFloat(numbers2[i+1].s, 64)
			   if err != nil {
				   panic(err)
			   }
			   finally10Number := fmt.Sprintf("%f", first10Number+second10Number)
			   Equall(sign + integer + "." + fractional, len(fractional), int(math.Max(float64(numbers2[i].lenAccuracy), float64(numbers2[i+1].lenAccuracy))))

			   _, err = file3.Write([]byte("Сумма чисел в четверичной системе счисления: " + first4Number.integerNumber+first4Number.fractionalNumber[1:] + " и " + second4Number.integerNumber+second4Number.fractionalNumber[1:] + " = " + sign + integer+"."+ fractional +
				   "\nСумма чисел в десятичной системе счисления: " + numbers2[i].s + " и " +  numbers2[i+1].s + " = " + finally10Number + "\n"))

		   } else {
			   first4Number := Numbers2{numbers2[i].PerformIntegerNumbers(), numbers2[i].PerformFractionalNumber(), numbers2[i].sign}
			   second4Number := Numbers2{numbers2[i+1].PerformIntegerNumbers(), numbers2[i+1].PerformFractionalNumber(), numbers2[i+1].sign}
			   fmt.Println(first4Number, second4Number)
			   if first4Number.integerNumber < second4Number.integerNumber {
			   	first4Number, second4Number = second4Number, first4Number
			   }
			   fractional, addPoint := InSumFractionalNumber(first4Number.fractionalNumber, second4Number.fractionalNumber)
			   fmt.Println(fractional, addPoint)
			   integer := InSumIntegerNumber(first4Number.integerNumber, second4Number.integerNumber, addPoint)
			   fmt.Println(integer)
			   sign := Sign(first4Number, second4Number)
			   first10Number, err := strconv.ParseFloat(numbers2[i].s, 64)
			   if err != nil {
			   	fmt.Println("ошибка")
			   }
			   second10Number, err := strconv.ParseFloat(numbers2[i+1].s, 64)
			   if err != nil {
			   	panic(err)
			   }
			   finally10Number := fmt.Sprintf("%f", first10Number+second10Number)
			   Equall(sign + integer + "." + fractional, 4, 2)
			   _, err = file3.Write([]byte("Сумма чисел в четверичной системе счисления: " + first4Number.integerNumber+first4Number.fractionalNumber[1:] + " и " + second4Number.integerNumber+second4Number.fractionalNumber[1:] + " = " + sign + integer+"."+ fractional +
				   "\nСумма чисел в десятичной системе счисления: " + numbers2[i].s + " и " +  numbers2[i+1].s + " = " + finally10Number + "\n"))
		   }
	}
}


/*if numbers2[i].sign == 0 && numbers2[i+1].sign == 0 {
			s := numbers2[i].PerformIntegerNumbers() + (numbers2[i].PerformFractionalNumber())[1:]
			s1 := numbers2[i+1].PerformIntegerNumbers() + (numbers2[i+1].PerformFractionalNumber())[1:]
			lastNumber, addit := SumFractionalNumber(numbers2[i].PerformFractionalNumber(), numbers2[i+1].PerformFractionalNumber())
			fmt.Println(lastNumber, addit)
			firstNumber := SumIntegerNumber(numbers2[i].PerformIntegerNumbers(), numbers2[i+1].PerformIntegerNumbers(), addit)
			finallyNumbers := firstNumber + "." + lastNumber
			_, err = file3.Write([]byte("Сумма чисел в четверичной системе счисления: " + s + " и " + s1 + " = " + finallyNumbers +
				"\nСумма чисел в десятичной системе счисления: "))
			fmt.Println(finallyNumbers)
			fmt.Println(Equall(finallyNumbers, 2, 1))
		} else {
		a, b := numbers2[i], numbers2[i+1]
		s := ""
		if a.sign == 1 {s = "-"}
		if b.sign == 1 {s = "-"}
        if a.integerPartNumber < b.integerPartNumber {
        	a, b = b, a

		} else if math.Abs(float64(a.integerPartNumber)) ==  math.Abs(float64(b.integerPartNumber)) {
			if math.Abs(float64(a.fractionalPartNumber)) <  math.Abs(float64(b.fractionalPartNumber)) {
				a, b = b, a
			}
		}
		fmt.Println(s, a.integerPartNumber, a.fractionalPartNumber, b.integerPartNumber, b.fractionalPartNumber)

		lastNumber, addit := InSumFractionalNumber(a.PerformFractionalNumber(), b.PerformFractionalNumber())
		fmt.Println(lastNumber, addit)
		firstNumber := InSumIntegerNumber(a.PerformIntegerNumbers(), b.PerformIntegerNumbers(), addit)
		finallyNumbers := firstNumber + "." + lastNumber
		fmt.Println(finallyNumbers)
		fmt.Println(Equall(finallyNumbers, 2, 1))
		}

 */
