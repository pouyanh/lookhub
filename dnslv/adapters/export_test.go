package adapters

type DNSClient = dnsClient

func NewDNSClient(server string) DNSClient {
	return newDNSClient(server)
}
