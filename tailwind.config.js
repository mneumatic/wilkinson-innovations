/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/*.html',
    '/public/js/*.js'
  ],
  theme: {
    extend: {
      container: {
        center: true,
        padding: {
          DEFAULT: '1rem',
          md: '0'
        }
      }
    }
  },
  plugins: []
}
