import {defineStore} from "pinia"

export interface IAlert {
    text: string
    type: string
    show: boolean
}

export const useAlertStore = defineStore("alert", {
    state: function () {
        return {
            text: "",
            type: "info",
            show: false
        }
    },
    getters: {
        getAlert: function (state): IAlert {
            return state
        }
    },
    actions: {
        showAlert: function (data: IAlert) {
            this.text = data.text
            this.type = data.type
            this.show = data.show
        },
    }
})