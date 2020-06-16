package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

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
	var err error
	var tallyMap = make(map[string]standings)

	scanner := bufio.NewScanner(r)
	// Read each line one at a time
	for scanner.Scan() {
		game := scanner.Text()
		// Ignore empty or commented lines
		if len(game) == 0 || strings.HasPrefix(game, "#") {
			continue
		}
		g := strings.Split(game, ";")
		if len(g) != 3 {
			return fmt.Errorf("input should be in the form team1;team2;outcome. got: %s", game)
		}
		team1, team2, outcome := g[0], g[1], g[2]
		a := tallyMap[team1]
		b := tallyMap[team2]

		a.team = team1
		b.team = team2
		// Matches played
		a.played++
		b.played++

		switch outcome {
		case "win":
			a.points += 3
			a.won++
			b.lost++
		case "loss":
			b.points += 3
			b.won++
			a.lost++
		case "draw":
			a.points++
			a.draw++
			b.points++
			b.draw++
		default:
			return fmt.Errorf("outcome should be one of win,loss,draw. got: %s", outcome)
		}
		tallyMap[team1] = a
		tallyMap[team2] = b
	}

	// Extract the standings into a slice
	tally := make([]standings, len(tallyMap))
	i := 0
	for _, v := range tallyMap {
		tally[i] = v
		i++
	}
	// Sort
	sort.Slice(tally, func(i, j int) bool {
		if tally[i].points == tally[j].points {
			return tally[i].team < tally[j].team
		}
		return tally[i].points > tally[j].points
	})
	// Format string and send it to writer
	_, err = fmt.Fprint(w, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	var table = "|  %d |  %d |  %d |  %d |  %d\n"
	for _, res := range tally {
		_, err = fmt.Fprintf(
			w, res.team+strings.Repeat(" ", 31-len(res.team))+table,
			res.played, res.won, res.draw, res.lost, res.points,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
