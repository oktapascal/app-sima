<script lang="ts">
    import {onMount} from "svelte";
    import {createForm} from "felte";
    import {AxiosError, type AxiosResponse} from "axios";
    import {alertStore} from "@/stores/alertStore";
    import {auth} from "@/stores/authStore";
    import type {IProfileRequest, IProfileResponse, IAlert, IUploadResponse, IErrorMessage} from "@/types";
    import {IconCamera, InputDefault, ButtonDefault, ErrorMessage} from "@/components";
    import {Validators} from "@/libs/validator";
    import instance from "@/libs/instance";

    let uploadFile;
    let avatar;
    let loading: boolean;

    const {form, errors, touched, setData, data} = createForm<IProfileRequest>({
        initialValues: {
            name: null,
            email: null,
            no_telp: null,
            address: null,
        },
        validate: (values) => {
            const errors: IProfileRequest = {
                name: null,
                no_telp: null,
                email: null,
                address: null,
            };

            if (Validators.required(values.name)) {
                errors.name = "Nama tidak boleh kosong";
            }

            if (Validators.required(values.email)) {
                errors.email = "Email tidak boleh kosong";
            }

            if (Validators.required(values.address)) {
                errors.address = "Alamat tidak boleh kosong";
            }

            if (!Validators.required(values.email) && !Validators.isEmailValid(values.email)) {
                errors.email = `Email "${values.email} tidak valid`;
            }

            if (Validators.required(values.no_telp?.toString())) {
                errors.no_telp = "No. Telepon tidak boleh kosong";
            } else {
                if (Validators.min(values.no_telp?.toString(), 11)) {
                    errors.no_telp = "No. Telepon minimal 11 digit";
                }
                if (Validators.max(values.no_telp?.toString(), 12)) {
                    errors.no_telp = "No. Telepon maksimal 12 digit";
                }
            }

            return errors;
        },
        onSubmit: async (values) => {
            loading = true;

            const data = {
                ...values,
                no_telp: String(values.no_telp),
            };

            return await instance.post("/auth/profile", data);
        },
        onError: (errors, context) => {
            loading = false;
            if (errors instanceof AxiosError) {
                if (errors.response.status === 422) {
                    const {
                        param,
                        error_message,
                    } = errors.response.data.data as IErrorMessage;

                    context.setErrors(param, (value) => error_message);
                }
            } else {
                const alertState: IAlert = {
                    type: "error",
                    text: errors.response.statusText,
                    show: true,
                };

                alertStore.set(alertState);
            }
        },
        onSuccess: (response: AxiosResponse) => {
            loading = false;

            const alertState: IAlert = {
                show: true,
                type: "success",
                text: "Data profile berhasil diupdate",
            };

            alertStore.set(alertState);
        },
    });

    async function getDataProfile() {
        try {
            const response: AxiosResponse<IProfileResponse> = await instance.get("/auth/profile");

            const {name, address, email, no_telp} = response.data.data;
            setData({
                name,
                address,
                email,
                no_telp,
            });
        } catch (error: AxiosError) {
            const alertState: IAlert = {
                type: "error",
                text: error.response.statusText,
                show: true,
            };

            alertStore.set(alertState);
        }
    }

    async function uploadPhotoUser(file: File) {
        loading = true;
        try {
            const response: AxiosResponse<IUploadResponse> = await instance.post("/auth/upload-photo", {
                photo: file,
            }, {
                headers: {
                    "Content-Type": "multipart/form-data",
                },
            });

            auth.update((state) => {
                return {...state, photo: response.data.data.file_name};
            });

        } catch (error: AxiosError) {
            const alertState: IAlert = {
                type: "error",
                text: error.response.statusText,
                show: true,
            };

            alertStore.set(alertState);
        } finally {
            loading = false;
        }
    }

    async function onUploadPhotoUser(event) {
        const image = event.target.files[0];
        const allowedExtensions = /(\.jpg|\.jpeg|\.png|\.gif)$/i;

        if (!allowedExtensions.exec(image.name)) {
            const alertState: IAlert = {
                type: "error",
                text: "Tipe file tidak valid untuk melakukan proses ini",
                show: true,
            };

            alertStore.set(alertState);

            return;
        }

        const fileSize = (image.size / 1024 / 1024).toFixed(2);

        if (parseFloat(fileSize) > 1) {
            const alertState: IAlert = {
                type: "error",
                text: "Ukuran file tidak valid untuk melakukan proses ini",
                show: true,
            };

            alertStore.set(alertState);

            return;
        }

        const reader = new FileReader();
        reader.readAsDataURL(image);
        reader.onload = e => avatar = e.target.result;

        await uploadPhotoUser(image);
    }

    function onTriggerFileInput() {
        uploadFile.click();
    }

    onMount(() => {
        async function load() {
            await getDataProfile();
        }

        load();
    });
</script>

<div class="mt-2.5 h-full w-full bg-white shadow-md lg:rounded-lg lg:shadow-lg lg:border lg:border-gray-200 dark:bg-gray-700">
    <div class="p-2.5">
        <h6 class="text-lg font-bold text-gray-500 dark:text-white">Profile</h6>
    </div>
    <hr class="border-t border-gray-300"/>
    <div class="flex flex-col flex-wrap p-2.5 lg:flex-row">
        <div class="flex justify-center lg:basis-1/2">
            <div class="relative">
                <div class="rounded-full border-2 border-gray-500">
                    {#if avatar}
                        <img alt="profile" src="{avatar}"
                             class="rounded-full h-28 w-28">
                    {:else }
                        <img alt="profile" src="{import.meta.env.VITE_API_URL_AVATAR}/{$auth.photo}"
                             class="rounded-full h-28 w-28">
                    {/if}
                </div>
                <input type="file" class="hidden" bind:this={uploadFile} on:change={onUploadPhotoUser}/>
                <button type="button"
                        class="absolute top-20 right-0 p-2 rounded-full bg-gray-300/60 focus:outline-none focus:ring-2 focus:ring-gray-500 dark:bg-gray-600/60"
                        on:click={onTriggerFileInput}
                >
                    <IconCamera className="h-5 w-5 text-gray-500 dark:text-white"/>
                </button>
            </div>
        </div>
        <div class="flex flex-col flex-1 my-4 lg:my-0">
            <form use:form>
                <div class="my-2.5">
                    <InputDefault label="Nama" name="nama" placeholder="Nama" type="text" value={$data.name}
                                  readonly={loading} isError={$errors.name && $touched.name}/>
                    {#if !!$errors.name}
                        <ErrorMessage>
                            {$errors.name}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="my-2.5">
                    <InputDefault label="Email" name="email" placeholder="Email" type="email" value={$data.email}
                                  readonly={loading} isError={$errors.email && $touched.email}/>
                    {#if !!$errors.email}
                        <ErrorMessage>
                            {$errors.email}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="my-2.5">
                    <InputDefault label="No. Telepon" name="no_telp" placeholder="No. Telepon" type="tel"
                                  pattern="[0-9]*"
                                  value={$data.no_telp} readonly={loading}
                                  isError={$errors.no_telp && $touched.no_telp}/>
                    {#if !!$errors.no_telp}
                        <ErrorMessage>
                            {$errors.no_telp}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="my-2.5">
                    <InputDefault label="Alamat" name="address" placeholder="Alamat" type="text" value={$data.address}
                                  readonly={loading} isError={$errors.address && $touched.address}/>
                    {#if !!$errors.address}
                        <ErrorMessage>
                            {$errors.address}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="mt-6 w-full">
                    <ButtonDefault type="submit" label="Save" fullWidth={true} disabled={loading}/>
                </div>
            </form>
        </div>
    </div>
</div>