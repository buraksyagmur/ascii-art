package main

import (
	"asciiart"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colormap = map[string]int{
	"black":    0,
	"orange":   208,
	"red":      9,
	"green":    10,
	"yellow":   11,
	"blue":     12,
	"purple":   5,
	"turqoise": 14,
	"white":    15,
	"pink":     198,
}

func printcolor(s string, color int) {
	ansi := "\033[38;5;%dm%s\033[0m"
	fmt.Printf(ansi, color, s)
}

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

	firstsecondinput []string
	firstsecondslice []string
	firsttempslice   []string

	firstinput []string
	firstslice []string

	clr = strings.Split(os.Args[3], "=")
	// s   = SelectedColor(clr[1])
)

func main() {
	tempmap := make(map[int][]string)
	spacecounter := 0
	leninput := len(os.Args[1])
	if len(os.Args) == 4 {
		asciiart.AsciiArt(os.Args[1])
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
				firstinput = (asciiart.AsciiArt(letters[i]))
				for j, line := range firstinput {
					tempmap[j] = append(tempmap[j], line)
				}
				firstslice = append(firstslice, firstinput...)
			}
			firsttempslice = firstslice

			for k := ourNumber5; k < len(os.Args[1]); k++ {
				firstsecondinput = (asciiart.AsciiArt(letters[k]))

				for j, line := range firstsecondinput {
					tempmap[j] = append(tempmap[j], line)
				}
				firstsecondslice = append(firstsecondslice, firstsecondinput...)
			}
			firsttempslice = append(firsttempslice, firstsecondslice...)

			// s := printresult(firsttempslice, ourNumber5)
			// printcolor(s, colormap[clr[1]])
			// fmt.Println(red(printresult(firstslice, ourNumber5)) + printresult(firstsecondslice, (len(os.Args[1])-ourNumber5)))

			// fmt.Print(printresult(firsttempslice, (len(os.Args[1]) - ourNumber5)))

			for k := 0; k < 8; k++ {
				for m := 0; m < ourNumber5; m++ {
					printcolor(tempmap[k][m], colormap[clr[1]])
					// fmt.Print(tempmap[k][m])
				}
				for j := ourNumber5; j < (len(os.Args) - 1); j++ {
					fmt.Print(tempmap[k][j])
				}

				fmt.Println()
			}
		} else if os.Args[4] == "last" {
			for i := 0; i < (leninput - ourNumber5); i++ {
				asciiart.AsciiArt(letters[i])
			}
			for k := (leninput - ourNumber5); k < len(os.Args[1]); k++ {
				asciiart.AsciiArt(letters[k])
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
				asciiart.AsciiArt(os.Args[1])
			} else {
				for i := 0; i < (ourNumber4 - 1); i++ {
					asciiart.AsciiArt(letters[i])
				}
				for k := ourNumber4 - 1; k < ourNumber6; k++ {
					asciiart.AsciiArt(letters[k])
					if letters[k] == " " {
						spacecounter++
					}

				}
				if spacecounter > 0 {
					for i := 0; i < spacecounter; i++ {
						asciiart.AsciiArt(letters[ourNumber6+i])
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
				asciiart.AsciiArt(os.Args[1])
			} else {
				for i := 0; i < (ourNumber4 - 1); i++ {
					asciiart.AsciiArt(letters[i])
				}
				asciiart.AsciiArt(letters[ourNumber4-1])

				for m := ourNumber4; m < ourNumber6-1; m++ {
					asciiart.AsciiArt(letters[m])
				}
				asciiart.AsciiArt(letters[ourNumber6-1])
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
		asciiart.AsciiArt(letters[ourNumber-1])
		for m := ourNumber; m < len(os.Args); m++ {
			asciiart.AsciiArt(letters[m])
		}
	}
}

func printresult(a []string, counter int) string {
	str := ""
	str2 := ""
	str3 := ""
	str4 := ""
	str5 := ""
	str6 := ""
	str7 := ""
	str8 := ""
	for j := 0; j <= ((counter * 8) - 1); j += 8 {
		str += a[j]
	}

	for j := 1; j <= ((counter * 8) - 1); j += 8 {
		str2 += a[j]
	}

	for j := 2; j <= ((counter * 8) - 1); j += 8 {
		str3 += a[j]
	}

	for j := 3; j <= ((counter * 8) - 1); j += 8 {
		str4 += a[j]
	}

	for j := 4; j <= ((counter * 8) - 1); j += 8 {
		str5 += a[j]
	}

	for j := 5; j <= ((counter * 8) - 1); j += 8 {
		str6 += a[j]
	}

	for j := 6; j <= ((counter * 8) - 1); j += 8 {
		str7 += a[j]
	}

	for j := 7; j <= ((counter * 8) - 1); j += 8 {
		str8 += a[j]
	}

	allstr := (str + "\n" + str2 + "\n" + str3 + "\n" + str4 + "\n" + str5 + "\n" + str6 + "\n" + str7 + "\n" + str8 + "\n")

	return allstr
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
