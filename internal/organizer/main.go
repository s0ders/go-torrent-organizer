package organizer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	PTN "github.com/middelink/go-parse-torrent-name"
)

var (
	mediaExtensions = []string{"mkv", "mp4"}
)

// Organize takes a torrent and create a Kodi friendly folder
// then moves the torrent files inside the previously created
// folder.
func Organize(destinationPath, contentPath, name string) error {

	name = strings.TrimSpace(name)

	info, err := PTN.Parse(name)
	if err != nil {
		return err
	}

	isMovie := info.Season == 0
	mediaName := fmt.Sprintf("%s (%d)", info.Title, info.Year)

	var mediaPath string

	// First, create Kodi folder
	if isMovie == 0 {
		mediaPath = filepath.Join(destinationPath, "Movies", mediaName)
	} else {
		seasonName := fmt.Sprintf("Season %d", info.Season)

		showPath := filepath.Join(destinationPath, "TV Shows", mediaName)
		mediaPath = filepath.Join(showPath, seasonName)
	}
	
	if err := os.MkdirAll(mediaPath, 0750); err != nil {
		return err
	}

	// Second, move torrent file(s) into Kodi's folder
	medias, err := filepath.Glob(fmt.Sprintf("%s/*.[%s]", contentPath, strings.Join(mediaExtensions, "|")))
	if err != nil {
		return err
	}

	fmt.Println(medias)

	for _, media := range medias {
		



		if err := os.Rename(media, )
	}

	return nil
}


func renameMedia(name, ext string, isMovie bool, info *PTN.TorrentInfo) string {
	if isMovie {
		return fmt.Sprintf("%s (%d).%s", info.Title, info.Year, ext)
	}

	return fmt.Sprintf("%s (%d) S%02dE%02d.%s", info.Title, info.Year, info.Season, ext)
}