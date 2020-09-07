package randomua

import (
	"encoding/json"
	"github.com/mroth/weightedrand"
	"log"
	"math/rand"
	"time"
)

var ua weightedrand.Chooser

//data from intoli

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	var uas []weightedrand.Choice
	var luas []struct{
		UserAgent string `json:"userAgent"`
		Weight int64 `json:"weight"`
	}

	err := json.Unmarshal([]byte(load), &luas)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range luas {
		uas = append(uas, weightedrand.NewChoice(v.UserAgent, uint(v.Weight)))
	}

	ua = weightedrand.NewChooser(uas...)
}

func Random() string {
	return ua.Pick().(string)
}