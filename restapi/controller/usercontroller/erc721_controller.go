package usercontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"net/http"
	"strconv"
	"time"
)

// @Title Mint
// @Summary Mint
// @Description Mint
// @Tags erc721
// @Accept  json
// @Produce  json
// @Success 200	{object} responses.ResponseCommonSingle
// @Failure 403 not found
// @router /erc721/mint [post]
func Mint(c *gin.Context) {
	client, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	tokenID := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := client.Execute(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "MintWithTokenURI",
		Args: [][]byte{
			[]byte(tokenID),
			[]byte(fmt.Sprintf("https://example.com/nft%s.json", tokenID)),
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	resp, err := GetInfoToken(client, tokenID)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":     string(resp.Payload),
		"token_id": tokenID,
	})
	return
}

// @Title Transfer
// @Summary Transfer
// @Description Transfer
// @Tags erc721
// @Accept  json
// @Produce  json
// @Success 200	{object} responses.ResponseCommonSingle
// @Failure 403 not found
// @router /erc721/transfer [post]
func Transfer(c *gin.Context) {
	client, _ := channel.New(getContext("mychannel", "admin", "Org1"))
	tokenID := strconv.FormatInt(time.Now().Unix(), 10)

	_, err := client.Execute(channel.Request{
		ChaincodeID: "token_erc721",
		Fcn:         "MintWithTokenURI",
		Args: [][]byte{
			[]byte(tokenID),
			[]byte(fmt.Sprintf("https://example.com/nft%s.json", tokenID)),
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	minter := GetClientIDAcc("mychannel", "admin", "org1")
	recipient := GetClientIDAcc("mychannel", "admin", "org2")

	_, err = client.Execute(channel.Request{
		ChaincodeID: "token_erc721",
		//Fcn:         "Initialize",
		Fcn: "TransferFrom",
		//Fcn: "ClientAccountBalance",
		Args: [][]byte{
			[]byte(minter),
			[]byte(recipient),
			[]byte(tokenID),
		},
	})

	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	token, err := GetOwnerOfToken(client, tokenID)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data":      string(token.Payload),
		"token_id":  tokenID,
		"recipient": recipient,
		"minter":    minter,
	})

	return
}
