package main

import (
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/jinzhu/gorm"
)

type Post struct {
	*gorm.Model
	Title    string
	Content  string
	AuthorID int32
}

type ResolvedPost struct {
	value *Post
}

func (r *ResolvedPost) ID() graphql.ID {
	id := strconv.Itoa(int(r.value.ID))
	return graphql.ID(id)
}
func (r *ResolvedPost) Title() string {
	return r.value.Title
}
func (r *ResolvedPost) Content() string {
	return r.value.Content
}

func (r *ResolvedPost) Author() *ResolvedUser {
	id := r.value.AuthorID
	var user User
	err := DB.First(&user, id).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedUser{value: &user}
}
