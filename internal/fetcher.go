package internal

import (
    "bufio"
    "net/http"
)

func FetchDomains(url string) ([]string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var domains []string
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        domains = append(domains, scanner.Text())
    }
    return domains, scanner.Err()
}
