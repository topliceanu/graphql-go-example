package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.ID, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.Title, nil
				}
				return nil, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if comment, ok := p.Source.(*Comment); ok == true {
					return comment.Body, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CommentType.AddFieldConfig("user", &graphql.Field{
		Type: UserType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if comment, ok := p.Source.(*Comment); ok == true {
				return GetUserByID(comment.UserID)
			}
			return nil, nil
		},
	})
	CommentType.AddFieldConfig("post", &graphql.Field{
		Type: PostType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Post ID",
				Type:        graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			i := p.Args["id"].(string)
			id, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			return GetPostByID(id)
		},
	})
}
