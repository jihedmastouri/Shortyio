-- name: GetBlock :many
select id, author, created_at, updated_at from blocks.Blocks;

-- name: GetBlockByID :one
select b.version_number,
from blocks b
where b.id = $1;

-- name: GetBlockImages :many
select bi.title, bi.alt, bi.file
from block_image bi
inner join block_lang bl
on block_lang.id = bi.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockText :many
select bt.content, bt.hint, bt.name
from block_text bt
inner join block_lang bl
on block_lang.id = bt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockRichText :many
select brt.name, brt.hint, brt.content
from block_rich_text brt
inner join block_lang bl
on block_lang.id = brt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockCategories :many
SELECT name, descr
  FROM categories
  INNER JOIN block_categ
  ON categories.id = block_categ.categ_id
  INNER JOIN blocks
  ON block_categ.block_id = block.id;

-- name: GetBlockTags :many
SELECT name, descr
  FROM categories
  INNER JOIN block_tags
  ON categories.id = block_categ.tag_id
  INNER JOIN blocks
  ON block_categ.block_id = block.id;

-- name: CreateBlock :exec
INSERT INTO block (
    has_likes, has_comments, type, comment_type
) VALUE ($1,$2,$3,$4);

-- name: AddTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUE ($1,$2);

-- name: AddCategToBlock :exec
INSERT INTO block_categ(
    block_id, categ_id
) VALUE ($1,$2);

-- name: AddCommentRules :exec
INSERT INTO comment_type(
    nested, has_likes, editable, max_nested
) VALUE ($1,$2,$3,$4) RETURNING id;

-- name: DeleteBlock :exec
DELETE FROM blocks
    WHERE id = $1
    RETURNING id;

-- name: DeleteBlockLang :exec
DELETE FROM block_lang
    WHERE block_id = $1
    AND lang_name = $2
    RETURNING id;
