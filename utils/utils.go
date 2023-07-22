package utils

import (
	"github.com/jjjosephhh/solitaire-thirteens/card"
	"github.com/jjjosephhh/solitaire-thirteens/constants"
)

func MatchExists(inPlay []*card.Card) bool {
	if len(inPlay) == 0 {
		return true
	}
	seen := make(map[int]bool)
	for _, c := range inPlay {
		dif := constants.TARGET_NUM - c.Num
		if dif == 0 {
			return true
		}
		if _, ok := seen[dif]; ok {
			return true
		}
		seen[c.Num] = true
	}
	return false
}

func IsMatch(c1, c2 *card.Card) []*card.Card {
	var matches []*card.Card
	if c1 == nil && c2 == nil {
		return matches
	} else if c1 == nil && c2.Num == 13 {
		matches = append(matches, c2)
	} else if c2 == nil && c1.Num == 13 {
		matches = append(matches, c1)
	} else if c1 != nil && c2 != nil && c1.Num+c2.Num == 13 {
		matches = append(matches, c1, c2)
	}
	return matches
}
