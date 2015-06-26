package main

//func new(Type) *Type
//内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针
//func make(Type, size IntegerType) Type
/*
内建函数 make 用来为 slice，map 或 chan 类型分配内存和初始化一个对象(注意：只能用在这三种类型上)，跟 new 类似，第一个参数也是一个类型而不是一个值，跟 new 不同的是，make 返回类型的引用而不是指针，而返回值也依赖于具体传入的类型，具体说明如下：

*/

