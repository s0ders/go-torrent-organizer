package organizer

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/s0ders/go-torrent-organizer/internal/parser"
)

var (
	mediaExtensions = []string{"mkv", "mp4"}
	logger          = slog.New(slog.NewJSONHandler(os.Stderr, nil))
)

// Organize takes a torrent and create a Kodi friendly folder
// then moves the torrent files inside the previously created
// folder.
func Organize(destinationPathRoot, contentPath string) error {

	torrentName := strings.TrimSpace(filepath.Base(contentPath))
	torrentInfo, err := parser.Parse(torrentName)
	if err != nil {
		return err
	}

	logger.Info("TORRENT", "infos", fmt.Sprintf("%+v", torrentInfo))

	contentIsMovie := torrentInfo.Season == 0
	contentIsDirectory := filepath.Ext(contentPath) == ""

	var organizedName string
	var organizedPath string

	if torrentInfo.Year != 0 {
		organizedName = fmt.Sprintf("%s (%d)", torrentInfo.Title, torrentInfo.Year)
	} else {
		organizedName = torrentInfo.Title
	}

	// Create Kodi folder
	if contentIsMovie {
		organizedPath = filepath.Join(destinationPathRoot, "Movies", organizedName)
	} else {
		season := fmt.Sprintf("Season %d", torrentInfo.Season)
		organizedPath = filepath.Join(destinationPathRoot, "TV Shows", organizedName, season)
	}

	if err := os.MkdirAll(organizedPath, 0750); err != nil {
		return err
	}

	if !contentIsDirectory {
		if err := moveFile(contentPath, organizedPath, contentIsMovie); err != nil {
			return err
		}
	}

	globPath := filepath.Join(contentPath, "*.mkv")
	globFiles, _ := filepath.Glob(globPath)

	logger.Info("FILES", "files", fmt.Sprintf("%v", globFiles))

	for _, file := range globFiles {
		// TODO: Handle when no info Season nor Episode, use index for episode and default to S01
		if err := moveFile(file, organizedPath, contentIsMovie); err != nil {
			return err
		}
	}

	return nil
}

func renameFile(name, ext string, movie bool, info *parser.TorrentInfo) string {
	if movie {
		return fmt.Sprintf("%s (%d)%s", info.Title, info.Year, ext)
	}

	return fmt.Sprintf("%s (%d) S%02dE%02d%s", info.Title, info.Year, info.Season, info.Episode, ext)
}

func moveFile(file, destination string, movie bool) error {
	ext := filepath.Ext(file)
	name := strings.ReplaceAll(filepath.Base(file), ext, "")

	info, err := parser.Parse(name)
	if err != nil {
		return err
	}

	newFileName := renameFile(name, ext, movie, info)
	newFilePath := filepath.Join(destination, newFileName)

	if err := os.Rename(file, newFilePath); err != nil {
		return err
	}

	return nil
}
