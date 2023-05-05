package store

import "context"

func (s *PgStore) CreateAtReceipt(ctx context.Context, statusCode uint, messageId string) error {
	_, err := s.db.Exec(ctx, s.queries.CreateAtReceipt, statusCode, messageId)
	if err != nil {
		return err
	}

	return nil
}

func (s *PgStore) CreateTgReceipt(ctx context.Context, messageId int) error {
	_, err := s.db.Exec(ctx, s.queries.CreateTgReceipt, messageId)
	if err != nil {
		return err
	}

	return nil
}

func (s *PgStore) SetAtDelivered(ctx context.Context, messageId string) error {
	_, err := s.db.Exec(ctx, s.queries.SetAtDelivered, messageId)
	if err != nil {
		return err
	}

	return nil
}
