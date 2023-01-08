package main

import (
	"Proxy/components/modules"
	"Proxy/components/utils"
	"fmt"
	"strings"
	"time"

	"github.com/zenthangplus/goccm"
)

func main() {
	utils.PrintLogo()
	utils.LoadConfig()
	utils.SetTitle("")

	if len(utils.Config.Filter.Country) == 0 {
		utils.Config.Filter.Country[0] = "*"
	}

	if utils.Config.Options.Scrape {
		modules.Scrape()
	}

	if utils.Config.Options.CheckScrapedProxies {
		proxies, err := utils.ReadLines("proxies.txt")
		if utils.HandleError(err) {
			return
		}

		proxies = utils.RemoveDuplicateStr(proxies)
		utils.Log(fmt.Sprintf("Loaded %d proxies", len(proxies)))

		StartTime := time.Now()
		c := goccm.New(utils.Config.Options.Threads)

		for _, proxy := range proxies {
			c.Wait()

			if strings.Contains(proxy, "http") && !utils.Config.Filter.HTTP || strings.Contains(proxy, "socks4") && !utils.Config.Filter.Socks4 || strings.Contains(proxy, "socks5") && !utils.Config.Filter.Socks5 {
				continue
			}

			go func(proxy string) {
				modules.CheckProxy(proxy)
				utils.SetTitle(fmt.Sprintf("Checker - %fs - HTTP: %d, SOCKS4: %d, SOCKS5: %d, DEAD: %d BAD: %d, REMANING: %d", time.Since(StartTime).Seconds(), utils.Http, utils.Socks4, utils.Socks5, utils.Dead, utils.Bad, len(proxies)-(utils.Http+utils.Socks4+utils.Socks5+utils.Dead+utils.Bad)))
				c.Done()
			}(proxy)
		}

		c.WaitAllDone()
		utils.Log(fmt.Sprintf("Checked %d proxies in %fs | HTTP: %d, SOCKS4: %d, SOCKS5: %d, DEAD: %d", len(proxies), time.Since(StartTime).Seconds(), utils.Http, utils.Socks4, utils.Socks5, utils.Dead))
	}
}
