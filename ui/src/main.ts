import {createApp} from "vue";
import {createPinia} from "pinia";
import {ErrorMessage, Field, Form, configure, defineRule} from "vee-validate";

import App from "./App.vue";
import router from "./router";

import "flowbite";
import "./assets/main.css";

// Validation configuration
// This code configures the validation settings for the application
configure({
    // Validate input when user leaves the field
    validateOnBlur: true,
    // Validate input every time it's changed
    validateOnChange: true,
    // Do not validate input as the user is typing
    validateOnInput: false,
});

// Defining "required" validation rule
// This rule checks that the input field is not empty and returns an error message if it is
defineRule("required", (value: string) => {
    // Check if the field is empty
    if(!value || value.length === 0) {
        // Return error message
        return "Field ini wajib diisi"
    }

    // Return true if field is not empty
    return true
});

// Defining "min" validation rule
// This rule checks that the input field has a minimum length and returns an error message if it doesn't
defineRule("min", (value: string, [min]:[number]) => {
    // Check if the length of the field is less than the minimum
    if(value.length < min) {
        // Return error message
        return `Field ini harus terdiri dari minimal ${min} karakter`
    }

    // Return true if field meets minimum length requirement
    return true
});

// Defining "max" validation rule
// This rule checks that the input field has a maximum length and returns an error message if it exceeds
defineRule("max", (value: string, [max]:[number]) => {
    // Check if the length of the field is greater than the maximum
    if(value.length > max) {
        // Return error message
        return `Field ini harus terdiri dari minimal ${max} karakter`
    }

    // Return true if field meets maximum length requirement
    return true
});

// Defining "email" validation rule
// This rule checks that the input field is a valid email address and returns an error message if it's not
defineRule("email", (value: string) => {
    // Test if email address is valid using regular expression
    const isEmailValid = (/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/).test(value)

    // If email address is invalid
    if(!isEmailValid) {
        // Return error message
        return "Field ini diisi dengan email tidak valid"
    }
    // Return true if the email is valid
    return true
});

// Create application instance
const app = createApp(App);

// Registering additional components
app.component("VForm", Form).component("VField", Field).component("VErrorMessage", ErrorMessage);

// Use Pinia plugin
app.use(createPinia());

// Use router for handling routing
app.use(router);

// Mount the application to the element with id "#app"
app.mount("#app");

