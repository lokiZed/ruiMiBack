package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"ruiMiBack2/database"
	"ruiMiBack2/internal/util"
	"ruiMiBack2/models/challenge"
	"ruiMiBack2/models/challengeUser"
	"ruiMiBack2/models/user"
	"time"

	"ruiMiBack2/internal/config"
	"ruiMiBack2/internal/handler"
	"ruiMiBack2/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/ruimi.yaml", "the config file")

func main() {
	flag.Parse()

	c := config.GetConfig()

	conf.MustLoad(*configFile, c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(*c)
	handler.RegisterHandlers(server, ctx)

	if err := database.InitDatabase(c.DataBase); err != nil {
		log.Fatal(err.Error())
	}

	generateUsers(40)
	generateChallenge(1, 5, 10)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func generateChallenge(gameId, gameLevel, canNum int64, account ...string) {
	db := database.GetMysqlConn()
	ctx := context.Background()
	cache := map[string]struct{}{}

	for _, ac := range account {
		cache[ac] = struct{}{}
	}

	challengeModel := challenge.NewChallengeModel(db)
	newChallenge := &challenge.Challenge{
		GameId:    sql.NullInt64{Int64: gameId, Valid: true},
		GameLevel: sql.NullInt64{Int64: gameLevel, Valid: true},
	}
	result, _ := challengeModel.Insert(ctx, newChallenge)
	newChallengeId, _ := result.LastInsertId()

	userModel := user.NewUserModel(db)
	allUsers, _ := userModel.FindMany(ctx)

	challengeUserModel := challengeUser.NewChallengeUserModel(db)
	for _, u := range allUsers {
		_, ok := cache[u.AccountName]
		if len(cache) != 0 && ok {
			continue
		}
		newChallengeUser := &challengeUser.ChallengeUser{
			UserId:      sql.NullInt64{Int64: u.Id, Valid: true},
			ChallengeId: sql.NullInt64{Int64: newChallengeId, Valid: true},
			LeaveNum:    sql.NullInt64{Int64: canNum, Valid: true},
		}
		_, _ = challengeUserModel.Insert(ctx, newChallengeUser)
	}
}

func generateUsers(num int) {
	filePath := fmt.Sprintf("./gen_user%s.txt", time.Now().Format("2006_01_01_15_04_05"))
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0333)
	if err != nil {
		return
	}
	defer file.Close()
	userModel := user.NewUserModel(database.GetMysqlConn())
	ctx := context.Background()
	for i := 0; i < num; i++ {
		newUser := &user.User{
			AccountName: util.GetUuid(),
			AccountPass: util.GetUuid(),
		}
		_, _ = userModel.Insert(ctx, newUser)
		_, _ = file.WriteString(fmt.Sprintf("%s----%s\n", newUser.AccountName, newUser.AccountPass))
	}
}
