package config

import (
	"ALBUMS/album_data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDB() (*gorm.DB, error) {
	var err error
	stmt, err := Connect()
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.Open(stmt), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err = db.AutoMigrate(&album_data.Album{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SelectAll() []album_data.Album {
	db, err := OpenDB()
	if err != nil {
		return nil
	}
	var records []album_data.Album
	datas := db.Find(&records)
	if datas.Error != nil {
		panic(datas.Error)
		return nil
	}
	return records

}
