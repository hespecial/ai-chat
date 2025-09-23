/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        brand: {
          50: '#eef7ff',
          100: '#d9ecff',
          200: '#b9dcff',
          300: '#8cc7ff',
          400: '#57acff',
          500: '#2f91ff',
          600: '#1f78e6',
          700: '#1a60b8',
          800: '#1b5192',
          900: '#1b4475',
        },
      },
      fontFamily: {
        sans: ['Poppins', 'Noto Sans SC', 'Inter', 'ui-sans-serif', 'system-ui', 'sans-serif'],
        display: ['Poppins', 'Noto Sans SC', 'Inter', 'ui-sans-serif', 'system-ui', 'sans-serif'],
      },
      container: { center: true, padding: '1rem' },
    },
  },
  plugins: [],
}
