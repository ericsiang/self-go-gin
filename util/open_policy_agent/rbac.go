package opa

import (
	"context"
	_ "embed"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/rego"
)

//go:embed rbac.rego
var policy []byte
var defaultResult rego.ResultSet

func readPolicy() ([]byte, error) {
	return policy, nil
}

func GetQueryResult(c *gin.Context) (rego.ResultSet, error) {
	policy, err := readPolicy()
	if err != nil {
		return defaultResult, errors.New("failed to readPolicy : ")
	}
	// prepare rego query
	query, err := rego.New(
		rego.Query("data.rbac.allow"),
		rego.Module("rbac.rego", string(policy)),
	).PrepareForEval(c)

	if err != nil {
		err1 := errors.New("failed to prepare rbac policy : ")
		err = errors.Join(err1, err)
		return defaultResult, err
	}

	// evaluate rego query by supplying values extracted from header
	result, err := query.Eval(context.Background(), rego.EvalInput(map[string]interface{}{
		"role":     c.Request.Header.Get("role"),
		"action":   c.Request.Header.Get("action"),
		"resource": c.Request.Header.Get("resource"),
	}))

	if err != nil {
		err1 := errors.New("failed to query eval : ")
		err = errors.Join(err1, err)
		return defaultResult, err
	}

	return result, nil
}
