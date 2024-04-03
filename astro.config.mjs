import { defineConfig } from "astro/config";

export default defineConfig({
  site: "https://gocon.github.io",
  base: "/2024",
  i18n: {
    defaultLocale: "ja",
    locales: ["ja"],
    routing: {
      prefixDefaultLocale: false,
    },
  },
});
