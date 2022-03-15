package cron

import (
	"context"
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/api"
	"github.com/Ptt-official-app/go-openbbsmiddleware/mand"
	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/sirupsen/logrus"
)

func RetryLoadManArticles(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			logrus.Infof("RetryLoadManArticles: to LoadManArticles")
			_ = LoadManArticles()
			select {
			case <-ctx.Done():
				return nil
			default:
				logrus.Infof("RetryLoadManArticles: to sleep 1 min")
				time.Sleep(1 * time.Minute)
			}
		}
	}
}

func LoadManArticles() (err error) {
	nextBrdname := ""
	count := 0

	for {
		boardIDs, err := schema.GetBoardIDs(nextBrdname, false, N_BOARDS+1, true)
		if err != nil {
			return err
		}

		newNextBrdname := ""
		if len(boardIDs) == N_BOARDS+1 {
			newNextBoardID := boardIDs[N_BOARDS]
			newNextBrdname = newNextBoardID.Brdname
			boardIDs = boardIDs[:N_BOARDS]
		}

		for _, each := range boardIDs {
			err = loadManArticles(each.BBoardID)
			if err == nil {
				count++
			}
		}

		if newNextBrdname == "" {
			logrus.Infof("cron.LoadManArticle: load %v boards", count)
			return nil

		}

		nextBrdname = newNextBrdname
	}
}

func loadManArticles(boardID bbs.BBoardID) (err error) {
	nextIdx := int32(0)

	count, err := loadManArticlesCore(boardID, "")
	if err != nil {
		logrus.Errorf("cron.loadManArticles: unable to loadManArticles: nextIdx: %v e: %v", nextIdx, err)
		return err
	}

	logrus.Infof("cron.loadManArticles: bid: %v count: %v", boardID, count)

	return nil
}

func loadManArticlesCore(boardID bbs.BBoardID, parentID types.ManArticleID) (count int, err error) {
	// backend load-general-articles
	brdname := boardID.ToBrdname()
	ctx := context.Background()
	req := &mand.ListRequest{
		BoardName: brdname,
		Path:      string(parentID),
	}
	resp, err := mand.Cli.List(ctx, req)
	if err != nil {
		return 0, err
	}

	entries := resp.Entries

	// update to db
	updateNanoTS := types.NowNanoTS()
	articleSummaries, err := api.DeserializePBManArticlesAndUpdateDB(boardID, parentID, entries, updateNanoTS)
	if err != nil {
		return 0, err
	}

	count += len(articleSummaries)

	// recursively loop through dirs.
	for _, each := range articleSummaries {
		if !each.IsDir {
			continue
		}
		eachCount, err := loadManArticlesCore(boardID, each.ArticleID)
		if err != nil {
			logrus.Errorf("cron.loadManArticlesCore: unable to loadManArticlesCore: boardID: %v path: %v each: %v e: %v", boardID, parentID, each.ArticleID, err)
		}

		count += eachCount
	}

	return count, nil
}
