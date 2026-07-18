package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go Launch: %s\n", t)

	n := time.Now()
	fmt.Printf("Current Time: %s\n", n)
	fmt.Printf("The Object Type: %T\n", n)

	fmt.Printf("Today: %s", n.Format(time.ANSIC)+"\n")

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow: %s\n", tomorrow.Format(time.ANSIC))

	format := "Mon 2006-02-01"
	fmt.Printf("Format: %s\n", tomorrow.Format(format))
}
