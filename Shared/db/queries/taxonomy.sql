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

-- name: AddTag :one
INSERT INTO tags (
    name, descr
) VALUES ($1,$2) RETURNING name;

-- name: AddCateg :one
INSERT INTO categories (
    name, descr
) VALUES ($1,$2) RETURNING name;

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

------------------
-- 3- Joins
------------------

------------------
-- 5- Updates
------------------

------------------
-- 6- Deletions
------------------

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
