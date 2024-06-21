package designPattern

import (
	"fmt"
	"testing"
	"time"
)

type User struct {
	Name string
	Age  int
}

func TestDemo(t *testing.T) {
	// Test IndexFunc
	nums := make([]int, 2)
	index := IndexFunc(nums, func(u int) bool {
		return u == 0
	})
	fmt.Println(index)

	// Test Min
	num := Min(1, 2)
	fmt.Println(num)

	// Test Sort
	users := []User{
		{"Tom", 20},
		{"Jerry", 18},
		{"Alice", 22},
	}
	Sort(users, func(a, b User) bool {
		return a.Age >= b.Age
	})
	fmt.Println(users)
}

func TestDemo1(t *testing.T) {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()

	go func() {
		for {
			fmt.Println("a:", <-ch)
		}
	}()

	time.Sleep(2 * time.Second)
}
