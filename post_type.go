package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if post, ok := p.Source.(*Post); ok == true {
					return post.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	PostType.AddFieldConfig("user", &graphql.Field{
		Type: graphql.NewNonNull(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				return GetUserByID(post.UserID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comment", &graphql.Field{
		Type: CommentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetCommentByIDAndPost(id, post.ID)
			}
			return nil, nil
		},
	})
	PostType.AddFieldConfig("comments", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CommentType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if post, ok := p.Source.(*Post); ok == true {
				return GetCommentsForPost(post.ID)
			}
			return []Comment{}, nil
		},
	})
}
