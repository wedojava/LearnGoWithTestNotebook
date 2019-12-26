package main

import "fmt"

const spanish = "Spanish"
const chinese = "Chinese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const chineseHelloPrefix = "你好, "
const suffix = "!"

// Hello just a begin of Go programing
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + suffix
}

// 公共函数( 例如上面的 Hello() )以大写字母开始，私有函数以小写字母开头。
func greetingPrefix(language string) (prefix string) { // 命名返回值 (prefix string) 将在函数中创建一个名为 prefix 的变量
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return // 只通过调用 return 而不是 return prefix 来返回值。
}

func main() {
	fmt.Println(Hello("world", ""))
}
