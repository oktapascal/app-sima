/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "./node_modules/flowbite/**/*.js"
  ],
  theme: {
    screens: {
      "sm": "640px",
      "md": "768px",
      "lg": "1024px",
      "xl": "1280px",
      "2xl": "1536px",
    },
    minWidth: {
      8: '8rem',
      15: '15rem'
    },
    minHeight: {
      2: '2rem',
      3: '3rem',
      15: '15rem'
    },
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}
