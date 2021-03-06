package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	"github.com/gin-gonic/gin"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
)

type userBoardInfo struct {
	Read  bool
	Stat  ptttype.BoardStatAttr
	NUser int
}

func deserializeBoardsAndUpdateDB(userID bbs.UUserID, boardSummaries_b []*bbs.BoardSummary, updateNanoTS types.NanoTS) (boardSummaries []*schema.BoardSummary, userBoardInfoMap map[bbs.BBoardID]*userBoardInfo, err error) {

	if len(boardSummaries_b) == 0 {
		return nil, nil, nil
	}

	boardSummaries = make([]*schema.BoardSummary, 0, len(boardSummaries_b))
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}
		each := schema.NewBoardSummary(each_b, updateNanoTS)

		boardSummaries = append(boardSummaries, each)
	}

	err = schema.UpdateBoardSummaries(boardSummaries, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	userReadBoards := make([]*schema.UserReadBoard, 0, len(boardSummaries_b))
	userBoardInfoMap = make(map[bbs.BBoardID]*userBoardInfo)
	for _, each_b := range boardSummaries_b {
		if each_b.Reason != 0 {
			continue
		}

		userBoardInfoMap[each_b.BBoardID] = &userBoardInfo{
			Read:  each_b.Read,
			Stat:  each_b.StatAttr,
			NUser: int(each_b.NUser),
		}

		if each_b.Read {
			each_db := &schema.UserReadBoard{
				UserID:       userID,
				BBoardID:     each_b.BBoardID,
				UpdateNanoTS: updateNanoTS,
			}
			userReadBoards = append(userReadBoards, each_db)
		}
	}

	err = schema.UpdateUserReadBoards(userReadBoards, updateNanoTS)
	if err != nil {
		return nil, nil, err
	}

	return boardSummaries, userBoardInfoMap, err
}

func isBoardValidUser(boardID bbs.BBoardID, c *gin.Context) (isValid bool, statusCode int, err error) {

	var result_b *pttbbsapi.IsBoardValidUserResult

	urlMap := map[string]string{
		"bid": string(boardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.IS_BOARD_VALID_USER_R)
	statusCode, err = utils.BackendGet(c, url, nil, nil, &result_b)
	if err != nil || statusCode != 200 {
		return false, statusCode, err
	}
	if !result_b.IsValid {
		return false, 403, ErrInvalidUser
	}

	return true, 200, nil

}
