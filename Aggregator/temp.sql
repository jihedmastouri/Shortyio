SELECT
    json_build_object(
        'blocks', (
            SELECT json_agg(
                json_build_object(
                    'id', b.id,
                    'has_comments', b.has_comments,
                    'has_likes', b.has_likes,
                    'version_number', b.version_number,
                    'created_at', b.created_at,
                    'updated_at', b.updated_at,
                    'author', b.author,
                    'block_type', bt.name,
                    'comments_type', ct.name,
                    'langs', (
                        SELECT json_agg(
                            json_build_object(
                                'id', bl.id,
                                'lang_name', bl.lang_name,
                                'lang_code', bl.lang_code,
                                'images', (
                                    SELECT json_agg(
                                        json_build_object(
                                            'id', bi.id,
                                            'file', bi.file,
                                            'alt', bi.alt,
                                            'title', bi.title
                                        )
                                    )
                                    FROM block_images bi
                                    WHERE bi.block_lang_id = bl.id
                                ),
                                'texts', (
                                    SELECT json_agg(
                                        json_build_object(
                                            'id', bt.id,
                                            'content', bt.content,
                                            'name', bt.name,
                                            'hint', bt.hint
                                        )
                                    )
                                    FROM block_texts bt
                                    WHERE bt.block_lang_id = bl.id
                                ),
                                'rich_texts', (
                                    SELECT json_agg(
                                        json_build_object(
                                            'id', brt.id,
                                            'content', brt.content,
                                            'name', brt.name,
                                            'hint', brt.hint
                                        )
                                    )
                                    FROM block_rich_texts brt
                                    WHERE brt.block_lang_id = bl.id
                                )
                            )
                        )
                        FROM block_langs bl
                        WHERE bl.block_id = b.id
                    )
                )
            )
            FROM blocks b
            LEFT JOIN block_types bt ON b.block_type = bt.id
            LEFT JOIN comment_types ct ON b.comments_type = ct.id
        ),
        'tags', (
            SELECT json_agg(
                json_build_object(
                    'id', t.id,
                    'name', t.name,
                    'descr', t.descr
                )
            )
            FROM tags t
            JOIN block_tags bt ON t.id = bt.tag_id
        ),
        'categories', (
            SELECT json_agg(
                json_build_object(
                    'id', c.id,
                    'name', c.name,
                    'descr', c.descr
                )
            )
            FROM categories c
            JOIN block_categ bc ON c.id = bc.categ_id
        )
    ) AS aggregated_json;
