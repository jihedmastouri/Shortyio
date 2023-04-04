CREATE SCHEMA blocks;

set schema 'blocks';

create table if not exists block_type
(
    id           uuid    default gen_random_uuid() not null
        primary key,
    has_comments boolean default false,
    has_likes    boolean default false,
    name         varchar(20)
);

create table if not exists comment_type
(
    id        uuid     default gen_random_uuid() not null
        primary key,
    nested    boolean  default false,
    max_nest  smallint default 3,
    has_likes boolean  default true,
    editable  boolean  default false
);

create table if not exists block
(
    id  uuid default gen_random_uuid() not null primary key,

    block_type_id  uuid
        references block_type,

    comments_type  uuid
        references comment_type,

    has_comments   boolean,
    has_likes      boolean,
    created_at     timestamp default now(),
    updated_at     timestamp default now(),
    author         uuid                                not null
);

create table if not exists block_lang
(
    id  serial primary key,
    block_id  uuid  not null references block,
    lang_name varchar(20) not null,
    lang_code varchar(10) not null
);

create table if not exists block_image
(
    id  serial primary key,
    block_lang_id serial references block_lang,
    file          varchar(100) not null,
    alt           varchar(100),
    title         varchar(50)  not null
);

create table if not exists block_text
(
    id  serial primary key,
    block_lang_id serial references block_lang,
    content       text        not null,
    name          varchar(50) not null,
    hint          varchar(200)
);


create table if not exists block_rich_text
(
    id  serial primary key,
    block_lang_id serial references block_lang,
    content       text        not null,
    name          varchar(50) not null,
    hint          varchar(200)
);

create table if not exists block_nested
(
    parent uuid references block,
    child  uuid references block
);

create table if not exists comments
(
    id  serial primary key,

    block_id  uuid not null references block,
    parent_id serial references comments,

    user_id   uuid not null,
    content   text,
    image     varchar(250)
);

create table if not exists likes
(
    user_id    uuid not null,
    created_at timestamp default now(),

    block_id   uuid not null references block,
    comment_id serial references comments
);


create table if not exists tags
(
    id    uuid        not null primary key,
    name  varchar(20) not null,
    descr varchar(200)
);

create table if not exists block_tags
(
    block_id uuid references block,

    tag_id   uuid references tags
);

create table if not exists categories
(
    id    uuid        not null primary key,
    name  varchar(20) not null,
    descr varchar(200)
);

create table if not exists block_categ
(
    block_id uuid references block,
    categ_id uuid references categories
);

create or replace function set_default_block() returns trigger
    language plpgsql
as
$$
DECLARE
    temp BOOLEAN;
BEGIN
    IF (NEW.has_comments IS NULL) THEN
        Select has_comments as temp
        from blocks."block_type" as bt
        where bt.id = NEW.id;
        NEW.has_likes := temp;
    elsif(NEW.has_likes IS NUll) THEN
        Select has_comments as temp
        from blocks."block_type" as bt
        where bt.id = NEW.id;
        NEW.has_comments := temp;
    end if;
    return NEW;
END;
$$;

create trigger set_default_values_block_trigger
    before insert
    on block
    for each row
    when (new.has_comments IS NULL OR new.has_likes IS NULL)
execute procedure set_default_block();
