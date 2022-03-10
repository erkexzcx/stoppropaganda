package targets

import (
	"net/url"
	"strings"
	"testing"
)

func TestWebsitesLinks(t *testing.T) {
	for k := range TargetWebsites {

		if strings.HasSuffix(k, "/") {
			t.Errorf("Invalid website '%s': ends with slash", k)
		}

		parsedURL, err := url.ParseRequestURI(k)

		if err != nil {
			t.Errorf("Invalid website '%v':", err)
		}

		if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
			t.Errorf("unknown or missing scheme '%v':", err)
		}

		// Find duplicates
		if strings.Contains(k, "://www.") {
			newStr := strings.Replace(k, "://www.", "://", 1) // Treat www.example.com and example.com as duplicates
			newStrHTTP := strings.Replace(newStr, "http://", "https://", 1)
			newStrHTTPS := strings.Replace(newStr, "https://", "http://", 1)

			_, fhttp := TargetWebsites[newStrHTTP]
			_, fhttps := TargetWebsites[newStrHTTPS]

			if fhttp {
				t.Errorf("duplicate websites '%v' and '%v'", k, newStrHTTP)
			}
			if fhttps {
				t.Errorf("duplicate websites '%v' and '%v'", k, newStrHTTPS)
			}
		}

	}
}
