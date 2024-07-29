// Given the daily rate calculates the monthly pay
package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	// Use Decimal for money
	var hourlyRate decimal.Decimal

	hourlyRate, _ = decimal.NewFromString("10.07")

	var workDays = weekdays(time.Now())
	var cost = decimal.NewFromInt(int64(workDays)).Mul(hourlyRate)

	fmt.Printf("workdays %d\ncost %s\n", workDays, cost.StringFixedBank(2))
}
