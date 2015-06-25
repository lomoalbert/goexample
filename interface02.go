package main
/*
任何实现某接口的所有方法的对象,均可作为接口参数运行相关函数
IPizzaCooker接口包含prepare bake cut 三个方法,任何支持这三个方法的对象均可作为IPizzaCooker接口参数执行cookOnePizza
*/

type Pizza struct {
}

type IPizzaCooker interface {
    Prepare(*Pizza)
    Bake(*Pizza)
    Cut(*Pizza)
}

func cookOnePizza(ipc IPizzaCooker) *Pizza {
    p := new(Pizza)
    ipc.Prepare(p)
    ipc.Bake(p)
    ipc.Cut(p)
    return p
}

type PizzaDefaultCooker struct {
...
}


func (this *PizzaDefaultCooker) Prepare(*Pizza) {
    //....default prepare pizza
}
func (this *PizzaDefaultCooker) Bake(*Pizza) {
    //....default bake pizza
}
func (this *PizzaDefaultCooker) Cut(*Pizza) {
    //....default cut pizza
}



type MyPizzaCooker PizzaDefaultCooker


func (this *MyPizzaCooker) Prepare(*Pizza) {
    //....bake pizza use my style
}
func (this *MyPizzaCooker) Bake(*Pizza) {
    //....bake pizza use my style
}
func (this *MyPizzaCooker) Cut(*Pizza) {
    //....bake pizza use my style
}

func main() {
    var cooker PizzaDefaultCooker
    p := cookOnePizza(cooker)
    var mycooker MyPizzaCooker
    m := cookOnePizza(mycooker)
    //....
}