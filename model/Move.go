package model

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

type Move struct {
	ID       uint `gorm:"column: id; NOT NULL; autoIncrement"json:"id"`
	PlayerID uint `gorm:"column:player_id;NOT NULL" json:"player_id""`
	GameID   uint `gorm:"column:game_id;NOT NULL" json:"game_id"`
	StartPos int  `gorm:"column:start_pos;default 0" json:"start_pos"`
	EndPos   int  `gorm:"column:end_pos;default 0" json:"end_pos"`
	Step     int  `gorm:"column:step;default 0" json:"step"`
	Episode  int  `gorm:"column:episode;default 0" json:"episode"`
}

func PlayerMove(playerID uint, gameID uint) (Move, bool, error) {
	//generate step

	step := randomStep()

	var game Game
	db.First(&game, gameID)
	if game.ID == 0 {
		return Move{}, false, errors.New("invalid game")
	}
	if game.Status != "END" {
		//load map
		var mapp Map
		db.First(&mapp, game.MapID)
		var rule Rule
		err := json.Unmarshal([]byte(mapp.Rules), &rule)
		if err != nil {
			return Move{}, false, errors.New("load rule fail")
		}
		var curMove Move
		db.Where("game_id = ? and episode = ? and player_id = ?", gameID, game.Episode, playerID).First(&curMove)
		if curMove.ID == 0 {
			return Move{}, false, errors.New("get current move fail")
		}
		var nextMove Move
		switch playerID {
		case game.Player1:
			//check if player1 already in a move
			db.Where("game_id = ? and episode = ? and player_id = ?", gameID, game.Episode+1, playerID).First(&nextMove)
			if nextMove.ID > 0 {
				return Move{}, false, errors.New("already moved")
			} else if nextMove.ID == 0 {
				//create new move
				currentPos := curMove.EndPos
				nextPos := moveInMove(currentPos, step, rule)
				//record new move
				newEpisode := game.Episode + 1
				nextMove, err = rercordMove(playerID, gameID, currentPos, nextPos, newEpisode, step)
				if err != nil {
					return nextMove, false, err
				}
				if nextPos == 100 {
					game.Episode += 1
					now := time.Now()
					game.EndTime = &now
					game.Status = "END"
					db.Save(&game)
					return nextMove, true, nil
				}
			}
		case game.Player2:
			//check if player2 already in a move
			db.Where("game_id = ? and episode = ? and player_id = ?", gameID, game.Episode+1, game.Player1).First(&nextMove)
			if nextMove.ID > 0 {
				//create new move
				currentPos := curMove.EndPos
				nextPos := moveInMove(currentPos, step, rule)
				//record new move
				newEpisode := game.Episode + 1

				nextMove, err = rercordMove(playerID, gameID, currentPos, nextPos, newEpisode, step)
				if err != nil {
					return nextMove, false, err
				}
				if nextPos == 100 {
					now := time.Now()
					game.Episode = newEpisode
					game.EndTime = &now
					game.Status = "END"
					db.Save(&game)
					return nextMove, true, nil
				}
				//update game
				game.Episode = newEpisode
				db.Save(&game)
			} else if nextMove.ID == 0 {
				return Move{}, false, errors.New("player2 must wait player1 finished")
			}
		default:
			return Move{}, false, errors.New("player not in game")
		}
		return nextMove, false, nil
	} else {
		return Move{}, false, errors.New("game end")
	}

}

func moveInMove(currPos int, step int, rule Rule) int {
	nextStep := currPos + step
	if nextStep > 100 {
		return 200 - nextStep
	} else {
		if rule.Ladder[nextStep] > 0 {
			nextStep = rule.Ladder[nextStep]
		} else if rule.Snake[nextStep] > 0 {
			nextStep = rule.Snake[nextStep]
		}
		return nextStep
	}
}

func rercordMove(playerID uint, gameID uint, startPos int, endPos int, episode int, step int) (Move, error) {
	var newMove Move
	newMove = Move{
		PlayerID: playerID,
		GameID:   gameID,
		StartPos: startPos,
		EndPos:   endPos,
		Step:     step,
		Episode:  episode,
	}

	err = db.Create(&newMove).Error
	if err != nil {
		return newMove, err
	}

	return newMove, nil

}

func randomStep() int {
	rand.Seed(time.Now().UnixNano())
	max := 7
	min := 1
	return rand.Intn(max-min) + min
}
