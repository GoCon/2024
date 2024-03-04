import { z, defineCollection } from "astro:content";

const sponsorsCollection = defineCollection({
  type: "data",
  schema: z.array(
    z.object({
      id: z.number(),
      name: z.string(),
      class: z.enum(["Platinum Gold", "Gold", "Silver", "Bronze"]),
      description: z.string(),
      image: z.string(),
    }),
  ),
});

const staffCollection = defineCollection({
  type: "data",
  schema: z.array(
    z.object({
      id: z.number(),
      name: z.string(),
     thumbnail: z.string().url(),
      organizationName: z.string(),
      twitterId: z.string(),
      githubId: z.string(),
      otherLink: z.string().url(),
    }),
  ),
});

export const collections = {
  sponsors: sponsorsCollection,
  staff: staffCollection,
};
