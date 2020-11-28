package logger

//
//import (
//	"errors"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type Syslog struct {
//	NetworkType string `yaml:"network type" json:"network_type" toml:"network_type"`
//	Host        string `yaml:"host" json:"host" toml:"host"`
//	Severity    string `yaml:"severity" json:"severity" toml:"severity"`
//	Facility    string `yaml:"facility" json:"facility" toml:"facility"`
//	Port        string `yaml:"port" json:"port" toml:"port"`
//}
//
//func sysloggerHook(config *Syslog) (log.Hook, error) {
//	return nil, errors.New("no windows support")
//}
