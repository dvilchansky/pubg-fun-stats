package dbplayers

import (
	"github.com/dvilchansky/gopubg/models/player"
	"pubg-fun-stats/db"
)

type Player struct {
	ID         int    `json:"id"`
	PlayerId   string `json:"player_id"`
	PlayerName string `json:"player_name"`
}

func SelectPlayerById(id int) (*Player, error) {
	p := new(Player)
	err := db.Connection.
		QueryRow("SELECT * FROM players WHERE id = ?", id).
		Scan(&p.ID, &p.PlayerId, &p.PlayerName)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func SelectPlayerByName(name string) (*Player, error) {
	p := new(Player)
	err := db.Connection.
		QueryRow("SELECT * FROM players WHERE player_name = ?", name).
		Scan(&p.ID, &p.PlayerId, &p.PlayerName)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func Insert(players []*player.Player) error {
	for _, p := range players {
		insert, err := db.Connection.Query("INSERT IGNORE INTO players (player_id, player_name) VALUES (?,?)", p.ID, p.Name)
		if err != nil {
			return err
		}
		insert.Close()
	}
	return nil
}
