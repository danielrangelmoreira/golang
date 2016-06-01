package igprotocol

import (
	"fmt"
	"testing"
	"encoding/json"
	 testvariables "bitbucket.org/bhbosman/golangtradingig/igprotocol/test" 
	

)

func TestMarketNavigationAll(t *testing.T) {
	test := MyTestingT{t}

	ctx := NewIGContextForTesting(
		testvariables.TestAccountIdentifier,
		testvariables.TestAccountPassword,
		testvariables.TestAccountAPIKey,
		t)

	_, err := ctx.Login()
	test.CheckErrorWithMessage(err, "IG Context did not connect")

	defer func() {
		_, err = ctx.Logout()
		test.CheckErrorWithMessage(err, "IG Context did not disconnect")
	}()

	keys := ctx.ConnectionContext.CreateConnectionContext()
	response, err := SendMarketNavigationRequest(
		&ctx.ConnectionContext,
		MarketNavigationRequest{},
		keys)
	test.CheckErrorWithMessage(err, "A")
	test.CheckBool(response.Header.Success, response.Header.ErrorCode)
}

func TestMarketNavigationPerInstance(t *testing.T) {
	test := MyTestingT{t}

	ctx := NewIGContextForTesting(
		testvariables.TestAccountIdentifier,
		testvariables.TestAccountPassword,
		testvariables.TestAccountAPIKey,
		t)

	_, err := ctx.Login()
	test.CheckErrorWithMessage(err, "IG Context did not connect")

	// Defer logout
	defer func() {
		_, err = ctx.Logout()
		test.CheckErrorWithMessage(err, "IG Context did not disconnect")
	}()

	result, err := ctx.SendMarketNavigationRequest("")
	test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
	test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
	test.CheckBool(err == nil, "Error must be assigned")

	test.CheckBool(len(result.Data.Nodes) > 0, "Node count > 0")
	for i, node := range result.Data.Nodes {
		test.Logf("Instance %d: NodeID: %s, Node Name: %s", i, node.ID, node.Name)
		//
		result, err = ctx.SendMarketNavigationRequest(node.ID)
		//
		test.CheckBool(result.Header.Success, fmt.Sprintf("Must be Successfull(%s, %s)", result.Header.ErrorCode, result.Header.ErrorDescription))
		test.CheckIntWithMessage(200, result.Header.StatusCode, "Status code must be 200")
		test.CheckBool(err == nil, "Error must be assigned")
	}
}



func TestGetMarkets(t *testing.T) {
	test := MyTestingT{t}

	ctx := NewIGContextForTesting(
		testvariables.TestAccountIdentifier,
		testvariables.TestAccountPassword,
		testvariables.TestAccountAPIKey,
		t)

	_, err := ctx.Login()
	test.CheckErrorWithMessage(err, "IG Context did not connect")

	// Defer logout
	defer func() {
		_, err = ctx.Logout()
		test.CheckErrorWithMessage(err, "IG Context did not disconnect")
	}()

	result, err := ctx.GetMarkets("", true)
	t.Log(json.Marshal(result))
	
}
