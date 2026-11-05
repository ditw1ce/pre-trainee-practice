package accesslog

import (
	"slices"
	"strconv"
	"strings"
)

type Line struct {
	IP     string
	Method string
	Path   string
	Status int
	Bytes  int
}

type PathCount struct {
	Path  string
	Count int
}

type Summary struct {
	Parsed    int
	Skipped   int
	UniqueIPs int
	ByStatus  map[int]int
	TopPaths  []PathCount
}

func ParseLine(s string) (Line, bool) {
	s = strings.TrimSpace(s)
	if s == "" || strings.HasPrefix(s, "#") {
		return Line{}, false
	}

	fields := strings.Split(s, " ")
	if len(fields) != 5 || fields[1] == "" || fields[2] == "" || !strings.HasPrefix(fields[2], "/") {
		return Line{}, false
	}

	status, err := strconv.Atoi(fields[3])
	if err != nil {
		return Line{}, false
	}
	bytes, err := strconv.Atoi(fields[4])
	if err != nil {
		return Line{}, false
	}

	return Line{
		IP:     fields[0],
		Method: fields[1],
		Path:   fields[2],
		Status: status,
		Bytes:  bytes,
	}, true
}

func Summarize(text string, topN int) Summary {
	summary := Summary{
		ByStatus: make(map[int]int),
		TopPaths: []PathCount{},
	}
	if text == "" {
		return summary
	}

	uniqueIPs := make(map[string]struct{})
	pathCounts := make(map[string]int)
	lines := strings.Split(text, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	for _, rawLine := range lines {
		line, ok := ParseLine(rawLine)
		if !ok {
			summary.Skipped++
			continue
		}

		summary.Parsed++
		uniqueIPs[line.IP] = struct{}{}
		summary.ByStatus[line.Status]++
		pathCounts[line.Path]++
	}

	summary.UniqueIPs = len(uniqueIPs)
	if topN <= 0 {
		return summary
	}

	summary.TopPaths = make([]PathCount, 0, len(pathCounts))
	for path, count := range pathCounts {
		summary.TopPaths = append(summary.TopPaths, PathCount{Path: path, Count: count})
	}
	slices.SortFunc(summary.TopPaths, func(a, b PathCount) int {
		if a.Count != b.Count {
			if a.Count > b.Count {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Path, b.Path)
	})
	if topN < len(summary.TopPaths) {
		summary.TopPaths = summary.TopPaths[:topN]
	}

	return summary
}
