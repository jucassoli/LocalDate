package main

import (
	"fmt"
	"time"

	"github.com/jucassoli/localdate"
)

func main() {
	now := time.Now()

	res := localdate.MakeFormatter(&now)
	res.AppendToken(localdate.LongMonthToken).AppendToken(localdate.DayToken).AppendToken(localdate.LongYearToken)
	res.GenerateWithAllBetween("-")
	result := res.Convert()

	fmt.Printf("-- %v  -- > %T", result, result)

}
