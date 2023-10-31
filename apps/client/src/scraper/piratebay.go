package scraper

import (
	"log"
	"net/url"

	"github.com/gocolly/colly/v2"
)

type SourcePiratebay struct {
	Source
	BaseUrl   string
	Torrents  []Torrent
	Collector *colly.Collector
}

type PiratebaySearchResponse []struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	InfoHash string `json:"info_hash"`
	Leechers string `json:"leechers"`
	Seeders  string `json:"seeders"`
	NumFiles string `json:"num_files"`
	Size     string `json:"size"`
	Username string `json:"username"`
	Added    string `json:"added"`
	Status   string `json:"status"`
	Category string `json:"category"`
	Imdb     string `json:"imdb"`
}

func MakePiratebaySource() *SourcePiratebay {
	return &SourcePiratebay{
		BaseUrl: "https://apibay.org",
	}
}

func (s *SourcePiratebay) RequiresCollector() bool {
	return false
}

func (s *SourcePiratebay) SetCollector(c *colly.Collector) {
	s.Collector = c
}

func (s *SourcePiratebay) Name() string {
	return "ThePirateBay"
}

func (s *SourcePiratebay) Query(query string) string {
	return s.BaseUrl + "/q.php?q=" + url.QueryEscape(query)
}

func (s *SourcePiratebay) CreateMagnet(hash string, name string) string {
	return "magnet:?xt=urn:btih:" + hash + "&dn=" + url.QueryEscape(name) + s.GetTrackers()
}

func (s *SourcePiratebay) GetTrackers() string {
	tr := "&tr=" + url.QueryEscape("udp://tracker.coppersurfer.tk:6969/announce")
	tr += "&tr=" + url.QueryEscape("udp://tracker.openbittorrent.com:6969/announce")
	tr += "&tr=" + url.QueryEscape("udp://tracker.bittor.pw:1337/announce")
	tr += "&tr=" + url.QueryEscape("udp://tracker.opentrackr.org:1337")
	tr += "&tr=" + url.QueryEscape("udp://bt.xxx-tracker.com:2710/announce")
	tr += "&tr=" + url.QueryEscape("udp://public.popcorn-tracker.org:6969/announce")
	tr += "&tr=" + url.QueryEscape("udp://eddie4.nl:6969/announce")
	tr += "&tr=" + url.QueryEscape("udp://tracker.torrent.eu.org:451/announce")
	tr += "&tr=" + url.QueryEscape("udp://p4p.arenabg.com:1337/announce")
	tr += "&tr=" + url.QueryEscape("udp://tracker.tiny-vps.com:6969/announce")
	tr += "&tr=" + url.QueryEscape("udp://open.stealth.si:80/announce")
	return tr
}

func (s *SourcePiratebay) Parse(url string) {
	var results PiratebaySearchResponse

	if err := GetJson(url, &results); err != nil {
		log.Println("Failed to get json from", url, err)
		return
	}

	for _, result := range results {
		if result.ID == "0" {
			continue
		}

		s.Torrents = append(s.Torrents, Torrent{
			Name:       result.Name,
			Link:       "https://thepiratebay.org/torrent/" + result.ID,
			Magnet:     s.CreateMagnet(result.InfoHash, result.Name),
			Seeders:    DefaultInt(result.Seeders, 0),
			Leechers:   DefaultInt(result.Leechers, 0),
			Size:       DefaultBytes(result.Size, 0),
			Source:     s.Name(),
			UploadDate: TimestampToString(result.Added),
		})
	}
}

func (s *SourcePiratebay) Result() []Torrent {
	return s.Torrents
}
