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
                23: "5.5rem",
            },
        },
    },
    plugins: [
        require("flowbite/plugin"),
    ],
};
