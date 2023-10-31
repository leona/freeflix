# freeflix

Search, download and watch torrents from a single place. Support for Wireguard with a built-in killswitch.

![image](https://i.imgur.com/gYa04jx.png)

![image](https://i.imgur.com/u0O7OA9.png)

## Features

- [Jellyfin](https://jellyfin.org/) media player & user management
- React frontend for search/downloads with support for direct download
- Go [torrent client](https://github.com/anacrolix/torrent/)
- 1337x & ThePirateBay scrapers (more planned)
- [Docker Wireguard](https://github.com/leona/docker-wireguard) with automatic Mullvad configuration
- [Caddy](https://caddyserver.com/) reverse proxy & automatic SSL

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

Put your Wireguard config into `config/client/*.conf` or disable Wireguard.

The client container has the following environment variables you can pass

```
OUTPUT_PATH=/data
API_PORT=80
MAX_DOWNLOAD_AGE=7 # Days before files are deleted
MULLVAD_ACCOUNT=
MULLVAD_COUNTRIES=nl,germany
```

If you pass your [Mullvad](https://mullvad.net) account ID it will automatically download your config files.

Config files will be randomly selected each time the client container starts.

### 4. Start

`docker-compose --profile build up && docker-compose --profile serve up -d`

### 5. Setup Jellyfin

Go to `localhost:8096` and setup a user and media library. Use /media as the folder and disable meta info checking as it breaks stuff.

Also create an API key from Administration > Dashboard > Advanced > API Keys and drop that into the `.env` file.

You can get clients for your phone, TV etc. to connect directly to this container.

### 6. Good to go!

If you visit `localhost` you should be able to access the frontend. The same user/pass combo as Jellyfin is used here, so you can add users through there.

### Test leaks

https://ipleak.net/

## Development

`docker-compose -f docker-compose.dev.yml up`

## Todo

- Scan libraries when torrents complete
