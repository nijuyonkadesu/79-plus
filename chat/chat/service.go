package chat

import (
	"chat/domain"
	"context"
)

func (s *UserStore) Create(ctx context.Context, user *domain.User) error {
	s.db.ExecContext(ctx, "INSERT INTO users (username, alias, bio) VALUES (?, ?, ?)",
		user.Username, user.Alias, user.Bio)

	return nil
}

func (s *UserStore) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	rows.Scan(&user.ID, &user.Username, &user.Alias, &user.Bio)
	rows.Close()
	return &user, err
}


/*
Plan for ChatRoom
type: dm, channel, group

*/

/* 
Plan for Message Service
fk to a chatroom!? really? nahh... need a better model

0. need DM / Channel Model first
1. Message table that functions like a crude queue -> that shd be part of separate-db
lemme think tmrw...

*/
