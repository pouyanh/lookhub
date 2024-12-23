package user

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/janstoon/toolbox/bricks"

	"gitlab.snapp.ir/pouyanh/lookhub/drivers/api/user/restful/restapi/operations/dnslv"
)

func (s server) handleDnslvLookup(params dnslv.DnslvLookupParams) middleware.Responder {
	tfm := s.transformer()

	//  return dnslv.NewDnslvLookupOK().
	//	  WithPayload()

	return tfm.errorResponse(params.HTTPRequest.Context(), dnslv.NewDnslvLookupDefault(0), bricks.ErrUnimplemented)
}
