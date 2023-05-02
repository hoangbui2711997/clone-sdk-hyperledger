package integration

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

// work?
func TestExecMint(t *testing.T) {
	c, err := channel.New(getContext("mychannel", "admin", "Org1"))
	logrus.Infoln(c, err)

	resp, err := c.Execute(channel.Request{
		ChaincodeID: "token_erc721",
		//Fcn:         "Initialize",
		Fcn: "MintWithTokenURI",
		//Fcn: "ClientAccountBalance",
		Args: [][]byte{
			[]byte("102"),
			[]byte("https://example.com/nft101.json"),
		},
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	})

	_ = resp
	logrus.Infoln(resp.ChaincodeStatus, resp.TxValidationCode, len(resp.Responses), len(resp.Payload), string(resp.Payload), err, ": test/integration/env_test.go:130")
	println()
	for _, resp := range resp.Responses {
		logrus.Infoln(
			resp.Response.Message,
			resp.Response.Status,
			string(resp.Response.Payload),
		)
	}

	//TestQueryClientID(t)
}

func TestExecTransfer(t *testing.T) {
	minter := GetClientIDAcc("mychannel", "admin", "org1")
	recipient := GetClientIDAcc("mychannel", "admin", "org2")

	c, err := channel.New(getContext("mychannel", "admin", "org1"))
	logrus.Infoln(c, err)
	resp, err := getInfoToken(c, "101")
	assert.Equal(t, nil, err)

	logrus.Infoln(string(resp.Payload))

	resp, err = c.Execute(channel.Request{
		ChaincodeID: "token_erc721",
		//Fcn:         "Initialize",
		Fcn: "TransferFrom",
		//Fcn: "ClientAccountBalance",
		Args: [][]byte{
			[]byte(minter),
			[]byte(recipient),
			[]byte("101"),
		},
		TransientMap:    nil,
		InvocationChain: nil,
		IsInit:          false,
	})

	_ = resp
	assert.Equal(t, nil, err)
	s, _ := status.FromError(err)
	logrus.Infoln(s.Code, s.Message)
	logrus.Infoln(resp.ChaincodeStatus)
	logrus.Infoln(resp.TxValidationCode)
	logrus.Infoln(err, ": test/integration/env_test.go:130")
	println()
	for _, resp := range resp.Responses {
		logrus.Infoln(
			resp.Response.Message,
			resp.Response.Status,
			string(resp.Response.Payload),
		)
	}

	//TestQueryClientID(t)
}

func getInfoToken(c *channel.Client, tokenID string) (channel.Response, error) {
	return c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "TokenURI",
		Args: [][]byte{
			//[]byte(strconv.FormatInt(time.Now().Unix(), 10)),
			[]byte(tokenID),
		},
	})
}

func getOwnerOfToken(c *channel.Client, tokenID string) (channel.Response, error) {
	return c.Query(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "OwnerOf",
		Args: [][]byte{
			[]byte(tokenID),
		},
	})
}
