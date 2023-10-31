package scraper

import (
	"log"
	"time"
	"sort"

	"github.com/gocolly/colly/v2"
)

type Source interface {
	Name() string
	Parse(*colly.Collector)
	Query(string) string
	Result() []Torrent
}

type Torrent struct {
	Name       string `json:"name"`
	Link 		 string `json:"link"`
	Magnet     string `json:"magnet"`
	Seeders    int    `json:"seeders"`
	Leechers   int    `json:"leechers"`
	Size       string `json:"size"`
	UploadDate string `json:"uploadDate"`
	Source     string `json:"source"`
}

type Scraper struct {
}

func MakeScraper() *Scraper {
	scraper := &Scraper{}

	return scraper
}

func (s *Scraper) GetSources() []Source {
	return[]Source{
		Make1337xSource(),
	}
}

func (s *Scraper) Sort(torrents []Torrent) []Torrent {
	sort.Slice(torrents, func(i, j int) bool {
		return torrents[i].Seeders > torrents[j].Seeders
	})

	return torrents
}

func (s *Scraper) Query(query string) ([]Torrent, error) {
	var torrents []Torrent
	sources := s.GetSources()

	for _, source := range sources {
		c := colly.NewCollector(
			colly.MaxDepth(2),
			colly.Async(),
		)

		c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 6})
		startTime := time.Now()
		source.Parse(c)
		path := source.Query(query)
		c.Visit(path)
		c.Wait()
		torrents = append(torrents, source.Result()...)
		log.Println(source.Name(), "Total torrents found:", len(torrents), "for:", path, "Time taken: ", time.Since(startTime))
	}

	torrents = s.Sort(torrents)
	return torrents, nil
}
