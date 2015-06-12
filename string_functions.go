package main

import s "strings"
import "fmt"

//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println
//string中是以byte计算的,如果涉及到unicode的计算,需要转换为rune数组
func main() {
    //Here’s a sample of the functions available in strings. Note that these are all functions from package, not methods on the string object itself. This means that we need pass the string in question as the first argument to the function.
    p("Contains:  ", s.Contains("test", "es"))
    p("ContainsAny:  ", s.ContainsAny("test", "abce"))
    p("ContainsRune:  ", s.ContainsRune("test", 'e'))
    p("ContainsRune:  ", s.ContainsRune("test中文", '中'))
    p("Count:     ", s.Count("test", "t"))
    p("Equalold:  ", s.EqualFold("ABC中文", "abC中文"))
    //以空格split字符串
    p("Fields:    ", s.Fields("  foo bar  baz   "))
    //以函数返回true的rune作为分隔符 分割字符串
    p("FieldsFunc:", s.FieldsFunc("abcdefg", func(r rune) bool {
        if r=='d' {
            return true
        }else {
            return false
        }}))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    //该函数主要判断chars集中任意的一个字符在s串中第一次出现的位置，如果不存在返回-1
    p("IndexAny:  ", s.IndexAny("test", "abcde"))
    //该函数主要判断s中的每一个字符传入函数f，如果符合，那么返回该字符的位置，如果都不符合则返回-1
    p("IndexFunc: ", s.IndexFunc("test", func(r rune) bool {
        if r=='s' {
            return true
        }else {
            return false
        }
    }))
    //该函数主要判断unicode r在s串中第一次出现的位置，如果不存在返回-1
    p("IndexRunc: ", s.IndexRune("中华人民共和国", '人'))
    p("Join:      ", s.Join([]string{"ab", "cd"}, "-"))
    //该函数主要判断sep串在s串中最后一次出现的位置，如果不存在返回-1
    p("LastIndex: ", s.LastIndex("chickenkenken", "ken"))
    //该函数主要判断chars集中任意的一个字符在s串中最后一次出现的位置，如果不存在返回-1
    p("LastIndexAny:", s.LastIndexAny("chickenkenkenken", "iken"))
    //该函数主要判断s中的每一个字符传入函数f，返回符合函数f的最后一个字符的位置，如果都不符合则返回-1
    p("LastIndexFunc: ", s.LastIndexFunc("testtesttest", func(r rune) bool {
        if r=='s' {
            return true
        }else {
            return false
        }
    }))
    //该函数以此读取s中的字符，传入mapping函数，然后返回的字符链接起来，说白了就是字符串的每一个字符通过mapping函数的处理，最后返回处理好的字符串，如果处理不正确，那么就抛弃该字符
    p("Map:       ", s.Map(func(r rune) rune {return r+1}, "abcd"))
    //该函数返回一个s的重复count字数的字符串
    p("Repeat:    ", s.Repeat("a", 5))
    //该函数实现在s中把old替换为new字符串，替换次数为n，如果n小于0，那么就全部替换
    p("Replace:   ", s.Replace("foooo", "o", "O", -1))
    p("Replace:   ", s.Replace("foooo", "o", "O", 2))
    //该函数s根据sep分割，返回分割之后子字符串的slice，如果sep为空，那么每一个字符都分割
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()
    //You can find more functions in the strings package docs.
    //Not part of strings but worth mentioning here are the mechanisms for getting the length of a string and getting a character by index.
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
/*
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       ab-cd
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:  5
Char: 101
*/