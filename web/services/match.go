package services

import (
	"pubg-fun-stats/parser"
	"pubg-fun-stats/parser/models/match"
	"sync"
)

type MatchService interface {
	RequestPlayerMatches(userName string, lim int) ([]*match.Match, error)
}

func NewMatchService(api *gopubg.API) MatchService {
	return &matchService{
		api: api,
	}
}

type matchService struct {
	api *gopubg.API
}

func (ms *matchService) RequestPlayerMatches(userName string, lim int) ([]*match.Match, error) {
	p, err := ms.api.RequestPlayerByName(userName)
	if err != nil {
		return nil, err
	}
	result := make([]*match.Match, 0)
	var wg sync.WaitGroup
	lenM := len(p.Matches)
	for i := -1; i < lenM; {
		for k := 0; k < 20; k++ {
			i++
			if i >= lim || i >= lenM {
				break
			}
			wg.Add(1)
			go func(matchID string) {
				defer wg.Done()
				m, _ := ms.api.RequestMatch(matchID)
				if m.MapName == `Range_Main` {
					return
				}
				result = append(result, m)
			}(p.Matches[i].ID)
		}
		wg.Wait()
		if i >= lim || i >= lenM {
			break
		}
	}

	return result, nil
}
