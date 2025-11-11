package model

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

func InitDB(dsn string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    return db
}