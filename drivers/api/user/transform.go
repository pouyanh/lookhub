package user

import (
	"context"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/janstoon/toolbox/handywares"
	"github.com/janstoon/toolbox/tricks"
	"go.opentelemetry.io/otel/trace"

	"github.com/pouyanh/lookhub/drivers/api/user/restful/models"
)

type transformer struct{}

func (s server) transformer() transformer {
	return transformer{}
}

// ----------------------
//         Error
// ----------------------

type errorResponder interface {
	middleware.Responder
	SetPayload(err *models.Error)
	SetStatusCode(code int)
}

func (t *transformer) errorResponse(ctx context.Context, rw errorResponder, err error) errorResponder {
	rw.SetStatusCode(t.errorToGwStatusCode(err))
	rw.SetPayload(tricks.ApplyOptions(tricks.ValPtr(t.errorToGw(err)),
		t.ctxToGwErrorOption(ctx)))

	return rw
}

func (t *transformer) ctxToGwErrorOption(ctx context.Context) tricks.InPlaceOption[models.Error] {
	traceId := trace.SpanFromContext(ctx).SpanContext().TraceID().String()

	return func(src *models.Error) {
		src.TraceID = traceId
	}
}

func (t *transformer) errorToGwStatusCode(err error) int {
	return handywares.BricksErrorToHttpStatusMapper(err)
}

func (t *transformer) errorToGw(src error) models.Error {
	return models.Error{
		Code:    "",
		Message: strings.ReplaceAll(src.Error(), "\n", ". "),
	}
}
