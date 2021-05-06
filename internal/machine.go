package internal

import "github.com/Chermenator/coffe_machin/models"

type Machine struct {
	Water int
	Milk int
	CoffeeBeans int
	Cups int
	Money int

	CoffeTypes map[string]models.Coffe
}


func NewMachine(wat, milk, beans, cups, mon int) *Machine {
	return &Machine{
		Water: wat,
		Milk: milk,
		CoffeeBeans: beans,
		Cups: cups,
		Money: mon,
	}
}
func (m *Machine) InitCoffeeTypes() {
	coffeTypes := make(map[string]models.Coffe)
	coffeTypes["cappuchino"] = models.Coffe{
		Cups:        1,
		CoffeeBeans: 16,
		Water:       200,
		Milk:        100,
		Cash:        140,
		Name:        "Cappuchino",
	}
	coffeTypes["latte"] = models.Coffe{
		Cups:        1,
		CoffeeBeans: 20,
		Water:       300,
		Milk:        76,
		Cash:        110,
		Name:        "Latte",
	}
	coffeTypes["espresso"] = models.Coffe{
		Cups:        1,
		CoffeeBeans: 16,
		Water:       250,
		Milk:        0,
		Cash:        60,
		Name:        "Espresso",
	}

	m.CoffeTypes = coffeTypes

}
//func (m *Machine) Fill(plusWater, plusMilk, plusBeans, plusCups int) {
//	m.Water += plusWater
//	m.Milk += plusMilk
//	m.CoffeeBeans += plusBeans
//	m.Cups += plusCups
//}
//func (m *Machine) Buy(minusWater, minusMilk, minusBeans, minusCups, plusMoney int) {
//	m.Water -= minusWater
//	m.Milk -= minusMilk
//	m.CoffeeBeans -= minusBeans
//	m.Cups -= minusCups
//	m.Money += plusMoney
//}
