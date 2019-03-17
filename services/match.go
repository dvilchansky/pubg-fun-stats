package services

import (
	"github.com/dvilchansky/gopubg"
	"github.com/dvilchansky/gopubg/models/match"
	"github.com/dvilchansky/gopubg/models/player"
	"pubg-fun-stats/models"
	"pubg-fun-stats/repositories"
	"runtime"
	"strings"
	"sync"
)

type MatchService interface {
	Fetch(num int64) ([]*models.Match, error)
	RequestPlayerMatches(userName string, lim int) ([]*models.Match, error)
	Store(m *models.Match) error
}

func NewMatchService(repo repositories.MatchRepository, api *gopubg.API) MatchService {
	return &matchService{
		api:  api,
		repo: repo,
	}
}

type matchService struct {
	repo repositories.MatchRepository
	api  *gopubg.API
}

func (ms *matchService) Fetch(num int64) ([]*models.Match, error) {
	return ms.repo.Fetch(num)
}

func (ms *matchService) Store(m *models.Match) error {
	return ms.repo.Store(m)
}

func (ms *matchService) prepareMatchesForCheck(matches []*player.Match) string {
	var preparedMatches []string
	for _, m := range matches {
		preparedMatches = append(preparedMatches, `"`+m.ID+`"`)
	}
	return strings.Join(preparedMatches, ",")
}

func (ms *matchService) RequestPlayerMatches(userName string, lim int) ([]*models.Match, error) {
	p, err := ms.api.RequestPlayerByName(userName)
	if err != nil {
		return nil, err
	}
	matchIDs := ms.prepareMatchesForCheck(p.Matches)
	matchesInDB, err := ms.repo.FetchByIDs(matchIDs)
	result := make([]*models.Match, 0)
	if err != nil {
		return nil, err
	}
	concurencyLevel := runtime.NumCPU() * 8
	var wg sync.WaitGroup
	lenM := len(p.Matches)
	for i := -1; i < lenM; {
		for k := 0; k < concurencyLevel; k++ {
			i++
			if i >= lim || i >= lenM {
				break
			}
			val, ok := matchesInDB[p.Matches[i].ID]
			if ok {
				result = append(result, val)
				continue
			}
			wg.Add(1)
			go func(matchID string) {
				defer wg.Done()
				m := ms.RequestMatchAndStore(matchID)
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

func NewMatchModelFromMatchStruct(m *match.Match) *models.Match {
	return &models.Match{
		MatchID:   m.ID,
		ShardID:   m.ShardID,
		CreatedAt: m.CreatedAt,
		Duration:  m.Duration,
		GameMode:  m.GameMode,
		MapName:   m.MapName,
	}
}

func (ms *matchService) RequestMatchAndStore(matchID string) *models.Match {
	m, _ := ms.api.RequestMatch(matchID)
	m_ := NewMatchModelFromMatchStruct(m)
	ms.repo.Store(m_)
	return m_
}
