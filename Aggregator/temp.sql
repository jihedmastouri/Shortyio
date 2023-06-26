SELECT json_agg(
 json_build_object(
         'id', b.id,
         'has_comments', b.has_comments,
         'has_likes', b.has_likes,
         'version_number', bl.version_number,
         'created_at', b.created_at,
         'updated_at', bl.updated_at,
         'author', b.author,
         'block_type', bt.name,
         'lang_name', bl.lang_name,
         'lang_code', bl.lang_code,
         'content', (
             json_build_object(
             'elements', (SELECT array_agg(
                     json_build_object(
                         'media',
                         json_build_object(
                                 'title',
                                 bi.title,
                                 'type',
                                 'MEDIA_TYPE_IMAGE',
                                 'file',
                                 bi.file,
                                 'alt', bi.alt
                             )
                         )
                     )
                  FROM block_images bi
                  WHERE bi.block_lang_id = bl.id) ||
                 (SELECT array_agg(
                     json_build_object(
                         'text',
                         json_build_object(
                                 'name',
                                 brt.name,
                                 'type',
                                 'TEXTUAL_TYPE_HTML',
                                 'content',
                                 brt.content,
                                 'hint',
                                 brt.hint
                             )
                         )
                     )
                  FROM block_rich_texts brt
                  WHERE brt.block_lang_id = bl.id) ||
                 (SELECT array_agg(
                         json_build_object(
                             'text',
                             json_build_object(
                                 'name',
                                 bt.name,
                                 'type',
                                 'TEXTUAL_TYPE_TEXT',
                                 'content',
                                 bt.content,
                                 'hint', bt.hint
                             )
                         )
                     )
              FROM block_texts bt
              WHERE bt.block_lang_id = bl.id)
            )
        )
     )
)
FROM blocks b
       INNER JOIN block_langs bl ON b.id = bl.block_id
       INNER JOIN block_types bt ON b.type = bt.id
WHERE b.id = '71e9706d-b03c-4aff-a032-31f075dae868'
AND bl.lang_code = 'en_US'
LIMIT 1;
