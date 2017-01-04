package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.ID, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Email, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	UserType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetPostByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("posts", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PostType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetPostsForUser(user.ID)
			}
			return []Post{}, nil
		},
	})
	UserType.AddFieldConfig("follower", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Follower ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetFollowerByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followers", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetFollowersForUser(user.ID)
			}
			return []User{}, nil
		},
	})
	UserType.AddFieldConfig("followee", &graphql.Field{
		Type: UserType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Followee ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetFolloweeByIDAndUser(id, user.ID)
			}
			return nil, nil
		},
	})
	UserType.AddFieldConfig("followees", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(UserType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if user, ok := p.Source.(*User); ok == true {
				return GetFolloweesForUser(user.ID)
			}
			return []User{}, nil
		},
	})
}
