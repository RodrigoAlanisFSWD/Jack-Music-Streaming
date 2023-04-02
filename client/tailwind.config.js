/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue"
  ],
  theme: {
    extend: {
      colors: {
        "primary": "#B18CFF",
        "secondary": "#181818",
        "light": "#666",
        "primaryLight": "rgb(177,140,255,25)"
      }
    }
      
  },
  plugins: [],
}
