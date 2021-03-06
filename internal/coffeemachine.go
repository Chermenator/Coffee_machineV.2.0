package internal

import (
	"fmt"
	"os"
)

func (m *Machine) Fill(plusWater, plusMilk, plusBeans, plusCups int) {
	m.Water += plusWater
	m.Milk += plusMilk
	m.CoffeeBeans += plusBeans
	m.Cups += plusCups
}
func (m *Machine) Buy(minusWater, minusMilk, minusBeans, minusCups, plusMoney int) {
	m.Water -= minusWater
	m.Milk -= minusMilk
	m.CoffeeBeans -= minusBeans
	m.Cups -= minusCups
	m.Money += plusMoney
}

func Start(cm *Machine) {
	count := 0
	moneyCount := 0
	cupCount := 0
	latCount := 0
	espCount := 0
LOOP:
	for {
		fmt.Println("\nInput command: buy, fill, take, exit")
		command := ""
		fmt.Fscan(os.Stdin, &command)

	SWIT:
		switch command {
		case "buy":
			inp := 0
			fmt.Println("What do you want to buy?\n1 - cuppuchino\n2 - latte\n3 - espresso\n4 - main menu")
			fmt.Fscan(os.Stdin, &inp)

			//coffeMass := [5]int{cm.Money,cm.CoffeeBeans,cm.Milk,cm.Cups,cm.Water}

			switch inp {
			case 1:
				if cm.Water < 200 {
					fmt.Println("Вам не хватает воды!!!")
					break SWIT
				}
				if cm.Milk < 100 {
					fmt.Println("Вам не хватает молока!!!")
					break SWIT
				}
				if cm.CoffeeBeans < 16 {
					fmt.Println("Вам не хватает кофе зёрен!!!")
					break SWIT
				}
				if cm.Cups == 0 {
					fmt.Println("Вам не хватает стаканчиков!!!")
					break SWIT
				}
				if cm.Money+140 > 9000 {
					fmt.Printf("Касса переполнена!!! Сейчас в ней %d", cm.Money)
					break SWIT
				}
				cm.Buy(200, 100, 16, 1, 140)
				count++
				moneyCount += 140
				cupCount++
				break
			case 2:
				if cm.Water < 300 {
					fmt.Println("Вам не хватает воды!!!")
					break SWIT
				}
				if cm.Milk < 76 {
					fmt.Println("Вам не хватает молока!!!")
					break SWIT
				}
				if cm.CoffeeBeans < 20 {
					fmt.Println("Вам не хватает кофе зёрен!!!")
					break SWIT
				}
				if cm.Cups == 0 {
					fmt.Println("Вам не хватает стаканчиков!!!")
					break SWIT
				}
				if cm.Money+110 > 9000 {
					fmt.Printf("Касса переполнена!!! Сейчас в ней %d", cm.Money)
					break SWIT
				}
				cm.Buy(300, 76, 20, 1, 110)
				count++
				moneyCount += 110
				latCount++
				break
			case 3:
				if cm.Water < 250 {
					fmt.Println("Вам не хватает воды!!!")
					break SWIT
				}
				if cm.CoffeeBeans < 16 {
					fmt.Println("Вам не хватает кофе зёрен!!!")
					break SWIT
				}
				if cm.Cups == 0 {
					fmt.Println("Вам не хватает стаканчиков!!!")
					break SWIT
				}
				if cm.Money+60 > 9000 {
					fmt.Printf("Касса переполнена!!! Сейчас в ней %d", cm.Money)
					break SWIT
				}
				cm.Buy(250, 0, 16, 1, 60)
				count++
				moneyCount += 60
				espCount++
				break
			case 4:
				break SWIT
			default:
				fmt.Println("Такого кофе нет в нашем меню.")
				break
			}
			fmt.Println("Ok. Coffee will be ready soon\n")
		case "fill":
			k := 0

			fmt.Println("\nHow much water should you add?")
			fmt.Fscan(os.Stdin, &k)
			cm.Fill(k, 0, 0, 0)
			if cm.Water > 5000 {
				cm.Water = 5000
			}

			fmt.Println("\nHow much milk should you add?")
			fmt.Fscan(os.Stdin, &k)
			cm.Fill(0, k, 0, 0)
			if cm.Milk > 1000 {
				cm.Milk = 1000
			}

			fmt.Println("\nHow many coffee beans should you add?")
			fmt.Fscan(os.Stdin, &k)
			cm.Fill(0, 0, k, 0)
			if cm.CoffeeBeans > 900 {
				cm.CoffeeBeans = 900
			}
			fmt.Println("\nHow many cups should you add?")
			fmt.Fscan(os.Stdin, &k)
			cm.Fill(0, 0, 0, k)
			if cm.Cups > 50 {
				cm.Cups = 50
			}
			break
		case "take":
			fmt.Printf("Выдано %d рублей", cm.Money)
			cm.Money = 0
			break
		case "stat":
			fmt.Println("В кофемашине:")
			fmt.Printf("%d мл воды\n%d мл молока\n%d грамм кофейных зёрен\n%d  чашек\n%d рублей\n\n", cm.Water, cm.Milk, cm.CoffeeBeans, cm.Cups, cm.Money)
			fmt.Printf("Всего продано %d напитков на %d рублей", count, moneyCount)
			fmt.Printf("\n%d-Капучино\n%d-Латте\n%d-Эспрессо\n\n", cupCount, latCount, espCount)
			break
		case "exit":
			break LOOP
		default:
			fmt.Println("Wrong command!!!")
			continue LOOP
		}
	}

}
