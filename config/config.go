package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var MAX_ROUNDS = 20
var CONFIG_FILE = "config.yml"
var HERO_NAME = "Hero"
var VILLAIN_NAME = "Villain"
var WAR_STORY = fmt.Sprintf("%s walks the whimsical forests of the Terminal Valley and he encounters a nefarious %s.\n", HERO_NAME, VILLAIN_NAME)

var HERO_RESILIENCE_NUM_TURNS = 2

type Config struct {
	Hero struct {
		Common struct {
			Health struct {
				HealthStart int `yaml:"health_start"`
				HealthEnd   int `yaml:"health_end"`
			} `yaml:"health"`
			Strength struct {
				StrengthStart int `yaml:"strength_start"`
				StrengthEnd   int `yaml:"strength_end"`
			} `yaml:"strength"`
			Defence struct {
				DefenceStart int `yaml:"defence_start"`
				DefenceEnd   int `yaml:"defence_end"`
			} `yaml:"defence"`
			Speed struct {
				SpeedStart int `yaml:"speed_start"`
				SpeedEnd   int `yaml:"speed_end"`
			} `yaml:"speed"`
			Luck struct {
				LuckStart float64 `yaml:"luck_start"`
				LuckEnd   float64 `yaml:"luck_end"`
			} `yaml:"luck"`
		} `yaml:"common"`
		Special struct {
			CriticalStrike struct {
				StrikeTwiceProbability           float64 `yaml:"strike_twice_probability"`
				StrikeThirdGivenTwiceProbability float64 `yaml:"strike_third_given_twice_probability"`
			} `yaml:"critical_strike"`
			Resilience struct {
				HalfDamageResilienceProbability float64 `yaml:"half_damage_resilience_probability"`
			} `yaml:"resilience"`
		} `yaml:"special"`
	} `yaml:"hero"`
	Villain struct {
		Common struct {
			Health struct {
				HealthStart int `yaml:"health_start"`
				HealthEnd   int `yaml:"health_end"`
			} `yaml:"health"`
			Strength struct {
				StrengthStart int `yaml:"strength_start"`
				StrengthEnd   int `yaml:"strength_end"`
			} `yaml:"strength"`
			Defence struct {
				DefenceStart int `yaml:"defence_start"`
				DefenceEnd   int `yaml:"defence_end"`
			} `yaml:"defence"`
			Speed struct {
				SpeedStart int `yaml:"speed_start"`
				SpeedEnd   int `yaml:"speed_end"`
			} `yaml:"speed"`
			Luck struct {
				LuckStart float64 `yaml:"luck_start"`
				LuckEnd   float64 `yaml:"luck_end"`
			} `yaml:"luck"`
		} `yaml:"common"`
	} `yaml:"villain"`
}

func ReadConfig() (*Config, error) {
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatalf("Could not open config file %s\n", CONFIG_FILE)
		return nil, errors.New("Could not open config file")
	}

	var config *Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Could not decode config\n")
		return nil, errors.New("Could not decode config")
	}

	return config, nil
}

type GameOutcome int

const (
	HERO_WINS GameOutcome = iota
	VILLAIN_WINS
	DRAW
)

func (o GameOutcome) String() string {
	return [...]string{"HERO WINS", "VILLAIN WINS", "DRAW"}[o]
}
