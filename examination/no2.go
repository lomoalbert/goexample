package main

import "fmt"
import "math"

/*程序 02：阶梯利润分配

题目：企业发放的奖金根据利润提成。
利润(I)低于或等于10万元时，奖金可提成10%；
利润高于10万元，低于20万元，低于10万元的部分按10%提成，高于10万元的部分，可提成7.5%；
20万到40万之间时，高于20万元的部分，可提成5%；
40万到60万之间时高于40万元的部分，可提成3%；
60万到100万之间时，高于60万元的部分，可提成1.5%;
高于100万元时，超过100万元的部分按1%提成，从键盘输入当月利润I，求应发放奖金总数？
1.程序分析：请利用数轴来分界，定位。注意定义时需把奖金定义成长整型。
2.程序源代码：*/
func main() {
    type rul struct {
        I float64
        L float64
    }

    var yuan float64
    var result float64
    var costed float64



    fmt.Println("请输入奖金")
    fmt.Scanf("%f\n", &yuan)

    rull := []rul{rul{1e5, 0.1},
        rul{2e5, 0.075},
        rul{4e5, 0.05},
        rul{6e5, 0.03},
        rul{1e6, 0.015},
        rul{math.MaxFloat64, 0.01},
    }
    for _, i := range rull {
        if yuan > costed {
            result += (math.Min(yuan, i.I)-costed)*i.L
            costed = i.I
        }else {
            break
        }
    }
    fmt.Println(result)
}