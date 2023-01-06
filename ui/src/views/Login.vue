<script lang="ts" setup>
import { ref, computed } from "vue"
import {useRouter} from "vue-router";
import * as zod from "zod"
import type { FormActions } from "vee-validate";
import { toFormValidator } from "@vee-validate/zod"
import instance from "@/api/instance";
import {AxiosError, type AxiosResponse} from "axios";
import {useAlertStore, type IAlert} from "@/stores/alert";
import {useAuthStore, type IAuth} from "@/stores/auth";
import BoxDefault from "@/components/box/BoxDefault.vue";
import InputDefault from "@/components/input/InputDefault.vue";
import InputAppend from "@/components/input/InputAppend.vue";
import IconEye from "@/components/icon/IconEye.vue";
import IconEyeOff from "@/components/icon/IconEyeOff.vue";
import ButtonDefault from "@/components/button/ButtonDefault.vue";

interface LoginRequest {
  username: string
  password: string
  [key: string]: any
}

interface Response<T> {
  data: T
}

interface AuthResponse extends Response<{
  id_user: string
  kode_lokasi: string
  role: string
  isAuthenticated: boolean
}> {}

const router = useRouter()

const alertStore = useAlertStore()
const authStore = useAuthStore()

const initialValues: LoginRequest = { username: "", password: "" }


const loading = ref<boolean>(false)
const isSecret = ref<boolean>(true)

const typeField = computed(() => {
  return isSecret.value ? "password" : "text"
})

const validationSchema = toFormValidator(
  zod.object({
    username: zod.string().min(1, "Wajib Diisi"),
    password: zod.string().min(1, "Wajib Diisi")
  }))

function togglePassword() {
  isSecret.value = !isSecret.value
}

async function onSubmit(values: LoginRequest, actions: FormActions<LoginRequest>) {
  loading.value = true
  try {
    const response: AxiosResponse<AuthResponse> = await instance.post("/auth/login", values)

    const data: IAuth = {
      id_user: response.data.data.id_user,
      kode_lokasi: response.data.data.kode_lokasi,
      role: response.data.data.role,
      isAuthenticated: true
    }

    authStore.setUserSession(data)

    await router.push({ path: "/protected/dashboard" })
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (error.response?.status === 422) {
        const field = error.response.data?.data.Param
        const message = error.response.data?.data.ErrorMessage

        actions.setFieldError(field, message)
      }
    } else {
      const alert: IAlert = {
        show: true,
        type: "error",
        text: "Terjadi kesalahan tidak diketahui"
      }

      alertStore.showAlert(alert)
    }
  } finally {
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