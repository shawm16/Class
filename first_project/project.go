package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "math"
    "strconv"
)

func checker(e error){
  if e != nil{
    fmt.Println("ERROR")
  }
}

func calc_dist(lat_one float64, lat_two float64, long_one float64, long_two float64) float64 {
  var r float64
  r = 6371 + 6371
  var x float64
  x = r * (math.Asin(math.Sqrt((math.Sin(math.Pow((lat_two - lat_one)/(2), 2))) + (math.Cos(lat_one) * math.Cos(lat_two) * (math.Sin(math.Pow((long_two - long_one)/(2),2)))))))
  return x
}


func main() {
    // Open file and create scanner on top of it
    file, err := os.Open("conties_lat_long.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)

    // Default scanner is bufio.ScanLines. Lets use ScanWords.
    // Could also use a custom function of SplitFunc type
    scanner.Split(bufio.ScanWords)

    // Scan for next token.
    success := true
    if success == false {
        // False on error or EOF. Check error
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        } else {
            log.Fatal(err)
        }
    }

    // Get data from scan with Bytes() or Text()
    // fmt.Println("First word found:", scanner.Text())
    // Call scanner.Scan() again to find next token
    Name := make(map[int]string)
    lat := make(map[int]string)
    long := make(map[int]string)
    i := 1
    for scanner.Scan() {
      Name[i] = scanner.Text()
      for scanner.Scan() {
        lat[i] = scanner.Text()
        for scanner.Scan() {
          long[i] = scanner.Text()
          i++
        }
      }
    }

    var total_distance float64
    var lat_one, lat_two, long_one, long_two float64
    for x := 1; x <= i; x = x + 2 {
      lat_one, err = strconv.ParseFloat(lat[i], 64)
      checker(err)
      lat_two, err = strconv.ParseFloat(lat[i+1], 64)
      checker(err)
      long_one, err = strconv.ParseFloat(long[i], 64)
      checker(err)
      long_two, err = strconv.ParseFloat(long[i+1], 64)
      checker(err)
      total_distance =+ calc_dist(lat_one, lat_two, long_one, long_two)
    }
    fmt.Println(total_distance)
}
