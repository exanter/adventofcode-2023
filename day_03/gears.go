package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "unicode"
)

// This is the parts schematic given by the elves
var schematic []string

func Substring(str string, start, end int) string {
    return strings.TrimSpace(str[start:end])
}

type PartNum struct {
    row, col, numdigits int
    valid bool
}

func (p *PartNum) convertDigit() int {
    result := 0

    substr := Substring(schematic[p.row], p.col, p.col+p.numdigits)
    result, _ = strconv.Atoi(substr)
    return result
}

func (p *PartNum) printPart() {
    //fmt.Printf("row: %d, col: %d, %d digits, num: %d, valid: %t\n", 
     //   p.row, p.col, p.numdigits, p.convertDigit(), p.valid)
    fmt.Printf("valid group:  %d  (row: %d)\n", p.convertDigit(), p.row)
    return
}

type Symbol struct {
    row, col int
    stype byte
    pgear bool
}

func main() {
    var rowct, colct, finalSum, finalProduct int = 0, 0, 0, 0
    var parts []PartNum
    var symbols []Symbol

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var line = filescanner.Text()
        schematic = append(schematic, line)
        rowct++
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    colct = len(schematic[0])
    for x:=0; x<rowct; x++ {
        inpart := false
        var newpart PartNum

        for y := 0; y<colct; y++ {
            isdig := unicode.IsDigit(rune(schematic[x][y]))
            if isdig {
                if ! inpart {
                    inpart = true
                    newpart = PartNum{row: x, col: y, numdigits: 1, valid: false}
                } else {
                    newpart.numdigits++
                }

                // Catch a number that ends at the end of a line
                if y == (colct-1) { parts = append(parts, newpart) }
            } else {
                if inpart {
                    inpart = false
                    parts = append(parts, newpart)
                } 

                if rune(schematic[x][y]) != '.' {
                    newsym := Symbol{row: x, col: y, stype: schematic[x][y], pgear: false}
                    if rune(schematic[x][y]) == '*' { newsym.pgear = true }

                    symbols = append(symbols, newsym)
                }
            }
        }
    }

    gears := map[Symbol][]int {}
    pct := 0
    for pidx, part := range(parts) {
        for _, sym := range(symbols) {
            if sym.row >= (part.row-1) && sym.row <= (part.row+1) {
                if sym.col >= (part.col-1) && sym.col <= (part.col+part.numdigits) {
                    parts[pidx].valid = true
                    finalSum += parts[pidx].convertDigit()
                    if sym.stype == '*' {
                        gears[sym] = append(gears[sym], parts[pidx].convertDigit())
                    }
                }
            }
        }

        pct++
    }

    for _, g := range(gears) {
        if len(g) == 2 {
            finalProduct += (g[0] * g[1])
        }
    }

    fmt.Printf("Part 1 - FinalSum: %d\n", finalSum)
    fmt.Printf("Part 2 - FinalProduct: %d\n", finalProduct)
}
