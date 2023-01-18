<script lang="ts" setup>
import {ref, computed} from "vue";
import {useRouter} from "vue-router";
import type {FormActions} from "vee-validate";
import instance from "@/api/instance";
import {AxiosError, type AxiosResponse} from "axios";
import {useAlertStore} from "@/stores/alert";
import {useAuthStore} from "@/stores/auth";
import type {IAuthResponse, ILoginRequest, IAlert, IAuth} from "@/types";
import {BoxDefault, InputDefault, InputAppend, IconEye, IconEyeOff, ButtonDefault} from "@/components";

const router = useRouter();
const alertStore = useAlertStore();
const authStore = useAuthStore();

const initialValues: ILoginRequest = {username: "", password: ""};

const loading = ref<boolean>(false);

const isSecret = ref<boolean>(true);

const typeField = computed(() => {
  return isSecret.value ? "password" : "text";
});

const validationSchema = {
  username: "required",
  password: "required",
};

function togglePassword() {
  isSecret.value = !isSecret.value;
}

async function onSubmit(values: ILoginRequest, actions: FormActions<ILoginRequest>) {
  loading.value = true;
  try {
    const response: AxiosResponse<IAuthResponse> = await instance.post("/auth/login", values);

    const data: IAuth = {
      id_user: response.data.data.id_user,
      kode_lokasi: response.data.data.kode_lokasi,
      role: response.data.data.role,
      isAuthenticated: true,
    };

    authStore.setUserSession(data);

    await router.push({path: "/protected/dashboard"});
  } catch (error: unknown) {
    if (error instanceof AxiosError) {
      if (error.response?.status === 422) {
        const field = error.response.data?.data.Param;
        const message = error.response.data?.data.ErrorMessage;

        actions.setFieldError(field, message);
      }
    } else {
      const alert: IAlert = {
        show: true,
        type: "error",
        text: "Terjadi kesalahan tidak diketahui",
      };

      alertStore.showAlert(alert);
    }
  } finally {
    loading.value = false;
  }

}
</script>

<template>
  <div class="flex justify-center mt-16">
    <BoxDefault>
      <template #box-header>
        <img alt="logo" class="h-16 lg:h-20" src="/images/logo.jpeg"/>
      </template>
      <template #box-body>
        <VForm :initial-values="initialValues" :validation-schema="validationSchema" v-slot="{ meta: formMeta }"
               @submit="onSubmit">
          <div class="flex flex-col w-full mb-4">
            <InputDefault name="username" placeholder="Username..." label="Username" type="text" :readonly="loading"/>
          </div>
          <div class="flex flex-col w-full mb-4">
            <InputAppend name="password" placeholder="Password..." label="Password" :type="typeField"
                         :readonly="loading" v-slot="slotProps"
                         @onClickIcon="togglePassword">
              <IconEye v-if="isSecret" className="icon h-7 w-7 text-gray-400 dark:text-gray-100" :class="{
              'error': slotProps.isError
            }"/>
              <IconEyeOff v-else className="icon h-7 w-7 fill-current dark:text-gray-100" :class="{
              'error': slotProps.isError
            }"/>
            </InputAppend>
          </div>
          <div class="flex flex-col w-full py-4">
            <ButtonDefault label="Sign In" type="submit"
                           className="bg-blue-700 text-white hover:bg-blue-800 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus-gray-500"
                           disabledClassName="cursor-not-allowed bg-blue-500 dark:bg-gray-500" :disabled="loading"/>
          </div>
        </VForm>
      </template>
    </BoxDefault>
  </div>
</template>