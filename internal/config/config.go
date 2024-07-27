package config

import (
	"fmt"
	"os"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	DB      DB      `env:"db"`
	Key     Key     `env:"key"`
	Logging Logging `env:"logging"`
	Server  Server  `env:"server"`
}

type DB struct {
	DbHost    string `required:"true"`
	DbUser    string `required:"true"`
	DbPass    string `required:"true"`
	DbPort    string `required:"true"`
	DbName    string `required:"true"`
	DbSslmode string `required:"true"`
	DbTz      string `required:"true"`
}

type Key struct {
	JwtKey string `required:"true"`
}
type Logging struct {
	GormLevel   int `default:"4"`
	LogrusLevel int `default:"4"`
}

type Server struct {
	App     string `default:"compass"`
	Env     string `default:"development"`
	Port    string `required:"true"`
	Host    string `required:"true"`
	Scheme  string `required:"true"`
	Version string `default:"1.0.0"`
}

var Config *Configuration = &Configuration{}

func Load(password string) error {
	file := os.Getenv("CONFIG")
	path := fmt.Sprintf("/run/secrets/%s", file)
	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		path = fmt.Sprintf("./run/secrets/%s", file)
	}

	fmt.Println(os.Getenv("ENVIRONMENT"))

	if password == "" {
		password = "213123asd"
	}

	if os.Getenv("ENVIRONMENT") != "LOCAL" {

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		message, err := helper.DecryptMessageWithPassword([]byte(password), string(b))
		if err != nil {
			return err
		}

		path = path + ".yml"

		err = os.WriteFile(path, []byte(message), 0o666)
		if err != nil {
			return err
		}

		defer func() {
			err := os.Remove(path)
			if err != nil {
				logrus.Fatal(err)
			}
		}()
	}

	err := configor.Load(Config, path)
	if err != nil {
		return err
	}

	return nil
}

func LoadForTest(password string) error {
	file := os.Getenv("CONFIG")
	path := fmt.Sprintf("/run/secrets/%s", file)
	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		path = fmt.Sprintf("../../../run/secrets/%s", file)
	}

	fmt.Println(os.Getenv("ENVIRONMENT"))

	if password == "" {
		password = "213123asd"
	}

	if os.Getenv("ENVIRONMENT") != "LOCAL" {

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		message, err := helper.DecryptMessageWithPassword([]byte(password), string(b))
		if err != nil {
			return err
		}

		path = path + ".yml"

		err = os.WriteFile(path, []byte(message), 0o666)
		if err != nil {
			return err
		}

		defer func() {
			err := os.Remove(path)
			if err != nil {
				logrus.Fatal(err)
			}
		}()
	}

	err := configor.Load(Config, path)
	if err != nil {
		return err
	}

	return nil
}

func Get() *Configuration {
	return Config
}
