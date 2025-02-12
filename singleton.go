package main

import (
	"fmt"
	"sync"
)

/*
	*

- Tuy nhiên đối với đoạn code bên dưới trong môi trường concurrent không có cơ chế synchronize đúng, thì nhiều goroutine có thể detect và cố tạo những instance mới at the same time
----> Và kết quả là nhiều instance được khởi tạo.

- Để giải quyết vấn đề này nó có thể sử dụng cơ chế để đảm bảo chỉ tạo một instance trong môi trường concurrent.
---> Trong package sync của Go có đối tượng sync.Once --> Thực hiện cơ chế thực thi một lần.
*/
type singleton struct {
	value string
}

var instance *singleton

func getInstance1() *singleton {
	if instance == nil {
		instance = &singleton{value: "unique value"}
	}

	return instance
}

/*
*
- Apply sync.Once
*/
var once sync.Once

func initInstance() {
	instance = &singleton{value: "unique value"}
}

// init only one time.
func getInstance() *singleton {
	if instance == nil {
		once.Do(initInstance)
	}
	return instance
}

func main() {
	firstInit := getInstance()

	fmt.Println(firstInit.value)

	anotherInstance := getInstance()
	if anotherInstance == firstInit {
		fmt.Println("Both instances are the same")
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singletonInstance := getInstance()
			fmt.Println(singletonInstance.value)
		}()
	}
	wg.Wait()
}
