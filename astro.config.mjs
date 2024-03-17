import { defineConfig } from "astro/config";

const base = import.meta.env.BASE_URL === "/" ? "/" : "/2024";

// https://astro.build/config
export default defineConfig({
  site: "https://gocon.github.io",
  base: base,
  i18n: {
    defaultLocale: "ja",
    locales: ["ja"],
    routing: {
      prefixDefaultLocale: false,
    },
  },
});
