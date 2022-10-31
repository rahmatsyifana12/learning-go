package main

import (
	"errors"
	"fmt"
)

func main() {
	// nums := [...]int{ 1, 2, 3, 4, 5, 6 }
	// nums1 := nums[2:5]
	// var nums2 [5]int
	// copy(nums2[:], nums1)
	// nums2[0] = 50

	// var nums3 = append(nums1, 12)
	// // nums3 = append(nums3, 12)
	// nums3[2] = 100

	// fmt.Println(nums)
	// fmt.Println(nums1)
	// fmt.Println(nums3)
	// fmt.Println(nums2)

	// arr := [...]int{1, 2, 3, 4, 5}
	// arr2 := arr[1:3]
	// arr3 := append(arr2, 7)
	// arr3[0] = 100
	// fmt.Println(arr)
	// fmt.Println(arr2)
	// fmt.Println(arr3)
	// fmt.Println(cap(arr2))

	// arr := [...]int{ 1, 2, 3 }
	// arr1 := arr
	// arr1[1] = 4
	// fmt.Println(arr)
	// fmt.Println(arr1)

	// slice := []int{ 1, 2, 3 }
	// slice1 := slice
	// slice1[1] = 100

	// fmt.Println(slice)
	// fmt.Println(slice1)

	// person := map[string]int{ "hi": 1, "to": 2}
	// person1 := person
	// person1["hi"] = 12
	// fmt.Println(person)
	// fmt.Println(person1)

	// var book = make(map[string]string)
	// var slice = make([]string, 3)
	// var arr [5]int
	// arr[0] = 5
	// newArr := arr
	// newArr[0] = 1
	// slice[0] = "aan"
	// var a = slice
	// a[0] = "rahmat"
	// book["title"] = "Jordan"
	// // fmt.Println(slice)
	// // fmt.Println(a)
	// // fmt.Println(arr)
	// // fmt.Println(newArr)
	// fmt.Println(len(book))

	// i := 0
	// s := true
	// for s && i < 10 {
	// 	i++
	// 	if i == 10 {
	// 		s = false
	// 	}
	// }
	// fmt.Println(i)

	// i = 0
	// for j := 0; j < 5; j++ {
	// 	fmt.Println(j)
	// }

	// arr := [...]int{ 1, 2, 3, 4, 5 }
	// slice := []int{ 3, 2, 9, 0 }

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// for idx, sl := range slice {
	// 	fmt.Println(idx, sl)
	// }

	// name := "rahmat"
	// fmt.Println(name)
	// newName := changeName(&name)
	// fmt.Println(newName)

	// firstName, lastName := getName()
	// fmt.Println(firstName, lastName)
	// a := sumAll(1, 2, 3, 4);
	// fmt.Println(a)

	// getHello1 := getHello
	// getHello1("Rahmat")

	// name := "aan"
	// sayHelloWithFilter(name, filterSpam)

	// name := "anjing"
	// register(name, func (name string) bool {
	// 	return name == "anjing"
	// });

	// res := factorial(5)
	// fmt.Println(res)

	// name := "rahmat"

	// count := func() {
	// 	name = "aan"
	// 	fmt.Println(&name)
	// }
	// count()

	// fmt.Println(&name)
	// defer endApp()
	// res := bagi(2, 0)
	// fmt.Println(res)

	// slice := make([]int, 5, 6)
	// slice1 := slice
	// fmt.Printf("%p - %p\n", &slice, &slice1)
	// slice1[1] = 2
	// slice1 = append(slice1, 2, 1)
	// fmt.Printf("%p\n", &slice[0])
	// fmt.Printf("%p\n", &slice1[0])
	// fmt.Printf("%p - %p\n", &slice, &slice1)

	// arr := [...]int{ 1, 2, 3 }
	// arr1 := arr
	// fmt.Printf("%p %p\n", &arr[0], &arr1[0])
	// fmt.Println(arr)
	// fmt.Println(arr1)

	// slice1 := make([]int, 5)
	// slice2 := slice1

	// fmt.Printf("%p - %p\n", &slice1, &slice2)
	// fmt.Printf("%p - %p\n", &slice1[1], &slice2[1])
	// A := rune('A')
	// fmt.Println("A:", A)

	// text := "hello world"
    // for _, char := range text {
    //     fmt.Printf("%v ", string(char))
    // }

	// my_char := 'a'
	// fmt.Println(string(my_char))

	// res, err := getBagi(7, 0)
	// if (err != nil) {
	// 	panic(err.Error())
	// }
	// fmt.Println(res)

	// create new pointer variable without any data in it
	// val_ptr := new(int)
	// *val_ptr = 5
	// val := *val_ptr
	// fmt.Println(val)
}

func getBagi(num1 int, num2 int) (res int, err error) {
	if (num2 == 0) {
		return 0, errors.New("Pembagian tidak dapat dibagi 0")
	}
	return num1/num2, nil
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
	fmt.Println("Aplikasi berhenti")
}

func bagi(num1 int, num2 int) int {
	return num1/num2
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

type BlackList func(string) bool

func register(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are being black list")
	} else {
		fmt.Println("Successfully registered")
	}
}

func changeName(name* string) string {
	*name = "aan"
	return *name
}

func getName() (string, string) {
	return "Rahmat", "syifana"
}

func sumAll(num int, nums ...int) int {
	sum := 0
	fmt.Println(nums)
	nums = append(nums, 12)
	nums[0] = 9
	fmt.Println(nums)
	for _, val := range nums {
		sum += val
	}
	return sum
}

func getHello(name string) {
	fmt.Println("Hello", name)
}

func sayHelloWithFilter(name string, filter func(*string)) {
	filter(&name)
	fmt.Println("Hello", name)
}

func filterSpam(name* string) {
	if *name == "anjing" {
		*name = "..."
	}
}
