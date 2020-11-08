//Package randomua generates User-Agent strings
package randomua

import (
	"encoding/json"
	"github.com/mroth/weightedrand"
	"log"
	"math/rand"
	"time"
)

var ua *weightedrand.Chooser

type choice struct {
	UserAgent string `json:"userAgent"`
	Weight    int64  `json:"weight"`
}

func init() {
	var choices []weightedrand.Choice
	var loadChoices []choice

	rand.Seed(time.Now().UTC().UnixNano())

	err := json.Unmarshal([]byte(load), &loadChoices)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range loadChoices {
		choices = append(choices, weightedrand.NewChoice(v.UserAgent, uint(v.Weight)))
	}

	ua, err = weightedrand.NewChooser(choices...)
	if err != nil {
		log.Fatal(err)
	}
}

//Random will give you a random User-Agent string
func Random() string {
	return ua.Pick().(string)
}
