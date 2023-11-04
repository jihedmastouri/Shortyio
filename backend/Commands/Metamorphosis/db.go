package main

import (
	"context"

	db "github.com/shorty-io/go-shorty/Shared/db"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func updateBlock(id uuid.UUID, lang, content, changeLog string) error {
	tx, err := Conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := db.New(Conn)
	qtx := q.WithTx(tx)

	ctx := context.Background()
	params := db.DeleteBlockTextParams{
		BlockID:  id,
		LangCode: lang,
	}
	qtx.DeleteBlockText(ctx, params)

	return nil
}
