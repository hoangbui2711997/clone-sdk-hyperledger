package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
)

// @Title Mint ERC20
// @Summary Mint ERC20
// @Description Mint ERC20
// @Tags erc20
// @Accept  json
// @Produce  json
// @Success 200	{object} responses.ResponseCommonSingle
// @Failure 403 not found
// @router /erc20/mint [post]
func MintERC20(c *gin.Context) {
	client, _ := channel.New(getContext("mychannel", "admin", "Org1"))

	_, err := client.Execute(channel.Request{
		ChaincodeID: "token_erc20",
		Fcn:         "Mint",
		Args: [][]byte{
			[]byte("5000"),
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	minter := GetClientIDAccErc20("mychannel", "admin", "Org1")

	res, err := client.Query(channel.Request{
		ChaincodeID: "token_erc20",
		Fcn:         "BalanceOf",
		Args: [][]byte{
			[]byte(minter),
		},
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": string(res.Payload),
	})
	return
}

// @Title Transfer ERC20
// @Summary Transfer ERC20
// @Description Transfer ERC20
// @Tags erc20
// @Accept  json
// @Produce  json
// @Success 200	{object} responses.ResponseCommonSingle
// @Failure 403 not found
// @router /erc20/transfer [post]
func TransferERC20(c *gin.Context) {
	client, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	recipient := GetClientIDAccErc20("mychannel", "admin", "Org2")

	_, err := client.Execute(channel.Request{
		ChaincodeID: "token_erc20",
		Fcn:         "Transfer",
		Args: [][]byte{
			[]byte(recipient),
			[]byte("1"),
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	res, err := client.Query(channel.Request{
		ChaincodeID: "token_erc20",
		Fcn:         "BalanceOf",
		Args: [][]byte{
			[]byte(recipient),
		},
	})

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": string(res.Payload),
	})
	return
}
