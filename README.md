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

- Map 查找的有趣特性: 它可以返回两个值。第二个值是一个布尔值，表示是否成功找到 key。此特性允许我们区分 key 不存在还是未定义。
- Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map，你可以在 [这里](https://en.wikipedia.org/wiki/Hash_table) 阅读有关 hash tables 的更多信息。
- Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。
- 如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。下面两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常。
    - 你永远不应该初始化一个 nil 的 map 变量：`var m map[string]string`
    - 初始化空 map，或使用 make 关键字创建 map：`var dictionary =  map[string]string{}` 或 `dictionary =  make(map[string]string)`




TODO: 

study continue: //TODO: https://learnku.com/docs/learn-go-with-tests/maps/6086#28b468