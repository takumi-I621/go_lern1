package main

import (
	"fmt"
	"math"
)

func main() {
	//体重と身長の変数定義

	BmiCalculation()

}

func BmiCalculation() {
	var Weight float64
	var Height float64
	var Bmi float64

	fmt.Println("身長(m)を入力してください")
	fmt.Scan(&Height)

	fmt.Println("体重(kg)を入力してください")
	fmt.Scan(&Weight)

	Bmi = Weight / (Height * Height)
	BmiResult := math.Round(Bmi*100) / 100

	BmiStr := "BMI ："

	fmt.Println(BmiStr, BmiResult)

}
