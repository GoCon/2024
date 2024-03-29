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
export const collections = {
  sponsors: sponsorsCollection,
};
