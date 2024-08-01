package configs

import (
	"fmt"

	"github.com/anandaadityap/algoquill-backend/internal/app/models/entity"
	"github.com/anandaadityap/algoquill-backend/internal/pkg/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := ENV.DATABASE_URL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.PanicIfError(err)
	}
	db.AutoMigrate(&entity.User{})

	DB = db
	fmt.Println("Database connected")

}
