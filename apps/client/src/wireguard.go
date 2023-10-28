package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/netip"
	"os"
	"strings"
	"time"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun/netstack"
)

type Wireguard struct {
	Config *WireguardConfig
	Tnet   *netstack.Net
}

type WireguardKeypair struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func MakeWireguard() *Wireguard {
	log.Println("Setting up wireguard")
	configPath, err := GetRandomFile("/app/config/client", "conf")

	if err != nil {
		log.Panic(err)
	}

	config, err := MakeWireguardConfigFromFile(configPath)

	if err != nil {
		log.Panic(err)
	}

	wireguard := &Wireguard{
		Config: config,
	}

	wireguard.Connect()
	go wireguard.TestTicker()
	return wireguard
}

func (w *Wireguard) Connect() {
	tun, tnet, err := netstack.CreateNetTUN(
		[]netip.Addr{netip.MustParseAddr(w.Config.Address)},
		[]netip.Addr{netip.MustParseAddr("8.8.8.8")},
		1420)
	if err != nil {
		log.Panic(err)
	}

	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, ""))
	config := w.Config.Serialize()
	log.Println("Wireguard connecting with:", config)
	err = dev.IpcSet(config)

	if err != nil {
		log.Panic(err)
	}

	err = dev.Up()

	if err != nil {
		log.Panic(err)
	}

	log.Println("Wireguard connected")
	w.Tnet = tnet
}

func (w *Wireguard) TestTicker() {
	interval := 1 * time.Minute
	w.Test()

	for range time.Tick(interval) {
		w.Test()
	}
}

func (w *Wireguard) Test() {
	client := http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: w.Tnet.DialContext,
		},
	}

	resp, err := client.Get("http://icanhazip.com/")

	if err != nil {
		log.Panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panic(err)
	}

	log.Println("Wireguard connected on IP:", string(body))
}

type WireguardConfig struct {
	PrivateKey string
	PublicKey  string
	AllowedIps string
	Endpoint   string
	Address    string
}

func MakeWireguardConfigFromFile(configPath string) (*WireguardConfig, error) {
	log.Println("Reading wireguard config from:", configPath)
	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	config := &WireguardConfig{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")

		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		otherParts := strings.Join(parts[1:], "=")
		value := strings.TrimSpace(otherParts)

		switch key {
		case "PrivateKey":
			config.PrivateKey = value
		case "PublicKey":
			config.PublicKey = value
		case "AllowedIPs":
			config.AllowedIps = strings.Split(value, ",")[0]
		case "Endpoint":
			config.Endpoint = value
		case "Address":
			config.Address = strings.Split(value, "/")[0]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func (w *WireguardConfig) Serialize() string {
	return fmt.Sprintf(`private_key=%s
public_key=%s
allowed_ip=%s
endpoint=%s`, base64ToHex(w.PrivateKey), base64ToHex(w.PublicKey), w.AllowedIps, w.Endpoint)
}
