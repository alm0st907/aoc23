package Day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gameRE = regexp.MustCompile(`Game (\d+): (.+)`)

type rgb struct {
	r int
	g int
	b int
}

func (r rgb) valid(maxR, maxG, maxB int) bool {
	return r.r <= maxR && r.g <= maxG && r.b <= maxB
}

type game struct {
	ID     int
	rounds []rgb
}

func (g game) maxvals() (result rgb) {
	for _, r := range g.rounds {
		if r.r > result.r {
			result.r = r.r
		}
		if r.g > result.g {
			result.g = r.g
		}
		if r.b > result.b {
			result.b = r.b
		}
	}
	return result
}

func parseRound(rstr string) rgb {
	round := rgb{}
	for _, r := range strings.Split(strings.TrimSpace(rstr), ",") {
		r = strings.TrimSpace(r)
		pair := strings.Split(r, " ")
		val, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		switch pair[1] {
		case "red":
			round.r = val
		case "green":
			round.g = val
		case "blue":
			round.b = val
		}
	}

	return round
}

func parseRounds(rstr string) (rounds []rgb) {
	rs := strings.Split(rstr, ";")
	for _, r := range rs {
		rounds = append(rounds, parseRound(r))
	}
	return rounds
}

func parseGame(line string) game {
	matches := gameRE.FindStringSubmatch(line)
	if len(matches) == 0 {
		panic("Invalid input, check text file for excess whitespace")
	}
	ID, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	rounds := parseRounds(matches[2])
	return game{ID: ID, rounds: rounds}
}

func part1(lines []string) (res int) {
	maxR, maxG, maxB := 12, 13, 14
	result := 0

	//this is a label being used to break a nested loop - I wonder if there is a better way to do this?
GAMELOOP:
	for _, line := range lines {
		g := parseGame(line)
		for _, r := range g.rounds {
			if !r.valid(maxR, maxG, maxB) {
				continue GAMELOOP
			}
		}
		result += g.ID
	}
	fmt.Println("Part 1:", result)
	return result
}

func part2(lines []string) (res int) {
	result := 0
	for _, line := range lines {
		g := parseGame(line)
		m := g.maxvals()
		result += m.r * m.g * m.b
	}
	fmt.Println("Part 2:", result)
	return result
}

func Day2Part1Driver(path string) (result int) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return part1(lines)
}

func Day2Part2Driver(path string) (result int) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return part2(lines)
}

func RunDay2(path string) {
	Day2Part1Driver(path)
	Day2Part2Driver(path)
}
