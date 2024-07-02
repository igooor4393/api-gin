package migrate

import (
	"api-gin/initializers"
	"api-gin/models"
	"fmt"
	"log"
)

// ----------------- Автомиграцию можно сделать как main пакет, для миграции баз руками -----------------
//func init() {
//	config, err := initializers.LoadConfig(".")
//	if err != nil {
//		log.Fatal("Could not load environment variables", err)
//	}
//
//	initializers.ConnectDB(&config)
//}

//func main() {
//	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
//	initializers.DB.AutoMigrate(&models.Post{})
//	fmt.Println("👍 Migration complete")
//}

func Migrate() {

	err := initializers.DB.AutoMigrate(&models.Program{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"") // создание расширения позволяющего генерировать в самой бд вместо primary key UUID v4

	fmt.Println("👍 Migration complete")
}
