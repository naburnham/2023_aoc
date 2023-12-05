package main

import (
    "testing"
)

func TestTableGetTwoDigitNumber(t *testing.T) {
    testTable := []struct{
        name string
        line string
        expected int
    }{
        {"First string", "1abc2", 12},
        {"Second string", "pqr3stu8vwx", 38},
        {"Third string", "a1b2c3d4e5f", 15},
        {"Fourth string", "treb7uchet", 77},
    } 

    for _, tt := range testTable {
        t.Run(tt.name, func(t *testing.T) {
            got := getTwoDigitNumberFromTextString(tt.line)
            if got != tt.expected {
                t.Errorf("got %d, expected %d", got, tt.expected)
            }
        })
    }
}
