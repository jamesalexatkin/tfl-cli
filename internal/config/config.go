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
	DefaultAppID               = "TODO"
	DefaultAppKey              = "TODO"
	DefaultDepartureBoardWidth = 70
	DefaultHomeStation         = "Charing Cross"
	DefaultWorkStation         = "Liverpool Street"
)

// Config encapsulates configuration variables for the app.
type Config struct {
	AppID               string `json:"app_id"`
	AppKey              string `json:"app_key"`
	DepartureBoardWidth int    `json:"departure_board_width"`
	HomeStation         string `json:"home_station"`
	WorkStation         string `json:"work_station"`
}

func (c Config) toMap() map[string]string {
	return map[string]string{
		"app_id":                c.AppID, // TODO: extract keys for these into constants
		"app_key":               c.AppKey,
		"departure_board_width": strconv.FormatInt(int64(c.DepartureBoardWidth), 10),
		"home_station":          c.HomeStation,
		"work_station":          c.WorkStation,
	}
}

// Validate checks the config to ensure it is valid.
func (c Config) Validate() error {
	if c.AppID == "" || c.AppID == DefaultAppID {
		return FieldInvalidError{Field: "app_id"}
	}

	if c.AppKey == "" || c.AppKey == DefaultAppKey {
		return FieldInvalidError{Field: "app_key"}
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

	appID := os.Getenv("APP_ID")
	appKey := os.Getenv("APP_KEY")
	departureBoardWidth := getenvInt("DEPARTURE_BOARD_WIDTH", DefaultDepartureBoardWidth)
	homeStation := os.Getenv("HOME_STATION")
	workStation := os.Getenv("WORK_STATION")

	return &Config{
		AppID:               appID,
		AppKey:              appKey,
		DepartureBoardWidth: departureBoardWidth,
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
