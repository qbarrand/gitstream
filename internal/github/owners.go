package github

import "math/rand"

var reviewers = map[string]string{
	"Brett Thurber":         "bthurber",
	"Chris Procter":         "chr15p",
	"Enrique Belarte Luque": "enriquebelarte",
	"Ezequiel Russo":        "erusso7",
	"Fabien Dupont":         "fabiendupont",
	"Michail Resvanis":      "mresvanis",
	"Pablo Iranzo Gómez":    "iranzo",
	"Quentin Barrand":       "qbarrand",
	"Yevgeny Shnaidman":     "yevgeny-shnaidman",
	"Yoni Bettan":           "ybettan",
}

var DefaultReviewers = []string{
	"chr15p",
	"enriquebelarte",
	"mresvanis",
	"iranzo",
	"qbarrand",
	"yevgeny-shnaidman",
	"ybettan",
}

func GetAssignee(commitAuthor string) string {

	var res string

	if assignee, ok := reviewers[commitAuthor]; ok {
		res = assignee
	} else {
		randomIndex := rand.Intn(len(DefaultReviewers))
		res = DefaultReviewers[randomIndex]
	}

	return res
}
