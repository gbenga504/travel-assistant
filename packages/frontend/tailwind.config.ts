import type { Config } from "tailwindcss";

export default {
  content: ["./app/**/{**,.client,.server}/**/*.{js,jsx,ts,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        sans: [
          "Inter",
          "ui-sans-serif",
          "system-ui",
          "sans-serif",
          "Apple Color Emoji",
          "Segoe UI Emoji",
          "Segoe UI Symbol",
          "Noto Color Emoji",
        ],
      },

      colors: {
        primary: "#2C6BED",
        primaryDark: "#1851B4",
        primaryOpacity1: "rgba(44, 107, 237, 0.1)",
        primaryOpacity5: "rgba(44, 107, 237, 0.5)",
      },
    },
  },
  plugins: [],
} satisfies Config;
