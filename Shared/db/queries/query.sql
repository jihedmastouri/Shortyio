--   ____                  _
--  / __ \__  _____  _____(_)__  _____
-- / / / / / / / _ \/ ___/ / _ \/ ___/
--/ /_/ / /_/ /  __/ /  / /  __(__  )
--\___\_\__,_/\___/_/  /_/\___/____/
--------------------------------------------
-- This File Contains all Queries on the Main Database.
-- Refer to Sqlc for more information https://docs.sqlc.dev/en/stable/
--
-- The File Includes 6 Section:
--
-- 1- `Selections` following: Get? / Get?By?
-- 2- `Adding` Inserts following Create?
-- 3- `Joins` insert for ManyToMany, following Add?To?
-- 4- `Updates` following: Update?
-- 5- `Deletions` following: Delete?
-- 6- `Counts` following: Count?
--
-- Please use PascalCase for naming.

------------------
-- 1- Selections
------------------

-- name: GetBlock :many
SELECT id, author, created_at, name
FROM blocks
LIMIT $1
OFFSET $2;

-- name: GetBlockByID :one
SELECT id, author, created_at, name
FROM blocks b
WHERE b.id = $1;

-- name: GetTypeByName :one
SELECT id
FROM block_types bt
WHERE bt.name = $1;

-- name: GetBlockImages :many
select bi.title, bi.alt, bi.file
from block_images bi
inner join block_langs bl
on block_lang.id = bi.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockText :many
select bt.content, bt.hint, bt.name
from block_texts bt
inner join block_langs bl
on bl.id = bt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockRichText :many
select brt.name, brt.hint, brt.content
from block_rich_texts brt
inner join block_langs bl
on bl.id = brt.block_lang_id
where bl.lang_code = $1 and  bl.block_id = $2;

-- name: GetBlockCategories :many
SELECT name, descr
  FROM categories c
  INNER JOIN block_categ bc
  ON c.id = bc.categ_id
  INNER JOIN blocks b
  ON bc.block_id = b.id;

-- name: GetBlockTags :many
SELECT name, descr
  FROM tags t
  INNER JOIN block_tags bt
  ON t.id = bt.tag_id
  INNER JOIN blocks b
  ON bt.block_id = b.id;

------------------
-- 2- Adding
------------------

-- name: AddBlock :exec
INSERT INTO blocks (author, nested, has_likes, has_comments, comments_max_nest,
                           comments_has_likes, comment_editable, rules_name, type)
VALUES ($1, $2, $3, $4, $5, $6 $7, $8, $9);

-- name: AddTag :exec
INSERT INTO Tags (
    name, descr
) VALUES ($1,$2);

-- name: AddCateg :exec
INSERT INTO Tags (
    name, descr
) VALUES ($1,$2);

-- name: AddBlockRules :exec
INSERT INTO comment_types(
    name, nested, has_likes, editable, max_nest
) VALUES ($1,$2,$3,$4,$5) RETURNING id;

------------------
-- 3- Joins
------------------

-- name: JoinTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUES ($1,$2);

-- name: JoinCategToBlock :exec
INSERT INTO block_categ(
    block_id, categ_id
) VALUES ($1,$2);


-- name: JoinChildToBlock :exec
INSERT INTO block_nested(
    child, parent
) VALUES ($1,$2);

------------------
-- 5- Updates
------------------

-- name: UpdateBlockRules :exec
Update blocks
    SET rules_name = $2
        nested = $3
        has_likes = $4
        has_comments = $5
        comments_max_nest = $6
        comments_has_likes = $7
        comment_editable = $8
WHERE name = $1


------------------
-- 5- Deletions
------------------

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE id = $1;

-- name: DeleteBlockLang :exec
DELETE FROM block_langs
WHERE block_id = $1
    AND lang_name = $2;

-- name: DeleteBlockText :exec
DELETE FROM block_texts bt
INNER JOIN block_langs bl
ON bt.block_lang_id = bl.id
WHERE bl.block_id = $1
    AND bl.lang_name = $2

-- name: DeleteBlockText :exec
DELETE FROM block_rich_texts brt
INNER JOIN block_langs bl
ON brt.block_lang_id = bl.id
WHERE bl.block_id = $1
    AND bl.lang_name = $2

-- name: DeleteBlockText :exec
DELETE FROM block_images bi
INNER JOIN block_langs bl
ON bi.block_lang_id = bl.id
WHERE bl.block_id = $1
    AND bl.lang_name = $2

-- name: DeleteBlockCategs :exec
DELETE FROM block_categ bc
INNER JOIN blocks b
ON bc.block_id = b.id
WHERE b.id = $1;

-- name: DeleteBlockTags :exec
DELETE  FROM block_tags bt
INNER JOIN blocks b
ON bt.block_id = b.id
WHERE b.id = $1;

-- name: DeleteCategByID :many
DELETE FROM categories c
WHERE c.id = $1;

-- name: DeleteTagByID :many
DELETE FROM tags t
WHERE t.id = $1;
