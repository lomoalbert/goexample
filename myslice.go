package main

func main() {
    myslice := make([]int, 5)
    for range [21]int{} {
        println(myslice)
        println(len(myslice))
        println(cap(myslice))
        println("***********")
        myslice = append(myslice, 0)
    }

    myslice=make([]int, 5, 7)
    for range [18]int{} {
        println(myslice)
        println(len(myslice))
        println(cap(myslice))
        println("===========")
        myslice = append(myslice, 0)
    }
}
