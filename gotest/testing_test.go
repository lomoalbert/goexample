package gotest

import (
    "testing"
)

/*
使用原则

1. 文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
2. 你必须import testing这个包
3. 所有的测试用例函数必须是Test开头
4. 测试用例会按照源代码中写的顺序依次执行
5. 测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是
   测试状态
6. 测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，
   但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
7. 函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，
   说明测试不通过，调用Log方法用来记录测试的信息。
*/


//功能测试
func Test_Division(t *testing.T) {
    if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
        t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
    } else {
        t.Log("第一个测试通过了") //记录一些你期望记录的信息
    }
}

//性能测试
func Benchmark_Division(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Division(6, 2)
    }
}

/*
$ go test
PASS
ok      _/home/albert/code/go/goexample/gotest  0.002s
$ go test -bench=.
PASS
BenchmarkDivision       200000000                6.84 ns/op
ok      _/home/albert/code/go/goexample/gotest  2.079s
$ go test -bench=. -v
=== RUN Test_Division
--- PASS: Test_Division (0.00s)
        testing_test.go:12: 第一个测试通过了
PASS
Benchmark_Division      200000000                6.77 ns/op
ok      _/home/albert/code/go/goexample/gotest  2.059s
*/