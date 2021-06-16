package main

//check doc
// go doc pakageName (ex: go doc fmt, go doc fmt.Println, go doc time.Now)

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
)

type Point struct {
	x int
	y int
}

type Circle struct {
	radius float32
	center *Point
}

type Student struct {
	name   string
	grades []int
	age    int
}

//method of struct
func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main() {
	/*var name string = "Hello Tim";
	var number uint16=260;
	var number1=260.33;
	number2 := 26;

	fmt.Printf("%T%T",number1,number2);
	fmt.Println("Hello world!"+name,number);
	var k string=fmt.Sprintf("Hello world!"+name,number);
	fmt.Printf(k);

	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	input, _ := strconv.ParseInt(scanner.Text(),10,64);
	fmt.Printf("%d",input)

	num1 := 9.6
	num2 := 3
	ans := num1/float64(num2)
	fmt.Println(ans)

	x:=5
	y:=6
	val:=x>=y
	fmt.Printf("%t",val)*/

	fmt.Println("Slice")

	a := [5]int{1, 2, 3, 4, 5}
	b := a[:]
	b[0] = 5

	fmt.Println(a, b)

	arr := [...]int{1, 2, 3, 4, 5}
	sl := arr[1:3]

	fmt.Println(sl, len(sl), cap(sl))

	sl = append(sl, 10, 100)
	fmt.Println("sl: ", sl, len(sl), cap(sl))
	fmt.Println(arr)

	fmt.Println(sl[1:len(sl)])
	fmt.Println(sl[:len(sl)-1])
	fmt.Println(append(sl[:2], sl[3:]...))

	for i, el := range sl {
		fmt.Println(i, el)
	}

	randArr := []int{5, 8, 3, 4, 6, 9, 4, 2, 5, 22}
	fmt.Println(randArr)

	for i, el := range randArr {
		for j := i + 1; j < len(randArr); j++ {
			el2 := randArr[j]
			if el == el2 {
				fmt.Println(el)
			}
		}
	}

	//map
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	mp2 := make(map[string]int)
	mp2["tim"] = 900
	delete(mp, "apple")
	fmt.Println(mp, mp2, mp["apple"])

	//check key ton tai trong map
	val, ok := mp["apple"]
	fmt.Println(val, ok)

	val2, ok2 := mp["pear"]
	fmt.Println(val2, ok2)

	fmt.Println(len(mp))

	ans1, ans2 := c(5, 6)
	fmt.Println(ans1, ans2)

	ans3, ans4 := add(5, 6)
	fmt.Println(ans3, ans4)

	re(589)(845)

	test := func() {
		fmt.Println("test")
	}
	test()

	fmt.Printf("%T \n", re)

	abc(func() int { return 10 })

	//struct
	p1 := &Point{4, 5}
	change(p1)
	fmt.Println(p1.y)

	c1 := Circle{4.56, &Point{5, 7}}
	fmt.Println(c1, c1.center.x)

	slice := []Point{Point{5, 6}, Point{7, 8}}
	fmt.Println(slice)

	//method struct
	s1 := Student{"Tim", []int{70, 80, 90, 85}, 19}
	s1.getAge()
	s1.setAge(7)
	fmt.Println(s1, s1.getAverageGrade(), s1.getMaxGrade())
}

func change(pt *Point) {
	pt.x = 2
}

func re(x int) func(y int) {
	return func(y int) { fmt.Println(x, y) }
}

//return nhieu gia tri
func c(x, y int) (int, float64) {
	return x + y, 1.56
}

func add(x, y int) (z1, z2 int) {
	defer fmt.Println("After return") //thuc hien sau khi func return
	z1 = x + y
	z2 = x - y
	return
}

func abc(s func() int) {
	fmt.Println(s())
}
