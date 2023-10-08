/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./views/components/**/*.html",
        "./views/layout/**/*.html",
        "./views/partials/**/*.html",
        "./views/screens/**/*.html",
        "./svelte/**/*.html"
    ],
    theme: {
        colors: {
            'white': "#FFFFFF",
            'black': "#000000",
            'main': "#F8F8F8",
            'primary': "#3627E7",
            'disabled': "#535073",
            'hover': "#594DE7"
        },
        fontFamily: {
            sans: ['Inter', 'sans-serif']
        },
        extend: {},
    },
    plugins: [],
}

