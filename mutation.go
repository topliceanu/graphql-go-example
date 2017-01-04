package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Description: "New User Email",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				user := &User{
					Email: email,
				}
				err := InsertUser(user)
				return user, err
			},
		},
		"removeUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveUserByID(id)
				return (err == nil), err
			},
		},
		"follow": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"follower": &graphql.ArgumentConfig{
					Description: "ID of follower user",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"followee": &graphql.ArgumentConfig{
					Description: "ID of followee user",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["follower"].(string)
				followerID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				j := p.Args["followee"].(string)
				followeeID, err := strconv.Atoi(j)
				if err != nil {
					return nil, err
				}
				err = Follow(followerID, followeeID)
				return (err == nil), err
			},
		},
		"unfollow": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"follower": &graphql.ArgumentConfig{
					Description: "ID of follower user",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"followee": &graphql.ArgumentConfig{
					Description: "ID of followee user",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["follower"].(string)
				followerID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				j := p.Args["followee"].(string)
				followeeID, err := strconv.Atoi(j)
				if err != nil {
					return nil, err
				}
				err = Unfollow(followerID, followeeID)
				return (err == nil), err
			},
		},
		"createPost": &graphql.Field{
			Type: PostType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Description: "Id of user creating the new post",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "New post title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "New post body",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["user"].(string)
				userID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				post := &Post{
					UserID: userID,
					Title:  title,
					Body:   body,
				}
				err = InsertPost(post)
				return post, err
			},
		},
		"removePost": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Post ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemovePostByID(id)
				return (err == nil), err
			},
		},
		"createComment": &graphql.Field{
			Type: CommentType,
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Description: "Id of user creating the new comment",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"post": &graphql.ArgumentConfig{
					Description: "Id of post to attach the comment to",
					Type:        graphql.NewNonNull(graphql.ID),
				},
				"title": &graphql.ArgumentConfig{
					Description: "New comment title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"body": &graphql.ArgumentConfig{
					Description: "New comment body",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["user"].(string)
				userID, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				j := p.Args["post"].(string)
				postID, err := strconv.Atoi(j)
				if err != nil {
					return nil, err
				}
				title := p.Args["title"].(string)
				body := p.Args["body"].(string)
				comment := &Comment{
					UserID: userID,
					PostID: postID,
					Title:  title,
					Body:   body,
				}
				err = InsertComment(comment)
				return comment, err
			},
		},
		"removeComment": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Comment ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveCommentByID(id)
				return (err == nil), err
			},
		},
	},
})
