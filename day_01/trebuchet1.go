package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
    "strconv"
)


func main() {
    finalSum := 0

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var d1 = 0
        var d2 = 0
        var line = filescanner.Text()

        for _, r := range line {
            r_s := string(r) 
            isdig := unicode.IsDigit(r)
            if(isdig == true) {
               if(d1 == 0) {
                   d1, _ = strconv.Atoi(r_s)
               } else {
                   d2, _ = strconv.Atoi(r_s)
               }
            }
        }

        // If there is only one digit, it is both first and last
        if(d2 == 0) {
            d2 = d1
        }

        if(d1 > 0 && d2 > 0) {
            tnum := d1*10 + d2
            finalSum += tnum
            d1 = 0
            d2 = 0
        }
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Final Sum: %d\n", finalSum)
}
