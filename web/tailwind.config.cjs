/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        "primary": "#439A97",
        secondary: "#62B6B7",
        minor: "#97DECE",
        major: "#CBEDD5"
      }
    },
  },
  plugins: [],
};
