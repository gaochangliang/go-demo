package math

import "testing"

func TestAdder(t *testing.T) {
	res := adder(4, 7)
	if res != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}

}

//测试go文件 go test -v math_test.go  math.go

//go test -v  math_test.go 因为math_test.go 文件中可能引入了其他的文件所以需要带上引入的函数go文件
