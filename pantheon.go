package pantheon

import (
	"math/rand"

	"github.com/ironarachne/naminglanguage"
)

var (
	descriptors = []string{"is parent to", "hates", "loves", "fears", "respects", "is amused by", "is chagrined by"}
	genders     = []string{"male", "female", "none"}
	traits		= getTraits()
)

type Deity struct {
	Name              string
	Domains           []string
	Appearance        string
	Gender            string
	PersonalityTraits []string
}

type Pantheon struct {
	Deities       []Deity
	Relationships []Relationship
}

type Relationship struct {
	Origin     Deity
	Target     Deity
	Descriptor string
}

func GenerateDeity() Deity {
	deity := Deity{}
	numberOfDomains := rand.Intn(2) + 1
	i := 0
	for i < numberOfDomains {
		domain := domains[rand.Intn(len(domains))]

		// Only add domain if it isn't already in Domains slice
		if !inSlice(domain, domains) {
			deity.Domains = append(deity.Domains, domain)
			i++
		}
	}
	deity.Appearance = appearances[rand.Intn(len(appearances))]
	deity.Gender = genders[rand.Intn(len(genders))]

	numTraits := 0
	for numTraits < 2 {
		// Only add a trait if it isn't already in the PersonalityTraits slice
		trait := traits[rand.Intn(len(traits))]
		if !inSlice(trait, deity.PersonalityTraits) {
			deity.PersonalityTraits = append(deity.PersonalityTraits, trait)
			numTraits++
		}
	}
	deity.Name = naminglanguage.GeneratePersonName()
	return deity
}

func GeneratePantheon() Pantheon {
	numberOfDeities := rand.Intn(28) + 3
	pantheon := Pantheon{}

	for i := 0; i < numberOfDeities; i++ {
		deity := GenerateDeity()
		pantheon.Deities = append(pantheon.Deities, deity)
	}
	for _, deity := range pantheon.Deities {
		target := pantheon.Deities[rand.Intn(len(pantheon.Deities))]
		descriptor := descriptors[rand.Intn(len(descriptors))]
		relationship := Relationship{Origin: deity, Target: target, Descriptor: descriptor}
		if deity.Name != target.Name {
			pantheon.Relationships = append(pantheon.Relationships, relationship)
		}
	}

	return pantheon
}

func inSlice(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}
