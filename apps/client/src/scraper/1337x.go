package scraper

import (
	"net/url"

	"github.com/gocolly/colly/v2"
)

type Source1337x struct {
	Source
	BaseUrl   string
	Torrents  []Torrent
	Collector *colly.Collector
}

func Make1337xSource() *Source1337x {
	return &Source1337x{
		BaseUrl: "https://1337x.to",
	}
}

func (s *Source1337x) Name() string {
	return "1337x"
}

func (s *Source1337x) RequiresCollector() bool {
	return true
}

func (s *Source1337x) SetCollector(c *colly.Collector) {
	s.Collector = c
}

func (s *Source1337x) Query(query string) string {
	return s.BaseUrl + "/search/" + url.QueryEscape(query) + "/1/"
}

func (s *Source1337x) Parse(url string) {
	s.Collector.OnHTML("table tbody tr td.name", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a:nth-child(2)", "href")
		e.Request.Visit(link)
	})

	s.Collector.OnHTML(".torrent-detail-page", func(e *colly.HTMLElement) {
		torrent := Torrent{
			Name:       e.ChildText(".box-info-heading h1"),
			Link:       e.Request.URL.String(),
			Magnet:     e.ChildAttr("div div ul:nth-child(1) li:nth-child(1) a", "href"),
			Seeders:    DefaultInt(e.ChildText(".seeds"), 0),
			Leechers:   DefaultInt(e.ChildText(".leeches"), 0),
			Size:       e.ChildText("ul:nth-child(2) li:nth-child(4) span"),
			Source:     s.Name(),
			UploadDate: e.ChildText("div div ul:nth-child(3) li:nth-child(3) span"),
		}

		s.Torrents = append(s.Torrents, torrent)
	})
}

func (s *Source1337x) Result() []Torrent {
	return s.Torrents
}
