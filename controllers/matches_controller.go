package controllers

import (
	"database/sql"
	"fmt"
	"github.com/dvilchansky/gopubg"
	"github.com/dvilchansky/gopubg/models/player"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"pubg-fun-stats/funstats"
	"strings"
	"sync"
)

var (
	wg                     sync.WaitGroup
	DBW                    *sql.DB
	StmtMain               *sql.Stmt
	SQLMatches             []*funstats.SQLMatch
	SQLInsertQuery         = "INSERT IGNORE INTO `matches` SET `match_id` = ?,`shard_id` = ?,`created_at` = ?,`duration` = ?,`game_mode` = ?, `map_name` = ?"
	SQLInsertTrainingQuery = "INSERT IGNORE INTO `matches_training` SET `match_id` = ?"
	SQLSelectQuery         = "SELECT `match_id`, `shard_id`, `created_at`, `duration`, `game_mode`, `map_name` FROM `matches` WHERE `match_id` IN (%s)"
	SQLSelectTrainingQuery = "SELECT `match_id` FROM `matches_training` WHERE `match_id` IN (%s)"
	errDbw, errStmt        error
)

func ProcessMatches(cl int, a *gopubg.API, matches []*player.Match) []*funstats.SQLMatch {
	DBW, errDbw = sql.Open("mysql", "root:root@tcp(127.0.0.1:3307)/pubg_fun_stats?parseTime=true")
	if errDbw != nil {
		log.Fatalln(errDbw)
	}
	defer DBW.Close()
	var ms []string
	for _, m := range matches {
		ms = append(ms, `"`+m.ID+`"`)
	}
	mstring := strings.Join(ms, ",")
	result := FetchMatches(mstring)
	training_matches := FetchTrainingMatches(mstring)
	DBW.SetMaxIdleConns(cl)
	StmtMain, errStmt = DBW.Prepare(SQLInsertQuery)
	if errStmt != nil {
		log.Fatalln(errStmt)
	}
	lenM := len(matches)
	for i := 0; i < lenM; {
		for k := 0; k < cl; k++ {
			i++
			if i > 10 || i >= lenM {
				break
			}
			val, ok := result[matches[i].ID]
			if ok {
				SQLMatches = append(SQLMatches, val)
				continue
			}
			_, ok = training_matches[matches[i].ID]
			if ok {
				continue
			}
			wg.Add(1)
			go Store(a, matches[i].ID)
		}
		wg.Wait()
		if i > 10 || i >= lenM {
			break
		}
	}
	return SQLMatches
}

func Store(a *gopubg.API, matchID string) {
	defer wg.Done()
	m, err := a.RequestMatch("steam", matchID)
	if err != nil {
		panic(err.Error())
	}
	if m.MapName == "Range_Main" {
		insert, err := DBW.Query(SQLInsertTrainingQuery, m.ID)
		if err != nil {
			panic(err.Error())
		}
		insert.Close()
		return
	}
	SQLMatches = append(SQLMatches, &funstats.SQLMatch{
		MatchID:   m.ID,
		ShardID:   m.ShardID,
		CreatedAt: m.CreatedAt,
		Duration:  m.Duration,
		GameMode:  m.GameMode,
		MapName:   m.MapName,
	})
	_, err = StmtMain.Exec(m.ID, m.ShardID, m.CreatedAt, m.Duration, m.GameMode, m.MapName)
	if err != nil {
		panic(err.Error())
	}
}

func FetchMatches(ms string) map[string]*funstats.SQLMatch {
	rows, err := DBW.Query(fmt.Sprintf(SQLSelectQuery, ms))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := make(map[string]*funstats.SQLMatch)
	for rows.Next() {
		sqlMatch := new(funstats.SQLMatch)
		err = rows.Scan(&sqlMatch.MatchID, &sqlMatch.ShardID, &sqlMatch.CreatedAt, &sqlMatch.Duration, &sqlMatch.GameMode, &sqlMatch.MapName)
		if err != nil {
			panic(err)
		}
		result[sqlMatch.MatchID] = sqlMatch
	}
	return result
}

func FetchTrainingMatches(ms string) map[string]string {
	rows, err := DBW.Query(fmt.Sprintf(SQLSelectTrainingQuery, ms))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := make(map[string]string)
	var mID string
	for rows.Next() {
		if err != nil {
			panic(err)
		}
		rows.Scan(&mID)
		result[mID] = mID
	}
	return result
}
