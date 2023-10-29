package main

import (
	"github.com/gin-gonic/gin"
	//"time"
)

var torrentClient *Torrentclient
var wireguardClient *Wireguard
var config *Config

func main() {
	config = MakeConfig()
	go cleanup()
	gin.SetMode(gin.ReleaseMode)

	if config.MullvadAccount != "" {
		_ = MakeMullvad(config.MullvadAccount)
	}

	if config.WireguardEnable {
		wireguardClient = MakeWireguard()
	}

	torrentClient = MakeTorrentclient()

	// run every 30 seconds
	//	for range time.Tick(5 * time.Second) {
	torrentClient.Add("magnet:?xt=urn:btih:0A62A7B53706DC91589763FBA8CD431C4371FE53&dn=160+Excel+Exercises+With+Solutions+And+Comments+-+Excel+Workbook&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.open-internet.nl%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2850%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2720%2Fannounce&tr=udp%3A%2F%2F9.rarbg.me%3A2810%2Fannounce&tr=udp%3A%2F%2F9.rarbg.to%3A2890%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.trackerfix.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.fatkhoala.org%3A13790%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Fopentracker.i2p.rocks%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce")
	//	torrentClient.Add("magnet:?xt=urn:btih:ba6991a33859e509b0629f9eb89adc0b3d61e0df&tr=https://ipleak.net/announce.php%3Fh%3Dba6991a33859e509b0629f9eb89adc0b3d61e0df&dn=ipleak.net+torrent+detection")
	//}
	ServeApi()
}
