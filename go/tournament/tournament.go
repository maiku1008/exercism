package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

const (
	top   = "Team                           | MP |  W |  D |  L |  P\n"
	table = "|  %d |  %d |  %d |  %d |  %d\n"
)

type match struct {
	team1   string
	team2   string
	outcome string
}

type standings struct {
	team   string
	played int
	won    int
	draw   int
	lost   int
	points int
}

// Tally tallies the results of a football tournament
func Tally(r io.Reader, w io.Writer) error {
	// Read as bytes
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// Process the results
	var tally []standings
	tally, err = processResults(b)
	if err != nil {
		return err
	}

	// Build the string
	output := top
	for _, res := range tally {
		teamName := res.team + strings.Repeat(" ", 31-len(res.team))
		points := fmt.Sprintf(
			table, res.played, res.won, res.draw, res.lost, res.points,
		)
		output += teamName + points
	}

	// Write the output to w
	_, err = fmt.Fprint(w, output)
	if err != nil {
		return err
	}

	return nil
}

// helper function that processes the input and returns a sorted slice of standings
func processResults(input []byte) ([]standings, error) {
	var tallyMap = make(map[string]*standings)

	// Create results slice
	games := strings.Split(string(input), "\n")
	for _, game := range games {
		// Ignore empty or commented lines
		if len(game) == 0 || strings.HasPrefix(game, "#") {
			continue
		}

		// Split game string into slice
		g := strings.Split(game, ";")
		if len(g) != 3 {
			return nil, fmt.Errorf("invalid input: %s", game)
		}

		// Init match to name each element instead of just using index
		m := match{g[0], g[1], g[2]}

		// Ensure teams are in tallyMap
		if _, ok := tallyMap[m.team1]; !ok {
			tallyMap[m.team1] = &standings{}
			tallyMap[m.team1].team = m.team1
		}

		if _, ok := tallyMap[m.team2]; !ok {
			tallyMap[m.team2] = &standings{}
			tallyMap[m.team2].team = m.team2
		}

		// Award points based on standings and save them in tallyMap
		err := award(m.team1, m.team2, m.outcome, tallyMap)
		if err != nil {
			return nil, err
		}
	}

	// extract the standings to sort them afterwards
	tally := make([]standings, len(tallyMap))
	i := 0
	for _, v := range tallyMap {
		tally[i] = *v
		i++
	}

	// Sort alphabetically
	sort.Slice(tally, func(i, j int) bool {
		return tally[i].team < tally[j].team
	})

	// Sort by points
	sort.Slice(tally, func(i, j int) bool {
		return tally[i].points > tally[j].points
	})

	return tally, nil
}

// Award points according to match outcome
func award(team1, team2, outcome string, tm map[string]*standings) error {
	// Matches played
	tm[team1].played++
	tm[team2].played++

	switch outcome {
	case "win":
		tm[team1].points += 3
		tm[team1].won++
		tm[team2].lost++
	case "loss":
		tm[team2].points += 3
		tm[team2].won++
		tm[team1].lost++
	case "draw":
		tm[team1].points++
		tm[team1].draw++
		tm[team2].points++
		tm[team2].draw++
	default:
		return errors.New("incorrect outcome")
	}
	return nil
}
