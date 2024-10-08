/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["ui/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
  daisyui: {
    themes: ["light"],
  },
};
