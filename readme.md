<h1 align="center">XProxy</h1>

`Powerfull proxy scraper and checker.`

### Features:
- scrape proxies from url/public proxy list.
- Keep only proxies from one country.
- check for proxy anonymity level.
- Support `http, socks4, socks5`.
- match custom url page.
- scrape any webpage.
- remove dead links.
- ultra fast.

### Speed:
- 100K+ scrape under 5s.
- (800 goroutines) Checked 55k proxies in 56s.

### Requirements:
- golang 1.18+
- python 3.X

### Download:
- download the lasted compiled version [here](https://github.com/0xVichy/XProxy/releases/tag/lasted).
- build from src using `go build .`

### Known issue:
- [FIXED] Checker crash if there is invalid format into proxies file, this can happen if you are using scraper.
- [FIXED] `fixed: "bufio.Scanner: token too long"`, this error happen when you are loading large proxies file (100K+)

### Utils:
- If you want to split checked proxies into socks4, socks5, http files you can use `parser.py` file.

---

<p align="center">
    <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/VichyGopher/XProxy?style=for-the-badge&logo=stylelint&color=black">
    <img alt="GitHub top language" src="https://img.shields.io/github/languages/top/VichyGopher/XProxy?style=for-the-badge&logo=stylelint&color=black">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/VichyGopher/XProxy?style=for-the-badge&logo=stylelint&color=black">
</p>
