package scraper

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Source1337x struct {
	BaseUrl  string
	Torrents []Torrent
}

func Make1337xSource() *Source1337x {
	return &Source1337x{
		BaseUrl: "https://1337x.to",
	}
}

func (s *Source1337x) Name() string {
	return "1337x"
}

func (s *Source1337x) Query(query string) string {
	return s.BaseUrl + "/search/" + url.QueryEscape(query) + "/1/"
}

func (s *Source1337x) Parse(c *colly.Collector) {
	c.OnHTML("table tbody tr td.name", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a:nth-child(2)", "href")
		e.Request.Visit(link)
	})

	c.OnHTML(".torrent-detail-page", func(e *colly.HTMLElement) {
		size := e.ChildText("ul:nth-child(2) li:nth-child(4) span")
		seedersString := e.ChildText(".seeds")
		seeders, err := strconv.Atoi(seedersString)

		if err != nil {
			log.Println("Error converting seeders to int:", seedersString, err)
			return
		}

		leechesString := e.ChildText(".leeches")
		leeches, err := strconv.Atoi(leechesString)

		if err != nil {
			log.Println("Error converting leeches to int:", leechesString, err)
			return
		}

		torrent := Torrent{
			Name:       e.ChildText(".box-info-heading h1"),
			Link :      e.Request.URL.String(),
			Magnet:     e.ChildAttr("div div ul:nth-child(1) li:nth-child(1) a", "href"),
			Seeders:    seeders,
			Leechers:   leeches,
			Size:       size,
			Source:     s.Name(),
			UploadDate: e.ChildText("div div ul:nth-child(3) li:nth-child(3) span"),
		}

		s.Torrents = append(s.Torrents, torrent)
	})
}

func (s *Source1337x) Result() []Torrent {
	return s.Torrents
}
