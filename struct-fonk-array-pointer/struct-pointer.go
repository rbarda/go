package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("========14-02-2017=========")

	c1:=customer{1,"ramazan",12391238123}
	fmt.Println(c1.name,c1.tell)

	circle1:=circle{5}
	area:=circle1.area(circle1.radius)
	fmt.Println(area)

	//noNameFunction
	noNameFunction()

	//innerFunction
	addF:=innerFunctionAdd();
	fmt.Println(addF(4))
	fmt.Println(addF(9))
	fmt.Println(addF(412))
	addF2:=innerFunctionAdd();
	fmt.Println(addF2(14))
	fmt.Println(addF2(19))
	fmt.Println(addF2(1412))
	fmt.Println(addF(1))

	//Pointer
	p:=product{name:"Samsung",price:1000}
	discount(p,100)
	fmt.Printf("Pointer kullanmadan =>>%f\n",p.price)
	discountWithPointer(&p,100)
	fmt.Printf("Pointer kullandığımızda =>>%f",p.price)
}
type customer struct {
	id int
	name string
	tell int
}
type circle struct {
	radius float64
}
func(v circle)  area(r float64) float64{
	return 3.14*r*r
}
func noNameFunction (){
	f1 := func(r rune) rune {
		switch {
		case r == ' ':
			return '_'
		case r == 'b':
			return 'B'
		}
		return r
	}
	fmt.Println(strings.Map(f1, "bugun guzel bir gun"))
}
func innerFunctionAdd() func(int) int{
	total:=0
	return func(x int) int{
		total+=x
		return total
	}
}


////Pointer kullanımı icin
type product struct{
	name string
	price float32
}
func discountWithPointer(p *product,price float32){
	p.price=p.price-price
}
func discount(p product,price float32){
	p.price=p.price-price
}
