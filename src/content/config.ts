import { z, defineCollection } from "astro:content";

const sessions = defineCollection({
  type: "data",
  schema: z.array(
    z.object({
      id: z.number(),
      type: z.enum(["challenge", "long", "short", "lt"]),
      level: z.enum(["advanced", "all", "beginner", "intermediate"]),
      track: z.enum(["room-1", "room-2"]),
      title: z.string(),
      description: z.string(),
      speaker: z.object({
        avatar: z.string(),
        name: z.string(),
        company: z.string(),
      }),
    }),
  ),
});

const sponsors = defineCollection({
  type: "data",
  schema: z.array(
    z.object({
      id: z.number(),
      name: z.string(),
      class: z.enum([
        "Venue",
        "Platinum Gold",
        "Gold",
        "Silver",
        "Bronze",
        "Gopher",
      ]),
      description: z.string(),
      image: z.string(),
    }),
  ),
});

export const collections = {
  sessions,
  sponsors,
};
