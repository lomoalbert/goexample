package gotest

import (
    "testing"
)

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