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
		_, err := url.ParseRequestURI(k)
		if err != nil {
			t.Errorf("Invalid website '%v':", err)
		}
	}
}
