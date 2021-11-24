package main

import (
	"asciiart"
	"bytes"
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
	b   bytes.Buffer
	abc []string
	a   []string
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
				// )
				(asciiart.AsciiArt2(letters[i]))
			}
			for k := ourNumber5; k < len(os.Args[1]); k++ {
				abc = (asciiart.AsciiArt(letters[k]))
				a = append(abc, abc...)
			}
			fmt.Println(a[20])

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
