import {defineStore} from "pinia";

export const useBreakPointStore = defineStore("breakpoint", {
    // Define the state with a default value of isMobile being set to false
    state: () => ({
        isMobile: false,
    }),
    // Define a getter that returns the value of isMobile from the state
    getters: {
        checkIsMobile: (state): boolean => state.isMobile,
    },
    // Define an action that updates the value of isMobile in the state
    actions: {
        updateCheckMobile(value: boolean) {
            this.isMobile = value;
        },
    },
});