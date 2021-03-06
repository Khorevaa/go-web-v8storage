package services

import (
	"../datasource"
	//"../models"
	models "../datamodels"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type StorageService interface {
	GetAll() []models.Storage
	GetByID(id int64) models.Storage
	DeleteByID() []models.Storage
	UpdateByID() []models.Storage
	GetList(page int, pageSize int) (r []models.Storage, c int)
}

// NewMovieService returns the default movie service.
func NewStorageService() *storageService {
	return &storageService{
		db:        datasource.DB,
		tableName: datasource.TableName("Storage"),
	}
}

type storageService struct {
	db        *pg.DB
	tableName string
}

func (s *storageService) GetAll() (r []models.Storage) {

	s.db.Model(&r).Order("ID").Select(r)
	return
}

func (s *storageService) GetList(page int, pageSize int) (storages []models.Storage, count int) {

	offset := (page - 1) * pageSize

	//f := []datamodels.Storage{}

	//rows, err := s.db.Raw("SELECT * FROM storages").Rows()
	//
	//if err != nil  {
	//	panic(err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	////rows.Scan(&name, &age, &email)
	//}

	//s.db.Last(&f)

	count, err := orm.NewQuery(s.db, &storages).Order("id").Limit(pageSize).Offset(offset).Column("storage.*", "Crserver", "Project").SelectAndCount(&storages)

	//count, err:= s.db.Model(&storages).Order("id").Limit(pageSize).Offset(offset).Column("storage.*","Tags", "Crserver", "Crserver.Tags", "Project", "Project.Tags").SelectAndCount(&storages)
	//println(storages[2].Crserver.Tags[0].Text)
	if err != nil {
		panic(err)
	}

	return
}

func (s *storageService) GetByID(id int64) (storage models.Storage) {

	err:= s.db.Model(&storage).Order("id").Column("storage.*").Where("id =?", id).Select(&storage)

	if err != nil {
		panic(err)
	}
	return
}

func (s *storageService) Update(storage models.Storage) (err error) {

	_, err= s.db.Model(&storage).Update(&storage)

	return
}

func (s *storageService) DeleteByID() (r []models.Storage) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}

func (s *storageService) UpdateByID() (r []models.Storage) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}
