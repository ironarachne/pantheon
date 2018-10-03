package pantheon

import (
	"math/rand"

	"github.com/ironarachne/naminglanguage"
)

var (
	descriptors = []string{"is parent to", "hates", "loves", "fears", "respects", "is amused by", "is chagrined by"}
	domains     = []string{"life", "death", "the sun", "the moon", "nature", "hunting", "war", "balance", "chaos", "law"}
	genders     = []string{"male", "female", "none"}
	traits      = []string{"fierce", "proud", "humble", "quiet", "boisterous"}
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
	for i := 0; i < numberOfDomains; i++ {
		deity.Domains = append(deity.Domains, domains[rand.Intn(len(domains))])
	}
	deity.Appearance = appearances[rand.Intn(len(appearances))]
	deity.Gender = genders[rand.Intn(len(genders))]
	for i := 0; i < 2; i++ {
		deity.PersonalityTraits = append(deity.PersonalityTraits, traits[rand.Intn(len(traits))])
	}
	deity.Name = naminglanguage.GeneratePersonName()
	return deity
}

func GeneratePantheon() Pantheon {
	numberOfDeities := 12
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
