package main

import "fmt"

/*程序 01：数字排列组合

题目：有 1、2、3、4 个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
1.程序分析：可填在百位、十位、个位的数字都是 1、2、3、4。组成所有的排列后再去
掉不满足条件的排列。
2.程序源代码：*/
func main() {
    nolist := []int{1, 2, 3, 4}
    for _, x := range (nolist) {
        for _, y := range (nolist) {
            for _, z := range (nolist) {
                if x!=y && y!=z && x!=z {
                    fmt.Println(x, y, z)
                }
            }
        }
    }
}