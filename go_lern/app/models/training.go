package models

import (
	"log"
	"time"
)

//トレーニングのフィールドを定義した構造体
type Training struct {
	ID        int
	Tday      string
	Hukkin    int
	Udetate   int
	Haikin    int
	Working   string
	CreatedAt time.Time
}

//トレーニングのレコードを作成する関数。
func (t *Training) CreteTraining() (err error) {
	cmd := `insert into training(
		tday,
		hukkin,
		udetate,
		haikin,
		working,
		created_at) values (?, ?, ?, ?, ?, ?)`
	_, err = Db.Exec(cmd,
		t.Tday,
		t.Hukkin,
		t.Udetate,
		t.Haikin,
		t.Working,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//trainingテーブルからTable表示のために値を取得する関数
func (t *Training) SelectAllTraining() (err error) {
	cmd := `select 
	id, 
	tday, 
	hukkin, 
	udetate, 
	haikin, 
	working from training`
	_, err = Db.Exec(cmd)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
