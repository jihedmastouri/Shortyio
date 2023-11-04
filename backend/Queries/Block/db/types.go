package db

type Block struct {
	BlockID     string        `bson:"block_id"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Type        string        `bson:"type"`
	Tags        []string      `bson:"tags"`
	Categories  []string      `bson:"categories"`
	LangCode    string        `bson:"lang_code"`
	LangName    string        `bson:"lang_name"`
	Version     int32         `bson:"version"`
	CreatedAt   string        `bson:"created_at"`
	UpdatedAt   string        `bson:"updated_at"`
	Authors     []Author      `bson:"authors"`
	Content     []ElementType `bson:"content"`
	Rules       Rules         `bson:"rules"`
}

type BlockMeta struct {
	BlockID     string   `bson:"block_id"`
	Name        string   `bson:"name"`
	Description string   `bson:"description"`
	Type        string   `bson:"block_type"`
	Tags        []string `bson:"tags"`
	Categories  []string `bson:"categories"`
	// Authors     []Author `bson:"authors"`
}

type BlockContent struct {
	BlockID    string        `bson:"block_id"`
	Type       string        `bson:"type"`
	Tags       []string      `bson:"tags"`
	Categories []string      `bson:"categories"`
	LangCode   string        `bson:"lang_code"`
	LangName   string        `bson:"lang_name"`
	Version    int32         `bson:"version"`
	CreatedAt  string        `bson:"created_at"`
	UpdatedAt  string        `bson:"updated_at"`
	Content    []ElementType `bson:"content"`
}

type Rules struct {
	RuleName          string `bson:"rule_name"`
	Descr             string `bson:"descr"`
	Nested            bool   `bson:"nested"`
	HasLikes          bool   `bson:"has_likes"`
	HasComments       bool   `bson:"has_comments"`
	CommentsHasLike   bool   `bson:"comments_has_likes"`
	CommentsEditable  bool   `bson:"comments_editable"`
	CommentsMaxNested int32  `bson:"comments_max_nested"`
}

type ElementType struct {
	Element struct {
		Media struct {
			Title string `bson:"title"`
			// Type  pb.MediaType `bson:"type"`
			File string `bson:"file"`
			Alt  string `bson:"alt"`
		} `bson:"media"`
		Text struct {
			Name string `bson:"name"`
			// Type    pb.TextualType `bson:"type"`
			Content string `bson:"content"`
			Hint    string `bson:"hint"`
		} `bson:"text"`
	}
}

type Author struct {
	ID    string `bson:"id"`
	Name  string `bson:"name"`
	Image string `bson:"image"`
}
