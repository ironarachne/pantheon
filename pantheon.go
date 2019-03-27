package pantheon

import (
	"math/rand"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/random"
)

var (
	descriptors = []string{"is parent to", "hates", "loves", "fears", "respects", "is amused by", "is chagrined by"}
	genders     = []string{"male", "female", "none"}
)

// Deity is a fictional god or goddess
type Deity struct {
	Name              string
	Domains           []string
	Appearance        string
	Gender            string
	PersonalityTraits []string
}

// Pantheon is a nonhierarchical group of deities
type Pantheon struct {
	Deities       []Deity
	Relationships []Relationship
}

// Relationship is a unidirectional relationship between deities
type Relationship struct {
	Origin     Deity
	Target     Deity
	Descriptor string
}

// GenerateDeity generates a random deity
func GenerateDeity() Deity {
	var deity Deity
	var domain string

	numberOfDomains := rand.Intn(3) + 1

	for i := 0; i < numberOfDomains; i++ {
		domain = random.Item(domains)

		// Only add domain if it isn't already in Domains slice
		if !inSlice(domain, deity.Domains) {
			deity.Domains = append(deity.Domains, domain)
		}
	}

	deity.Appearance = random.Item(appearances)
	deity.Gender = random.Item(genders)

	for i := 0; i < 2; i++ {
		// Only add a trait if it isn't already in the PersonalityTraits slice
		trait := random.Item(traits)
		if !inSlice(trait, deity.PersonalityTraits) {
			deity.PersonalityTraits = append(deity.PersonalityTraits, trait)
		}
	}
	deity.Name = randomName()
	return deity
}

// GeneratePantheon creates a random pantheon of deities
func GeneratePantheon(maxSize int) Pantheon {
	var descriptor string
	var relationship Relationship
	var deity, target Deity
	var pantheon Pantheon

	numberOfDeities := rand.Intn(maxSize) + 1

	for i := 0; i < numberOfDeities; i++ {
		deity = GenerateDeity()
		pantheon.Deities = append(pantheon.Deities, deity)
	}

	for _, deity := range pantheon.Deities {
		target = pantheon.Deities[rand.Intn(len(pantheon.Deities))]
		descriptor = random.Item(descriptors)
		relationship = Relationship{Origin: deity, Target: target, Descriptor: descriptor}

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

func randomName() string {
	var person *naminglanguage.Person

	person = naminglanguage.GeneratePerson()

	return person.FirstName
}
