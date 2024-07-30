// Given the daily rate calculates the monthly pay
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/shopspring/decimal"
)

type Item struct {
	quantity    decimal.Decimal
	cost        decimal.Decimal
	description string
	subtotal    decimal.Decimal
}

func outputText(item Item) {
	fmt.Println("Description, quantity, rate, total")
	fmt.Printf("%s, %s, %s, %s\n", item.description, item.quantity.StringFixed(0), item.cost.StringFixedBank(2), item.subtotal.StringFixedBank(2))
}

func outputHtml(item Item) {
	fmt.Printf("<td>%s</td><td>%s</td><td>%s</td><td>%s</td>\n", item.description, item.quantity.StringFixed(0), item.cost.StringFixedBank(2), item.subtotal.StringFixedBank(2))
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

	item.description = args[0]

	item.cost, err = decimal.NewFromString(args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing the daily rate")
		return
	}

	item.quantity = decimal.NewFromInt(int64(workDays))
	item.subtotal = item.quantity.Mul(item.cost)

	if htmlOutput {
		outputHtml(item)
		return
	}

	outputText(item)
}
