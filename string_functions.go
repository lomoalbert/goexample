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
    //该函数s根据sep分割，返回分割之后子字符串的slice，返回的子串的长度如n的定义，如果sep为空，那么每一个字符都分割
    p("SplitN:    ", s.SplitN("a,b,c", ",", 2))
    //该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割
    p("SplitAfter:", s.SplitAfter("a,b,c", ","))
    //该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割.
    //不同的是仅分为N份,后续的将不再分割
    p("SplitAfterN:", s.SplitAfterN("a,b,c", ",", 2))
    //该函数把s字符串里面的每个单词首字母转化为大写
    p("Title:      ", s.Title("her royal highness"))
    //该函数把s字符串里面的每个单词转化为小写
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))

    //该函数把s字符串开头或者结尾里面包含字符集的字符全部过滤掉，返回过滤之后的字符串
    p("Trim:      ", s.Trim(" !!! Achtung !!! ", "! "))
    //该函数把s字符串里面开头和结尾部分字符传入f函数进行判断是否过滤，为真就过滤
    p("TrimFunc:   ", s.TrimFunc("abcdefg", func(r rune) bool {
        if r>'d' {
            return true
        }else {
            return false
        }
    }))
    //该函数把s字符串开头或者结尾里面空白符('\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))全部过滤掉，返回过滤之后的字符串
    p("TrimSpace:  ", s.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
    /*
    type Reader
    Reader是通过读取一个字符串之后实现了io.Reader, io.ReaderAt, io.Seeker, io.ByteScanner, 和io.RuneScanner 接口

    func NewReader(s string) *Reader
    参数列表

    s 读取的字符串
    返回值：

    *Reader 通过读取一个字符串之后返回Reader对象
    对象的方法列表：

    func (r *Reader) Len() int //返回未读取的字符串的长度
    func (r *Reader) Read(b []byte) (n int, err error) //读取数据到b中，返回读取的实际大小n，如果出错返回err，例如EOF或者b的长度为0
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) //按照指定的off位置开始读取内容到b，返回读取的实际大小n，如果出错返回err，例如off小于0或者大于本身的长度或者文件尾
    func (r *Reader) ReadByte() (b byte, err error) //读取一个byte的数据
    func (r *Reader) ReadRune() (ch rune, size int, err error) //读取一个rune的数据
    func (r *Reader) Seek(offset int64, whence int) (int64, error) //根据whence来移动offset，如果whence=0为直接移动offset位置，=1为移动到当前位置之后的offset，=2为移动到当前字符串长度之后的offset位置
    func (r *Reader) UnreadByte() error //当前读取的位置向前移一个byte
    func (r *Reader) UnreadRune() error //当前读取的位置向前移一个rune
    功能说明：

    该函数主要是通过把字符串读取Reader之后进行的一些读取操作
    */
    read := s.NewReader("I am asta谢")
    var b []byte
    fmt.Println(read.Len())  //12
    b = make([]byte, 8)
    n, err := read.Read(b)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b)   //[73 32 97 109 32 97 115 116]
    fmt.Println(n)   //8
    n, err = read.ReadAt(b, 3)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b)   //[109 32 97 115 116 97 232 176]
    fmt.Println(n)   //8
    bt, err := read.ReadByte()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bt)  //97
    rn, size, err := read.ReadRune()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(rn)    //35874
    fmt.Println(size)  //3
    /*
    type Replacer
这是一个字符串替换的对象

func NewReplacer(oldnew ...string) *Replacer
参数列表

oldnew是一个slice，是一个需要替换的字符串和新的字符串的配对出现
返回参数

Replacer返回一个替换对象
Replacer方法列表

func (r *Replacer) Replace(s string) string // 把字符串替换为oldnew定义的
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) //替换之后的字符串写入到w之中，返回写入的数量
应用示例，下面代码来自于beego的模板替换：
    */

    patterns := []string{"abc", "efg"}
    replacer := s.NewReplacer(patterns...)
    format := replacer.Replace("abc is abc is abc")
    fmt.Println(format)
    //efg is efg is efg
}
/*
Contains:   true
ContainsAny:   true
ContainsRune:   true
ContainsRune:   true
Count:      2
Equalold:   true
Fields:     [foo bar baz]
FieldsFunc: [abc efg]
HasPrefix:  true
HasSuffix:  true
Index:      1
IndexAny:   1
IndexFunc:  2
IndexRunc:  6
Join:       ab-cd
LastIndex:  10
LastIndexAny: 15
LastIndexFunc:  10
Map:        bcde
Repeat:     aaaaa
Replace:    fOOOO
Replace:    fOOoo
Split:      [a b c d e]
SplitN:     [a b,c]
SplitAfter: [a, b, c]
SplitAfterN: [a, b,c]
Title:       Her Royal Highness
ToLower:    test
ToUpper:    TEST
Trim:       Achtung
TrimFunc:    abcd
TrimSpace:   a lone gopher
12
[73 32 97 109 32 97 115 116]
8
[109 32 97 115 116 97 232 176]
8
97
35874
3
*/