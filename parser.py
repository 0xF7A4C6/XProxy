# This script is helpfull to parse checked proxies.

with open('./script/socks4.txt', 'w+') as socks4_fd:
    with open('./script/socks5.txt', 'w+') as socks5_fd:
        with open('./script/http.txt', 'w+') as http_fd:
            for proxy in list(set(open('./data/checked.txt', 'r+').read().splitlines())):
                if 'socks4' in proxy:
                    socks4_fd.write(f'{proxy.split("://")[1]}\n')
                
                if 'socks5' in proxy:
                    socks5_fd.write(f'{proxy.split("://")[1]}\n')
                
                if 'http' in proxy:
                    http_fd.write(f'{proxy.split("://")[1]}\n')