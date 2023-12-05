package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "unicode"
    //"slices"
)

func IntPow(n, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}

func processCard(cardparts []string) ([]int, []int) {
    var wn = []int{}
    var pn = []int{}

    left := strings.Split(cardparts[0], " ")
    right := strings.Split(cardparts[1], " ")

    for _, strl := range(left) {
        elem1, _ := strconv.Atoi(strl)
        wn = append(wn, elem1)
    }

    for _, strr := range(right) {
        elem2, _ := strconv.Atoi(strr)
        pn = append(pn, elem2)
    }

    return wn, pn
}

func computeCard(wn []int, pn []int, idx int) int {
    result, wnums := 0, 0
    seen := make(map[int]int)

    for pidx := 0; pidx < len(pn); pidx++ {
        for widx := 0; widx < len(wn); widx++ {
            if wn[widx] == pn[pidx] {
                seen[pn[pidx]] = 1
            } 
        }
    }

    wnums = len(seen)-1
    if wnums > 0 {
        var exp = wnums-1
        result = IntPow(2, exp)
    }

    return result
}

func main() {
    cardstotal := 0
    winningNums := [][]int{}
    playerNums := [][]int{}

    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var line = filescanner.Text()

        lparts := strings.Split(line, ":")
        cardline := strings.TrimLeftFunc(lparts[1], unicode.IsSpace)
        numparts := strings.Split(cardline, " | ")
        wn, pn := processCard(numparts)
        winningNums = append(winningNums, wn)
        playerNums = append(playerNums, pn)
    }

    if err := filescanner.Err(); err != nil {
        log.Fatal(err)
    }

    nr := len(winningNums)
    for idx := 0; idx<nr; idx++ {
        fmt.Printf("Calling computecard with idx %d\n", idx)
        pts := computeCard(winningNums[idx], playerNums[idx], idx) 
        fmt.Printf("Card %d: pts: %d\n", idx, pts)
        cardstotal += pts
    }

    fmt.Printf("Part 1 - Scratchcards worth: %d\n", cardstotal)
}
