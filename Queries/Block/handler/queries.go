package handler

import (
	"context"
	"log"

	"github.com/shorty-io/go-shorty/queries/db"

	pb "github.com/shorty-io/go-shorty/Shared/proto"
)

type Queries struct {
	pb.UnimplementedQueriesServer
}

func (q *Queries) GetBlock(ctx context.Context, rq *pb.BlockRequest) (*pb.Block, error) {
	block, error := db.GetBlock(ctx, rq)
	if error != nil {
		return nil, error
	}

	res := &pb.Block{
		BlockId:    block.BlockID,
		Name:       block.Name,
		Type:       block.Type,
		Lang:       block.LangCode,
		Version:    block.Version,
		Tags:       block.Tags,
		Categories: block.Categories,
		Content:    []*pb.ElementType{},
		// Children:   []*pb.BlockContent{},
		Rules:     getRules(block),
		UpdatedAt: block.UpdatedAt,
		CreatedAt: block.CreatedAt,
	}

	log.Println("content at 1:", block.Content[0])

	for _, author := range block.Authors {
		res.Authors = append(res.Authors, &pb.Author{
			Id:    author.ID,
			Name:  author.Name,
			Image: author.Image,
		})
	}

	for _, content := range block.Content {
		if content.Media.Title != "" {
			res.Content = append(res.Content, &pb.ElementType{
				Element: &pb.ElementType_Media{
					Media: &pb.Media{
						Title: content.Media.Title,
						// Type:  content.Media.Type,
						File: content.Media.File,
						Alt:  content.Media.Alt,
					},
				},
			})
		} else {
			res.Content = append(res.Content, &pb.ElementType{
				Element: &pb.ElementType_Text{
					Text: &pb.Textual{
						Name:    content.Text.Name,
						Content: content.Text.Content,
						// Type:    content.Text.Type,
						Hint: content.Text.Hint,
					},
				},
			})

		}
	}

	return res, nil
}

func (q *Queries) GetBlockMeta(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockMeta, error) {
	return db.GetBlockMeta(ctx, rq)
}

func (q *Queries) GetBlockRules(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockRules, error) {
	block, err := db.GetBlock(ctx, rq)
	if err != nil {
		return nil, err
	}
	return getRules(block), nil
}

// Fix this later
func (q *Queries) GetBlockContent(ctx context.Context, rq *pb.BlockRequest) (*pb.BlockContent, error) {
	_, err := db.GetBlock(ctx, rq)
	if err != nil {
		return nil, err
	}

	return &pb.BlockContent{}, nil
}

func (q *Queries) GetVersions(ctx context.Context, rq *pb.VersionsRequest) (*pb.VersionResponse, error) {
	return db.GetVersions(ctx, rq)
}

func (q *Queries) GetLanguages(ctx context.Context, rq *pb.LanguageRequest) (*pb.LanguageList, error) {
	return db.GetLanguages(ctx, rq)
}

func (q *Queries) Search(ctx context.Context, rq *pb.SearchRequest) (*pb.BlockList, error) {
	return db.Search(ctx, rq)
}
