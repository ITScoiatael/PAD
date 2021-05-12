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
            colors: {
                emerald: colors.emerald,
                gray: colors.trueGray,
                primary: '#9A9385',
                'light-blue': colors.lightBlue,
                cyan: colors.cyan,
            },
            w:{
                '10': '10px',
                '20': '20px',
                '100p': '100%',
              },
              h:{
                '10': '10px',
                '20': '20px',
                '5': '5px'
              },
              mr:{
                '10': '10px'
              },
              mb:{
                '10': '10px'
              },
              md:{
                'md5': '5px',
                'md10': '10px'
              }
        },
    },
    variants: {
        extend: {},
    },
    plugins: [
        require("@tailwindcss/forms")
    ],
}