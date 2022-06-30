package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type word struct {
	word  string
	count int
}

var (
	re0 = regexp.MustCompile(`["!?,.:;']`)
	re1 = regexp.MustCompile(`( - )|(- )|(-$)`)
	re2 = regexp.MustCompile(`["]*([а-яА-Яa-zA-Z]+)([",.:!?;'])*\s+`)
)

func prepare(s string) string {
	s = re0.ReplaceAllString(s, " ")
	s = re1.ReplaceAllString(s, " ")
	s = re2.ReplaceAllString(s, "$1 ")
	s = strings.ToLower(s)

	return s
}

func Top10(s string) []string {
	output := []string{}
	preparedStr := prepare(s)

	fields := strings.Fields(preparedStr)

	if len(fields) == 0 {
		return nil
	}

	words := map[string]int{}

	for _, v := range fields {
		words[v]++
	}

	w := make([]word, 0, len(words))

	for k, v := range words {
		w = append(w, word{word: k, count: v})
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i].count > w[j].count || ((w[i].count == w[j].count) && (w[i].word < w[j].word))
	})

	if len(w) > 10 {
		w = w[:10]
	}

	for _, v := range w {
		output = append(output, v.word)
	}

	return output
}
