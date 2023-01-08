package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/BurntSushi/toml"
)

var (
	Config   = ConfigStruct{}
	ActualIp string
	Valid    = 0

	Socks4 = 0
	Socks5 = 0
	Http   = 0
	Dead   = 0
	Bad    = 0
)

type ConfigStruct struct {
	Filter struct {
		Timeout       int      `toml:"timeout"`
		ScrapeTimeout int      `toml:"scrape_timeout"`
		HTTP          bool     `toml:"http"`
		Socks4        bool     `toml:"socks4"`
		Socks5        bool     `toml:"socks5"`
		Country       []string `toml:"country"`
		URLCustom     string   `toml:"url_custom"`
		Match         string   `toml:"match"`
	} `toml:"filter"`
	Options struct {
		Scrape              bool `toml:"scrape"`
		Threads             int  `toml:"threads"`
		ScrapeThreads       int  `toml:"scrape_threads"`
		SaveTransparent     bool `toml:"save_transparent"`
		ShowDeadProxies     bool `toml:"show_dead_proxies"`
		RemoveURLOnError    bool `toml:"remove_url_on_error"`
		CheckScrapedProxies bool `toml:"check_scraped_proxies"`
		EnableCustomURL     bool `toml:"enable_custom_url"`
	} `toml:"options"`
	Dev struct {
		Debug bool `toml:"debug"`
	} `toml:"dev"`
}

func GetActualIp() string {
	res, err := http.Get("https://api.ipify.org")
	if HandleError(err) {
		return ""
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if HandleError(err) {
		return ""
	}

	return string(content)
}

func LoadConfig() {
	if _, err := toml.DecodeFile("script/config.toml", &Config); err != nil {
		panic(err)
	}
	ActualIp = GetActualIp()
}
