package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"grocery/pkg/config"
)

const (
	appConfigEnvVarKey = "GROCERY_CONFIG_PATH"
)

type Conf struct {
	ListenAddr string `yaml:"listen_addr"`
	Debug      bool   `yaml:"debug"`
	Db         struct {
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

func (c Conf) isValid() error {
	if c.ListenAddr == "" {
		return errors.New("empty ListenAddr")
	}
	return nil
}

func defaultConf() Conf {
	c := Conf{
		ListenAddr: ":8080",
		Debug:      true,
	}
	c.Db.Host = "localhost"
	c.Db.Name = "grocery_db"
	c.Db.User = "grocery_user"
	c.Db.Password = "123456"
	return c
}

func fatalIf(err error) {
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}

func AppConfig() (Conf, error) {
	p := strings.TrimSpace(os.Getenv(appConfigEnvVarKey))
	if p == "" {
		return Conf{}, errors.New("must set env variable: GROCERY_CONFIG_PATH")
	}
	var c Conf
	if err := config.ParseYAML(p, &c); err != nil {
		return Conf{}, err
	}
	if err := c.isValid(); err != nil {
		return Conf{}, err
	}
	return c, nil
}

func GenerateConfigs(p string) {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		fatalIf(os.MkdirAll(p, os.ModePerm))
	} else {
		fatalIf(err)
	}
	appConfigPath := filepath.Join(p, "app.yaml")
	appConf := defaultConf()
	for _, c := range []struct {
		filename string
		content  interface{}
	}{
		{appConfigPath, appConf},
	} {
		var bs []byte
		var err error
		switch c.content.(type) {
		case string:
			bs = []byte(c.content.(string))
		case []byte:
			bs = c.content.([]byte)
		default:
			bs, err = yaml.Marshal(c.content)
		}

		fatalIf(err)
		fatalIf(ioutil.WriteFile(c.filename, bs, 0666))
	}
}
