from collections import defaultdict


def tally(rows: list) -> list:
    # transform the row data into actual results
    results = defaultdict(lambda: defaultdict(int))
    for row in rows:
        team_a, team_b, outcome = row.split(";")

        results[team_a]["matches_played"] += 1
        results[team_b]["matches_played"] += 1

        if outcome == "win":
            results[team_a]["wins"] += 1
            results[team_b]["losses"] += 1
            results[team_a]["points"] += 3
        if outcome == "loss":
            results[team_b]["wins"] += 1
            results[team_a]["losses"] += 1
            results[team_b]["points"] += 3
        if outcome == "draw":
            results[team_a]["draws"] += 1
            results[team_a]["points"] += 1
            results[team_b]["draws"] += 1
            results[team_b]["points"] += 1

    # sort the results by points(descending) and alphabetically
    sorted_results = sorted(results, key=lambda team: (-results[team]["points"], team))

    # use the sorted results to generate the final output
    tally_table = ["{:<31}".format("Team") + "| MP |  W |  D |  L |  P"]
    for team in sorted_results:
        points = "|  {MP} |  {W} |  {D} |  {L} |  {P}".format(
            MP=results[team]["matches_played"],
            W=results[team]["wins"],
            D=results[team]["draws"],
            L=results[team]["losses"],
            P=results[team]["points"],
        )
        tally_table.append("{:<31}".format(team) + points)

    return tally_table
