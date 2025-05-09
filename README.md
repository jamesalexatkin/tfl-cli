# 🚇 TfL CLI
[![Go Reference](https://pkg.go.dev/badge/github.com/jamesalexatkin/tfl-cli.svg)](https://pkg.go.dev/github.com/jamesalexatkin/tfl-cli)
[![GitHub License](https://img.shields.io/github/license/jamesalexatkin/tfl-cli)](https://github.com/jamesalexatkin/tfl-cli/blob/main/LICENSE)
[![Build](https://github.com/jamesalexatkin/tfl-cli/actions/workflows/go.yml/badge.svg)](https://github.com/jamesalexatkin/tfl-cli/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jamesalexatkin/tfl-cli)](https://goreportcard.com/report/github.com/jamesalexatkin/tfl-cli)

> A fast and minimal terminal app to check Transport for London (TfL) status and live departures, written in Go.

![Status](assets/status_screenshot.png)



## ✨ Features

- Live status of TfL lines (Tube, Overground, Elizabeth Line, DLR)
- Real-time departures for any TfL station
- Config management to set preferences, such as home and work stations


## 📦 Installation

```bash
go install github.com/jamesalexatkin/tfl-cli@latest
```


## 🚀 Usage

### Line status (all lines)

```bash
tfl-cli status
```

### TODO - Line status (single line)

```bash
tfl-cli status victoria
```

### Live departures from a station

```bash
tfl-cli station 'Liverpool Street'
```

### Help

```bash
tfl-cli h
tfl-cli help
tfl-cli -h
tfl-cli --help
```


## ⚙️ Configuration

The first time you run the app, a default `.tfl.env` file will be generated.

You will need to configure API credentials here before fetching any data. These can be generated by signing up for a free account on [TfL's API portal](https://api-portal.tfl.gov.uk/).

### Options

| Key                     | Description                                                                                              |
|-------------------------|----------------------------------------------------------------------------------------------------------|
| `app_id`                | **(Required)** Your TfL API App ID for increased rate limits.                                            |
| `app_key`               | **(Required)** Your TfL API Key. Used together with `app_id`.                                            |
| `departure_board_width` | Sets the character width for the departure board display. Handy for wide or narrow terminals.            |
| `num_departures`        | Sets the number of departures to display on a departure board.                                           |
| `home_station`          | Sets your default "home" station. Use with the `station` command.                                        |
| `work_station`          | Sets your default "work" station. Use with the `station` command.                                        |


## 🛠️ Built with

- [Go](https://golang.org/)
- [urfave/cli](https://github.com/urfave/cli/v3) (for CLI)
- [fatih/color](https://github.com/fatih/color) (for terminal colours)
- [jamesalexatkin/tfl-golang](https://github.com/jamesalexatkin/tfl-golang) (for transport status)


<!-- ## 🧪 Development

```bash
go build
```

To run tests:

```bash
go test ./...
``` -->


## 🤝 Contributing

Pull requests welcome! If you have ideas or bug reports, feel free to open an issue.


## 📜 License

MIT © 2025 James Atkin
