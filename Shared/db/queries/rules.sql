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

-- name: GetBlockRulesByName :one
SELECT nested, has_comments, has_likes,
    comments_max_nest, comments_has_likes, comment_editable
FROM block_rules
WHERE name = $1;

-- name: GetAllBlockRules :many
SELECT name, descr
FROM block_rules
LIMIT 100;

------------------
-- 2- Adding
------------------

-- name: CreateBlockRule :one
INSERT INTO block_rules (name, nested, descr, has_likes, has_comments, comments_max_nest,
        comments_has_likes, comment_editable)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING name;

------------------
-- 3- Joins
------------------

------------------
-- 5- Updates
------------------

-- name: UpdateBlockRules :exec
Update block_rules
    SET nested = $2,
        has_likes = $3,
        has_comments = $4,
        comments_max_nest = $5,
        comments_has_likes = $6,
        comment_editable = $7,
        descr = $8
WHERE name = $1;

------------------
-- 6- Deletions
------------------

-- name: DeleteBlockRuleById :exec
DELETE FROM block_rules WHERE id = $1;

-- name: DeleteBlockRule :exec
DELETE FROM block_rules WHERE name = $1;
