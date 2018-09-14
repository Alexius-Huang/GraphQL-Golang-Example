package main

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	*gorm.Model
	Name  string
	Age   int32
	Posts []Post `gorm:"foreignkey:AuthorID"`
}

type ResolvedUser struct {
	value *User
}

func (r *ResolvedUser) ID() graphql.ID {
	id := strconv.Itoa(int(r.value.ID))
	return graphql.ID(id)
}
func (r *ResolvedUser) Name() string {
	return r.value.Name
}
func (r *ResolvedUser) Age() int32 {
	return r.value.Age
}
func (r *ResolvedUser) Posts() []*ResolvedPost {
	user := r.value
	var posts []Post
	err := DB.Model(&user).Related(&posts, "posts").Error
	if err != nil {
		panic(err)
	}

	results := make([]*ResolvedPost, 0, len(posts))
	for i := range posts {
		results = append(results, &ResolvedPost{value: &posts[i]})
	}

	return results
}
