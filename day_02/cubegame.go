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


func checkGameValues(red int, green int, blue int, game []BagSample) bool {
    var result bool = true

    fmt.Printf("checkGameValues, game: %+v\n", game)
    for _, sample := range game {
        if red < sample.r {
            return false
        }

        if green < sample.g {
            return false
        }

        if blue < sample.b {
            return false
        }
    }

    return result
}


func main() {
    finalIDSum := 0
    currentID := 0

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var line = filescanner.Text()
        var parts = strings.Split(line, ":")
        var hdr = parts[0]

        currentID = parseHeader(hdr)
        sample := parseGameLine(parts[1])

        fmt.Printf("gameid: %d  sample: %+v\n", currentID, sample)
        if checkGameValues(12, 13, 14, sample) == true {
            finalIDSum += currentID
        }
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Final Sum: %d\n", finalIDSum)
}
