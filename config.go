package omgo

import (
	"fmt"
)

type MgoDBCfg struct {
	User string
	Pass string
	Host string
	DB   string
}

var (
	MgDBConfigs = map[string]MgoDBCfg{}
)

func getMgUrl(config *MgoDBCfg) string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", config.User, config.Pass, config.Host, config.DB)
}
