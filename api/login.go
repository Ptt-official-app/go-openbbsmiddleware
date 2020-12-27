package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOGIN_R = "/account/login"

type LoginParams struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func NewLoginParams() *LoginParams {
	return &LoginParams{}
}

type LoginResult struct {
	UserID      bbs.UUserID `json:"user_id"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
}

func LoginWrapper(c *gin.Context) {
	params := NewLoginParams()
	JSON(Login, params, c)
}

func Login(remoteAddr string, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {
	theParams, ok := params.(*LoginParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	if !isValidClient(theParams.ClientID, theParams.ClientSecret) {
		return nil, 400, ErrInvalidParams
	}

	//backend login
	theParams_b := &pttbbsapi.LoginParams{
		Username: theParams.Username,
		Passwd:   theParams.Password,
	}

	var result_b *pttbbsapi.LoginResult

	url := pttbbsapi.LOGIN_R
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update db
	accessToken, err := serializeAccessTokenAndUpdateDB(result_b.UserID, result_b.Jwt)
	if err != nil {
		return nil, 500, err
	}

	//result
	result = NewLoginResult(accessToken)

	return result, 200, nil
}

func NewLoginResult(accessToken_db *schema.AccessToken) *LoginResult {
	return &LoginResult{
		UserID:      accessToken_db.UserID,
		AccessToken: accessToken_db.AccessToken,
		TokenType:   "bearer",
	}
}
