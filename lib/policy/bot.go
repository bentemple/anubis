package policy

import (
	"fmt"

	"github.com/bentemple/anubis/internal"
	"github.com/bentemple/anubis/lib/policy/checker"
	"github.com/bentemple/anubis/lib/policy/config"
)

type Bot struct {
	Rules     checker.Impl
	Challenge *config.ChallengeRules
	Weight    *config.Weight
	Name      string
	Action    config.Rule
}

func (b Bot) Hash() string {
	return internal.FastHash(fmt.Sprintf("%s::%s", b.Name, b.Rules.Hash()))
}
