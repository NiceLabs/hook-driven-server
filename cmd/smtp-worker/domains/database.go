package domains

import (
	"embed"
	_ "embed"
	"strings"

	"golang.org/x/net/publicsuffix"
)

//go:embed *.txt.gz
var fs embed.FS

var domains = make(map[string]string)

func init() {
	putDomainList("disposable.txt.gz", "DISPOSABLE_MAIL")
	putDomainList("free.txt.gz", "FREE_MAIL")
	putDomainList("swot.txt.gz", "SWOT_MAIL")
	putDomainList("ddns.txt.gz", "DDNS")
}

func Type(domain string) string {
	domain, err := publicsuffix.EffectiveTLDPlusOne(strings.ToLower(domain))
	if err != nil {
		return ""
	}
	if code, ok := domains[domain]; ok {
		return code
	}
	return ""
}

func putDomainList(filename string, code string) {
	data, _ := fs.ReadFile(filename)
	for _, domain := range strings.Split(string(data), "\n") {
		domain = strings.ToLower(domain)
		if len(domain) == 0 {
			continue
		}
		domains[domain] = code
	}
	return
}
