package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/ironarachne/pantheon"
	"github.com/ironarachne/random"
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
			fmt.Println("- " + domain)
		}
	}

	fmt.Println("Relationships")
	for _, relationship := range pantheon.Relationships {
		fmt.Println("- " + relationship.Origin.Name + " " + relationship.Descriptor + " " + relationship.Target.Name)
	}
}

func main() {
	randomSeed := flag.String("s", "now", "Optional random generator seed")

	flag.Parse()

	if *randomSeed == "now" {
		rand.Seed(time.Now().UnixNano())
	} else {
		random.SeedFromString(*randomSeed)
	}

	newPantheon := pantheon.GeneratePantheon()
	displayPantheon(newPantheon)
}
