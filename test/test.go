package main

import (
	"asciiart"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leninput := len(os.Args[1])
	if len(os.Args) == 4 {
		asciiart.AsciiArt(os.Args[1])
	} else if len(os.Args) == 6 {
		letters := strings.Split(os.Args[1], "")
		words := strings.Split(os.Args[1], " ")
		ourNumber, err := strconv.Atoi(os.Args[5])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if os.Args[4] == "first" {
			for i := 0; i < ourNumber; i++ {
				asciiart.AsciiArt2(letters[i])
			}
			for k := ourNumber; k < len(os.Args[1]); k++ {
				asciiart.AsciiArt(letters[k])
			}
		} else if os.Args[4] == "last" {
			for i := 0; i < (leninput - ourNumber); i++ {
				asciiart.AsciiArt(letters[i])
			}
			for k := (leninput - ourNumber); k < len(os.Args[1]); k++ {
				asciiart.AsciiArt2(letters[k])
			}
		} else if os.Args[4] == "word" {
			if os.Args[5] != "1" {
				for i := 0; i < (ourNumber - 1); i++ {
					asciiart.AsciiArt(words[i])
				}
				asciiart.AsciiArt2(words[ourNumber-1])
				for k := ourNumber - 1; k < len(words); k++ {
					asciiart.AsciiArt(words[k])
				}
			}
		}
	}
}
