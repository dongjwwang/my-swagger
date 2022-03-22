/**
@Author: dongjwwang@Tecent.com
@Time: 2022/3/22 22:26
@File: policy
@Project: my-swagger
**/

package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"my-swagger/src/pkg/v1/models"
	"my-swagger/src/pkg/v1/server/operations/policy"
)

type PolicyHandler interface {
	AddPolicy(params policy.AddPolicyParams, principal interface{}) middleware.Responder
}

type policyHandler struct {
}

func NewPolicyHandle() PolicyHandler {
	return &policyHandler{}
}

func (p *policyHandler) AddPolicy(params policy.AddPolicyParams, principal interface{}) middleware.Responder {
	if *params.XRequestID == "500" {
		return policy.NewAddPolicyInternalServerError().
			WithXRequestID(*params.XRequestID).
			WithPayload(&models.Errors{
				Errors: []*models.Error{
					{
						Code:    "500",
						Message: "InternalError",
					},
				},
			})
	}
	return policy.NewAddPolicyInternalServerError().
		WithXRequestID(*params.XRequestID)
}
