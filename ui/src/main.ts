import {createApp} from "vue";
import {createPinia} from "pinia";
import {ErrorMessage, Field, Form, configure} from "vee-validate";

import App from "./App.vue";
import router from "./router";

import "flowbite";
import "./assets/main.css";

configure({
    validateOnBlur: true,
    validateOnChange: true,
    validateOnInput: false,
});

const app = createApp(App);
app.component("VForm", Form).component("VField", Field).component("VErrorMessage", ErrorMessage);

app.use(createPinia());
app.use(router);

app.mount("#app");
