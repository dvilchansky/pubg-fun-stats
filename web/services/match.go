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
	Fetch(num int64) ([]*match.Match, error)
	RequestPlayerMatches(userName string, lim int) ([]*match.Match, error)
	Store(m *match.Match) error
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

func (ms *matchService) Fetch(num int64) ([]*match.Match, error) {
	return ms.repo.Fetch(num)
}

func (ms *matchService) Store(m *match.Match) error {
	return ms.repo.Store(m)
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
	//matchIDs := ms.prepareMatchesForCheck(p.Matches)
	//matchesInDB, err := ms.repo.FetchByIDs(matchIDs)
	result := make([]*match.Match, 0)
	if err != nil {
		return nil, err
	}
	concurencyLevel := /*runtime.NumCPU() * 8*/ 20
	var wg sync.WaitGroup
	lenM := len(p.Matches)
	for i := -1; i < lenM; {
		for k := 0; k < concurencyLevel; k++ {
			i++
			if i >= lim || i >= lenM {
				break
			}
			wg.Add(1)
			go func(matchID string) {
				defer wg.Done()
				m, _ := ms.api.RequestMatch(matchID)
				//_, ok := matchesInDB[p.Matches[i].ID]
				//if !ok {
				//	ms.repo.Store(m)
				//}
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
