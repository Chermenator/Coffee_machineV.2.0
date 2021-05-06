package main

import (

	"github.com/Chermenator/coffe_machin/internal"


)

func main() {
	cm := internal.NewMachine(540, 400, 120, 9, 390)
	cm.InitCoffeeTypes()
	internal.Start(cm)
}