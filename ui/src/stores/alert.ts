import {defineStore} from "pinia"

export interface IAlert {
    text: string
    type: string
    show: boolean
}

// Export the useAlertStore function, which is defined using the defineStore function
export const useAlertStore = defineStore("alert", {
    // Declare the state object
    state: function () {
        return {
            text: "", // The text of the alert
            type: "info", // The type of the alert
            show: false // Whether the alert is shown or not
        }
    },
    // Declare the getters object
    getters: {
        // Declare the getAlert getter that takes a state object and returns an IAlert object
        getAlert: function (state): IAlert {
            return state
        }
    },
    // Declare the actions object
    actions: {
        // Declare the showAlert action that takes an IAlert object and updates the state properties accordingly
        showAlert: function (data: IAlert) {
            this.text = data.text
            this.type = data.type
            this.show = data.show
        },
    }
})