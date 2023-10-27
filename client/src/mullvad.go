package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type Mullvad struct {
	Account string
	BaseUrl string
}

func MakeMullvad(account string) *Mullvad {
	return &Mullvad{
		Account: account,
		BaseUrl: "https://api.mullvad.net/",
	}
}

func (m *Mullvad) Get(endpoint string, response interface{}) error {
	resp, err := http.Get(m.BaseUrl + endpoint)

	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panic(err)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Panic(err)
	}

	return nil
}

type MullvadServers struct {
	Countries []struct {
		Name   string `json:"name"`
		Code   string `json:"code"`
		Cities []struct {
			Name      string  `json:"name"`
			Code      string  `json:"code"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Relays    []struct {
				Hostname     string `json:"hostname"`
				Ipv4AddrIn   string `json:"ipv4_addr_in"`
				Ipv6AddrIn   string `json:"ipv6_addr_in"`
				PublicKey    string `json:"public_key"`
				MultihopPort int    `json:"multihop_port"`
			} `json:"relays"`
		} `json:"cities"`
	} `json:"countries"`
}

func (m *Mullvad) GetServers() {
	var servers MullvadServers

	if err := m.Get("public/relays/wireguard/v1/", servers); err != nil {
		log.Panic(err)
	}

	spew.Dump(servers.Countries)

}
