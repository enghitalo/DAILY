package databases

import (
	"daily_project_module/src/modules/user/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateDBConnection() error {
	// Close the existing connection if open
	if dbConn != nil {
		CloseDBConnection(dbConn)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true, /// AutoMigrate creates database foreign key constraints automatically, you can disable this feature during initialization, for example:
		},
	)

	/// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes. It will change existing column’s type if its size, precision, nullable changed. It WON’T delete unused columns to protect your data.
	// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&models.User{})
	db.AutoMigrate(
		&models.User{},
		// &models.Product{},
		// &models.Order{},
	)

	if err != nil {
		fmt.Println("Error occurred while connecting with the database %w", err)
	}
	
	dbConn = db
	return err

}

func GetDatabaseConnection() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

func CloseDBConnection(conn *gorm.DB) {
	sqlDB, err := conn.DB()
	if err != nil {
		fmt.Println("Error occurred while closing a DB connection")
	}
	defer sqlDB.Close()
}
