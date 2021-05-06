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
                gray: colors.trueGray
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