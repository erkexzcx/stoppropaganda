package targets

import (
	"strings"
	"testing"
)

func TestDNSTargets(t *testing.T) {
	for k := range TargetDNSServers {
		if !strings.HasSuffix(k, ":53") {
			t.Errorf("Invalid DNS target '%s'", k)
		}
	}
}
