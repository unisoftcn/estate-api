package entity

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/vuuvv/errors"
	"github.com/vuuvv/orca/config"
	"github.com/vuuvv/orca/id"
	"github.com/vuuvv/orca/orm"
	"gorm.io/gorm"
	"testing"
)

func getDb(t *testing.T) *gorm.DB {
	loader := config.MustLoad("../../estate/resources/application.yaml")
	ormConfig := &orm.Config{}
	err := loader.Unmarshal(ormConfig, "database")
	_, _ = id.NewGenerator()
	if err != nil {
		t.Error(err)
	}
	db, err := orm.New(ormConfig)
	if err != nil {
		t.Error(err)
	}
	return db
}

func TestMigration(t *testing.T) {
	db := getDb(t)

	err := db.AutoMigrate(
		&orm.Sequence{},
		&Building{},
		&Floor{},
		&Room{},
	)
	if err != nil {
		t.Error(err)
	}
}

func JsonParseG[T any](json string) (*T, error) {
	var ret T
	err := jsoniter.Unmarshal([]byte(json), &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &ret, nil
}