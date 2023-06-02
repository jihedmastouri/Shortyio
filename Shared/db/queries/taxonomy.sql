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

-- name: CreateTag :one
INSERT INTO tags (
    name, descr
) VALUES ($1,$2) RETURNING name;

-- name: CreateCateg :one
INSERT INTO categories (
    name, descr
) VALUES ($1,$2) RETURNING name;

-- name: AddTagToBlock :exec
INSERT INTO block_tags(
    block_id, tag_id
) VALUES ($1,$2);

-- name: AddCategToBlock :exec
INSERT INTO block_categs(
    block_id, categ_id
) VALUES ($1,$2);

-- name: UpdateTag :one
Update tags
    SET name = $1,
        descr = $2
WHERE name = $1
RETURNING id;

-- name: UpdateCategory :one
Update categories
    SET name = $1,
        descr = $2
WHERE name = $1
RETURNING id;

-- name: UpdateTagById :exec
Update tags
    SET name = $2,
        descr = $3
WHERE id = $1;

-- name: UpdateCategoryById :exec
Update categories
    SET name = $2,
        descr = $3
WHERE id = $1;

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

-- name: DeleteCateg :one
DELETE FROM categories
WHERE name = $1
RETURNING id;

-- name: DeleteTag :one
DELETE FROM tags
WHERE name = $1
RETURNING id;
