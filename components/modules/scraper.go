package modules

import (
	"Proxy/components/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/zenthangplus/goccm"
)

var (
	proxyRegex = regexp.MustCompile("([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}):([0-9]{1,5})")
)

func RemoveUrl(Url string, ProxyType string) {
	utils.Log(fmt.Sprintf("Dead link: %s", Url))
	utils.RemoveLine("url.csv", fmt.Sprintf("%s,%s", ProxyType, Url))
}

func ScrapeUrl(Url string, ProxyType string) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(utils.Config.Filter.ScrapeTimeout),
	}

	res, err := client.Get(Url)
	if utils.HandleError(err) {
		if utils.Config.Options.RemoveURLOnError {
			RemoveUrl(Url, ProxyType)
		}
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 403 || res.StatusCode == 404 || res.StatusCode == 401 {
		RemoveUrl(Url, ProxyType)
		return
	}

	content, err := ioutil.ReadAll(res.Body)
	if utils.HandleError(err) {
		return
	}

	for _, proxy := range proxyRegex.FindAllString(string(content), -1) {
		utils.AppendFile("proxies.txt", fmt.Sprintf("%s://%s", ProxyType, proxy))

		switch ProxyType {
		case "http":
			utils.Http++
		case "socks4":
			utils.Socks4++
		case "socks5":
			utils.Socks5++
		}
	}
}

func Scrape() {
	url_list, err := utils.ReadLines("url.csv")
	if utils.HandleError(err) {
		return
	}

	StartTime, c, crawled := time.Now(), goccm.New(utils.Config.Options.ScrapeThreads), 0

	for i, url := range url_list {
		c.Wait()

		// * type,url
		s := strings.Split(url, ",")

		if s[0] == "http" && !utils.Config.Filter.HTTP || s[0] == "socks4" && !utils.Config.Filter.Socks4 || s[0] == "socks5" && !utils.Config.Filter.Socks5 {
			return
		}

		go func(u string, t string, n int) {
			ScrapeUrl(u, t)
			crawled++

			utils.Log(fmt.Sprintf("Scraped page #%d (%d/%d)", n, crawled, len(url_list)))
			utils.SetTitle(fmt.Sprintf("Scraper - %fs - HTTP: %d, SOCKS4: %d, SOCKS5: %d", time.Since(StartTime).Seconds(), utils.Http, utils.Socks4, utils.Socks5))
			c.Done()
		}(s[1], s[0], i)
	}

	c.WaitAllDone()
	utils.Log(fmt.Sprintf("Scraped %d urls in %fs | HTTP: %d, SOCKS4: %d, SOCKS5: %d", len(url_list), time.Since(StartTime).Seconds(), utils.Http, utils.Socks4, utils.Socks5))

	// reset counter for checking lel
	utils.Http, utils.Socks4, utils.Socks5 = 0, 0, 0
}
