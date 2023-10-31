package scraper

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Source interface {
	Name() string
	Parse(string)
	Query(string) string
	Result() []Torrent
	RequiresCollector() bool
	SetCollector(*colly.Collector)
}

type Torrent struct {
	Name       string `json:"name"`
	Link       string `json:"link"`
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
	return []Source{
		Make1337xSource(),
		MakePiratebaySource(),
	}
}

func (s *Scraper) Sort(torrents []Torrent) []Torrent {
	sort.Slice(torrents, func(i, j int) bool {
		return torrents[i].Seeders > torrents[j].Seeders
	})

	return torrents
}

func (s *Scraper) CreateCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 6})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	return c
}

func (s *Scraper) Query(query string) ([]Torrent, error) {
	torrents := s.SourceMap(func(source Source, mu *sync.Mutex) []Torrent {
		var c *colly.Collector

		if source.RequiresCollector() {
			c = s.CreateCollector()
			source.SetCollector(c)
		}

		path := source.Query(query)
		source.Parse(path)

		if source.RequiresCollector() {
			c.Visit(path)
			c.Wait()
		}

		return source.Result()
	})

	torrents = s.Sort(torrents)
	log.Println("Total torrents found:", len(torrents))
	return torrents, nil
}

func (s *Scraper) SourceMap(callback func(Source, *sync.Mutex) []Torrent) []Torrent {
	sources := s.GetSources()
	mu := &sync.Mutex{}
	var wg sync.WaitGroup
	torrents := []Torrent{}

	for _, item := range sources {
		wg.Add(1)

		go func(itm Source) {
			defer wg.Done()
			log.Println("Searching", itm.Name())
			startTime := time.Now()
			output := callback(itm, mu)
			mu.Lock()
			torrents = append(torrents, output...)
			mu.Unlock()
			log.Println(itm.Name(), "Total torrents found:", len(output), "Time taken: ", time.Since(startTime))
		}(item)
	}

	wg.Wait()
	return torrents
}
