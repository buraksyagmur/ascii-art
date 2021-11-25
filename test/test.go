package main

import (
	"asciiart"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lengthwithoutspace(s string) int {
	space := 0
	split := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		if split[i] == " " {
			space++
		}
	}
	return (len(s) - space)
}

func foundspace(s string) int {
	space := 0
	split := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		if split[i] == " " {
			space++
		}
	}
	return space
}

var (
	black    = Color("\033[1;30m%s\033[0m")
	red      = Color("\033[38;5;196m%s\033[0m")
	green    = Color("\033[1;32m%s\033[0m")
	yellow   = Color("\033[1;33m%s\033[0m")
	blue     = Color("\033[1;34m%s\033[0m")
	purple   = Color("\033[1;35m%s\033[0m")
	turqoise = Color("\033[1;36m%s\033[0m")
	white    = Color("\033[1;37m%s\033[0m")
	orange   = Color("\033[38;5;208m%s\033[0m")
	pink     = Color("\033[38;5;198m%s\033[0m")

	abc        []string
	abcd       []string
	counter1   int
	counter2   int
	firstinput []string
	firstslice []string

	clr = strings.Split(os.Args[3], "=")
	// s   = SelectedColor(clr[1])
)

func main() {
	spacecounter := 0
	leninput := len(os.Args[1])
	if len(os.Args) == 4 {
		asciiart.AsciiArt2(os.Args[1])
	} else if len(os.Args) == 6 {
		letters := strings.Split(os.Args[1], "")
		// words := strings.Split(os.Args[1], " ")
		ourNumber5, err := strconv.Atoi(os.Args[5])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if os.Args[4] == "first" {
			for i := 0; i < ourNumber5; i++ {
				firstinput = (asciiart.AsciiArt2(letters[i]))
				firstslice = append(firstslice, firstinput...)
				counter2++

			}

			for k := ourNumber5; k < len(os.Args[1]); k++ {
				abc = (asciiart.AsciiArt(letters[k]))
				abcd = append(abcd, abc...)
				counter1++

			}
			printresult(abcd, counter1)
			printcolorresult(firstslice, counter2)
			// fmt.Println(b.String())
			// for m := 0; m < 16; m++ {
			// 	fmt.Println(abcde[m])
			// }

		} else if os.Args[4] == "last" {
			for i := 0; i < (leninput - ourNumber5); i++ {
				asciiart.AsciiArt(letters[i])
			}
			for k := (leninput - ourNumber5); k < len(os.Args[1]); k++ {
				asciiart.AsciiArt2(letters[k])
			}
		}
	} else if len(os.Args) == 7 {
		letters := strings.Split(os.Args[1], "")
		ourNumber4, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		ourNumber6, err := strconv.Atoi(os.Args[6])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if os.Args[5] == "-" {
			if ourNumber6 == lengthwithoutspace(os.Args[1]) && ourNumber4 == 1 {
				asciiart.AsciiArt2(os.Args[1])
			} else {
				for i := 0; i < (ourNumber4 - 1); i++ {
					asciiart.AsciiArt(letters[i])
				}
				for k := ourNumber4 - 1; k < ourNumber6; k++ {
					asciiart.AsciiArt2(letters[k])
					if letters[k] == " " {
						spacecounter++
					}

				}
				if spacecounter > 0 {
					for i := 0; i < spacecounter; i++ {
						asciiart.AsciiArt2(letters[ourNumber6+i])
					}
					if spacecounter > 0 {
						for m := ourNumber6 + spacecounter; m < (len(os.Args[1])); m++ {
							asciiart.AsciiArt(letters[m])
						}
					}
				} else {
					for m := ourNumber6; m < (len(os.Args[1])); m++ {
						asciiart.AsciiArt(letters[m])
					}
				}
			}
		} else if os.Args[5] == "and" {
			if ourNumber6 == lengthwithoutspace(os.Args[1]) && ourNumber4 == 1 {
				asciiart.AsciiArt2(os.Args[1])
			} else {
				for i := 0; i < (ourNumber4 - 1); i++ {
					asciiart.AsciiArt(letters[i])
				}
				asciiart.AsciiArt2(letters[ourNumber4-1])

				for m := ourNumber4; m < ourNumber6-1; m++ {
					asciiart.AsciiArt(letters[m])
				}
				asciiart.AsciiArt2(letters[ourNumber6-1])
				for k := ourNumber6; k < len(os.Args[1]); k++ {
					asciiart.AsciiArt(letters[k])
				}
			}
		}
	} else if len(os.Args) == 5 {
		letters := strings.Split(os.Args[1], "")
		ourNumber, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		for i := 0; i < (ourNumber - 1); i++ {
			asciiart.AsciiArt(letters[i])
		}
		asciiart.AsciiArt2(letters[ourNumber-1])
		for m := ourNumber; m < len(os.Args); m++ {
			asciiart.AsciiArt(letters[m])
		}
	}
}

func printresult(a []string, counter int) {
	for j := 0; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 1; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 2; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 3; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 4; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 5; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 6; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
	for j := 7; j <= ((counter * 8) - 1); j += 8 {
		fmt.Print(a[j])
	}
	fmt.Println(" ")
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func printcolorresult(a []string, counter int) {
	if clr[1] == "red" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(red(a[j]))
			fmt.Println(" ")
		}
	} else if clr[1] == "black" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(black(a[j]))
		}
	} else if clr[1] == "green" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(green(a[j]))
		}
	} else if clr[1] == "yellow" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(yellow(a[j]))
		}
	} else if clr[1] == "blue" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(blue(a[j]))
		}
	} else if clr[1] == "purple" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(purple(a[j]))
		}
	} else if clr[1] == "turqoise" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(turqoise(a[j]))
		}
	} else if clr[1] == "white" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(white(a[j]))
		}
	} else if clr[1] == "pink" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(pink(a[j]))
		}
	} else if clr[1] == "orange" {
		for j := 0; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 1; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 2; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 3; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 4; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 5; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 6; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
		fmt.Println(" ")
		for j := 7; j <= ((counter * 8) - 1); j += 8 {
			fmt.Print(orange(a[j]))
		}
	}
}
