package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const fileName = ".tfl.env"

const (
	// DefaultAppID is the default value for app ID.
	DefaultAppID = "TODO"
	// DefaultAppKey is the default value for app key.
	DefaultAppKey = "TODO"
	// DefaultDepartureBoardWidth is the default value for the departure board width.
	DefaultDepartureBoardWidth = 70
	// DefaultNumDepartures is the default values for the number of departures to display.
	DefaultNumDepartures = 4
	// DefaultHomeStation is the default value for the home station.
	DefaultHomeStation = "Charing Cross"
	// DefaultWorkStation is the default value for the work station.
	DefaultWorkStation = "Liverpool Street"
)

// Key represents a key for a config value.
type Key string

const (
	// AppIDConfigKey represents the key for the `app_id` field.
	AppIDConfigKey Key = "app_id"
	// AppKeyConfigKey represents the key for the `app_key` field.
	AppKeyConfigKey Key = "app_key"
	// DepartureBoardWidthConfigKey represents the key for the `departure_board_width` field.
	DepartureBoardWidthConfigKey Key = "departure_board_width"
	// NumDeparturesConfigKey represents the key for the `num_departures` field.
	NumDeparturesConfigKey Key = "num_departures"
	// HomeStationConfigKey represents the key for the `home_station` field.
	HomeStationConfigKey Key = "home_station"
	// WorkStationConfigKey represents the key for the `work_station` field.
	WorkStationConfigKey Key = "work_station"
)

// Config encapsulates configuration variables for the app.
type Config struct {
	AppID               string `json:"app_id"`
	AppKey              string `json:"app_key"`
	DepartureBoardWidth int    `json:"departure_board_width"`
	NumDepartures       int    `json:"num_departures"`
	HomeStation         string `json:"home_station"`
	WorkStation         string `json:"work_station"`
}

func (c Config) toMap() map[string]string {
	return map[string]string{
		string(AppIDConfigKey):               c.AppID,
		string(AppKeyConfigKey):              c.AppKey,
		string(DepartureBoardWidthConfigKey): strconv.FormatInt(int64(c.DepartureBoardWidth), 10),
		string(NumDeparturesConfigKey):       strconv.FormatInt(int64(c.NumDepartures), 10),
		string(HomeStationConfigKey):         c.HomeStation,
		string(WorkStationConfigKey):         c.WorkStation,
	}
}

// Validate checks the config to ensure it is valid.
func (c Config) Validate() error {
	if c.AppID == "" || c.AppID == DefaultAppID {
		return FieldInvalidError{Field: AppIDConfigKey}
	}

	if c.AppKey == "" || c.AppKey == DefaultAppKey {
		return FieldInvalidError{Field: AppKeyConfigKey}
	}

	return nil
}

// LoadConfig looks for a config file to load into memory.
// If there isn't one already, a new one is created with default values.
func LoadConfig() (*Config, error) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No existing config found, creating now")

		config := Config{
			AppID:               DefaultAppID,
			AppKey:              DefaultAppKey,
			DepartureBoardWidth: DefaultDepartureBoardWidth,
			NumDepartures:       DefaultNumDepartures,
			HomeStation:         DefaultHomeStation,
			WorkStation:         DefaultWorkStation,
		}

		err = godotenv.Write(config.toMap(), fileName)
		if err != nil {
			return nil, err
		}
	}

	err := godotenv.Load(fileName)
	if err != nil {
		return nil, err
	}

	appID := os.Getenv(string(AppIDConfigKey))
	appKey := os.Getenv(string(AppKeyConfigKey))
	departureBoardWidth := getenvInt(string(DepartureBoardWidthConfigKey), DefaultDepartureBoardWidth)
	numDepartures := getenvInt(string(NumDeparturesConfigKey), DefaultNumDepartures)
	homeStation := os.Getenv(string(HomeStationConfigKey))
	workStation := os.Getenv(string(WorkStationConfigKey))

	return &Config{
		AppID:               appID,
		AppKey:              appKey,
		DepartureBoardWidth: departureBoardWidth,
		NumDepartures:       numDepartures,
		HomeStation:         homeStation,
		WorkStation:         workStation,
	}, nil
}

func getenvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
