/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/**/*.{html,svelte}"
    ],
    theme: {
        extend: {
          colors: {
            'main': "#F8F8F8",
            'primary': "#3627E7",
            'disabled': "#535073",
            'hover': "#594DE7"
        },
        },
    },
    plugins: [],
}

