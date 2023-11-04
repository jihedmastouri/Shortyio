package handler

import (
	"github.com/shorty-io/go-shorty/queries/db"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

func getRules(b *db.Block) *pb.BlockRules {
	return &pb.BlockRules{
		RuleName:          b.Rules.RuleName,
		Nested:            b.Rules.Nested,
		HasLikes:          b.Rules.HasLikes,
		HasComments:       b.Rules.HasComments,
		CommentsHasLikes:  b.Rules.CommentsHasLike,
		CommentsEditable:  b.Rules.CommentsEditable,
		CommentsMaxNested: b.Rules.CommentsMaxNested,
	}
}
