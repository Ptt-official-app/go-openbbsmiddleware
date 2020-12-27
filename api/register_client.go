package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const REGISTER_CLIENT_R = "/client/register"

type RegisterClientParams struct {
	ClientID string `json:"client_id"`
}

type RegisterClientResult struct {
	ClientSecret string `json:"client_secret"`
}

func NewRegisterClientParams() *RegisterClientParams {
	return &RegisterClientParams{}
}

func RegisterClientWrapper(c *gin.Context) {
	params := NewRegisterClientParams()
	LoginRequiredJSON(RegisterClient, params, c)
}

func RegisterClient(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	RegisterClientParams, ok := params.(*RegisterClientParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if userID != SYSOP {
		return nil, 401, ErrInvalidToken
	}

	client, err := deserializeClientAndUpdateDB(RegisterClientParams.ClientID, remoteAddr)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = NewRegisterClientResult(client)
	return result, 200, nil
}

func deserializeClientAndUpdateDB(clientID string, remoteAddr string) (client *schema.Client, err error) {

	client = schema.NewClient(clientID, remoteAddr)

	err = schema.UpdateClient(client)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewRegisterClientResult(client *schema.Client) *RegisterClientResult {
	return &RegisterClientResult{
		ClientSecret: client.ClientSecret,
	}
}
