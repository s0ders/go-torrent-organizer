package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Data struct {
	Results []Result `json:"results"`
}

type Result struct {
	Date string `json:"first_air_date"`
}

var (
	tmdbSearchURL = "https://api.themoviedb.org/3/search/tv?query={show}&include_adult=true&language=en-US&page=1"
)

func QueryYear(s string) (year int, err error) {

	show := url.QueryEscape(s)
	query := strings.ReplaceAll(tmdbSearchURL, "{show}", show)

	req, _ := http.NewRequest("GET", query, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("tmdb")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer func() {
		err = res.Body.Close()
	}()

	data := &Data{
		Results: []Result{},
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	if len(data.Results) > 0 {
		releaseDate := data.Results[0].Date

		t, err := time.Parse("2006-01-02", releaseDate)
		if err != nil {
			return 0, err
		}

		year = t.Year()
		return year, nil
	}

	return
}
