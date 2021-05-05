package schema

import (
	"context"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	IsTest = false
	lock   sync.Mutex
)

func SetIsTest() {
	lock.Lock()
	IsTest = true

	logrus.Info("schema: SetIsTest")

	err := Init()
	if err != nil {
		logrus.Errorf("schema.Init: unable to init: e: %v", err)
	}

	testResetDB()
}

func UnsetIsTest() {
	Close()

	logrus.Info("schema: UnsetIsTest")

	IsTest = false
	lock.Unlock()
}

func testResetDB() {
	_ = AccessToken_c.Drop()
	_ = Article_c.Drop()
	_ = Board_c.Drop()
	_ = BoardBanuser_c.Drop()
	_ = BoardFriend_c.Drop()
	_ = Client_c.Drop()
	_ = Comment_c.Drop()
	_ = User_c.Drop()
	_ = UserAloha_c.Drop()
	_ = UserFavorites_c.Drop()
	_ = UserFavoritesMeta_c.Drop()
	_ = UserFriend_c.Drop()
	_ = UserReadArticle_c.Drop()
	_ = UserReadBoard_c.Drop()
	_ = UserReject_c.Drop()
	_ = UserIDEmail_c.Drop()
	_ = UserEmail_c.Drop()
	_ = Rank_c.Drop()

	ctx, cancel := context.WithTimeout(context.Background(), REDIS_TIMEOUT_MILLI_TS*time.Millisecond)
	defer func() {
		cancel()
	}()

	_, _ = rdb.FlushDB(ctx).Result()
}
