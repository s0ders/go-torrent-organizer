package parser

import (
	"testing"
)

func TestCleanTorrentName(t *testing.T) {
	type test struct {
		name     string
		expected string
	}

	tests := []test{
		{"Rick.and.Morty.", "Rick and Morty"},
		{"Grrr...", "Grrr..."},
		{" Jujutsu Kaisen ", "Jujutsu Kaisen"},
	}

	for _, tt := range tests {
		if got := cleanTorrentName(tt.name); got != tt.expected {
			t.Errorf("got: %s, want: %s", got, tt.expected)
		}
	}
}

func TestTitle(t *testing.T) {
	type test struct {
		name     string
		expected string
	}

	tests := []test{
		{"Blade Runner 2049 2017 + Bonus BR EAC3 VFF VFQ VO 1080p x265 10Bits T0M (Blade Runner 2)", "Blade Runner 2049"},
		{"Club.Soly.2021.S03.FRENCH.1080p.WEB.DD5.1.H264-TFA", "Club Soly"},
		{"1917.2019.2160p.TWN.UHD.Blu-ray.HEVC.TrueHD.7.1-nLiBRA", "1917"},
		{"Jujutsu.Kaisen.S01.MULTi.COMPLETE.BLURAY-DRAGONFEU", "Jujutsu Kaisen"},
	}

	for _, tt := range tests {
		if got := title(tt.name); got != tt.expected {
			t.Errorf("got: %s, want: %s", got, tt.expected)
		}
	}
}

func TestYear(t *testing.T) {
	type test struct {
		name     string
		expected int
	}

	tests := []test{
		{"Blade Runner 2049 2017 + Bonus BR EAC3 VFF VFQ VO 1080p x265 10Bits T0M (Blade Runner 2)", 2017},
		{"Club.Soly.2021.S03.FRENCH.1080p.WEB.DD5.1.H264-TFA", 2021},
		{"1917.2019.2160p.TWN.UHD.Blu-ray.HEVC.TrueHD.7.1-nLiBRA", 2019},
		{"Escape.From.New.York.1981.VFi.1080P.mHD.X264.AC3-ROMKENT", 1981},
		{"Jujutsu.Kaisen.S01.MULTi.COMPLETE.BLURAY-DRAGONFEU", 0},
	}

	for _, tt := range tests {
		if got, _ := matchPatternAndReturnInt(tt.name, PatternYear); got != tt.expected {
			t.Errorf("got: %d, want: %d", got, tt.expected)
		}
	}
}

func TestSeason(t *testing.T) {
	type test struct {
		name     string
		expected int
	}

	tests := []test{
		{"Blade Runner 2049 2017 + Bonus BR EAC3 VFF VFQ VO 1080p x265 10Bits T0M (Blade Runner 2)", 0},
		{"Club.Soly.2021.S03.FRENCH.1080p.WEB.DD5.1.H264-TFA", 3},
		{"1917.2019.2160p.TWN.UHD.Blu-ray.HEVC.TrueHD.7.1-nLiBRA", 0},
		{"Jujutsu.Kaisen.S01.MULTi.COMPLETE.BLURAY-DRAGONFEU", 1},
		{"Game.of.Thrones.S05.MULTI.Extra.Bonus.1080p.Bluray.HEVC.AAC.5.1.x265 ", 5},
		{" 	NCIS.2003.S12.MULTi.1080p.AMZN.WEB-DL.DDP5.1.H264-BONBON", 12},
	}

	for _, tt := range tests {
		if got, _ := matchPatternAndReturnInt(tt.name, PatternSeason); got != tt.expected {
			t.Errorf("got: %d, want: %d", got, tt.expected)
		}
	}
}

func TestEpisode(t *testing.T) {
	type test struct {
		name     string
		expected int
	}

	tests := []test{
		{"Jujutsu Kaisen S01 E13 VOSTFR WebRIP 1080P X264-BLV", 13},
		{"Rick.and.Morty.S07E03.1080p.WEB.H264-NHTFS[TGx]", 3},
		{"South Park S26E12 720p HDTV x264-SYNCOPY", 12},
		{"Jujutsu.Kaisen.S01.MULTi.COMPLETE.BLURAY-DRAGONFEU", 0},
		{"1917.2019.2160p.TWN.UHD.Blu-ray.HEVC.TrueHD.7.1-nLiBRA", 0},
	}

	for _, tt := range tests {
		if got, _ := matchPatternAndReturnInt(tt.name, PatternEpisode); got != tt.expected {
			t.Errorf("got: %d, want: %d", got, tt.expected)
		}
	}
}

func TestParse(t *testing.T) {

	type test struct {
		name  string
		infos *TorrentInfo
	}

	tests := []test{
		{"Jujutsu Kaisen S01 E13 VOSTFR WebRIP 1080P X264-BLV", &TorrentInfo{"Jujutsu Kaisen", 0, 1, 13}},
		{"South Park S26E12 720p HDTV x264-SYNCOPY", &TorrentInfo{"South Park", 0, 26, 12}},
		{"Blade Runner 2049 2017 + Bonus BR EAC3 VFF VFQ VO 1080p x265 10Bits T0M (Blade Runner 2)", &TorrentInfo{"Blade Runner 2049", 2017, 0, 0}},
		{"1917.2019.2160p.TWN.UHD.Blu-ray.HEVC.TrueHD.7.1-nLiBRA", &TorrentInfo{"1917", 2019, 0, 0}},
		{"Club.Soly.2021.S03.FRENCH.1080p.WEB.DD5.1.H264-TFA", &TorrentInfo{"Club Soly", 2021, 3, 0}},
	}

	for _, tt := range tests {
		if got, _ := Parse(tt.name); *got != *tt.infos {
			t.Errorf("got: %+v, want: %+v", got, tt.infos)
		}
	}
}
