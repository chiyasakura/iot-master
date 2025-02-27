package db

import (
	"github.com/zgwit/iot-master/v4/config"
	"github.com/zgwit/iot-master/v4/pkg/env"
	"strconv"
)

// Options 参数
type Options struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Debug    bool   `json:"debug,omitempty"`
	LogLevel int    `json:"log_level"`
}

func Default() Options {
	return Options{
		Type:     "mysql",
		URL:      "root:root@tcp(localhost:3306)/master?charset=utf8",
		Debug:    true,
		LogLevel: 2,
	}
}

var options Options = Default()
var configure = "database"

const ENV = config.ENV_PREFIX + "DATABASE_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	//首先加载环境变量
	options.FromEnv()
}

func (options *Options) FromEnv() {
	options.Type = env.Get(ENV+"TYPE", options.Type)
	options.URL = env.Get(ENV+"URL", options.URL)
	options.Debug = env.GetBool(ENV+"DEBUG", options.Debug)
	options.LogLevel = env.GetInt(ENV+"LOG_LEVEL", options.LogLevel)
}

func (options *Options) ToEnv() []string {
	ret := []string{ENV + "TYPE=" + options.Type, ENV + "URL=" + options.URL}
	if options.Debug {
		ret = append(ret, ENV+"DEBUG=TRUE")
	}
	if options.LogLevel > 0 {
		ret = append(ret, ENV+"LOG_LEVEL="+strconv.Itoa(options.LogLevel))
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
