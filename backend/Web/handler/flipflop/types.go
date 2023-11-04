package flipflop

type BlockRq struct {
	Name        string `json:"name"`
	BlockType   string `json:"blockType"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Rules       struct {
		RuleName          string `json:"ruleName"`
		Nested            bool   `json:"nested"`
		HasLikes          bool   `json:"hasLikes"`
		HasComments       bool   `json:"hasComments"`
		CommentsNested    bool   `json:"commentsNested"`
		CommentsHasLikes  bool   `json:"commentsHasLikes"`
		CommentsEditable  bool   `json:"commentsEditable"`
		CommentsMaxNested int    `json:"commentsMaxNested"`
	}
}

type BlockMetaRq struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Categories  []string `json:"categories"`
}

type BlockRulesRq struct {
	RuleName          string `json:"ruleName"`
	Nested            bool   `json:"nested"`
	HasLikes          bool   `json:"hasLikes"`
	HasComments       bool   `json:"hasComments"`
	CommentsNested    bool   `json:"commentsNested"`
	CommentsHasLikes  bool   `json:"commentsHasLikes"`
	CommentsEditable  bool   `json:"commentsEditable"`
	CommentsMaxNested int    `json:"commentsMaxNested"`
}

type LangRq struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	PreviousLang *string `json:"previousLang"`
}

type RuleRq struct {
	RuleName          string `json:"ruleName"`
	Descr             string `json:"description"`
	Nested            bool   `json:"nested"`
	HasLikes          bool   `json:"hasLikes"`
	HasComments       bool   `json:"hasComments"`
	CommentsNested    bool   `json:"commentsNested"`
	CommentsHasLikes  bool   `json:"commentsHasLikes"`
	CommentsEditable  bool   `json:"commentsEditable"`
	CommentsMaxNested int    `json:"commentsMaxNested"`
}

type TaxonomyRq struct {
	Name  string `json:"name"`
	Descr string `json:"descr"`
}
