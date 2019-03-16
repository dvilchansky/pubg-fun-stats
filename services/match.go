package services

import (
	"github.com/dvilchansky/gopubg"
	match_ "github.com/dvilchansky/gopubg/models/match"
	"github.com/dvilchansky/gopubg/models/player"
	"pubg-fun-stats/models"
	"pubg-fun-stats/repositories"
	"runtime"
	"strings"
	"sync"
)

type MatchService interface {
	Fetch(num int64) ([]*models.Match, error)
	//FetchManyByIds(matchIDs string) ([]*models.Match, error)
	//GetByID(id string) (*models.Match, error)
	RequestPlayerMatches(userName string) ([]*models.Match, error)
	Store(m *models.Match) error
}

// NewMovieService returns the default movie service.
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

//func (ms *matchService) GetByID(id string) (*models.Match, error) {
//	return ms.repo.GetByID(id)
//}
//
func (ms *matchService) Store(m *models.Match) error {
	return ms.repo.Store(m)
}

func (ms *matchService) requestMatch(matchID string) (*match_.Match, error) {
	return ms.api.RequestMatch(matchID)
}

func (ms *matchService) prepareMatchesForCheck(matches []*player.Match) string {
	var preparedMatches []string
	for _, m := range matches {
		preparedMatches = append(preparedMatches, `"`+m.ID+`"`)
	}
	return strings.Join(preparedMatches, ",")
}

func (ms *matchService) requestPlayer(userName string) (*player.Player, error) {
	return ms.api.RequestPlayerByName(userName)
}

func (ms *matchService) RequestPlayerMatches(userName string) ([]*models.Match, error) {
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
	concurencyLevel := runtime.NumCPU() * 22
	var wg sync.WaitGroup
	lenM := len(p.Matches)
	println(lenM)
	for i := 0; i < lenM; {
		for k := 0; k < concurencyLevel; k++ {
			i++
			if i >= lenM {
				break
			}
			val, ok := matchesInDB[p.Matches[i].ID]
			if ok {
				result = append(result, val)
				continue
			}
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				m, _ := ms.requestMatch(p.Matches[i].ID)
				m_ := &models.Match{
					MatchID:   m.ID,
					ShardID:   m.ShardID,
					CreatedAt: m.CreatedAt,
					Duration:  m.Duration,
					GameMode:  m.GameMode,
					MapName:   m.MapName,
				}
				err = ms.repo.Store(m_)
				result = append(result, m_)
			}(i)
		}
		wg.Wait()
		if i >= lenM {
			break
		}
	}

	return result, nil
}
