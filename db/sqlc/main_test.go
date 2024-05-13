package db

import (
	"GoBank/util"
	"context"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *pgx.Conn

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()
	// 建立连接
	testDB, err = pgx.Connect(ctx, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer func(testDB *pgx.Conn, ctx context.Context) {
		err := testDB.Close(ctx)
		if err != nil {

		}
	}(testDB, ctx)

	// 绑定测试
	testQueries = New(testDB)

	os.Exit(m.Run())
}
