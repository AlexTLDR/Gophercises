package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string
	var zero = placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	var one = placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	var two = placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	var three = placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	var four = placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	var five = placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	var six = placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	var seven = placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	var eight = placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	var nine = placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	var colon = placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	screen.Clear()
	for {
		screen.MoveTopLeft()
		t := time.Now()
		hour, minute, second := t.Hour(), t.Minute(), t.Second()
		digits := [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}
		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[minute/10], digits[minute%10],
			colon,
			digits[second/10], digits[second%10],
		}

		for space := range clock[0] {
			for i, digit := range clock {
				blink := clock[i][space]
				if digit == colon && second%2 != 0 {
					blink = "   "
				}
				fmt.Print(blink, "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
