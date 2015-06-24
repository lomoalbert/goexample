package main

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
}

func (this *PizzaDefaultCooker) CookOnePizza() *Pizza {
    return cookOnePizza(this)
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



type MyPizzaCooker struct {
    PizzaDefaultCooker
}

func (this *MyPizzaCooker) CookOnePizza() *Pizza {
    return cookOnePizza(this)
}
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
    var cooker MyPizzaCooker
    p := cooker.CookOnePizza()
    //....
}