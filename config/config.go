package config

import (
	"github.com/iunsuccessful/generator/path"
	"github.com/magiconair/properties"
	"log"
	"path/filepath"
	"strings"
)

func GetFileConfig() config {
	path := path.GetBasePath()
	p := properties.MustLoadFile(filepath.Join(path, "config.properties"), properties.UTF8)
	var cfg config
	if err := p.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
	// java jdbc.url 转换成 golang url
	// jdbc.url = jdbc:mysql://test.mysql.rds.aliyuncs.com/demo
	// jdbc.url = tcp(test.mysql.rds.aliyuncs.com)/demo
	s := strings.TrimPrefix(cfg.Url, "jdbc:mysql://")
	lastIndex := strings.LastIndex(s, "/")
	cfg.Url = "tcp(" + s[0:lastIndex] + ")" + s[lastIndex:]
	return cfg
}

type config struct {
	Url      string `properties:"jdbc.url"`
	Username string `properties:"jdbc.username"`
	Password string `properties:"jdbc.password"`
}
