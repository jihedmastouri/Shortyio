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

-- name: GetBlockRulesByName :one
SELECT nested, has_comments, has_likes,
    comments_max_nest, comments_has_likes, comment_editable
FROM block_rules
WHERE name = $1;

-- name: GetAllBlockRules :many
SELECT name, descr
FROM block_rules
LIMIT 100;

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
SELECT c.name
FROM categories c
INNER JOIN block_categs bc
ON c.id = bc.categ_id
WHERE bc.block_id = $1
LIMIT 100;

-- name: GetAllCategories :many
SELECT name, descr
FROM categories
LIMIT 100;

-- name: GetBlockTags :many
SELECT t.name
FROM tags t
INNER JOIN block_tags bt
ON t.id = bt.tag_id
WHERE bt.block_id = $1
LIMIT 100;

-- name: GetAllTags :many
SELECT name, descr
FROM tags
LIMIT 100;

------------------
-- 2- Adding
------------------

-- name: AddBlock :one
INSERT INTO blocks (author, name, nested, has_likes, has_comments, comments_max_nest,
                           comments_has_likes, comment_editable, rules_name, type)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;

-- name: AddTag :one
INSERT INTO tags (
    name, descr
) VALUES ($1,$2) RETURNING name;

-- name: AddCateg :one
INSERT INTO categories (
    name, descr
) VALUES ($1,$2) RETURNING name;

-- name: AddBlockRules :one
INSERT INTO block_rules (
    name, nested, has_likes, has_comments, comments_max_nest,
    comments_has_likes, comment_editable)
VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING name;

------------------
-- 3- Joins
------------------

-- name: JoinTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUES ($1,$2);

-- name: JoinCategToBlock :exec
INSERT INTO block_categs(
    block_id, categ_id
) VALUES ($1,$2);


-- name: JoinChildToBlock :exec
INSERT INTO block_nested (
    child, parent
) VALUES ($1,$2);

------------------
-- 5- Updates
------------------

-- name: UpdateBlockRules :exec
Update blocks
    SET rules_name = $2,
        nested = $3,
        has_likes = $4,
        has_comments = $5,
        comments_max_nest = $6,
        comments_has_likes = $7,
        comment_editable = $8
WHERE name = $1;


------------------
-- 6- Deletions
------------------

-- name: DeleteBlock :exec
DELETE FROM blocks WHERE id = $1;

-- name: DeleteBlockLang :exec
DELETE FROM block_langs
WHERE block_id = $1 AND lang_name = $2;

-- name: DeleteBlockText :exec
DELETE FROM block_texts
WHERE block_lang_id = (
        SELECT id
        FROM  block_langs
        WHERE block_id = $1 AND lang_name = $2
    );

-- name: DeleteBlockRichText :exec
DELETE FROM block_rich_texts
WHERE block_lang_id = (
        SELECT id
        FROM  block_langs
        WHERE block_id = $1 AND lang_name = $2
    );

-- name: DeleteBlockImages :exec
DELETE FROM block_images
WHERE block_lang_id = (
        SELECT id
        FROM  block_langs
        WHERE block_id = $1 AND lang_name = $2
    );

-- name: DeleteBlockCateg :exec
DELETE FROM block_categs
WHERE block_id = $1 AND
    categ_id = (
        SELECT id
        FROM categories
        WHERE name = $2
    );

-- name: DeleteAllBlockCategs :exec
DELETE FROM block_categs
WHERE block_id = $1;

-- name: DeleteBlockTag :exec
DELETE FROM block_tags
WHERE block_id = $1 AND
    tag_id = (
        SELECT id
        FROM tags
        WHERE name = $2
    );

-- name: DeleteAllBlockTags :exec
DELETE FROM block_tags
WHERE block_id = $1;

-- name: DeleteCategByID :exec
DELETE FROM categories
WHERE id = $1;

-- name: DeleteTagByID :exec
DELETE FROM tags
WHERE id = $1;
