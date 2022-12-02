package main

import (
	"os"
	"strings"
)

type RockPaperScissor int8

const (
	Rock RockPaperScissor = iota
	Paper
	Scissor
)

type Pair struct {
	opponent RockPaperScissor
	mine     string
}

var (
	rockPaperScissorMap = map[string]RockPaperScissor{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
	}
)

var (
	rockPaperScissorScore = map[RockPaperScissor]int{
		Rock:    1,
		Paper:   2,
		Scissor: 3,
	}
)

func main() {
	input, _ := os.ReadFile("input.txt")

	rounds := []Pair{}
	for _, round := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		x := strings.Split(round, " ")
		rounds = append(rounds, Pair{
			opponent: rockPaperScissorMap[x[0]],
			mine:     x[1],
		})
	}

	score := play(rounds, map[string]RockPaperScissor{
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	})

	println("Solution 1:", score)

	score = playWinLoose(rounds)
	println("Solution 2:", score)
}

func playWinLoose(rounds []Pair) int {
	total := 0
	for _, round := range rounds {
		whatIPlay := winOrLooseMapping[round.mine](round.opponent)
		total += rockPaperScissorScore[whatIPlay]
		total += score(round.opponent, whatIPlay)
	}
	return total
}

var (
	winOrLooseMapping = map[string]func(RockPaperScissor) RockPaperScissor{
		"X": loose,
		"Y": draw,
		"Z": win,
	}
)

func play(rounds []Pair, mapping map[string]RockPaperScissor) int {
	total := 0
	for _, round := range rounds {
		mine := mapping[round.mine]
		total += rockPaperScissorScore[mine]
		total += score(round.opponent, mine)
	}
	return total
}

func score(a RockPaperScissor, b RockPaperScissor) int {
	if a == b {
		// Draw
		return 3
	}

	if a == Rock {
		if b == Paper {
			return 6
		} else {
			return 0
		}
	} else if a == Scissor {
		if b == Rock {
			return 6
		} else {
			return 0
		}
	} else {
		if b == Scissor {
			return 6
		} else {
			return 0
		}
	}
}

func win(a RockPaperScissor) RockPaperScissor {
	if a == Rock {
		return Paper
	} else if a == Paper {
		return Scissor
	} else {
		return Rock
	}
}

func loose(a RockPaperScissor) RockPaperScissor {
	if a == Rock {
		return Scissor
	} else if a == Paper {
		return Rock
	} else {
		return Paper
	}
}

func draw(a RockPaperScissor) RockPaperScissor {
	return a
}
