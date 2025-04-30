package dnslv_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pouyanh/lookhub/dnslv"
)

func TestCreateDomain(t *testing.T) {
	d, err := dnslv.DomainByName("")
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidDomainName)

	d, err = dnslv.DomainByName(".pouyan")
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidDomainName)

	d, err = dnslv.DomainByName("pouyan..dev")
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidDomainName)

	d, err = dnslv.DomainByName("pouyan.dev")
	require.NoError(t, err)
	assert.EqualValues(t, "pouyan.dev", d.Name())
}

func TestAddRecords(t *testing.T) {
	d, err := dnslv.DomainByName("pouyan.dev")
	require.NoError(t, err)

	record := dnslv.ResourceRecord{}
	err = d.AddRecords(record)
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidRecordType)
	assert.Empty(t, d.Records())

	record = dnslv.ResourceRecord{
		Type: "A",
		TTL:  -1,
	}
	err = d.AddRecords(record)
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidRecordTTL)
	assert.Empty(t, d.Records())

	record = dnslv.ResourceRecord{
		Type: "A",
		TTL:  dnslv.MaxTTL + 1,
	}
	err = d.AddRecords(record)
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidRecordTTL)
	assert.Empty(t, d.Records())

	record = dnslv.ResourceRecord{
		Type: "A",
		TTL:  300,
	}
	err = d.AddRecords(record)
	require.Error(t, err)
	require.ErrorIs(t, err, dnslv.ErrInvalidRecordValue)
	assert.Empty(t, d.Records())

	record = dnslv.ResourceRecord{
		Type:  "A",
		Value: "10.0.0.1",
		TTL:   300,
	}
	err = d.AddRecords(record)
	require.NoError(t, err)
	assert.NotEmpty(t, d.Records())
	assert.Contains(t, d.Records(), record)
}

func TestRemoveRecord(t *testing.T) {
	d, err := dnslv.DomainByName("pouyan.dev")
	require.NoError(t, err)

	// todo: remove a record
	_ = d
}

func TestUpdateRecord(t *testing.T) {
	d, err := dnslv.DomainByName("pouyan.dev")
	require.NoError(t, err)

	// todo: update a record
	_ = d
}
