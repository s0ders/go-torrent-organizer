package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	PatternTitle        = regexp.MustCompile(`(.*)((19[0-9]{2}|20[0-9]{2})|([S|s]([0-9]{1,2})))`)
	PatternYear         = regexp.MustCompile(`(19[0-9]{2}|20[0-9]{2})`)
	PatternSeason       = regexp.MustCompile(`[Ss]([0-9]{1,2})`)
	PatternEpisode      = regexp.MustCompile(`[Ss][0-9]{1,2}[Ee]([0-9]{1,2})`)
	PatternDotSeparator = regexp.MustCompile(`\.{2,}`)

	ErrNoTitle   = errors.New("title not found")
	ErrNoYear    = errors.New("year not found")
	ErrNoSeason  = errors.New("season not found")
	ErrNoEpisode = errors.New("episode not found")
)

type TorrentInfo struct {
	Title   string
	Year    int
	Season  int
	Episode int
}

func Parse(s string) (*TorrentInfo, error) {

	title, err := Title(s)
	if err != nil {
		return nil, err
	}

	year, err := MatchPatternAndReturnInt(s, PatternYear, ErrNoYear)
	if err != nil && !errors.Is(err, ErrNoYear) {
		return nil, err
	}

	season, err := MatchPatternAndReturnInt(s, PatternSeason, ErrNoSeason)
	if err != nil && !errors.Is(err, ErrNoSeason) {
		return nil, err
	}

	episode, err := MatchPatternAndReturnInt(s, PatternEpisode, ErrNoEpisode)
	if err != nil && !errors.Is(err, ErrNoEpisode) {
		return nil, err
	}

	infos := &TorrentInfo{
		Title:   title,
		Year:    year,
		Season:  season,
		Episode: episode,
	}

	return infos, nil
}

func Title(s string) (string, error) {
	matches := PatternTitle.FindStringSubmatch(s)

	title := matches[1]

	if title == "" {
		return "", ErrNoTitle
	}

	title = cleanTorrentName(title)
	return title, nil
}

// MatchPatternAndReturnInt tries to find the given pattern
// inside a given string and returns the last match capturing
// group converted to an integer, if no match are found,
// the given error is returned.
func MatchPatternAndReturnInt(s string, pattern *regexp.Regexp, notFoundError error) (int, error) {
	matches := pattern.FindAllStringSubmatch(s, -1)

	lastMatch := matches[len(matches)-1]

	if lastMatch[1] == "" {
		return 0, notFoundError
	}

	target, err := strconv.Atoi(lastMatch[1])
	if err != nil {
		return 0, err
	}

	return target, nil
}

func cleanTorrentName(s string) string {

	dots := PatternDotSeparator.FindString(s)
	replacer := strings.Repeat("_", len(dots))

	s = PatternDotSeparator.ReplaceAllLiteralString(s, replacer)
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, "_", ".")
	s = strings.TrimSpace(s)

	return s
}
