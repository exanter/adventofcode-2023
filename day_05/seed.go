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

var seeds = []int{}
var seedToSoil = [][]int{}
var soilToFertilizer = [][]int{}
var fertilizerToWater = [][]int{}
var waterToLight = [][]int{}
var lightToTemperature = [][]int{}
var temperatureToHumidity = [][]int{}
var humidityToLocation = [][]int{}

const (
    SeedToSoil int            = 0
    SoilToFertilizer int      = 1
    FertilizerToWater int     = 2
    WaterToLight int          = 3
    LightToTemperature int    = 4
    TemperatureToHumidity int = 5
    HumidityToLocation int    = 6
)

func numberlineToArray(line string) []int {
    var result []int

    numbers := strings.Split(line, " ")
    for _, snum := range(numbers) {
        tmp, _ := strconv.Atoi(snum)
        result = append(result, tmp)
    }

    return result
}

func processNumberLine(mode int, line string) {
    switch mode {
    case SeedToSoil:
        seedToSoil = append(seedToSoil, numberlineToArray(line)) 
    case SoilToFertilizer:
        soilToFertilizer = append(soilToFertilizer, numberlineToArray(line)) 
    case FertilizerToWater:
        fertilizerToWater = append(fertilizerToWater, numberlineToArray(line))
    case WaterToLight:
        waterToLight = append(waterToLight, numberlineToArray(line)) 
    case LightToTemperature:
        lightToTemperature = append(lightToTemperature, numberlineToArray(line))
    case TemperatureToHumidity:
        temperatureToHumidity = append(temperatureToHumidity, numberlineToArray(line))
    case HumidityToLocation:
        humidityToLocation = append(humidityToLocation, numberlineToArray(line))
    }
}


func main() {
    mode := 0
    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    filescanner := bufio.NewScanner(file)
    for filescanner.Scan() {
        var line = filescanner.Text()

        lparts := strings.Split(line, ":")
        switch lparts[0] {
        case "seeds":
            tline := strings.TrimLeftFunc(lparts[1], unicode.IsSpace)
            seeds = numberlineToArray(tline)
        case "seed-to-soil map":
            mode = SeedToSoil
        case "soil-to-fertilizer map":
            mode = SoilToFertilizer
        case "fertilizer-to-water map":
            mode = FertilizerToWater
        case "water-to-light map":
            mode = WaterToLight
        case "light-to-temperature map":
            mode = LightToTemperature
        case "temperature-to-humidity map":
            mode = TemperatureToHumidity
        case "humidity-to-location map":
            mode = HumidityToLocation
        case "": // blank line

        default:
            processNumberLine(mode, lparts[0])
        } 
    }

    fmt.Printf("Seeds: %+v\n", seeds)
    fmt.Printf("seed-to-soil: %+v\n", seedToSoil)
    fmt.Printf("soil-to-fertilizer: %+v\n", soilToFertilizer)
    fmt.Printf("fertilizer-to-water: %+v\n", fertilizerToWater)
    fmt.Printf("water-to-light: %+v\n", waterToLight)
    fmt.Printf("light-to-temperature: %+v\n", lightToTemperature)
    fmt.Printf("temperature-to-humidity: %+v\n", temperatureToHumidity)
    fmt.Printf("humidity-to-location: %+v\n", humidityToLocation)
}
