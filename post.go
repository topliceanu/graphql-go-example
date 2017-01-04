package main

type Post struct {
	ID     int
	UserID int
	Title  string
	Body   string
}

func InsertPost(post *Post) error {
	var id int
	err := db.QueryRow(`
		INSERT INTO posts(user_id, title, body)
		VALUES ($1, $2, $3)
		RETURNING id
	`, post.UserID, post.Title, post.Body).Scan(&id)
	if err != nil {
		return err
	}
	post.ID = id
	return nil
}

func RemovePostByID(id int) error {
	_, err := db.Exec("DELETE FROM posts WHERE id=$1", id)
	return err
}

func GetPostByID(id int) (*Post, error) {
	var (
		userID      int
		title, body string
	)
	err := db.QueryRow(`
		SELECT user_id, title, body
		FROM posts
		WHERE id=$1
	`, id).Scan(&userID, &title, &body)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

func GetPostByIDAndUser(id, userID int) (*Post, error) {
	var title, body string
	err := db.QueryRow(`
		SELECT title, body
		FROM posts
		WHERE id=$1
		AND user_id=$2
	`, id, userID).Scan(&title, &body)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:     id,
		UserID: userID,
		Title:  title,
		Body:   body,
	}, nil
}

func GetPostsForUser(id int) ([]*Post, error) {
	rows, err := db.Query(`
		SELECT p.id, p.title, p.body
		FROM posts AS p
		WHERE p.user_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		posts       = []*Post{}
		pid         int
		title, body string
	)
	for rows.Next() {
		if err = rows.Scan(&pid, &title, &body); err != nil {
			return nil, err
		}
		posts = append(posts, &Post{ID: id, UserID: id, Title: title, Body: body})
	}
	return posts, nil
}
