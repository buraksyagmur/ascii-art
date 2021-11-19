package asciiart

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	Black                = Color("\033[1;30m%s\033[0m")
	Red                  = Color("\033[1;31m%s\033[0m")
	Green                = Color("\033[1;32m%s\033[0m")
	Yellow               = Color("\033[1;33m%s\033[0m")
	Blue                 = Color("\033[1;34m%s\033[0m")
	Purple               = Color("\033[1;35m%s\033[0m")
	Turqoise             = Color("\033[1;36m%s\033[0m")
	White                = Color("\033[1;37m%s\033[0m")
	numberfornewline int = 0
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// checking for errors other than io.EOF
func checkErrNoEOF(e error) {
	if e != nil && e != io.EOF {
		log.Fatal(e)
	}
}

// scaning 8 lines starting from startLine
func scanChar(r io.Reader, startLine int) ([]string, error) {
	lineScanner := bufio.NewScanner(r)
	bigCharLines := []string{}
	curLine := 0
	linesAdded := 0
	for lineScanner.Scan() {
		curLine++ // coz the first line of txt is 1
		if curLine == startLine {
			// scan 8 lines
			for sc := 0; sc < 8; sc++ {
				bigCharLines = append(bigCharLines, lineScanner.Text())
				linesAdded++
				lineScanner.Scan() // advance to the next line
			}
		}
	}
	// fmt.Println(bigCharLines) // this will print stuff crumbled tgt in a single line
	return bigCharLines, io.EOF
}

func printBigChar(chMap *map[byte][]string, inpBSlice []byte) {
	for l := 0; l < 8; l++ {
		chLine := ""
		for ch := 0; ch < len(inpBSlice); ch++ {
			chLine += string((*chMap)[inpBSlice[ch]][l])
		}
		res1 := strings.Split(os.Args[3], "=")

		if res1[1] == "black" {
			fmt.Println(Black(chLine))
		} else if res1[1] == "red" {
			fmt.Println(Red(chLine))
		} else if res1[1] == "purple" {
			fmt.Println(Purple(chLine))
		} else if res1[1] == "blue" {
			fmt.Println(Blue(chLine))
		} else if res1[1] == "green" {
			fmt.Println(Green(chLine))
		} else if res1[1] == "yellow" {
			fmt.Println(Yellow(chLine))
		} else if res1[1] == "white" {
			fmt.Println(White(chLine))
		} else if res1[1] == "turqoise" {
			fmt.Println(Turqoise(chLine))
		}
	}
}

func AsciiArt(s string) {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Arg amount is not right")
	// }

	banner := os.Args[2]
	// read input str
	var inputStrSlices []string

	// fmt.Println("input: ", inputStr)
	inputrune := []rune(s)
	s1 := s
	// case when \n at the end of the input
	if len(os.Args[1]) == 4 {
		if inputrune[len(s)-2] == '\\' && inputrune[len(s)-1] == 'n' {
			if last := len(s1) - 1; last >= 0 && s1[last] == 'n' {
				s1 = s1[:last] // remove the last n
			}
			s2 := s1
			if last := len(s2) - 1; last >= 0 && s2[last] == '\\' {
				s2 = s2[:last] // remove the 2nd last \
			}
			// convert s2 (string) to slice of string (with only itself in the slice)
			inputStrSlices = strings.Split(s2, " ")
			numberfornewline++ // increment the counter
		} else {
			inputStrSlices = strings.Split(s, "\\n") // normal case
		}
	} else {
		inputStrSlices = strings.Split(s, "\\n")
	}
	for _, inputSlice := range inputStrSlices {
		// process the str
		inputBSlice := []byte(inputSlice)
		// fmt.Println(inputBSlice)
		charMap := make(map[byte][]string)

		for inp := 0; inp < len(inputBSlice); inp++ {
			// find the corresponding line num
			startLine := (int(inputBSlice[inp])-32)*9 + 2
			// fmt.Println("startLine: ", startLine)
			// scan the reqired lines in from the txt file
			fread, err := os.Open(banner + ".txt")
			checkErr(err)
			defer fread.Close()

			scanner := bufio.NewScanner(fread)

			scanner.Split(bufio.ScanBytes)

			// scan 8 lines starting from startLine from the txt file
			bigChar, err := scanChar(fread, startLine)
			checkErrNoEOF(err)
			if len(bigChar) != 8 {
				fmt.Println("Wrong number of lines read")
			}

			charMap[inputBSlice[inp]] = bigChar
		}
		// }
		printBigChar2(&charMap, inputBSlice)
	}
}

func AsciiArt2(s string) {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Arg amount is not right")
	// }
	banner := os.Args[2]
	// read input str
	var inputStrSlices []string
	// inputStr := os.Args[1]
	// // fmt.Println("input: ", inputStr)

	// case when \n at the end of the input

	inputStrSlices = strings.Split(s, "\\n") // normal case

	for _, inputSlice := range inputStrSlices {
		// process the str
		inputBSlice := []byte(inputSlice)
		// fmt.Println(inputBSlice)
		charMap := make(map[byte][]string)

		for inp := 0; inp < len(inputBSlice); inp++ {
			// find the corresponding line num
			startLine := (int(inputBSlice[inp])-32)*9 + 2
			// fmt.Println("startLine: ", startLine)
			// scan the reqired lines in from the txt file
			fread, err := os.Open(banner + ".txt")
			checkErr(err)
			defer fread.Close()

			scanner := bufio.NewScanner(fread)

			scanner.Split(bufio.ScanBytes)

			// scan 8 lines starting from startLine from the txt file
			bigChar, err := scanChar(fread, startLine)
			checkErrNoEOF(err)
			if len(bigChar) != 8 {
				fmt.Println("Wrong number of lines read")
			}

			charMap[inputBSlice[inp]] = bigChar
		}
		// }
		printBigChar(&charMap, inputBSlice)
	}
}

func printBigChar2(chMap *map[byte][]string, inpBSlice []byte) {
	for l := 0; l < 8; l++ {
		chLine := ""
		for ch := 0; ch < len(inpBSlice); ch++ {
			chLine += string((*chMap)[inpBSlice[ch]][l])
			fmt.Println(chLine)
		}
	}
}
