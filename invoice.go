// Given the daily rate calculates the monthly pay
package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/shopspring/decimal"
)

const templateFile = "./templates/simple.tmpl"

type Invoice struct {
	Date   string
	Serial string
	Item   Item
}

type Item struct {
	Quantity    decimal.Decimal
	Cost        decimal.Decimal
	Description string
	Subtotal    decimal.Decimal
}

func outputText(item Item) {
	fmt.Println("Description, quantity, rate, total")
	fmt.Printf("%s, %s, %s, %s\n",
		item.Description,
		item.Quantity.StringFixed(0),
		item.Cost.StringFixedBank(2),
		item.Subtotal.StringFixedBank(2))
}

func outputHtml(invoice Invoice) {
	tmpl, err := template.New(templateFile).ParseFiles(templateFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed opening the HTML template", err)
		return
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "simple.tmpl", invoice)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed reading the HTML template", err)
	}
}

func main() {
	var daysOutOfOffice int
	var htmlOutput bool

	flag.IntVar(&daysOutOfOffice, "out-of-office", 0, "Days not worked this month")
	flag.BoolVar(&htmlOutput, "html", false, "Output the invoice in HTML")

	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments")
		return
	}

	// Use Decimal for money
	var item Item
	var err error

	var workDays = weekdays(time.Now()) - daysOutOfOffice

	if workDays < 0 {
		workDays = 0
	}

	item.Description = args[0]

	item.Cost, err = decimal.NewFromString(args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing the daily rate")
		return
	}

	item.Quantity = decimal.NewFromInt(int64(workDays))
	item.Subtotal = item.Quantity.Mul(item.Cost)

	if htmlOutput {
		var invoice = Invoice{
			Item:   item,
			Date:   "2024-Jun-29",
			Serial: "Jun2024",
		}
		outputHtml(invoice)
		return
	}

	outputText(item)
}
