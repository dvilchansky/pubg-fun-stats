package services

import (
	"pubg-fun-stats/parser"
	"pubg-fun-stats/parser/models/match"
	"pubg-fun-stats/parser/models/player"
	"pubg-fun-stats/repositories"
	"strings"
	"sync"
)

type MatchService interface {
	RequestPlayerMatches(userName string, lim int) ([]*match.Match, error)
}

func NewMatchService(repo repositories.MatchRepository, api *gopubg.API) MatchService {
	return &matchService{
		repo: repo,
		api:  api,
	}
}

type matchService struct {
	repo repositories.MatchRepository
	api  *gopubg.API
}

func (ms *matchService) prepareMatchesForCheck(matches []*player.Match) string {
	var preparedMatches []string
	for _, m := range matches {
		preparedMatches = append(preparedMatches, `"`+m.ID+`"`)
	}
	return strings.Join(preparedMatches, ",")
}

func (ms *matchService) RequestPlayerMatches(userName string, lim int) ([]*match.Match, error) {
	p, err := ms.api.RequestPlayerByName(userName)
	if err != nil {
		return nil, err
	}
	if p.Matches == nil {
		return nil, nil
	}
	var result []*match.Match
	concurrencyLevel := /*runtime.NumCPU() * 8*/ 20
	var wg sync.WaitGroup
	lenM := len(p.Matches)
	for i := -1; i < lenM; {
		for k := 0; k < concurrencyLevel; k++ {
			i++
			if i >= lim || i >= lenM {
				break
			}
			wg.Add(1)
			go func(matchID string) {
				defer wg.Done()
				m, err := ms.api.RequestMatch(matchID)
				if err != nil {
					panic(err.Error())
				} else {
					result = append(result, m)
				}
			}(p.Matches[i].ID)
		}
		wg.Wait()
		if i >= lim || i >= lenM {
			break
		}
	}
	return result, nil
}
