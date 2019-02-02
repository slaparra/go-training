package main

import (
	"fmt"
	"time"
)

func main() {
	anExample()
	anotherExample()
}

func anExample() {
	today := time.Now()
	fmt.Println("Today is : ", today.Format("2006-01-02 03:04:00pm"))

	// travel back 1 day
	oneDay := time.Hour * -24 // minus 1 day

	yesterday := today.Add(oneDay) // minus 1 day
	fmt.Println("Yesterday is : ", yesterday.Format("2006-01-02 03:04:00pm"))

	// for easier calculation, use AddDate() function instead

	sevenDaysAgo := today.AddDate(0, 0, -7) // minus 7 days
	fmt.Println("Seven days ago is : ", sevenDaysAgo.Format("2006-01-02 03:04:00pm"))

	oneYearAgo := today.AddDate(-1, 0, 0) // minus one year
	fmt.Println("One year ago is : ", oneYearAgo.Format("2006-01-02 03:04:00pm"))

	// Today is :  2019-02-02 05:05:00pm
	// Yesterday is :  2019-02-01 05:05:00pm
	// Seven days ago is :  2019-01-26 05:05:00pm
	// One year ago is :  2018-02-02 05:05:00pm
}

func anotherExample() {
	now := time.Now()
	fmt.Println("now:", now)

	then := now.AddDate(0, -1, 0) //(y,m,d)
	fmt.Println("then:", then)

	//now: 2019-02-02 17:05:31.014166 +0100 CET m=+0.000583755
	//then: 2019-01-02 17:05:31.014166 +0100 CET
}
