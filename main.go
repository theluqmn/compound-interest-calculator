package main

import (
	// "github.com/shopspring/decimal"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	title := widget.NewLabelWithStyle("Compound Interest Calculator", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	principalEntry := widget.NewEntry()
	principalEntry.SetPlaceHolder("Principal amount")
	principalEntry.Validator = func(s string) error {
		for _, r := range s {
			if (r < '0' || r > '9') && r != '.' {
				return fmt.Errorf("only numbers and decimal point allowed")
			}
		}
		return nil
	}

	rateEntry := widget.NewEntry()
	rateEntry.SetPlaceHolder("Rate of interest")
	rateEntry.Validator = func(s string) error {
		for _, r := range s {
			if (r < '0' || r > '9') && r != '.' {
				return fmt.Errorf("only numbers and decimal point allowed")
			}
		}
		return nil
	}

	timesCompoundEntry := widget.NewEntry()
	timesCompoundEntry.SetPlaceHolder("Times compounded")
	timesCompoundEntry.Validator = func(s string) error {
		for _, r := range s {
			if (r < '0' || r > '9') && r != '.' {
				return fmt.Errorf("only numbers and decimal point allowed")
			}
		}
		return nil
	}

	yearsEntry := widget.NewEntry()
	yearsEntry.SetPlaceHolder("Years")
	yearsEntry.Validator = func(s string) error {
		for _, r := range s {
			if (r < '0' || r > '9') && r != '.' {
				return fmt.Errorf("only numbers and decimal point allowed")
			}
		}
		return nil
	}
	
	content := container.NewVBox(
		title,
		principalEntry,
		rateEntry,
		container.NewGridWithColumns(
			2,
			timesCompoundEntry,
			yearsEntry,
		),
		widget.NewButton("Calculate", func() {
			log.Println("Content was:", principalEntry.Text)
		}),
		
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	quit()
}

func quit() {
	log.Println("Exiting program...")
}