package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "unicode"
    "strconv"
)

func getTwoDigitNumberFromTextString(line string) int {
    var temp_dig, first_dig, second_dig rune
    first_dig = -1

    runes := []rune(line)

    for _, r := range runes {
        if unicode.IsDigit(r) {
            temp_dig = r
            if first_dig == -1 {
                first_dig = r
            }
        }
    }
    second_dig = temp_dig

    i, err := strconv.Atoi(string(first_dig)+string(second_dig))
    if err != nil {
        return 0
    }
    return i 
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)

    var sum int = 0

    for scanner.Scan() {
        sum += getTwoDigitNumberFromTextString(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println(sum)
}
