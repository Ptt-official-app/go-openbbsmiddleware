package api

import (
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const LOAD_GENERAL_BOARDS_BY_CLASS_R = "/boards/byclass"

func LoadGeneralBoardsByClassWrapper(c *gin.Context) {
	params := NewLoadGeneralBoardsParams()
	LoginRequiredQuery(LoadGeneralBoardsByClass, params, c)
}

func LoadGeneralBoardsByClass(remoteAddr string, userID bbs.UUserID, params interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	return loadGeneralBoardsCore(remoteAddr, userID, params, c, pttbbsapi.LOAD_GENERAL_BOARDS_BY_CLASS_R)
}
