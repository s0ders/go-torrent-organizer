package parser

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/s0ders/go-torrent-organizer/internal/tmdb"
)

var (
	PatternTitleYear    = regexp.MustCompile(`(.*)(19[0-9]{2}|20[0-9]{2})`)
	PatternTitleSeason  = regexp.MustCompile(`(.*)([Ss][0-9]{1,2})`)
	PatternYear         = regexp.MustCompile(`(19[0-9]{2}|20[0-9]{2})`)
	PatternSeason       = regexp.MustCompile(`[Ss]([0-9]{1,2})`)
	PatternEpisode      = regexp.MustCompile(`[Ss][0-9]{1,2}\s*[Ee]([0-9]{1,2})`)
	PatternDotSeparator = regexp.MustCompile(`\.{2,}`)
)

type TorrentInfo struct {
	Title   string
	Year    int
	Season  int
	Episode int
}

func Parse(s string) (*TorrentInfo, error) {

	s = strings.TrimSpace(s)

	title := title(s)

	year, err := matchPatternAndReturnInt(s, PatternYear)
	if err != nil {
		return nil, err
	}

	if year == 0 {
		if year, err = tmdb.QueryYear(title); err != nil {
			return nil, err
		}
	}

	season, err := matchPatternAndReturnInt(s, PatternSeason)
	if err != nil {
		return nil, err
	}

	episode, err := matchPatternAndReturnInt(s, PatternEpisode)
	if err != nil {
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

func ParseToJSON(s string) (string, error) {
	infos, err := Parse(s)
	if err != nil {
		return "", err
	}

	marshalled, err := json.Marshal(infos)
	if err != nil {
		return "", err
	}

	return string(marshalled), nil
}

func title(s string) string {
	var title string

	matches := PatternTitleYear.FindStringSubmatch(s)

	if len(matches) > 2 {
		title = matches[1]
	}

	if title == "" {
		matches = PatternTitleSeason.FindStringSubmatch(s)

		if len(matches) > 2 {
			title = matches[1]
		}
	}

	title = cleanTorrentName(title)
	return title
}

// matchPatternAndReturnInt tries to find the given pattern
// inside a given string and returns the last match capturing
// group converted to an integer, if no match are found,
// the given error is returned.
func matchPatternAndReturnInt(s string, pattern *regexp.Regexp) (int, error) {
	matches := pattern.FindAllStringSubmatch(s, -1)

	if len(matches) == 0 {
		return 0, nil
	}

	lastMatch := matches[len(matches)-1]

	if lastMatch[1] == "" {
		return 0, nil
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
