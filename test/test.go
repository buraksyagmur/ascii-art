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
	firstinput  []string
	secondinput []string
	thirdinput  []string
	clr         = strings.Split(os.Args[3], "=")
	// s   = SelectedColor(clr[1])
)

func main() {
	tempmap := make(map[int][]string)
	spacecounter := 0
	secondcounter := 0
	thirdcounter := 0
	_ = secondcounter
	_ = thirdcounter
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
			}
			for k := ourNumber5; k < len(os.Args[1]); k++ {
				secondinput = (asciiart.AsciiArt(letters[k]))
				for j, line := range secondinput {
					tempmap[j] = append(tempmap[j], line)
				}
			}
			for y := 0; y < ourNumber5; y++ {
				if letters[y] == " " {
					spacecounter++
				}
			}
			ourNumber5 += spacecounter
			for k := 0; k < 8; k++ {
				for m := 0; m < ourNumber5; m++ {
					printcolor(tempmap[k][m], colormap[clr[1]])
					// fmt.Print(tempmap[k][m])
				}
				for j := ourNumber5; j < (len(os.Args[1])); j++ {
					fmt.Print(tempmap[k][j])
				}

				fmt.Println()
			}

		} else if os.Args[4] == "last" {

			for i := 0; i < (leninput - ourNumber5); i++ {
				firstinput = (asciiart.AsciiArt(letters[i]))
				for j, line := range firstinput {
					tempmap[j] = append(tempmap[j], line)
				}
			}
			for k := (leninput - ourNumber5); k < len(os.Args[1]); k++ {
				secondinput = (asciiart.AsciiArt(letters[k]))
				for j, line := range secondinput {
					tempmap[j] = append(tempmap[j], line)
				}
			}
			for y := (leninput - ourNumber5); y < len(os.Args[1]); y++ {
				if letters[y] == " " {
					spacecounter++
				}
			}
			ourNumber5 += spacecounter
			for k := 0; k < 8; k++ {
				for m := 0; m < (leninput - ourNumber5); m++ {
					fmt.Print(tempmap[k][m])
					// fmt.Print(tempmap[k][m])
				}
				for j := (leninput - ourNumber5); j < (len(os.Args[1])); j++ {
					printcolor(tempmap[k][j], colormap[clr[1]])
				}

				fmt.Println()
			}
		} else if os.Args[4] == "word" {
			var firstinp string
			var secondinp string
			sliceinput := strings.Split(os.Args[1], " ")
			for i := 0; i < (ourNumber5 - 1); i++ {
				firstinp += sliceinput[i] + " "
			}
			colorword := sliceinput[ourNumber5-1] + " "
			_ = colorword
			for t := ourNumber5; t < len(sliceinput); t++ {
				secondinp += sliceinput[t] + " "
			}
			

			for i := 0; i < (len(os.Args[1])); i++ {
				firstinput = (asciiart.AsciiArt(letters[i]))
				for j, line := range firstinput {
					tempmap[j] = append(tempmap[j], line)
				}
			}
			
			
			
			for k := 0; k < 8; k++ {
				for m := 0; m < len(firstinp); m++ {
					 fmt.Print(tempmap[k][m])
					// fmt.Print(tempmap[k][m])
				}
				for j := len(firstinp); j < (len(os.Args[1]) - len(secondinp)); j++ {
					printcolor(tempmap[k][j], colormap[clr[1]])
				}
				for p := (len(os.Args[1]) - len(secondinp)); p < len(os.Args[1]); p++ {
					fmt.Print(tempmap[k][p])
				}

				fmt.Println()
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
					firstinput = (asciiart.AsciiArt(letters[i]))
					for j, line := range firstinput {
						tempmap[j] = append(tempmap[j], line)
					}
				}

				for k := ourNumber4 - 1; k < ourNumber6; k++ {
					secondinput = (asciiart.AsciiArt(letters[k]))
					for j, line := range secondinput {
						tempmap[j] = append(tempmap[j], line)
					}
				}

				for m := ourNumber6; m < (len(os.Args[1])); m++ {
					thirdinput = (asciiart.AsciiArt(letters[m]))
					for j, line := range thirdinput {
						tempmap[j] = append(tempmap[j], line)
					}
				}
				for y := ourNumber4 - 1; y <= ourNumber6+1; y++ {
					if letters[y] == " " {
						spacecounter++
					}
				}
				ourNumber6 += spacecounter
				for k := 0; k < 8; k++ {
					for m := 0; m < ourNumber4-1; m++ {
						fmt.Print(tempmap[k][m])
						// fmt.Print(tempmap[k][m])
					}
					for j := ourNumber4 - 1; j < (ourNumber6); j++ {
						printcolor(tempmap[k][j], colormap[clr[1]])
					}
					for p := ourNumber6; p < len(os.Args[1]); p++ {
						fmt.Print(tempmap[k][p])
					}

					fmt.Println()
				}

			}
		} else if os.Args[5] == "and" {
			if ourNumber4 > ourNumber6 {
				ourNumber4, ourNumber6 = ourNumber6, ourNumber4
			}
			for i := 0; i < (len(os.Args[1])); i++ {
				firstinput = (asciiart.AsciiArt(letters[i]))
				for j, line := range firstinput {
					tempmap[j] = append(tempmap[j], line)
				}
			}

			for q := 0; q < ourNumber4; q++ {
				if letters[q] == " " {
					spacecounter++
				}
			}

			if spacecounter > 0 {
				if letters[ourNumber4] == " " {
					spacecounter++
				}
			}

			ourNumber4 += spacecounter

			if spacecounter > 0 {
				if letters[ourNumber4-1] == " " {
					spacecounter++
					thirdcounter++
				}
			}
			ourNumber4 += thirdcounter

			for u := ourNumber4; u <= ourNumber6+spacecounter-1; u++ {
				if letters[u] == " " {
					secondcounter++
				}
			}
			if secondcounter > 0 {
				if letters[ourNumber6] == " " {
					secondcounter++
				}
			}

			ourNumber6 += spacecounter + secondcounter
			if secondcounter > 0 {
				if letters[ourNumber6-1] == " " {
					thirdcounter++
				}
				ourNumber6 += thirdcounter
			}

			for k := 0; k < 8; k++ {
				for m := 0; m < ourNumber4-1; m++ {
					fmt.Print(tempmap[k][m])
				}

				printcolor(tempmap[k][ourNumber4-1], colormap[clr[1]])

				for p := ourNumber4; p < ourNumber6-1; p++ {
					fmt.Print(tempmap[k][p])
				}
				printcolor(tempmap[k][ourNumber6-1], colormap[clr[1]])
				for i := ourNumber6; i < len(os.Args[1]); i++ {
					fmt.Print(tempmap[k][i])
				}
				fmt.Println()
			}

		}

	} else if len(os.Args) == 5 {
		letters := strings.Split(os.Args[1], "")
		ourNumber, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		for i := 0; i < (len(os.Args[1])); i++ {
			firstinput = (asciiart.AsciiArt(letters[i]))
			for j, line := range firstinput {
				tempmap[j] = append(tempmap[j], line)
			}
		}
		for q := 0; q < ourNumber; q++ {
			if letters[q] == " " {
				spacecounter++
			}
		}
		if spacecounter > 0 {
			if letters[ourNumber] == " " {
				spacecounter++
			}
		}
		ourNumber += spacecounter
		if spacecounter > 0 {
			if letters[ourNumber-1] == " " {
				secondcounter++
			}
		}
		ourNumber += secondcounter
		for k := 0; k < 8; k++ {
			for m := 0; m < ourNumber-1; m++ {
				fmt.Print(tempmap[k][m])
			}

			printcolor(tempmap[k][ourNumber-1], colormap[clr[1]])

			for i := ourNumber; i < len(os.Args[1]); i++ {
				fmt.Print(tempmap[k][i])
			}
			fmt.Println()
		}

	}
}
