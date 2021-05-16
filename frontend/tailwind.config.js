const colors = require('tailwindcss/colors');
module.exports = {
  purge: [
    './components/**/*.{vue,js}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './nuxt.config.{js,ts}',
  ],
  darkMode: false,
  theme: {
    extend: {
      backgroundImage: theme => ({
        'hero-image': "url('http://localhost:8080/static/slavesImage.png')"
      }),
      backgroundSize: {
        'auto': 'auto',
        '50%': '50%',
        '25%': '25%',
        '20%': '20%',
        '15%': '15%',
      },
      backgroundColor: theme => ({
        ...theme('colors'),
      }),
      backgroundPosition: {
       'center-right': 'center right 25rem',
      },
      fontFamily: {
        'gilroy': ['Gilroy'],
      },
      colors: {
        emerald: colors.emerald,
        gray: colors.trueGray,
        primary: '#CDC1C1',
        'light-blue': colors.lightBlue,
        cyan: colors.cyan,
        'light-blue': colors.lightBlue,
        cyan: colors.cyan,
        yellow:{
          'bgcol':'#F6F4EE',
          'buycol':'#E1E0DC',
          'bgbuycol':'#E5E5E5'
        }
      },
      w: {
        '10': '10px',
        '20': '20px',
        '100p': '100%',
        '100': '100px',
        '300':'400px',
        '1':'1px',
      },
      h: {
        '10': '10px',
        '20': '20px',
        '5': '5px'
      },
      mr: {
        '10': '10px',
        '10': '10px'
      },
      mb: {
        '10': '10px',
        '10': '10px'
      },
      md: {
        'md5': '5px',
        'md10': '10px',
        'md5': '5px',
        'md10': '10px'
      },
      height:{
        '330':'430px'
      },
      rounded:{
        'md': "10%",
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/forms")
  ],
}