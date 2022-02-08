package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eifzed/antre-app/lib/common"
	"github.com/eifzed/antre-app/lib/helper/toggle"
	"github.com/eifzed/antre-app/lib/utility/jwt"
	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v2"
)

type HTTP struct {
	Address        string `yaml:"address"`
	WriteTimeout   int    `yaml:"write_timeout"`
	ReadTimeout    int    `yaml:"read_timeout"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
}
type Server struct {
	Name      string `yaml:"name"`
	HTTP      HTTP   `yaml:"http"`
	Debug     int    `yaml:"debug"`
	PathVault string `yaml:"path_vault"`
	URL       string `yaml:"url"`
}

type Config struct {
	Secretes   *SecreteVault
	Server     *Server                   `yaml:"server"`
	Toggle     *toggle.Toggle            `yaml:"toggle"`
	RouteRoles map[string]jwt.RouteRoles `yaml:"route_roles"`
}

func GetConfig() (*Config, error) {
	env := "production"
	pathBase := ""

	if IsDevelopment() {
		env = "development"
		dir, _ := os.Getwd()
		pathBase = filepath.Join(dir, "files")

	}
	fileName := fmt.Sprintf("%s.%s.yaml", "antre-config", env)
	filePath := filepath.Join(pathBase, "/etc/antre-config", fileName)
	log.Infoln("reading config file from: ", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer common.SafelyCloseFile(f)

	cfg := &Config{}
	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
