// strings包的使用
package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 1 EqualFold
	ans := strings.EqualFold("zhangsan", "Zhangsan")
	fmt.Println(ans)
	ans = strings.EqualFold("LISI-123", "list-123 ")
	fmt.Println(ans)

	fmt.Println("==============EqualFold↑==============")

	// 2 TrimSpace
	str := "                  Zhang              San               "
	fmt.Printf("原字符串: \"%s\"\n", str)
	str = strings.TrimSpace(str)
	fmt.Printf("去除前后空格之后的字符串: \"%s\"\n", str)

	fmt.Println("==============TrimSpace↑==============")

	// 3 HasPrefix
	str = "Hello World"
	ans = strings.HasPrefix(str, "Hello")
	fmt.Println(ans)
	ans = strings.HasPrefix(str, "World")
	fmt.Println(ans)

	fmt.Println("==============HasPrefix↑==============")

	// 4 HasSuffix
	str = "Zhang San"
	ans = strings.HasSuffix(str, "san")
	fmt.Println(ans)
	ans = strings.HasSuffix(str, "San")
	fmt.Println(ans)

	fmt.Println("==============HasSuffix↑==============")

	// 5 Index
	str = "Do you like sugar? Do you like cat?"
	idx := strings.Index(str, "like")
	fmt.Println(idx)
	idx = strings.Index(str, "dog")
	fmt.Println(idx)

	fmt.Println("==============Index↑==============")

	// 6 LastIndex
	str = "Do you like sugar? Do you like cat?"
	idx = strings.LastIndex(str, "like")
	fmt.Println(idx)
	idx = strings.LastIndex(str, "dog")
	fmt.Println(idx)

	fmt.Println("==============LastIndex↑==============")

	// 7 IndexAny
	str = "Ambitious"
	idx = strings.IndexAny(str, "hello")
	fmt.Println(idx)
	idx = strings.IndexAny(str, "iti")
	fmt.Println(idx)

	fmt.Println("==============IndexAny↑==============")

	// 8 IndexByte
	str = "Unbelievable"
	idx = strings.IndexByte(str, 'e')
	fmt.Println(idx)
	idx = strings.IndexByte(str, 'b')
	fmt.Println(idx)

	fmt.Println("==============IndexByte↑==============")

	// 9 Replace
	str = "Hello sarah do oh do"
	fmt.Println(strings.Replace(str, "o", "aaa", 1))
	fmt.Println(strings.Replace(str, "o", "aaa", -1))

	fmt.Println("==============Replace↑==============")

	// 10 Title
	str = "when your talent can't support your ambition"
	fmt.Println(strings.Title(str))

	fmt.Println("==============Title↑==============")

	// 11 ToTitle
	str = "when your talent can't support your ambition"
	fmt.Println(strings.ToTitle(str))

	fmt.Println("==============ToTitle↑==============")

	// 12 ToLower
	str = "Hello World!"
	fmt.Println(strings.ToLower(str))

	fmt.Println("==============ToLower↑==============")

	// 13 ToUpper
	str = "when your talent can't support your ambition"
	fmt.Println(strings.ToUpper(str))

	fmt.Println("==============ToUpper↑==============")

	// 14 Contains
	str = "I am ambitious!"
	fmt.Println(strings.Contains(str, "ambitious"))

	fmt.Println("==============Contains↑==============")

	// 15 ContainsAny
	str = "You should calm down and learn!"
	fmt.Println(strings.ContainsAny(str, "abc"))

	fmt.Println("==============ContainsAny↑==============")

	// 16 IndexRune
	str = "Ambitious哈哈哈"
	fmt.Println(strings.IndexRune(str, 'i'))
	fmt.Println(strings.IndexRune(str, utf8.RuneError))

	fmt.Println("==============IndexRune↑==============")

	// 17 Count
	str = "Ambitious"
	fmt.Println(strings.Count(str, "i"))
	fmt.Println(strings.Count(str, ""))

	fmt.Println("==============Count↑==============")

	// 18 Repeat
	str = "Ambitious"
	fmt.Println(strings.Repeat(str, 3))
	// fmt.Println(strings.Repeat(str, -1))
	// fmt.Println(strings.Repeat(str, 999999999999999))

	fmt.Println("==============Repeat↑==============")

	// 19 Trim
	str = "00000000000000000000000Zhang00000000000San0000000000000"
	fmt.Printf("原字符串: \"%s\"\n", str)
	str = strings.Trim(str, "0")
	fmt.Printf("去除前后字符之后的字符串: \"%s\"\n", str)

	fmt.Println("==============Trim↑==============")

	// 20 TrimLeft
	str = "00000000000000000000000Zhang00000000000San0000000000000"
	fmt.Printf("原字符串: \"%s\"\n", str)
	str = strings.TrimLeft(str, "0")
	fmt.Printf("去除前边字符之后的字符串: \"%s\"\n", str)

	fmt.Println("==============TrimLeft↑==============")

	// 21 TrimRight
	str = "00000000000000000000000Zhang00000000000San0000000000000"
	fmt.Printf("原字符串: \"%s\"\n", str)
	str = strings.TrimRight(str, "0")
	fmt.Printf("去除后边边字符之后的字符串: \"%s\"\n", str)

	fmt.Println("==============TrimRight↑==============")

	// 22 TrimPrefix
	str = "When your ability can't realize your dream"
	fmt.Println(strings.TrimPrefix(str, "When"))
	fmt.Println(strings.TrimPrefix(str, "hen"))

	fmt.Println("==============TrimRight↑==============")

	// 23 Fields
	str = "You should practise with all your heart"
	fmt.Println(strings.Fields(str))

	fmt.Println("==============Fields↑==============")

	// 24 FieldsFunc
	str = "Nice to meet you.2Are you ok?"
	fmt.Println(strings.FieldsFunc(str, func(c rune) bool {
		return unicode.IsDigit(c)
	}))

	fmt.Println("==============FieldsFunc↑==============")

	// 25 ContainsRune
	str = "Ambitious"
	fmt.Println(strings.ContainsRune(str, 'A'))
	fmt.Println(strings.ContainsRune(str, 'z'))

	fmt.Println("==============ContainsRune↑==============")

	// 26 Split
	str = "I am coding now, please go out."
	fmt.Println(strings.Split(str, ","))

	fmt.Println("==============Split↑==============")

	// 27 Join
	slice := []string{"I", "am", "ambitious."}
	fmt.Println(strings.Join(slice, " "))

	fmt.Println("==============Join↑==============")

	// 28 SplitN
	str = "a,b,c,d,e,f"
	fmt.Println(strings.SplitN(str, ",", 3))
	fmt.Println(strings.SplitN(str, ",", -1))

	fmt.Println("==============SplitN↑==============")

	// 29 SplitAfter
	str = "a,b,c,d"
	fmt.Println(strings.SplitAfter(str, ","))
	fmt.Println(strings.SplitAfter(str, "b"))

	fmt.Println("==============SplitAfter↑==============")

	// 30 SplitAfterN
	str = "a,b,c,d,b,s"
	fmt.Println(strings.SplitAfterN(str, "b", 2))
	fmt.Println(strings.SplitAfterN(str, "b", -1))

	fmt.Println("==============SplitAfterN↑==============")

	// 31 Cut
	str = "127.0.0.1:8080"
	ip, port, ok := strings.Cut(str, ":")
	fmt.Printf("ip: %s\nport: %s\nfound: %t\n", ip, port, ok)
	ip, port, ok = strings.Cut(str, "m")
	fmt.Printf("ip: %s\nport: %s\nfound: %t\n", ip, port, ok)

	fmt.Println("==============Cut↑==============")

}
