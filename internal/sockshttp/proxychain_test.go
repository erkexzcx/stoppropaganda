package sockshttp

import (
	"testing"
)

func TestProxyChains(t *testing.T) {
	generatedChain := ParseProxyChain("direct,socks5://1.2.3.4:9050,http://user:pass@5.6.7.8:8080,socks4://7.7.7.7:1080,3.3.3.3:9051,direct")
	{
		generated := generatedChain[0]
		expected := Proxy{
			Addr:   "1.2.3.4:9050",
			Method: ProxyMethodSocks5,
		}
		if generated != expected {
			t.Errorf("Generated[0] chain does not equal expected: \n%v != %v", generated, expected)
		}
	}
	{
		generated := generatedChain[1]
		expected := Proxy{
			Addr:   "5.6.7.8:8080",
			Method: ProxyMethodHttp,
			Auth: Auth{
				User:     "user",
				Password: "pass",
			},
		}
		if generated != expected {
			t.Errorf("Generated[0] chain does not equal expected: \n%v != %v", generated, expected)
		}
	}
	{
		generated := generatedChain[2]
		expected := Proxy{
			Addr:   "7.7.7.7:1080",
			Method: ProxyMethodSocks4,
		}
		if generated != expected {
			t.Errorf("Generated[1] chain does not equal expected: \n%v != %v", generated, expected)
		}
	}
	{
		generated := generatedChain[3]
		expected := Proxy{
			Addr:   "3.3.3.3:9051",
			Method: ProxyMethodSocks5,
		}
		if generated != expected {
			t.Errorf("Generated[2] chain does not equal expected: \n%v != %v", generated, expected)
		}
	}

	if len(generatedChain) != 4 {
		t.Errorf("Generated proxy chain length was not 4")
	}
}

func TestDirectChain(t *testing.T) {
	generatedChain := ParseProxyChain("direct")
	if len(generatedChain) != 0 {
		t.Errorf("Generated direct (localhost) proxy chain length was not 0")
	}

}
