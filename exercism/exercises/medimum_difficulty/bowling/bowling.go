package main

// https://exercism.org/tracks/go/exercises/bowling

import "errors"

type RoundResult int8

const (
	Undefined RoundResult = -1
	Open                  = 0
	Spare                 = 1
	Strike                = 2
)

type Round struct {
	Result      RoundResult
	ThrowsAfter int
}

func (r *Round) Add(pins int) int {
	toAdd := 0
	r.ThrowsAfter++
	if int(r.Result) >= r.ThrowsAfter {
		toAdd += pins
	}
	return toAdd
}

type Game struct {
	rolls           int
	score           int
	throw           int
	roundScore      int
	currRound       int
	oneButLastRound Round
	lastRound       Round
	bonusRound      bool
	bonusThrows     int
	finished        bool
}

func NewGame() *Game {
	g := Game{}
	g.finished = false
	g.bonusRound = false
	g.oneButLastRound = Round{Undefined, 0}
	g.lastRound = Round{Undefined, 0}
	return &Game{}
}

func (g *Game) Roll(pins int) error {
	g.rolls++
	if g.finished {
		return errors.New("end of the game")
	} else if pins < 0 {
		return errors.New("negative amount of pins")
	}

	g.roundScore += pins
	if g.roundScore > 10 {
		return errors.New("too many pins")
	}

	if !g.bonusRound {
		g.score += pins + g.oneButLastRound.Add(pins) + g.lastRound.Add(pins)
	} else {
		g.score += g.oneButLastRound.Add(pins) + g.lastRound.Add(pins)
		if g.roundScore == 10 {
			g.roundScore = 0
		}
		g.bonusThrows--
		if g.bonusThrows == 0 {
			g.finished = true
		}
		return nil
	}

	g.throw++

	if g.roundScore == 10 {
		g.lastRound = g.oneButLastRound
		if g.currRound == 9 {
			g.roundScore = 0
			if g.throw == 1 {
				g.bonusRound = true
				g.bonusThrows = 2
				g.oneButLastRound = Round{Strike, 0}
			} else if g.throw == 2 {
				g.bonusRound = true
				g.bonusThrows = 1
				g.oneButLastRound = Round{Spare, 0}
			} else {
				g.bonusRound = false
				g.currRound++
			}
		} else {
			if g.throw == 1 {
				g.oneButLastRound = Round{Strike, 0}
			} else if g.throw == 2 {
				g.oneButLastRound = Round{Spare, 0}
			}

			g.throw = 0
			g.roundScore = 0
			g.currRound++
		}

	} else if g.throw >= 2 {
		g.lastRound = g.oneButLastRound
		g.oneButLastRound = Round{Open, 0}
		g.currRound++
		g.throw = 0
		g.roundScore = 0
	}

	if g.currRound == 10 {
		g.finished = true
	}

	return nil
}

func (g *Game) Score() (int, error) {
	if g.finished {
		return g.score, nil
	} else {
		return 0, errors.New("game not finished yet")
	}
}

func main() {
}
