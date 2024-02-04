/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../templates/html/*.html"],
  theme: {
    extend: {
      fontFamily: {},
      colors: {},
    },
  },
  plugins: [require("daisyui")],
};
