import {defineStore} from "pinia";

export const useBreakPointStore = defineStore("breakpoint", {
    state: () => ({
        isMobile: false,
    }),
    getters: {
        checkIsMobile: (state): boolean => state.isMobile,
    },
    actions: {
        updateCheckMobile(value: boolean) {
            this.isMobile = value;
        },
    },
});