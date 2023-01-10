import axios, { AxiosError, type AxiosResponse } from "axios"; // Import axios and its types
import { useAlertStore, type IAlert } from "./alert"; // Import useAlertStore hook and IAlert type from alert module
import { defineStore } from "pinia"; // Import defineStore function from pinia

// Define the Response interface with a single property, data, of generic type T
interface Response<T> {
    data: T
}

// Define the AuthResponse interface that extends the Response interface and defines the shape of the data property as an object with 3 properties
interface AuthResponse extends Response<{
    id_user: string
    kode_lokasi: string
    role: string
}> {}

// Define the IAuth interface with 4 properties: id_user, kode_lokasi, role, and isAuthenticated
export interface IAuth {
    id_user: string|null
    kode_lokasi: string|null
    role: string|null
    isAuthenticated: boolean
}

export const useAuthStore = defineStore("auth", {
    state: function () {
        // Declare the state object that contains 4 properties: id_user, kode_lokasi, role, and isAuthenticated
        return {
            id_user: null, // The user's ID
            kode_lokasi: null, // The user's location code
            role: null, // The user's role
            isAuthenticated: false // Whether the user is authenticated or not
        }
    },
    // Declare the getters object
    getters: {
        // Declare the getUserSession getter that takes a state object and returns it
        getUserSession: function (state: IAuth) {
            return state
        },
        // Declare the getStatusAuthenticated getter that takes a state object and returns the value of the isAuthenticated property
        getStatusAuthenticated: function (state: IAuth) {
            return state.isAuthenticated
        }
    },
    actions: {
        // Declare the setUserSession action that takes a data object
        setUserSession: function (data: IAuth) {
            // Set the id_user, kode_lokasi, role, and isAuthenticated properties of the state object to the corresponding values in the data object
            this.id_user = data.id_user
            this.kode_lokasi = data.kode_lokasi
            this.role = data.role
            this.isAuthenticated = data.isAuthenticated
        },
        // Declare the setUserSessionFromCredentials action
        setUserSessionFromCredentials: async function() {
            // Get the alert store
            const alertStore = useAlertStore()
            // Save the current context to a variable
            const self = this
            try {
                // Make a GET request to the specified URL with credentials
                const response: AxiosResponse<AuthResponse> = await axios.get("/api/auth/session/user-access", {
                    withCredentials: true
                })
                // Create a data object from the response data
                const data: IAuth = {
                    id_user: response.data.data.id_user,
                    role: response.data.data.role,
                    kode_lokasi: response.data.data.kode_lokasi,
                    isAuthenticated: true
                }
                // Call the setUserSession action with the data object
                self.setUserSession(data)
            } catch (error: unknown) {
                // If the error is an instance of AxiosError
                if(error instanceof AxiosError) {
                    // If the error's response status is 500
                    if(error.response?.status === 500) {
                        // Create an alert object
                        const alert: IAlert = {
                            show: true,
                            type: "error",
                            text: error.response.data?.statusMessage
                        }
                        // Call the showAlert action with the alert object
                        alertStore.showAlert(alert)
                    }
                }
            }
        }
    }
})