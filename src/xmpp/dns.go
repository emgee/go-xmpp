package xmpp

import (
	"fmt"
	"net"
	"strings"
)

const (
	// Standard port for XMPP clients to connect to.
	ClientPort = 5222
)

// Perform a DNS SRV lookup and return an ordered list of "host:port" TCP
// addresses for the JID's home server. If no SRV records are found then assume
// the JID's domain is also the home server.
func HomeServerAddrs(jid JID) (addr []string, err error) {

	// DNS lookup.
	_, addrs, _ := net.LookupSRV("xmpp-client", "tcp", jid.Domain)

	// If there's nothing in DNS then assume the JID's domain and the standard
	// port will work.
	if len(addrs) == 0 {
		addr = []string{fmt.Sprintf("%s:%d", jid.Domain, ClientPort)}
		return
	}

	// Build list of "host:port" strings.
	for _, a := range addrs {
		target := strings.TrimRight(a.Target, ".")
		addr = append(addr, fmt.Sprintf("%s:%d", target, a.Port))
	}
	return
}
