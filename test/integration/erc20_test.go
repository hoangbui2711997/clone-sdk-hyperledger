package integration

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// transfer erc20
func TestMintERC20(t *testing.T) {
	logrus.SetReportCaller(true)

	t.Run("transfer erc20", func(t *testing.T) {
		c, _ := channel.New(getContext("mychannel", "admin", "Org1"))

		resp, err := c.Execute(channel.Request{
			ChaincodeID: "token_erc20",
			Fcn:         "Mint",
			Args: [][]byte{
				[]byte("105000"),
			},
		})

		_ = resp
		assert.NotNil(t, err)
	})
}

// test get balance erc20
func TestTransferERC20(t *testing.T) {
	logrus.SetReportCaller(true)

	t.Run("TestTransferERC20", func(t *testing.T) {
		c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
		c1, _ := channel.New(getContext("mychannel", "admin", "Org2"))

		balBefore, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)
		balBefore1, _ := strconv.ParseFloat(GetBalanceERC20(c1), 64)
		//minter := GetClientIDAccErc20("mychannel", "admin", "Org1")
		recipient := GetClientIDAccErc20("mychannel", "admin", "Org2")

		resp, err := c.Execute(channel.Request{
			ChaincodeID: "token_erc20",
			Fcn:         "Transfer",
			Args: [][]byte{
				[]byte(recipient),
				[]byte("500"),
			},
		})
		_ = resp
		balAfter, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)
		balAfter1, _ := strconv.ParseFloat(GetBalanceERC20(c1), 64)

		assert.Nil(t, err)
		assert.Equal(t, float64(500), balBefore-balAfter)
		assert.Equal(t, float64(500), balAfter1-balBefore1)
	})
}

func TestTransferERC20WithInvalidAmount(t *testing.T) {
	logrus.SetReportCaller(true)

	t.Run("TestTransferERC20WithInvalidAmount", func(t *testing.T) {
		c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
		c1, _ := channel.New(getContext("mychannel", "admin", "Org2"))

		balBefore, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)
		balBefore1, _ := strconv.ParseFloat(GetBalanceERC20(c1), 64)
		//minter := GetClientIDAccErc20("mychannel", "admin", "Org1")
		recipient := GetClientIDAccErc20("mychannel", "admin", "Org2")

		resp, err := c.Execute(channel.Request{
			ChaincodeID: "token_erc20",
			Fcn:         "Transfer",
			Args: [][]byte{
				[]byte(recipient),
				[]byte("50000000"),
			},
		})
		_ = resp
		balAfter, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)
		balAfter1, _ := strconv.ParseFloat(GetBalanceERC20(c1), 64)

		assert.NotNil(t, err)
		assert.Equal(t, float64(0), balBefore-balAfter)
		assert.Equal(t, float64(0), balAfter1-balBefore1)
	})
}

func TestTransferERC20WithInvalidRecipient(t *testing.T) {
	logrus.SetReportCaller(true)

	t.Run("TestTransferERC20WithInvalidRecipient", func(t *testing.T) {
		c, _ := channel.New(getContext("mychannel", "admin", "Org1"))
		balBefore, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)

		resp, err := c.Execute(channel.Request{
			ChaincodeID: "token_erc20",
			Fcn:         "Transfer",
			Args: [][]byte{
				[]byte("123"),
				[]byte("500"),
			},
		})
		_ = resp
		balAfter, _ := strconv.ParseFloat(GetBalanceERC20(c), 64)

		logrus.Infoln("balBefore", balBefore, "balAfter", balAfter)
		assert.NotNil(t, err)
		assert.Equal(t, float64(500), balBefore-balAfter)
	})
}

// get balance erc20
func GetBalanceERC20(c *channel.Client) string {
	resp, _ := c.Query(channel.Request{
		ChaincodeID: "token_erc20",
		Fcn:         "ClientAccountBalance",
		Args:        [][]byte{},
	})

	return string(resp.Payload)
}
