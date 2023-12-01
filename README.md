![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/s0ders/go-torrent-organizer) ![GitHub workflow status](https://github.com/s0ders/go-torrent-organizer/actions/workflows/go.yml/badge.svg) ![Go Report Card](http://goreportcard.com/badge/github.com/s0ders/go-torrent-organizer) ![License](https://img.shields.io/github/license/s0ders/go-torrent-organizer)

# Go Torrent Organizer

A CLI tool meant to be used a post-download hook of your favorite torrent client to automatically organize your downloaded movies or TV shows in a clean way.



## Usage

#### Parse

This command will return a JSON containing the title, year, season and episode — if found — of a given torrent name. If not found, the field will default to their default type value (e.g. `0` for integers and `""` for strings).

If the title is successfully found but the year is not present, the CLI will attempt to make an API call to [TMDB](https://www.themoviedb.org) to retrieve the first air date. In order for the call to happen, you must have a [developer API key](https://www.themoviedb.org/settings/api/new/form?type=developer) defined in your `~/.torrent-organizer` configuration file under the `tmdb` key. This configuration file must be in YAML language.

```bash
$ go-torrent-organizer parse [name]
$ go-torrent-parser parse "Jujutsu.Kaisen.S02E12.MULTi.1080p.WEB.x264-AMB3R"
$ {"Title":"Jujutsu Kaisen","Year":2020,"Season":2,"Episode":12}
```



#### Organize

This command can be executed manually or as a post-download hook of your favorite torrent client (usually in the "option" section).

```bash
$ go-torrent-organizer organize [destination] [origin]
$ go-torrent-organizer organize /ssd/ ~/downloads/Jujutsu.Kaisen.S02E12.MULTi.1080p.WEB.x264-AMB3R/
$ go-torrent-organizer organize /ssd/ ~/downloads/Blade.Runner.2049.(2017).1080p.BluRay.x264.Full/
```

The `destination` is the directory under which your movies and TV shows will be moved and organized whereas the `origin` is the folder under which the content your just downloaded was stored. 
Going on with the example above, the TV show and movie would be organized as bellow:

```
.
└── ssd/
    ├── TV Shows/
    │   └── Jujutsu Kaisen (2020)/
    │       └── Season 2/
    │           └── Jujutsu Kaisen (2020) S02E12.mkv
    └── Movies/
        └── Blade Runner 2049 (2017)/
            └── Blade Runner 2049 (2017).mkv
```



### Configuration

To make API calls to TMDB, the CLI expects to find an API developer key in the configuration file located at `~/.torrent-organizer` as bellow

```yaml
# ~/.torrent-organizer
tmdb: {API_KEY}
```

