package model

import (
	"time"
)

type Game struct {
	ID        uint       `gorm:"column: id; NOT NULL; autoIncrement"json:"id"`
	StartTime time.Time  `gorm:"column:start_time;NOT NULL" json:"start_time"`
	EndTime   *time.Time `gorm:"column:end_time" json:"end_time"`
	Player1   uint       `gorm:"column:player_1;NOT NULL" json:"player_1"`
	Player2   uint       `gorm:"column:player_2;NOT NULL" json:"player_2"`
	Winner    uint       `gorm:"column:winner" json:"winner"`
	Episode   int        `gorm:"column:episode; default: 0" json:"episode"`
	Status    string     `gorm:"column:driving_code;type:enum('Processing', 'END', 'STOP','INITIAL');default: 'INITIAL'" json:"status"`
}

func PlayGame(player1 uint, player2 uint) (Game, error) {
	var game Game
	game.Player1 = player1
	game.Player2 = player2
	game.StartTime = time.Now()

	err := db.Create(&game).Error
	if err != nil {
		return game, err
	}

	var move1 Move
	var move2 Move
	move1.GameID = game.ID
	move1.PlayerID = player1
	move2.GameID = game.ID
	move2.PlayerID = player2

	err = db.Create(&move1).Error
	if err != nil {
		return game, err
	}
	err = db.Create(&move2).Error
	if err != nil {
		return game, err
	}

	return game, nil

}
