package main

import "fmt"
import "strconv"

var money, price, beerNum, capsule, bottles, tempBeerNum int

func main() {
	money = 10
	price = 2
	beerNum = 0 //啤酒数
	capsule = 0 //瓶盖数
	bottles = 0 //空瓶数
	tempBeerNum = 0

	beerNum = money / price
	capsule = beerNum
	bottles = beerNum

	fmt.Println(beerNum, capsule, bottles, strconv.Itoa(tempBeerNum))

	cisend := true
	bisend := true
	for {
		cisend = CalCapsule()
		bisend = CalBottles()

		if cisend == false && bisend == false {
			break
		} else {
			continue
		}
	}

	fmt.Println("消费啤酒总数：", strconv.Itoa(beerNum))
	fmt.Println("剩余瓶盖总数：", strconv.Itoa(capsule))
	fmt.Println("剩余空瓶总数：", strconv.Itoa(bottles))
}

func CalCapsule() bool {
	if capsule >= 4 {
		tempBeerNum = capsule % 4
		if tempBeerNum != 0 && tempBeerNum > 0 {

			beerNum += tempBeerNum
			bottles += tempBeerNum
			capsule += tempBeerNum
			capsule = capsule - tempBeerNum*4

			return true
		} else {
			tempBeerNum = capsule / 4
			beerNum += tempBeerNum
			bottles += tempBeerNum
			capsule += tempBeerNum
			capsule = capsule - tempBeerNum*4

			return true
		}

	}

	return false
}

func CalBottles() bool {
	if bottles >= 2 {
		tempBeerNum = bottles % 2
		if tempBeerNum != 0 && tempBeerNum > 0 {
			beerNum += tempBeerNum
			capsule += tempBeerNum
			bottles += tempBeerNum
			bottles = bottles - tempBeerNum*2
			return true
		} else {
			tempBeerNum = bottles / 2
			beerNum += tempBeerNum
			capsule += tempBeerNum
			bottles += tempBeerNum
			bottles = bottles - tempBeerNum*2
			return true

		}
	}

	return false
}
