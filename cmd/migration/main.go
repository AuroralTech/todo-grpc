package main

import (
	"fmt"
	"os"

	"github.com/AuroralTech/todo-grpc/config"
	"github.com/AuroralTech/todo-grpc/pkg/domain/model"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "down" {
		err := down()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if err := migrate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// cfg, err = config.LoadTestConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err := migrate(cfg); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
}

func migrate() error {
	db, err := config.NewSQLConnection()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Todo{}, &model.User{}, &model.Profile{}, &model.Follow{}, &model.Post{}, &model.Favorite{})
	if err != nil {
		return err
	}

	todo := model.Todo{Task: "test", IsCompleted: false}

	err = db.Create(&todo).Error
	if err != nil {
		return err
	}

	return nil
}

func down() error {
	db, err := config.NewSQLConnection()
	if err != nil {
		return err
	}

	err = db.Migrator().DropTable(&model.Todo{})
	if err != nil {
		return err
	}

	return nil
}
