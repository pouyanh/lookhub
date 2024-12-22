package adapters_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv/adapters"
)

func TestDNSClient_QueryAllDNSRecords(t *testing.T) {
	if testing.Short() {
		return
	}

	c := adapters.NewDNSClient("8.8.8.8")

	records, err := c.QueryAllDNSRecords(context.Background(), "pouyan.dev")
	require.NoError(t, err)
	assert.NotEmpty(t, records)
}
