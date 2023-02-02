<script lang="ts">
    import {createForm} from "felte";
    import {useNavigate} from "svelte-navigator";
    import {AxiosError, type AxiosResponse} from "axios";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import type {IAlert, IAuth, ILoginRequest, ILoginResponse, IErrorMessage} from "@/types";
    import {Validators} from "@/libs/validator";
    import {BoxDefault, ButtonDefault, IconEye, IconEyeOff, InputDefault, InputGroup, ErrorMessage} from "@/components";
    import instance from "@/libs/instance";

    export let registerFocus;

    const navigate = useNavigate();

    let loading: boolean;

    const {form, touched, errors, reset} = createForm<ILoginRequest>({
        initialValues: {
            username: "",
            password: "",
        },
        validate: (values) => {
            const errors: ILoginRequest = {
                username: "",
                password: "",
            };

            if (Validators.required(values.username)) {
                errors.username = "Username tidak boleh kosong";
            }

            if (Validators.required(values.password)) {
                errors.password = "Password tidak boleh kosong";
            }

            return errors;
        },
        onSubmit: async (values) => {
            loading = true;
            return await instance.post("/auth/login", values);
        },
        onError: (errors, context) => {
            loading = false;
            if (errors instanceof AxiosError) {
                errors;
                if (errors.response.status === 422) {
                    const {
                        param,
                        error_message,
                    } = errors.response.data.data as IErrorMessage;

                    context.setErrors(param, (value) => error_message);
                }
            }
        },
        onSuccess: (response: AxiosResponse<ILoginResponse>) => {
            reset();
            loading = false;

            const authState: IAuth = {
                nik: response.data.data.nik,
                role: response.data.data.role,
                id_location: response.data.data.id_location,
                isAuthenticated: true,
                photo: response.data.data.photo,
            };

            auth.set(authState);

            const alertState: IAlert = {
                show: true,
                type: "success",
                text: "Kamu berhasil login",
            };

            alertStore.set(alertState);

            navigate("/protected/dashboard", {replace: true});
        },
    });

    let isSecret: boolean = true;
    $: typeFieldPassword = isSecret ? "password" : "text";

    function togglePassword() {
        isSecret = !isSecret;
    }
</script>

<div class="flex justify-center mt-16">
    <BoxDefault>
        <svelte:fragment slot="box-header">
            <img use:registerFocus alt="logo" src="/images/sima.png" class="h-16 lg:h-20"/>
        </svelte:fragment>
        <svelte:fragment slot="box-body">
            <form use:form>
                <div class="flex flex-col w-full mb-4">
                    <InputDefault name="username" type="text" label="Username" placeholder="Username"
                                  isError={$errors.username && $touched.username} readonly={loading}/>
                    {#if !!$errors.username}
                        <ErrorMessage>
                            {$errors.username}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="flex flex-col w-full mb-4">
                    <InputGroup name="password" type="{typeFieldPassword}" label="Password" placeholder="Password"
                                className="pr-12"
                                isError={$errors.password && $touched.password} readonly={loading}>
                        <svelte:fragment slot="append">
                            <button type="button" on:click={togglePassword}>
                                {#if isSecret}
                                    <IconEye
                                            className="icon h-7 w-7 text-gray-500 dark:text-gray-100 {$errors.password && $touched.password ? 'error': ''}"/>
                                {:else }
                                    <IconEyeOff
                                            className="icon h-7 w-7 text-gray-500 dark:text-gray-100 {$errors.password && $touched.password ? 'error': ''}"/>
                                {/if}
                            </button>
                        </svelte:fragment>
                    </InputGroup>
                    {#if !!$errors.password}
                        <ErrorMessage>
                            {$errors.password}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="flex flex-col w-full py-4">
                    <ButtonDefault type="submit" label="Sign In" disabled={loading}/>
                </div>
            </form>
        </svelte:fragment>
    </BoxDefault>
</div>