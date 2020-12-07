package days

import (
	"github.com/lavinj87/advent/internal/lib"
)

func countAnswers(answers []string) int {
	distinctAnswers := map[rune]bool{}
	for _, individualAnswers := range answers {
		for _, answer := range individualAnswers {
			distinctAnswers[answer] = true
		}
	}
	return len(distinctAnswers)
}

func countAnswersFromFile(inputPath string) int {
	lines := []string{}
	lib.Lines(inputPath, func(line string) {
		lines = append(lines, line)
	})
	count := 0
	for _, answerLines := range lib.Partition(lines, func(line string) bool { return line == "" }) {
		count += countAnswers(lib.FilterEmpty(answerLines))
	}
	return count
}
