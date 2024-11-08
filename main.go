package main

import (
	"fmt"
	"log"
	"strconv"
	"errors"

	"github.com/shopspring/decimal"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Compound Interest Calculator")
	myWindow.Resize(fyne.NewSize(500, 400))
	
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

	// Create result labels that we can reference later
	totalReturnsLabel := widget.NewLabelWithStyle("N/A", fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})
	totalInterestLabel := widget.NewLabelWithStyle("N/A", fyne.TextAlignTrailing, fyne.TextStyle{Bold: false})

	result := container.NewGridWithRows(
		2,
		container.NewGridWithColumns(
			2,
			widget.NewLabelWithStyle("Total returns", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			totalReturnsLabel,
		),
		container.NewGridWithColumns(
			2,
			widget.NewLabelWithStyle("Total interest", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			totalInterestLabel,
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
			// Validate inputs and type conversions
			principal, err := strconv.ParseFloat(principalEntry.Text, 64)
			if err != nil {
				status.SetText("Invalid principal amount")
				return
			}
			rate, err := strconv.ParseFloat(rateEntry.Text, 64)
			if err != nil {
				status.SetText("Invalid rate of interest")
				return
			}
			timesCompound, err := strconv.ParseInt(timesCompoundEntry.Text, 10, 64)
			if err != nil {
				status.SetText("Invalid times compounded")
				return
			}
			years, err := strconv.ParseInt(yearsEntry.Text, 10, 64)
			if err != nil {
				status.SetText("Invalid years")
				return
			}

			total, interest, err := CompoundInterest(principal, rate, int(timesCompound), int(years))
			if err != nil {
				status.SetText(err.Error())
				return
			}
			totalReturnsLabel.SetText(fmt.Sprintf("%.2f", total))
			totalInterestLabel.SetText(fmt.Sprintf("%.2f", interest))
			status.SetText("Calculated")

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

// Check out the repo at https://github.com/theluqmn/go4finance

// Calculates the total returns of a compound interest investment, and the total interest earned.
// It uses the shopspring/decimal library to handle floating-point arithmetic accurately, making it suitable for financial applications.
//
// Parameters:
// 	- principal: The initial investment amount.
// 	- rate: The annual interest rate as a decimal.
// 	- timesCompounded: The number of times the interest is compounded per year.
// 	- years: The number of years the investment is held.
//
// Returns:
// 	- total: The total amount of money after the investment period.
// 	- interest: The total interest earned.
// 	- err: An error if any of the inputs are invalid.
//
// Made by https://theluqmn.github.io
func CompoundInterest(principal float64, rate float64, timesCompounded int, years int) (total float64, interest float64, err error) {
    // Convert to decimals
	P := decimal.NewFromFloat(principal)
	r := decimal.NewFromFloat(rate)
	n := decimal.NewFromInt(int64(timesCompounded))
	t := decimal.NewFromInt(int64(years))
	one := decimal.NewFromInt(1)

	// Validate inputs
	if principal <= 0 {
		return 0, 0, errors.New("principal must be greater than 0")
	}
	if rate <= 0 {
		return 0, 0, errors.New("rate must be greater than 0")
	}
	if timesCompounded <= 0 {
		return 0, 0, errors.New("timesCompounded must be greater than 0")
	}
	if years <= 0 {
		return 0, 0, errors.New("years must be greater than 0")
	}

	// Calculate compound interest
	CalcExponent := n.Mul(t)
	CalcBracket := one.Add(r.Div(n))
	CalcTotal := P.Mul(CalcBracket.Pow(CalcExponent))
	CalcInterest := CalcTotal.Sub(P)

	return CalcTotal.InexactFloat64(), CalcInterest.InexactFloat64(), nil
}