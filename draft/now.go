package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
)

func main() {

	time.Now().Sub()
	time.Now() // 2013-11-18 17:51:49.123456789 Mon
	fmt.Println()
	now.BeginningOfMinute()        // 2013-11-18 17:51:00 Mon
	now.BeginningOfHour()          // 2013-11-18 17:00:00 Mon
	now.BeginningOfDay()           // 2013-11-18 00:00:00 Mon
	now.BeginningOfWeek()          // 2013-11-17 00:00:00 Sun
	now.BeginningOfMonth()         // 2013-11-01 00:00:00 Fri
	now.BeginningOfQuarter()       // 2013-10-01 00:00:00 Tue
	now.BeginningOfYear()          // 2013-01-01 00:00:00 Tue

	now.WeekStartDay = time.Monday // Set Monday as first day, default is Sunday
	now.BeginningOfWeek()          // 2013-11-18 00:00:00 Mon
	now.EndOfMinute()              // 2013-11-18 17:51:59.999999999 Mon
	now.EndOfHour()                // 2013-11-18 17:59:59.999999999 Mon
	now.EndOfDay()                 // 2013-11-18 23:59:59.999999999 Mon
	now.EndOfWeek()                // 2013-11-23 23:59:59.999999999 Sat
	now.EndOfMonth()               // 2013-11-30 23:59:59.999999999 Sat
	now.EndOfQuarter()             // 2013-12-31 23:59:59.999999999 Tue
	now.EndOfYear()
}
