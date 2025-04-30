package adapters

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/janstoon/toolbox/bricks"
	"github.com/miekg/dns"
	"golang.org/x/sync/errgroup"

	"github.com/pouyanh/lookhub/dnslv"
)

type dnsClient struct {
	server   string
	resolver *dns.Client
}

func newDNSClient(server string) dnsClient {
	return dnsClient{
		server:   server + ":53",
		resolver: new(dns.Client),
	}
}

func (c dnsClient) QueryAllDNSRecords(ctx context.Context, domainName string) ([]dnslv.ResourceRecord, error) {
	var wg errgroup.Group

	chRecords := make(chan dnslv.ResourceRecord, 20)
	wg.Go(c.knownRecordsQuery(ctx, chRecords, domainName))

	records := make([]dnslv.ResourceRecord, 0)
	wg.Go(func() error {
		for {
			record, ok := <-chRecords
			if !ok {
				break
			}

			records = append(records, record)
		}

		return nil
	})

	err := wg.Wait()
	if err != nil && len(records) == 0 {
		return nil, err
	}

	return records, nil
}

func (c dnsClient) knownRecordsQuery(
	ctx context.Context, records chan<- dnslv.ResourceRecord, domainName string,
) func() error {
	return func() error {
		defer close(records)

		var wg errgroup.Group
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeA, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeNS, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeAAAA, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeMX, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeCNAME, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeTXT, records))
		wg.Go(c.recordQuery(ctx, domainName, dns.TypeSOA, records))

		return wg.Wait()
	}
}

func (c dnsClient) recordQuery(
	ctx context.Context, domainName string, t uint16, records chan<- dnslv.ResourceRecord,
) func() error {
	return func() error {
		m := new(dns.Msg)
		m.SetQuestion(dns.Fqdn(domainName), t)

		r, _, err := c.resolver.ExchangeContext(ctx, m, c.server)
		if err != nil {
			return err
		}

		if r.Rcode != dns.RcodeSuccess {
			return errors.Join(rcodeToBricksError[r.Rcode], fmt.Errorf("DNS query for %v failed with code %s(%v)",
				dns.TypeToString[t], dns.RcodeToString[r.Rcode], r.Rcode))
		}

		for _, ans := range r.Answer {
			records <- c.parseRR(ans)
		}

		return nil
	}
}

func (c dnsClient) parseRR(rr dns.RR) dnslv.ResourceRecord {
	value := rr.String()
	switch resource := rr.(type) {
	case *dns.A:
		value = resource.A.String()

	case *dns.AAAA:
		value = resource.AAAA.String()

	case *dns.CNAME:
		value = resource.Target

	case *dns.MX:
		value = fmt.Sprintf("%d %s", resource.Preference, resource.Mx)

	case *dns.NS:
		value = resource.Ns

	case *dns.SOA:
		value = fmt.Sprintf("%s %s %d %d %d %d %d",
			resource.Ns, resource.Mbox, resource.Serial, resource.Refresh, resource.Retry, resource.Expire, resource.Minttl)

	case *dns.TXT:
		value = strings.Join(resource.Txt, " ")
	}

	return dnslv.ResourceRecord{
		Type:  dnslv.ResourceRecordType(dns.TypeToString[rr.Header().Rrtype]),
		Value: value,
		TTL:   dnslv.ResourceRecordTTL(rr.Header().Ttl),
	}
}

var rcodeToBricksError = map[int]error{
	dns.RcodeSuccess:        nil,
	dns.RcodeFormatError:    bricks.ErrInvalidArgument,
	dns.RcodeServerFailure:  bricks.ErrInternal,
	dns.RcodeNameError:      bricks.ErrNotFound,
	dns.RcodeNotImplemented: bricks.ErrUnimplemented,
	dns.RcodeRefused:        bricks.ErrAborted,
	dns.RcodeYXDomain:       bricks.ErrInternal,
	dns.RcodeYXRrset:        bricks.ErrInternal,
	dns.RcodeNXRrset:        bricks.ErrNotFound,
	dns.RcodeNotAuth:        bricks.ErrPermissionDenied,
	dns.RcodeNotZone:        bricks.ErrNotFound,
	dns.RcodeBadSig:         bricks.ErrPermissionDenied,
	dns.RcodeBadKey:         bricks.ErrUnauthenticated,
	dns.RcodeBadTime:        bricks.ErrPermissionDenied,
	dns.RcodeBadMode:        bricks.ErrPermissionDenied,
	dns.RcodeBadName:        bricks.ErrAlreadyExists,
	dns.RcodeBadAlg:         bricks.ErrUnimplemented,
	dns.RcodeBadTrunc:       bricks.ErrDataLoss,
	dns.RcodeBadCookie:      bricks.ErrUnauthenticated,
}
