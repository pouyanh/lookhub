package dnslv

import (
	"errors"
	"fmt"
	"strings"

	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/tricks"
	"github.com/miekg/dns"
)

var (
	ErrInvalidDomainName  = errors.Join(bricks.ErrInvalidArgument, errors.New("invalid domain name"))
	ErrInvalidRecordType  = errors.Join(bricks.ErrInvalidArgument, errors.New("invalid resource record type"))
	ErrInvalidRecordTTL   = errors.Join(bricks.ErrInvalidArgument, errors.New("invalid resource record ttl"))
	ErrInvalidRecordValue = errors.Join(bricks.ErrInvalidArgument, errors.New("invalid resource record value"))
)

// Domain represents a domain and all its associated DNS information. It's the consistency boundary.
// Any changes to resource records should happen through the Domain aggregate root.
type Domain struct {
	name    DomainName       // Identity of the Aggregate
	records []ResourceRecord // Collection of resource records
}

// DomainByName creates a new Domain with specified name and empty records and NameServers
func DomainByName(name DomainName) (*Domain, error) {
	if err := name.validate(); err != nil {
		return nil, err
	}

	return &Domain{
		name:    name,
		records: make([]ResourceRecord, 0),
	}, nil
}

// AddRecords adds ResourceRecord(s) to Domain's records list. It's not concurrent-safe
func (d *Domain) AddRecords(records ...ResourceRecord) error {
	for _, record := range records {
		if err := record.validate(); err != nil {
			return err
		}
	}

	d.records = append(d.records, records...)

	return nil
}

// Name returns domain name
func (d *Domain) Name() DomainName {
	return d.name
}

func (d *Domain) Records() []ResourceRecord {
	return tricks.Copy(d.records)
}

// RecordsWithType returns Domain's DNS Records with ResourceRecord.Type matching t
func (d *Domain) RecordsWithType(t ResourceRecordType) []ResourceRecord {
	return tricks.Filter(d.records, func(src ResourceRecord) bool {
		return src.Type == t
	})
}

type DomainName string

func (dn DomainName) validate() error {
	if len(strings.TrimSpace(string(dn))) == 0 {
		return errors.Join(ErrInvalidDomainName, errors.New("name cannot be empty"))
	}

	if _, ok := dns.IsDomainName(string(dn)); !ok {
		return ErrInvalidDomainName
	}

	return nil
}

func (dn DomainName) equal(v DomainName) bool {
	return strings.ToLower(string(dn)) == strings.ToLower(string(v))
}

// ResourceRecord holds a single DNS record information, identified by combination of Type and Value.
type ResourceRecord struct {
	// Record type
	Type ResourceRecordType

	// Keeps the data associated with the record, e.g., an IP address for an A record, a hostname for a CNAME record
	Value string

	// Specifies how long (in seconds) a DNS resolver should cache the result of a DNS query.
	TTL ResourceRecordTTL
}

func (r ResourceRecord) validate() error {
	if err := r.Type.validate(); err != nil {
		return err
	}

	if err := r.TTL.validate(); err != nil {
		return err
	}

	if len(strings.TrimSpace(r.Value)) == 0 {
		return errors.Join(ErrInvalidRecordValue, errors.New("value cannot be empty"))
	}

	return nil
}

type ResourceRecordType string

func (t ResourceRecordType) validate() error {
	if len(strings.TrimSpace(string(t))) == 0 {
		return errors.Join(ErrInvalidRecordType, errors.New("type cannot be empty"))
	}

	return nil
}

// ResourceRecordTTL specifies how long (in seconds) a client should cache the result of a query.
type ResourceRecordTTL int

const (
	ZeroTTL ResourceRecordTTL = 0          // No caching
	MaxTTL  ResourceRecordTTL = 2147483647 // Approximately 68 years
)

func (ttl ResourceRecordTTL) validate() error {
	if ttl < ZeroTTL || ttl > MaxTTL {
		return errors.Join(ErrInvalidRecordTTL, fmt.Errorf("value %d not in range [%d, %d]", ttl, ZeroTTL, MaxTTL))
	}

	return nil
}
