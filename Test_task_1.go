package main

import (
	// "fmt"
	// "math/rand"

	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// var number uint64 = uint64(rand.Intn(1000))
	// fmt.Println("Hello Ivan and Valentin")
	// fmt.Println("Random number = ", number)
	// var count int = 100
	// for i := 0; i < count; i++ {
	// 	number *= 5
	// }
	// fmt.Println("New number = ", number)
	// fmt.Println("Random number % 7 = ", number%7)
	a := app.New()
	w := a.NewWindow("Golang test task 1")
	w.Resize(fyne.NewSize(400, 400))
	label1 := widget.NewLabel("Enter random integer number")
	entry := widget.NewEntry()
	label2 := widget.NewLabel("")
	answer := widget.NewLabel("")
	btn := widget.NewButton("OK", func() {
		number, err := strconv.ParseUint(entry.Text, 10, 64)
		if err != nil {
			answer.SetText("Input error")
		} else {
			for i := 0; i < 10; i++ {
				number *= 5
			}
			label2.SetText(fmt.Sprintf("New number = %d", number))
			number %= 7
			answer.SetText(fmt.Sprintf("Number %% 7 = %d", number))
		}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry,
		btn,
		label2,
		answer,
	))

	w.ShowAndRun()
}
