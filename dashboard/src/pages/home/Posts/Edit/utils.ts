import { z } from "zod";
import axios from "axios";
import { useQuery } from "@tanstack/react-query";

export const BlockMetaSchema = z.object({
  name: z.string(),
  description: z.string(),
  type: z.string(),
  tags: z.array(z.string()),
  categories: z.array(z.string()),
});

export const BlockRulesSchema = z.object({
  rulename: z.string().optional(),
  nested: z.boolean(),
  hasLikes: z.boolean(),
  hasComments: z.boolean(),
  commentsHasLikes: z.boolean(),
  commentsEditable: z.boolean(),
  commentsMaxNested: z.number(),
});

export const BlockSchema = z.object({
  id: z.string(),
  name: z.string(),
  description: z.string(),
  type: z.string(),
  tags: z.array(z.string()),
  categories: z.array(z.string()),
  rules: BlockRulesSchema,
  content: z.array(
    z.object({
      element: z.string(),
    })
  ),
});

export const BlockContentSchema = z.object({
  content: z.array(
    z.union([
      z.object({ text: z.object({ type: z.string() }) }),
      z.object({ media: z.object({ type: z.string() }) }),
    ])
  ),
});

export function useBlock(id: string, lang: string) {
  return useQuery(["block", id, lang], async () => {
    const res = await axios.get(`public/block/full/${lang}/${id}`);

    const validationResult = BlockSchema.safeParse(res.data);

    if (validationResult.success) {
      return validationResult.data;
    } else {
      throw new TypeError(validationResult.error.message);
    }
  });
}

export function useBlockRules(id: string) {
  return useQuery(["block-rules", id], async () => {
    const res = await axios.get(`public/block/rules/${id}`);

    const validationResult = BlockRulesSchema.safeParse(res.data);

    if (validationResult.success) {
      return validationResult.data;
    } else {
      throw new TypeError(validationResult.error.message);
    }
  });
}

export function useBlockMeta(id: string) {
  return useQuery(["block-meta", id], async () => {
    const res = await axios.get(`public/block/meta/${id}`);

    const validationResult = BlockMetaSchema.safeParse(res.data);

    if (validationResult.success) {
      return validationResult.data;
    } else {
      throw new TypeError(validationResult.error.message);
    }
  });
}

export function useBlockContent(id: string, lang: string) {
  return useQuery(["block-content", id, lang], async () => {
    const res = await axios.get(`public/block/content/${lang}/${id}`);

    const validationResult = BlockContentSchema.safeParse(res.data);

    if (validationResult.success) {
      return validationResult.data;
    } else {
      throw new TypeError(validationResult.error.message);
    }
  });
}
