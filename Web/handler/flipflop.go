package handler

import (

	"github.com/labstack/echo/v4"
	pb "github.com/shorty-io/go-shorty/flipFlop/proto"
)

func GetBlock(c echo.Context, client pb.FlipFlopClient) error {

	req = pb.CreateRequest{
		Name: "Testing",
		Type: "Post",
		Rules:  &pb.BlockRules{
			Rules: *pb.BlockRules_Rules{
				RuleName:          "",
				Nested:            false,
				HasLikes:          false,
				HasComments:       false,
				CommentsNested:    false,
				CommentsHasLikes:  false,
				CommentsEditable:  false,
				CommentsMaxNested: 0,
			},
		},
		Author: "Default",
	}

	res, err := client.CreateBlock(context.Background(), req)
}
