package main

import "fmt"
import "sort"

/*程序 05：数字排序

题目：输入三个整数 x,y,z，请把这三个数由小到大输出。
1.程序分析：我们想办法把最小的数放到 x 上，先将 x 与 y 进行比较，如果 x>y 则将 x 与 y
的值进行交换，然后再用 x 与 z 进行比较，如果 x>z 则将 x 与 z 的值进行交换，这样能使 x 最小。
2.程序源代码：*/


func main () {
    var x,y,z int
    fmt.Scanf("%d %d %d\n",&x,&y,&z)
    i:=[]int{x,y,z}
    sort.Sort(sort.IntSlice(i))
    fmt.Printf("%d<%d<%d",i[0],i[1],i[2])
}