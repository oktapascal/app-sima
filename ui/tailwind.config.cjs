/** @type {import("tailwindcss").Config} */
module.exports = {
    content: [
        "./index.html",
        "./src/**/*.{html,js,svelte,ts}",
    ],
    theme: {
        minWidth: {
            8: "8rem",
            15: "15rem",
        },
        minHeight: {
            2: "2rem",
            3: "3rem",
            15: "15rem",
        },
        extend: {
            width: {
                81: "21rem",
            },
            transitionProperty: {
                "width": "width",
            },
            spacing: {
                1.3: "1.375rem",
                1.6: "1.625rem",
                3.1: "3.125rem",
                5.5: "5.625rem",
                12.5: "12.5rem",
                23: "5.5rem",
            },
            fontSize: {
                "10xl": "5rem",
                "11xl": "8.75rem",
            },
        },
    },
    plugins: [
        require("flowbite/plugin"),
    ],
};
