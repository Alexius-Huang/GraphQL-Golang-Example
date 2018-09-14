package main

import (
	"github.com/jinzhu/gorm"
)

func Connection(path string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	db.DropTableIfExists(&User{}, &Post{})
	db.AutoMigrate(&User{}, &Post{})

	users := []User{
		{
			Name: "Max",
			Age:  18,
		},
		{
			Name: "Toby",
			Age:  32,
		},
		{
			Name: "Allen",
			Age:  24,
		},
	}

	for i := range users {
		if err := db.Create(&users[i]).Error; err != nil {
			panic(err)
		}
	}

	posts := []Post{
		{
			Title:    "How to pet a cat?",
			Content:  "Pet it nicely",
			AuthorID: 1,
		},
		{
			Title:    "How to poke a cat?",
			Content:  "Poke it wisely",
			AuthorID: 2,
		},
		{
			Title:    "How to train a cat?",
			Content:  "Play with it",
			AuthorID: 1,
		},
		{
			Title:    "How to teach code to a cat?",
			Content:  "Don't let it sleep on the MacBook Pro",
			AuthorID: 3,
		},
	}

	for i := range posts {
		if err := db.Create(&posts[i]).Error; err != nil {
			panic(err)
		}
	}

	return db
}
