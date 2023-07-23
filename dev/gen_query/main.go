package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	gormdb, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Errorf("error opening db: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./pkg/query",
		ModelPkgPath: "./pkg/models",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `models.User` following conventions
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `models.User` and `models.Company`
	// g.ApplyInterface(func(MessagesQuerier) {}, models.Message{})

	// Generate the code
	g.Execute()
}
