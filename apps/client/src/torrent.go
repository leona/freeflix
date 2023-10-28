package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/dustin/go-humanize"
)

type Torrentclient struct {
	Client             *torrent.Client
	SpeedCheckInterval time.Duration
	Speeds             map[string]int64
}

func MakeTorrentclient() *Torrentclient {
	clientConfig := torrent.NewDefaultClientConfig()
	clientConfig.DataDir = config.OutputPath
	clientConfig.Seed = false

	if config.WireguardEnable {
		if wireguardClient == nil {
			log.Panic("Wireguard is enabled but not configured")
		}

		clientConfig.HTTPDialContext = wireguardClient.Tnet.DialContext
		clientConfig.HTTPProxy = nil
		clientConfig.ClientTrackerConfig.TrackerDialContext = wireguardClient.Tnet.DialContext
	}

	client, err := torrent.NewClient(clientConfig)

	if err != nil {
		log.Panic(err)
	}

	torrentclient := &Torrentclient{
		Client:             client,
		SpeedCheckInterval: 3 * time.Second,
		Speeds:             map[string]int64{},
	}

	go torrentclient.WatchSpeeds()
	return torrentclient
}

func LinkToMagnet(link string) (string, error) {
	log.Println("Converting link to magnet")
	req, err := http.NewRequest("GET", link, nil)

	if err != nil {
		return "", err
	}

	client := new(http.Client)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)

	if response != nil && (response.StatusCode == http.StatusFound || response.StatusCode == http.StatusMovedPermanently) {
		url, _ := response.Location()
		return url.String(), nil
	} else {
		return "", err
	}
}

func (tc *Torrentclient) Add(magnet string) error {
	var err error

	if strings.HasPrefix(magnet, "http") {
		magnet, err = LinkToMagnet(magnet)

		if err != nil {
			log.Println("Failed to get magnet", err)
			return err
		}
	}

	log.Println("Adding torrent:", magnet)
	torrent, err := tc.Client.AddMagnet(magnet)

	if err != nil {
		log.Println("Failed to add torrent:", err)
		return err
	}

	log.Println("Added torrent:", torrent.Name())

	go func() {
		<-torrent.GotInfo()
		info := torrent.Info()
		log.Printf("Received metdata for: %s", info.Name)
		torrent.DownloadAll()
	}()

	return nil
}

func (tc *Torrentclient) Remove(hash string) {
	torrents := tc.Client.Torrents()

	for _, torrent := range torrents {
		if torrent.InfoHash().HexString() == hash {
			torrent.Drop()
			os.RemoveAll(filepath.Join(config.OutputPath, torrent.Name()))
			log.Println("Removed torrent:", torrent.Name())
			break
		}
	}
}

type TorrentStats struct {
	Name       string  `json:"name"`
	Hash       string  `json:"hash"`
	Complete   bool    `json:"complete"`
	Progress   float64 `json:"progress"`
	Size       string  `json:"size"`
	Downloaded string  `json:"downloaded"`
	CreatedAt  int64   `json:"createdAt"`
	Peers      int     `json:"peers"`
	Speed      string  `json:"speed"`
}

func MakeTorrentStats(torrent *torrent.Torrent) *TorrentStats {
	metainfo := torrent.Metainfo()
	progress := roundFloat64(float64(torrent.BytesCompleted())/float64(torrent.Length())*100, 2)
	complete := torrent.BytesCompleted() == torrent.Length()

	if torrent.Length() == 0 {
		progress = 0.0
		complete = false
	}

	stats := &TorrentStats{
		Name:       torrent.Name(),
		Hash:       torrent.InfoHash().HexString(),
		Progress:   progress,
		Size:       humanize.Bytes(uint64(torrent.Length())),
		Downloaded: humanize.Bytes(uint64(torrent.BytesCompleted())),
		Complete:   complete,
		CreatedAt:  metainfo.CreationDate,
		Peers:      torrent.Stats().TotalPeers,
		Speed:      "0",
	}

	return stats
}

func (tc *Torrentclient) WatchSpeeds() {
	lastStats := map[string]torrent.TorrentStats{}

	for range time.Tick(tc.SpeedCheckInterval) {
		for _, torrent := range tc.Client.Torrents() {
			if torrent.Length() == 0 {
				continue
			}

			hash := torrent.InfoHash().HexString()

			if torrent.BytesCompleted() == torrent.Length() {
				delete(tc.Speeds, hash)
				delete(lastStats, hash)
				torrent.Drop() // Race condition?
				log.Println("Torrent complete:", torrent.Name(), torrent.BytesCompleted(), torrent.Length())
				continue
			}

			lastStat := lastStats[hash]
			stats := torrent.Stats()
			byteRate := int64(time.Second)
			byteRate *= stats.BytesReadUsefulData.Int64() - lastStat.BytesReadUsefulData.Int64()
			byteRate /= int64(tc.SpeedCheckInterval)
			tc.Speeds[hash] = byteRate
			lastStats[hash] = torrent.Stats()
		}
	}
}

func (tc *Torrentclient) List() []TorrentStats {
	torrents := tc.Client.Torrents()
	torrentStats := []TorrentStats{}

	for _, torrent := range torrents {
		stats := MakeTorrentStats(torrent)
		stats.Speed = humanize.Bytes(uint64(tc.Speeds[stats.Hash]))

		if stats.Complete {
			log.Println("Dropping complete torrent:", torrent.Name())
			torrent.Drop()
			continue
		}

		torrentStats = append(torrentStats, *stats)
	}

	return torrentStats
}
