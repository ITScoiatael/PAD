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
        'hero-image': "url('http://localhost:8080/static/slavesImage.png')",
        'GrayImage': "url('http://localhost:8080/static/Background.png')",
      }),
      backgroundSize: {
        'auto': 'auto',
        '50%': '50%',
        '25%': '25%',
        '20%': '20%',
        '18%': '18%',
        '15%': '15%',
      },
      backgroundColor: theme => ({
        ...theme('colors'),
      }),
      backgroundPosition: {
        'center-right': 'center right 22%',
      },
      colors: {
        emerald: colors.emerald,
        gray: colors.trueGray,
        primary: '#CDC1C1',
        'light-blue': colors.lightBlue,
        cyan: colors.cyan,
        yellow: {
          'bgcol': '#F6F4EE',
          'buycol': '#E1E0DC',
          'bgbuycol': '#E5E5E5'
        }
      },
      screen: {
        'lg': '1024px'
      },
      placeholderColor: {
        'white': '#fff'
      },
      fontFamily: {
        'gilroy': 'Gilroy',
      },
      width: {
        '1':'1px',
        '100': '100px',
        '150': '150px',
        '200': '200px',
        '300':'400px',
        '400':'400px',
        '700': '700px',
        '1': '1px',
        '1':'1px',
      },
      height: {
        '150': '150px',
        '200': '200px',
        '400': '400px',
        '430': '430px'
      },
      minHeight: {
        '1/4': '25%',
        '1/2': '50%',
        '3/4': '75%',
        'full': '100%',
       },
      rounded: {
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