--    _____      __
--   / ___/_____/ /_  ___  ____ ___  ____ _
--   \__ \/ ___/ __ \/ _ \/ __ `__ \/ __ `/
--  ___/ / /__/ / / /  __/ / / / / / /_/ /
-- /____/\___/_/ /_/\___/_/ /_/ /_/\__,_/
--------------------------------------------

CREATE SCHEMA IF NOT EXISTS blocks;

SET SCHEMA 'blocks';

CREATE TABLE IF NOT EXISTS block_types
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(20) UNIQUE NOT NULL,
    descr VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS block_rules
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(20) NOT NULL,
    descr              VARCHAR(200) NOT NULL,
    nested             BOOLEAN  DEFAULT FALSE,
    has_comments       BOOLEAN  DEFAULT FALSE,
    has_likes          BOOLEAN  DEFAULT FALSE,
    comments_max_nest  SMALLINT DEFAULT 3,
    comments_has_likes BOOLEAN  DEFAULT TRUE,
    comment_editable   BOOLEAN  DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS blocks
(
    id                 UUID        DEFAULT gen_random_uuid() PRIMARY KEY,
    name               VARCHAR(20) NOT NULL,
    created_at         TIMESTAMP   DEFAULT NOW(),
    author             UUID        NOT NULL,

    -- General Rules:
    rules_name         VARCHAR(20) DEFAULT 'Custom',
    nested             BOOLEAN     NOT NULL,
    has_likes          BOOLEAN     NOT NULL,

    -- Comments Rules:
    has_comments       BOOLEAN     NOT NULL,
    comments_max_nest  SMALLINT    NOT NULL,
    comments_has_likes BOOLEAN     NOT NULL,
    comment_editable   BOOLEAN     NOT NULL,

    type               INTEGER     NOT NULL
        REFERENCES block_types (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block_langs
(
    id             SERIAL PRIMARY KEY,
    lang_name      VARCHAR(20) NOT NULL,
    lang_code      VARCHAR(10) NOT NULL,
    version_number INTEGER   DEFAULT 1,

    -- Timestamps
    created_at     TIMESTAMP DEFAULT NOW(),
    updated_at     TIMESTAMP DEFAULT NOW(),

    block_id       UUID        NOT NULL
        REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_block_langs_block_id_lang_name ON block_langs (block_id, lang_name);

CREATE TABLE IF NOT EXISTS block_images
(
    id            SERIAL PRIMARY KEY,
    file          VARCHAR(100) NOT NULL,
    alt           VARCHAR(100),
    title         VARCHAR(50)  NOT NULL,

    block_lang_id INTEGER      NOT NULL REFERENCES block_langs (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block_texts
(
    id            SERIAL PRIMARY KEY,
    content       TEXT        NOT NULL,
    name          VARCHAR(50) NOT NULL,
    hint          VARCHAR(200),

    block_lang_id INTEGER     NOT NULL
        REFERENCES block_langs (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block_rich_texts
(
    id            SERIAL PRIMARY KEY,
    content       TEXT        NOT NULL,
    name          VARCHAR(50) NOT NULL,
    hint          VARCHAR(200),

    block_lang_id INTEGER     NOT NULL
        REFERENCES block_langs (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block_nested
(
    parent UUID REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    child  UUID REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS tags
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(20) NOT NULL,
    descr VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS block_tags
(
    block_id UUID REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    tag_id   INTEGER REFERENCES tags (id) ON DELETE CASCADE,
    CONSTRAINT block_tags_pkey PRIMARY KEY (block_id, tag_id)
);

CREATE TABLE IF NOT EXISTS categories
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(20) NOT NULL,
    descr VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS block_categs
(
    block_id UUID REFERENCES blocks (id) ON DELETE CASCADE ON UPDATE CASCADE,
    categ_id INTEGER REFERENCES categories (id) ON DELETE CASCADE,
    CONSTRAINT block_categs_pkey PRIMARY KEY (block_id, categ_id)
);

---------------------------
-- Triggers
---------------------------

-- Increment block_langs version number on update
CREATE OR REPLACE FUNCTION increment_version_block_lang()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.version_number := OLD.version_number + 1;
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_increment_version_block_lang
    AFTER UPDATE
    ON block_langs
    FOR EACH ROW
EXECUTE FUNCTION increment_version_block_lang();


-- Update 'updated_at' timestamp on every change
CREATE OR REPLACE FUNCTION update_block_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at := NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_block_updated_at
    BEFORE UPDATE
    ON block_langs
    FOR EACH ROW
EXECUTE FUNCTION update_block_updated_at();

---------------------------
-- Insertions
---------------------------

INSERT INTO block_types (name, descr)
VALUES ('Post', 'Main block used for building blogs');

INSERT INTO block_rules (name, nested, has_comments, has_likes, comments_max_nest, comments_has_likes, comment_editable)
VALUES ('default', false, false, false, 0, false, false),
       ('Interactive', false, true, true, 3, true, true);

