package main

import (
	// "github.com/shopspring/decimal"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")
	myWindow.Resize(fyne.NewSize(400, 300))
	
	title := widget.NewLabelWithStyle("Compound Interest Calculator", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	status := widget.NewLabelWithStyle("Awaiting inputs", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})

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

	result := container.NewGridWithRows(
		2,
		container.NewGridWithColumns(
			2,
			widget.NewLabelWithStyle("Total returns", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("N/A", fyne.TextAlignTrailing, fyne.TextStyle{Bold: false}),
		),
		container.NewGridWithColumns(
			2,
			widget.NewLabelWithStyle("Total interest", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			widget.NewLabelWithStyle("N/A", fyne.TextAlignTrailing, fyne.TextStyle{Bold: false}),
		),
	)
	
	content := container.NewVBox(
		title,
		principalEntry,
		rateEntry,
		container.NewGridWithColumns(
			2,
			timesCompoundEntry,
			yearsEntry,
		),
		status,
		widget.NewButton("Calculate", func() {
			log.Println("Content was:", principalEntry.Text)
		}),
		result,
		layout.NewSpacer(),
		widget.NewRichTextFromMarkdown("Created by [theluqmn](https://theluqmn.github.io) using Go [Fyne](https://fyne.io). [Source code](https://github.com/theluqmn/compound-interest-calculator)"),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	quit()
}

func quit() {
	log.Println("Exiting program...")
}