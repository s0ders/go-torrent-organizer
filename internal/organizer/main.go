package organizer

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	PTN "github.com/middelink/go-parse-torrent-name"
)

var (
	mediaExtensions = []string{"mkv", "mp4"}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
)

// Organize takes a torrent and create a Kodi friendly folder
// then moves the torrent files inside the previously created
// folder.
func Organize(destinationPath, contentPath string) error {

	torrentName := strings.TrimSpace(filepath.Base(contentPath))

	logger.Info("torrent name", "name", torrentName)

	torrentInfo, err := PTN.Parse(torrentName)
	if err != nil {
		return err
	}

	logger.Info("torrent infos", "infos", fmt.Sprintf("%+v", torrentInfo))

	isMovie := torrentInfo.Season == 0
	mediaName := fmt.Sprintf("%s (%d)", torrentInfo.Title, torrentInfo.Year)

	var mediaPath string

	// First, create Kodi folder
	if isMovie {
		mediaPath = filepath.Join(destinationPath, "Movies", mediaName)
	} else {
		seasonName := fmt.Sprintf("Season %d", torrentInfo.Season)

		showPath := filepath.Join(destinationPath, "TV Shows", mediaName)
		mediaPath = filepath.Join(showPath, seasonName)
	}
	
	if err := os.MkdirAll(mediaPath, 0750); err != nil {
		return err
	}

	// Second, move torrent file(s) into Kodi's folder
	filesGlob := filepath.Join(contentPath, "*.mkv")

	files, err := filepath.Glob(filesGlob)
	if err != nil {
		return err
	}

	logger.Info("files found", "slice", fmt.Sprintf("%v", files))

	for _, file := range files {
	
		// Handle when no info Season nor Episode, use index for episode and default to S01
		ext := filepath.Ext(file)
		name := strings.Replace(filepath.Base(file), ext, "", -1)

		info, err := PTN.Parse(name)
		if err != nil {
			return err
		}

		newFile := renameFile(name, ext, isMovie, info)
		newFilePath := filepath.Join(mediaPath, newFile)
		
		if err := os.Rename(file, newFilePath); err != nil {
			return err
		}

	}

	return nil
}


func renameFile(name, ext string, isMovie bool, info *PTN.TorrentInfo) string {
	if isMovie {
		return fmt.Sprintf("%s (%d)%s", info.Title, info.Year, ext)
	}

	return fmt.Sprintf("%s (%d) S%02dE%02d%s", info.Title, info.Year, info.Season, ext)
}