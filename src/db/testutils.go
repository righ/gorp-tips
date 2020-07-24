package db

import (
	"context"
	"database/sql"
	"testing"

	"gorp-tips/models"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// TestingBlock テスト対象のブロック(関数)
type TestingBlock func(ctx context.Context, tx *NestableTx)

// Dependency 依存レコード(自身を含む)
type Dependency interface{}

func initDb(t *testing.T) *gorp.DbMap {
	db, err := sql.Open("mysql", "usr:pw@tcp(testing_mysql:3306)/db")
	if err != nil {
		t.Fatalf("Failed to connect db. %s", err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	models.MapStructsToTables(dbmap)
	return dbmap
}

// RunTest テストを実行する
func RunTest(ctx context.Context, t *testing.T, block TestingBlock, deps ...Dependency) {
	dbmap := initDb(t)
	defer dbmap.Db.Close()
	// トランザクション作成
	tx, err := dbmap.Begin()
	if err != nil {
		t.Errorf("Failed to start transaction. %w", err)
		return
	}
	ntx := &NestableTx{Transaction: tx}
	defer ntx.Rollback()

	// dependencies 投入
	for _, m := range deps {
		if err := ntx.Insert(m); err != nil {
			t.Fatalf("Failed to load dependencies. %s, %+v", err, m)
		}
	}
	// テスト実行
	block(ctx, ntx)
}
