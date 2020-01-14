# LearnGoWithTestNotebook


Reference: 
- [通过测试学习 Go 语言](https://learnku.com/docs/learn-go-with-tests) / [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

- [Go by Example](https://gobyexample.com/)

- [GoDoc](https://godoc.org/)

Tips:

- Local Document: run `godoc -http :8000` then view http://localhost:8000

## Chapters

### 01 hello

#### TDD 过程和步骤

 - 编写一个失败的测试，并查看失败信息，我们知道现在有一个为需求编写的相关的测试，并且看到它产生了易于理解的失败描述

 - 编写最少量的代码使其通过，以获得可以运行的程序

 - 然后重构，基于我们测试的安全性，以确保我们拥有易于使用的精心编写的代码


### 02 integers

为测试的运行编写最少量的代码并检查失败测试的输出

adder_test.go:
```
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	excepted := 4

	if sum != excepted {
		t.Errorf("excepted '%d' but got '%d'", excepted, sum)
	}
}

func ExampleAdd() {
	// Attension: without `//Output: 6` this case can not run by `go test -v`
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
adder.go:
```
package integers

func Add(x, y int) int {
	return x + y
}
```
cmd:
```
$ go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
```
请注意，如果删除注释 「//Output: 6」，示例函数会被编译,但不会执行!

---