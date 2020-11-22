package main

import (
    "fmt"
    "math/rand"
)

const secondsPerDay = 86400

func main() {
    distance := 62100000
    carrier := ""
    trip := ""

    fmt.Println("Spaceline         Days Trip type  Price")
    fmt.Println("=======================================")

    for count := 0; count < 10; count++{
        switch rand.Intn(3) {
            case 0:
                carrier = "SpaceX"
            case 1:
                carrier = "Space Adventures"
            case 2:
                carrier = "Virgin Galactic"
        }
        speed := rand.Intn(15) + 16
        trip_duration := (distance/speed)/secondsPerDay

        price := 20 + speed

        if rand.Intn(2) == 0 {
            trip = "Round Trip"
            price *= 2
        } else {
            trip = "One way"
        }
    fmt.Printf("%-16v %4v %-10v $%4v\n", carrier, trip_duration, trip, price) 
    }
}
