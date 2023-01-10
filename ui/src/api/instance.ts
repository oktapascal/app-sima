import axios from "axios"
import routes from "@/router"
import {useAlertStore, type IAlert} from "@/stores/alert";

// Import the base URL from the environment
const BASE_URL = import.meta.env.VITE_BASE_URL

// Create a new instance of axios with specified base URL
const instance = axios.create({
    withCredentials: true, // Include credentials in requests
    baseURL: BASE_URL // Set the base URL for all requests
})

// Add an interceptor to all requests made with this instance
// of axios to include credentials
instance.interceptors.request.use(function (config) {
    config.withCredentials = true // Include credentials in request

    return config
}, function (error) {
    return Promise.reject(error) // Reject the request if there is an error
})

// Set up an Axios response interceptor to handle HTTP errors
instance.interceptors.response.use(
    // If the request was successful, simply return the response
    function (response) {
        return response
    },
    // If the request returned an error, handle it here
    async function (error) {
        // Get a reference to the alert store
        const alertStore = useAlertStore()

        // If the error is a 401 Unauthorized error, redirect the user to the login page
        if(error.response.status === 401) {
            let redirectUrl = ""
            if(routes.currentRoute.value.path !== "/login") {
                redirectUrl = routes.currentRoute.value.path
            }

            await routes.push({
                path: "/login",
                query: {
                    redirect: redirectUrl
                }
            })
        }

        // If the error is a 404 Not Found error, show an error alert with the status message
        if(error.response.status === 404) {
            const alert: IAlert = {
                show: true,
                type: "error",
                text: error.response.data?.statusMessage
            }

            alertStore.showAlert(alert)
        }

        // If the error is a 500 Internal Server Error, show a generic error alert
        if (error.response?.status === 500) {
            const alert: IAlert = {
                show: true,
                type: "error",
                text: "Terjadi kesalahan pada server"
            }

            alertStore.showAlert(alert)
        }

        // Return the error so it can be handled by a catch block in the code that made the HTTP request
        return Promise.reject(error)
    }
)

// Export the Axios instance with the interceptor attached
export default instance
