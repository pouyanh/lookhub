package user

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/janstoon/toolbox/bricks"
	"github.com/janstoon/toolbox/tricks"
	"golang.org/x/sync/errgroup"

	"gitlab.snapp.ir/pouyanh/lookhub/dnslv"
	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user/restful/models"
	. "gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/dnslv"
	"gitlab.snapp.ir/pouyanh/lookhub/lutel"
)

func (s server) handleDNSLVLookup(params DnslvLookupParams) middleware.Responder {
	lutel.RequestCnt.Inc()

	tfm := s.transformer()

	ctx := params.HTTPRequest.Context()
	cmd := tricks.PtrVal(params.Lookup.Domain)
	results := make(chan *dnslv.Domain, len(s.apps))

	wg := errgroup.Group{}
	for _, app := range s.apps {
		func(app any) {
			handler, ok := app.(DNSLVLookup)
			if !ok {
				return
			}

			wg.Go(func() error {
				res, err := handler.Lookup(ctx, cmd)
				if err != nil {
					return err
				}

				results <- res

				return nil
			})
		}(app)
	}

	err := wg.Wait()
	close(results)
	if err != nil {
		return tfm.errorResponse(params.HTTPRequest.Context(), NewDnslvLookupDefault(0), err)
	}

	if len(results) == 0 {
		return tfm.errorResponse(params.HTTPRequest.Context(), NewDnslvLookupDefault(0), bricks.ErrUnavailable)
	}

	return NewDnslvLookupOK().
		WithPayload(&models.DnslvLookupRsp{
			Domain: tricks.PtrPtr(<-results /* only first result */, tfm.dnslvDomainToGW),
		})
}

func (t *transformer) dnslvDomainToGW(src dnslv.Domain) models.Domain {
	return models.Domain{
		Name: tricks.ValPtr(t.dnslvDomainNameToGW(src.Name())),
		ResourceRecords: tricks.Map(src.Records(), func(src dnslv.ResourceRecord) *models.ResourceRecord {
			return tricks.ValPtr(t.dnslvResourceRecordToGW(src))
		}),
	}
}

func (t *transformer) dnslvDomainNameToGW(src dnslv.DomainName) string {
	return src.String()
}

func (t *transformer) dnslvResourceRecordToGW(src dnslv.ResourceRecord) models.ResourceRecord {
	return models.ResourceRecord{
		Type:  tricks.ValPtr(t.dnslvResourceRecordTypeToGW(src.Type)),
		Value: tricks.ValPtr(src.Value),
		TTL:   tricks.ValPtr(t.dnslvResourceRecordTTLToGW(src.TTL)),
	}
}

func (t *transformer) dnslvResourceRecordTypeToGW(src dnslv.ResourceRecordType) string {
	return string(src)
}

func (t *transformer) dnslvResourceRecordTTLToGW(src dnslv.ResourceRecordTTL) int64 {
	return int64(src)
}
