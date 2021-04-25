package rulebasedconv

import (
	"sort"
	"strings"
)

type rule struct {
	match      string
	replace    string
	exceptions []exception
}

type exception struct {
	ifAllMatch  []matchCondition
	thenReplace string
}

type matchCondition struct {
	when  string
	is    string
	value string
	isNot bool
}

type rules []rule

func newRules() *rules {
	var rulesCopy = make(rules, len(avroClassicPhoneticRules))

	copy(rulesCopy, avroClassicPhoneticRules)

	rules := &rulesCopy

	// The converter algorithm depends on rules being sorted by descending order of match length.
	// It's hard to maintain that manually when we design avroClassicPhoneticRules, so we don't enforce it there and
	// do the sorting when we initialize the rules from the avroClassicPhoneticRules.
	rules.sortRulesByDescendingMatchLength()

	// Set the matchCondition.isNot boolean property based on the presence of ! character
	// in matchCondition.is to speed up conversion
	rules.updateNegativeConditions()

	return rules
}

func (rules *rules) sortRulesByDescendingMatchLength() {
	sort.Slice(*rules, func(i, j int) bool {
		return len((*rules)[i].match) > len((*rules)[j].match)
	})
}

func (rules *rules) updateNegativeConditions() {
	for i := 0; i < len(*rules); i++ {
		rule := (*rules)[i]

		for j := 0; j < len(rule.exceptions); j++ {
			exception := &rule.exceptions[j]

			for k := 0; k < len(exception.ifAllMatch); k++ {
				matchCondition := &exception.ifAllMatch[k]

				matchCondition.isNot = false
				if strings.HasPrefix(matchCondition.is, "!") {
					matchCondition.isNot = true
					matchCondition.is = matchCondition.is[1:]
				}
			}
		}
	}
}
