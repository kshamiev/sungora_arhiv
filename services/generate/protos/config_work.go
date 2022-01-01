package protos

import "sungora/services/mdsun"

func init() {
	GenerateConfig["mdsun"] = []interface{}{
		&mdsun.GooseDBVersion{},
		&mdsun.Order{},
		&mdsun.Role{},
		&mdsun.User{},
	}
}
