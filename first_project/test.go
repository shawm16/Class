package main

import (
    "fmt"
    "io/ioutil"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("conties_lat_long.txt")
    check(err)
    file_as_string := []string(dat)
    for i := 0; i < len(file_as_string); i++ {
      fmt.Print(file_as_string[i])
    }

}
