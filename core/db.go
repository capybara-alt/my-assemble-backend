package core

import (
	"fmt"

	"github.com/capybara-alt/my-assemble/config"
)

var (
	DB_DSN = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=require TimeZone=Asia/Tokyo",
		config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_HOST, config.DB_PORT)
)
