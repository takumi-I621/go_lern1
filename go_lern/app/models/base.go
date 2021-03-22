package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"go_lern/config"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

//テーブル作成のための変数宣言
var Db *sql.DB
var err error

//テーブル名の定義
const (
	tableNameUser     = "users"
	tableNameTodo     = "todos"
	tableNameSession  = "sessions"
	tableNameTraining = "training"
)

//テーブル作成
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//コマンドの作成 テーブル作成
	//%sはtableNameUserが入る。
	//id INTEGER・・・はテーブル定義の設定
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)
	//コマンドの実行
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
	Db.Exec(cmdS)

	//【オリジナル★】trainingテーブルを作成する
	cmdTR := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tday STRING,
		hukkin INTEGER,
		udetate INTEGER,
		haikin INTEGER,
		working , STRING,
		created_at DATETIME)`, tableNameTraining)
	Db.Exec(cmdTR)
}

//uuid作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

//password作成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
