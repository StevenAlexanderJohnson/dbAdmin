module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx,svelte}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'Courier-Prime': ['Courier Prime', 'monospace']
      },
      colors: {
        'background': '#051621',
        'primary': '#82c9ea',
        'secondary': '#183695',
        'accent': '#2d2fdd',
        'text': '#d7ebf8'
      }
    },
  },
  plugins: [],
}
