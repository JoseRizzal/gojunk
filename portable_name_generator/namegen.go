package namegen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type (
	NameGenerator struct {
		adjectives []string
		consonants string
		vowels     string
		nounLenght int
		separator  string
		random     *rand.Rand
	}
)

func CreateNameGenerator() *NameGenerator {
	return &NameGenerator{
		adjectives: []string{
			"good", "red", "blue", "misty", "gold", "empty", "dry", "dark", "green",
			"icy", "old", "quiet", "white", "cool", "young", "evil", "steel",
			"warm", "dumb", "titan", "thin", "dirty", "funny", "plush",
			"silly", "cold", "idiot", "rich", "frosty", "poor", "long", "late", "short",
			"bold", "little", "smart", "agile", "quick", "prime", "rough", "still", "boss",
			"super", "nice", "shy", "deluxe", "luxury", "wild", "black",
			"easy", "holy", "cheer", "cozy", "ruby", "snowy", "proud", "floral",
			"brutal", "sad", "ancient", "purple", "silver", "iron"},
		consonants: "bdfghklmnprstvwxz",
		vowels:     "aeiou",
		random:     rand.New(rand.NewSource(time.Now().UnixNano())),

		nounLenght: 5,
		separator:  "_",
	}
}

// Public Methods

func (c *NameGenerator) Generate() string {
	return fmt.Sprint(c.adjectives[c.random.Intn(len(c.adjectives))] + c.separator + c.createWordWithLenght(c.nounLenght))
}

func (c *NameGenerator) SetSeparator(s string) {
	c.separator = s
}

func (c *NameGenerator) SetNounLenght(i int) {
	c.nounLenght = i
}

func (c *NameGenerator) SetAjectives(l []string) {
	c.adjectives = l
}

func (c *NameGenerator) SetConsonants(s string) {
	c.consonants = s
}

func (c *NameGenerator) SetVowels(s string) {
	c.vowels = s
}

// Local Methods

func (c *NameGenerator) pickRandomCharFromCharset(charset string) string {
	return string(charset[c.random.Intn(len(charset))])
}

func (c *NameGenerator) createWordWithLenght(lenght int) string {
	word := make([]string, lenght)
	for id := range word {
		if id%2 == 0 {
			word[id] = c.pickRandomCharFromCharset(c.consonants)
		} else {
			word[id] = c.pickRandomCharFromCharset(c.vowels)
		}
	}
	return strings.Join(word, "")
}
