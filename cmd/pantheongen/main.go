package main

import (
	"github.com/ironarachne/pantheon"
	"math/rand"
	"flag"
	"fmt"
	"time"
)

func displayPantheon(pantheon pantheon.Pantheon) {
	fmt.Println("Deities")
	for _, deity := range pantheon.Deities {
		fmt.Println(deity.Name)
		fmt.Println("Gender: " + deity.Gender)
		fmt.Println("Appearance: " + deity.Appearance)

		fmt.Println("## Personality Traits")
		for _, trait := range deity.PersonalityTraits {
			fmt.Println(trait)
		}

		fmt.Println("## Domains")
		for _, domain := range deity.Domains {
			fmt.Println(domain)
		}

		fmt.Println("\n")
	}

	fmt.Println("Relationships\n")
	for _, relationship := range pantheon.Relationships {
		fmt.Println(relationship.Origin.Name + " " + relationship.Descriptor + " " + relationship.Target.Name)
	}
}

func main() {
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")

	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	pantheon := pantheon.GeneratePantheon()
	displayPantheon(pantheon)
}
