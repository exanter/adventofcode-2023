package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "unicode"
    "strconv"
    "strings"
)


func Substring(str string, start, end int) string {
    return strings.TrimSpace(str[start:end])
}


func reverseScanForDigit(line string) (int, int) {
    var result int = 0
    var strlen = len(line)
    var numberStrs = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    for i := strlen-1; i > 0; i-- {
        r := rune(line[i])
        isdig := unicode.IsDigit(r)
        if(isdig == true) {
            result, _ = strconv.Atoi(string(r))
            if i == 0 {
                i = 1
            }

            return result, i
        } else {
            slice := Substring(line, i, strlen)
            for num := 0; num < len(numberStrs); num++  {
                if(strings.HasPrefix(slice, numberStrs[num])) {
                    result = num+1
                    return result, (i+len(numberStrs[num]))
                }
            }
        }
    }

    return -1, strlen
}


func scanForDigit(line string) (int, int) {
    var result int = 0
    var strlen = len(line)
    var numberStrs = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    for i := 0; i < strlen; i++ {
        r := rune(line[i])
        isdig := unicode.IsDigit(r)
        if(isdig == true) {
            result, _ = strconv.Atoi(string(r))
            if i == 0 {
                i = 1
            }
            return result, i
        } else {
            slice := Substring(line, i, strlen)
            for num := 0; num < len(numberStrs); num++  {
                if(strings.HasPrefix(slice, numberStrs[num])) {
                    result = num+1
                    return result, (i+len(numberStrs[num]))
                }
            }
        }
    }

    return -1, strlen
}


func main() {
    finalSum := 0

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var d1, d2 = 0, 0
        var line = filescanner.Text()

        d1, _ = scanForDigit(line)
        d2, _ = reverseScanForDigit(line)
        
        // If there is only one digit, it is both first and last
        if(d2 == -1) {
            d2 = d1
        }

        tnum := d1*10 + d2
        finalSum += tnum
        d1, d2 = 0, 0
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Final Sum: %d\n", finalSum)
}
