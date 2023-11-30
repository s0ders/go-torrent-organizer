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
		if got, _ := Title(tt.name); got != tt.expected {
			t.Errorf("got: %s, want: %s", got, tt.expected)
		}
	}
}