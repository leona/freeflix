# freeflix

Search, download and watch torrents from a single place. Support for Wireguard with a built-in killswitch.

![image](https://i.imgur.com/rCIuwfg.png)

![image](https://i.imgur.com/u0O7OA9.png)

## Services

- [Prowlarr](https://github.com/Prowlarr/Prowlarr) - Torrent indexer proxy
- [Jellyfin](https://jellyfin.org/) - Media player & user management
- [Caddy](https://caddyserver.com/) - Reverse proxy & automatic SSL
- Navigator - Preact frontend for search/downloads
- API - JS API for the navigator
- Client - Go [torrent](https://github.com/anacrolix/torrent)+[Wireguard](https://github.com/WireGuard/wireguard-go) client & API

## Requirements
- docker + docker-compose
- Wireguard config file (optional)

## Setup
This was made to run on a remote machine with a domain setup for SSL, but can work fine locally too.

### 1. Clone the repo
`git clone https://github.com/leona/freeflix.git`

### 2. Setup .env

Copy the example.env file into .env and change the JWT_SECRET and PUBLIC_URL.

### 3. Wireguard config

Put your Wireguard config into `config/client/1.conf` or disable Wireguard.

The client container has the following environment variables you can pass

```
WIREGUARD_ENABLED=true
WIREGUARD_CONFIG_PATH=/config/1.conf
OUTPUT_PATH=/data
API_PORT=80
MAX_DOWNLOAD_AGE=7 # Days before files are deleted
```

### 3. Start

`docker-compose up`

### 4. Setup Prowlarr

Go to `localhost:9696` in order to setup a Prowlarr user and add at least 1 or 2 indexers.

### 5. Setup Jellyfin

Go to `localhost:8096` and setup a user and media library. Use /media as the folder and disable meta info checking as it breaks stuff.

Also create an API key from Administration > Dashboard > Advanced > API Keys and drop that into the `.env` file.

You can get clients for your phone, TV etc. to connect directly to this container.

### 6. Good to go!

If you visit `localhost` you should be able to access the frontend. The same user/pass combo as Jellyfin is used here, so you can add users through there.

### Test leaks
The Wireguard setup changes the transports of the torrent client, so it's worth checking if anything is leaking.
https://ipleak.net/

## Development
`docker-compose -f docker-compose.dev.yml up`