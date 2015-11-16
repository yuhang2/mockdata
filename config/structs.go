package config

import (
	"fmt"
	"sync"
)

var Config = &struct {
	RedShift RedShift `json:"redshift"`
}{}

type RedShift struct {
	Host           string `json:"host"`
	Port           int    `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	Db             string `json:"db"`
	Ssl            string `json:"ssl"`
	MaxConn        int    `json:"maxConn"`
	ConnectTimeout int    `json:"connectTimeout"`
	ReadTimeout    int    `json:"readTimeout"`
	WriteTimeout   int    `json:"writeTimeout"`

	stringer     string
	key          string
	stringerLock sync.Mutex
	keyLock      sync.Mutex
}

// String implements the Stringer interface
func (conf *RedShift) String() string {
	conf.stringerLock.Lock()
	if conf.stringer == "" {
		conf.stringer = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.Db,
			conf.Ssl,
		)
	}
	conf.stringerLock.Unlock()
	return conf.stringer
}
