package main

type Resolver struct{}

func (r *Resolver) Users(args struct{ Limit int32 }) *[]*ResolvedUser {
	var users []User
	var err error

	if args.Limit != -1 {
		err = DB.Limit(args.Limit).Find(&users).Error
	} else {
		err = DB.Find(&users).Error
	}
	// err := DB.Find(&users).Error
	if err != nil {
		panic(err)
	}

	result := make([]*ResolvedUser, 0, len(users))
	for i := range users {
		result = append(result, &ResolvedUser{value: &users[i]})
	}
	return &result
}

func (r *Resolver) User(args struct{ ID int32 }) *ResolvedUser {
	var user User
	err := DB.First(&user, args.ID).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedUser{value: &user}
}

func (r *Resolver) CreateUser(args struct {
	Name string
	Age  int32
}) *ResolvedUser {
	user := User{Name: args.Name, Age: args.Age}
	err := DB.Create(&user).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedUser{value: &user}
}

func (r *Resolver) DestroyUser(args struct{ ID int32 }) *ResolvedUser {
	var user User
	err := DB.First(&user, args.ID).Error
	if err != nil {
		panic(err)
	}

	err = DB.Delete(&user).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedUser{value: &user}
}

func (r *Resolver) Posts() *[]*ResolvedPost {
	var posts []Post
	err := DB.Find(&posts).Error
	if err != nil {
		panic(err)
	}

	result := make([]*ResolvedPost, 0, len(posts))
	for i := range posts {
		result = append(result, &ResolvedPost{value: &posts[i]})
	}
	return &result
}

func (r *Resolver) Post(args struct{ ID int32 }) *ResolvedPost {
	var post Post
	err := DB.First(&post, args.ID).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedPost{value: &post}
}

func (r *Resolver) CreatePost(args struct {
	Title    string
	Content  string
	AuthorID int32
}) *ResolvedPost {
	post := Post{Title: args.Title, Content: args.Content, AuthorID: args.AuthorID}
	err := DB.Create(&post).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedPost{value: &post}
}

func (r *Resolver) DestroyPost(args struct{ ID int32 }) *ResolvedPost {
	var post Post
	err := DB.First(&post, args.ID).Error
	if err != nil {
		panic(err)
	}

	err = DB.Delete(&post).Error
	if err != nil {
		panic(err)
	}

	return &ResolvedPost{value: &post}
}
