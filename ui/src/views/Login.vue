<script lang="ts" setup>
// Import various functions and types from the Vue.js library
import { ref, computed } from "vue"

// Import the useRouter hook from the vue-router library
import {useRouter} from "vue-router";

// Import the zod library for schema validation
import * as zod from "zod"

// Import the FormActions type from the vee-validate library
import type { FormActions } from "vee-validate";

// Import the toFormValidator function for converting a zod schema to a vee-validate form validator
import { toFormValidator } from "@vee-validate/zod"

// Import an instance of the axios library for making HTTP requests
import instance from "@/api/instance";

// Import various types from the axios library
import {AxiosError, type AxiosResponse} from "axios";

// Import the useAlertStore and IAlert types from the alert store
import {useAlertStore, type IAlert} from "@/stores/alert";

// Import the useAuthStore and IAuth types from the auth store
import {useAuthStore, type IAuth} from "@/stores/auth";

// Import various Vue components
import BoxDefault from "@/components/box/BoxDefault.vue";
import InputDefault from "@/components/input/InputDefault.vue";
import InputAppend from "@/components/input/InputAppend.vue";
import IconEye from "@/components/icon/IconEye.vue";
import IconEyeOff from "@/components/icon/IconEyeOff.vue";
import ButtonDefault from "@/components/button/ButtonDefault.vue";

// Define the shape of a login request object
interface LoginRequest {
  username: string
  password: string
  [key: string]: any
}

// Define a generic response interface with a data field of type T
interface Response<T> {
  data: T
}

// Define an interface for an authentication response, extending the Response interface and specifying the shape of the data field
interface AuthResponse extends Response<{
  id_user: string
  kode_lokasi: string
  role: string
  isAuthenticated: boolean
}> {}

// Use the useRouter hook from the vue-router library to access the Vue router instance
const router = useRouter()

// Use the useAlertStore custom hook to access the alert store in the component
const alertStore = useAlertStore()

// Use the useAuthStore custom hook to access the auth store in the component
const authStore = useAuthStore()

// Define the initial values for the login form fields
const initialValues: LoginRequest = { username: "", password: "" }

// Create a reactive variable to track whether the form is currently being submitted
const loading = ref<boolean>(false)

// Create a reactive variable to toggle the visibility of the password field in the form
const isSecret = ref<boolean>(true)

// Create a computed property that returns "password" or "text" based on the value of isSecret
const typeField = computed(() => {
  return isSecret.value ? "password" : "text"
})

// Use the toFormValidator function to convert a zod schema to a vee-validate form validator
const validationSchema = toFormValidator(
    zod.object({
      // Validate that the username field is a string with a minimum length of 1
      username: zod.string().min(1, "Wajib Diisi"),
      // Validate that the password field is a string with a minimum length of 1
      password: zod.string().min(1, "Wajib Diisi")
    }))

// Toggle the value of the isSecret reactive variable between true and false
function togglePassword() {
  isSecret.value = !isSecret.value
}

// Define an async function to handle the submission of the login form
async function onSubmit(values: LoginRequest, actions: FormActions<LoginRequest>) {
  // Set the loading reactive variable to true to indicate that the form is being submitted
  loading.value = true
  try {
    // Make a POST request to the /auth/login endpoint with the login form values as the request body
    const response: AxiosResponse<AuthResponse> = await instance.post("/auth/login", values)

    // Extract the relevant data from the response and store it in an IAuth object
    const data: IAuth = {
      id_user: response.data.data.id_user,
      kode_lokasi: response.data.data.kode_lokasi,
      role: response.data.data.role,
      isAuthenticated: true
    }

    // Store the user's authentication data in the auth store
    authStore.setUserSession(data)

    // Navigate to the protected dashboard route
    await router.push({ path: "/protected/dashboard" })
  } // Define the catch block of the try-catch statement in the onSubmit function
  catch (error: unknown) {
    // Check if the error is an instance of AxiosError
    if (error instanceof AxiosError) {
      // Check the status of the error's response
      if (error.response?.status === 422) {
        // Extract the field and message from the response data
        const field = error.response.data?.data.Param
        const message = error.response.data?.data.ErrorMessage

        // Display the error message in the form next to the appropriate field
        actions.setFieldError(field, message)
      }
    } else {
      // If the error is not an instance of AxiosError, show an error alert to the user
      const alert: IAlert = {
        show: true,
        type: "error",
        text: "Terjadi kesalahan tidak diketahui"
      }

      alertStore.showAlert(alert)
    }
  } // Reset the loading state in the finally block
  finally {
    loading.value = false
  }

}
</script>

<template>
  <div class="flex justify-center mt-16">
    <BoxDefault>
      <template #box-header>
        <img alt="logo" class="h-16 lg:h-20" src="/images/logo.jpeg" />
      </template>
      <template #box-body>
        <VForm :initial-values="initialValues" :validation-schema="validationSchema" v-slot="{ meta: formMeta }"
          @submit="onSubmit">
          <InputDefault name="username" placeholder="Username..." label="Username" type="text" :readonly="loading" />
          <InputAppend name="password" placeholder="Password..." label="Password" :type="typeField" :readonly="loading" v-slot="slotProps"
            @onClickIcon="togglePassword">
            <IconEye v-if="isSecret" className="icon h-7 w-7 fill-current dark:text-gray-100" :class="{
              'error': slotProps.isError
            }" />
            <IconEyeOff v-else className="icon h-7 w-7 fill-current dark:text-gray-100" :class="{
              'error': slotProps.isError
            }" />
          </InputAppend>
          <ButtonDefault label="Sign In" type="submit"
            className="bg-blue-700 text-white hover:bg-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500"
            disabledClassName="cursor-not-allowed bg-blue-500 dark:bg-gray-500" :disabled="loading" />
        </VForm>
      </template>
    </BoxDefault>
  </div>
</template>