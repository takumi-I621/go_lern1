package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	fmt.Println("--- START ---")
	for {
		line, err = reader.Read()
		if err != nil {
			fmt.Println("--- END ---")
			break
		} else if err != nil {
			fmt.Println("--- ERROR END ---")
			fmt.Println(err)
			break
		}
		fmt.Println(line)

	}

}

/*
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
*/
