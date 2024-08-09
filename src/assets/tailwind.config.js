/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,js}","../templates/**/*.html","./node_modules/preline/dist/*.js"],
  theme: {
    extend: {},
  },
  plugins: [
    require('preline/plugin'),
  ],
}

