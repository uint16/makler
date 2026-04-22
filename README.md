
# Makler

Know Your Representative — a portal for publicly available information on Members of Parliament of the United Republic of Tanzania.

## Running locally

### Prerequisites

- [Go 1.21+](https://golang.org/doc/install)
- GCC (required by the SQLite driver) — on macOS: `xcode-select --install`

### Setup

Clone the repo:

```sh
git clone https://github.com/uint16/makler
cd makler
```

Install dependencies:

```sh
go mod download
```

Run the app:

```sh
go run main.go
```

Open [http://localhost:8080](http://localhost:8080).

The app uses the bundled SQLite database at `database/scraperwiki0.sqlite` by default — no extra configuration needed.

### Environment variables

| Variable      | Required in prod | Description                          |
|---------------|-----------------|--------------------------------------|
| `PORT`        | No              | HTTP port (default: `8080`)          |
| `BEEGO_RUNMODE` | No            | Set to `prod` in production          |
| `DB_NAME`     | Yes (prod)      | PostgreSQL database name             |
| `PG_USER`     | Yes (prod)      | PostgreSQL user                      |
| `PG_PASSWORD` | Yes (prod)      | PostgreSQL password                  |
| `ASSETS_URL`  | No              | Base URL for member photo assets     |

When `DB_NAME` is not set the app falls back to the local SQLite database.

## Deploying to Heroku

```sh
heroku create
heroku config:set BEEGO_RUNMODE=prod DB_NAME=... PG_USER=... PG_PASSWORD=... ASSETS_URL=...
git push heroku master
```

## Tech stack

- [Go](https://golang.org/) + [Beego v2](https://github.com/beego/beego)
- [Bootstrap 5](https://getbootstrap.com/) (Bootswatch Simplex theme)
- SQLite (dev) / PostgreSQL (prod)
