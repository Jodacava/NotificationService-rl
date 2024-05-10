package postgres

import "github.com/jinzhu/gorm"

type DbRepository struct {
	DbClient *gorm.DB
}

type DbRepositoryBase interface {
	Save(data interface{}) error
	GetByConditions(entities interface{}, condition string, parameters []interface{})
}

func NewDbPRepository(DbClient *gorm.DB) DbRepositoryBase {
	return DbRepository{DbClient: DbClient}
}

func (r DbRepository) Save(data interface{}) error {
	return r.DbClient.Save(data).Error
}

func (r DbRepository) GetByConditions(entities interface{}, condition string, parameters []interface{}) {
	if r.DbClient != nil {
		r.DbClient.Where(condition, parameters...).Find(entities)
	}
}
