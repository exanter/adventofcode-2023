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

type BagSample struct {
    r int
    b int
    g int
}

func parseHeader(hdr string) int {
    var parts = strings.Split(hdr, " ")
    var result, _ = strconv.Atoi(parts[1])

    return result
}


func parseGameLine(line string) []BagSample {
    var result []BagSample
    var parts = strings.Split(line, ";")

    for _,part := range parts {
        var bs_elem BagSample

        sp := strings.TrimLeftFunc(part, unicode.IsSpace)
        items := strings.Split(sp, ",")
        for _,combo := range items {
            cb := strings.TrimLeftFunc(combo, unicode.IsSpace)
            qcolor := strings.Split(cb, " ") 

            switch color := qcolor[1]; color {
            case "red":
               bs_elem.r,_ = strconv.Atoi(qcolor[0])
            case "green":
               bs_elem.g,_ = strconv.Atoi(qcolor[0])
            case "blue":
               bs_elem.b,_ = strconv.Atoi(qcolor[0])
            default:
               fmt.Printf("Uhh, WHAT: qcolor: [%s] [%s]\n", qcolor[0], qcolor[1])
           }
        }

        result = append(result, bs_elem)
    }

    return result
}


func getMaxValues(game []BagSample) BagSample {
    var result BagSample 

    for _, sample := range game {
        if sample.r > result.r {
            result.r = sample.r
        }
        if sample.g > result.g {
            result.g = sample.g
        }
        if sample.b > result.b {
            result.b = sample.b
        }
    }

    return result
}


func main() {
    finalSum := 0
    //currentID := 0

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var line = filescanner.Text()
        var parts = strings.Split(line, ":")

        sample := parseGameLine(parts[1])
        var maxval = getMaxValues(sample)
        var power = maxval.r * maxval.g * maxval.b
        finalSum += power
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Final Sum: %d\n", finalSum)
}
