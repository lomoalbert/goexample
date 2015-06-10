package main

import (
    "fmt"
    "sort"
)

type point struct {
    x, y int
}

type points []point

func (p points) Len() int {return len(p)}
func (p points) Less(i,j int) bool {return p[i].x*p[i].y < p[j].x*p[j].y}
func (p points) Swap(i,j int) {p[i],p[j] = p[j],p[i]}

func main() {
    i := []int{4, 2, 57, 43, 38, 2, 507, 57, 100, 2}
    sort.Ints(i)
    fmt.Println(i)
    fmt.Println("已排序:",sort.IntsAreSorted(i))
    //搜索仅限于已排序的序列
    fmt.Println(sort.SearchInts(i, 38))
    sort.Sort(sort.Reverse(sort.IntSlice(i)))
    fmt.Println(i)

    f:=[]float64{1.1,1.2,1.5,1.32,1.3132}
    sort.Float64s(f)
    fmt.Println(f)
    //搜索仅限于已排序的序列
    fmt.Println("已排序:",sort.Float64sAreSorted(f))
    fmt.Println(sort.SearchFloat64s(f,1.32))

    s:=[]string{"abc","iop","efd","阿","波","次","的","zho","aec","wjd"}
    sort.Strings(s)
    fmt.Println(s)
    //搜索仅限于已排序的序列
    fmt.Println(sort.SearchStrings(s,"efd"))
    fmt.Println("已排序:",sort.StringsAreSorted(s))
    for _,i :=range s{
        fmt.Print([]byte(i))
    }

    p := points{point{3,2}, point{1,5}, point{5,7},point{5,2},point{1,7}}
    sort.Sort(p)
    fmt.Println("已排序:",sort.IsSorted(p))
    fmt.Println(p)
    point_min :=sort.Search(p.Len(),func(i int)bool{return p[i].x*p[i].y >5})
    fmt.Println("面积大于5的最小元素index:",point_min)
    sort.Sort(sort.Reverse(p))
    fmt.Println(p)
}