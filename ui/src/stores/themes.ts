import {defineStore} from "pinia";

export const useThemeStore = defineStore("themes", {
    // Define the state with a default value of mode being set to the value stored in local storage or "light" if not found
    state: () => ({
        mode: window.localStorage.getItem("theme") || "light",
    }),
    // Define getters for the current theme mode and whether the mode is dark
    getters: {
        isDark: (state) => state.mode !== "light",
        getMode: (state) => state.mode,
    },
    // Define an action that updates the value of mode in the state and in local storage
    actions: {
        toggleMode(value: string) {
            this.mode = value;
            window.localStorage.setItem("theme", value);
        },
    },
});
