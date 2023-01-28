<script lang="ts">
    import {onMount} from "svelte";
    import {createForm} from "felte";
    import {AxiosError, type AxiosResponse} from "axios";
    import {alert} from "@/stores/alertStore";
    import type {IProfileRequest, IProfileResponse, IAlert} from "@/types";
    import {IconCamera, InputDefault, ButtonDefault, ErrorMessage} from "@/components";
    import {Validators} from "@/libs/validator";
    import instance from "@/libs/instance";

    let uploadFile;
    let loading: boolean;

    const {form, errors, touched, setData, data} = createForm<IProfileRequest>({
        initialValues: {
            nama: "",
            email: "",
            no_telp: "",
            alamat: "",
            jabatan: "",
        },
        validate: (values) => {
            const errors: IProfileRequest = {
                nama: "",
                no_telp: "",
                email: "",
                alamat: "",
                jabatan: "",
            };

            if (Validators.required(values.nama)) {
                errors.nama = "Nama tidak boleh kosong";
            }

            if (Validators.required(values.jabatan)) {
                errors.jabatan = "Jabatan tidak boleh kosong";
            }

            if (Validators.required(values.email)) {
                errors.email = "Email tidak boleh kosong";
            }

            if (Validators.required(values.alamat)) {
                errors.alamat = "Alamat tidak boleh kosong";
            }

            if (!Validators.required(values.email) && !Validators.isEmailValid(values.email)) {
                errors.email = `Email "${values.email} tidak valid`;
            }

            if (Validators.required(values.no_telp.toString())) {
                errors.no_telp = "No. Telepon tidak boleh kosong";
            } else {
                if (Validators.min(values.no_telp.toString(), 11)) {
                    errors.no_telp = "No. Telepon minimal 11 digit";
                }
                if (Validators.max(values.no_telp.toString(), 12)) {
                    errors.no_telp = "No. Telepon maksimal 12 digit";
                }
            }

            return errors;
        },
        onSubmit: async (values) => {
            loading = true;
            console.log(values);
        },
    });

    async function getDataProfile() {
        try {
            const response: AxiosResponse<IProfileResponse> = await instance.get("/auth/profile");

            const {nama, alamat, email, jabatan, no_telp} = response.data.data;
            setData({
                nama,
                alamat,
                jabatan,
                email,
                no_telp,
            });
        } catch (error: AxiosError) {
            const alertState: IAlert = {
                type: "error",
                text: error.response.statusText,
                show: true,
            };

            alert.set(alertState);
        }
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
                    <img alt="profile" src="/images/avatars/avatar.jpg" class="rounded-full h-28 w-28">
                </div>
                <input type="file" class="hidden" bind:this={uploadFile}/>
                <button type="button"
                        class="absolute top-20 right-0 p-2 rounded-full bg-gray-300/60 focus:outline-none focus:ring-2 focus:ring-gray-500 dark:bg-gray-600/60"
                >
                    <IconCamera className="h-5 w-5 text-gray-500 dark:text-white"/>
                </button>
            </div>
        </div>
        <div class="flex flex-col flex-1 my-4 lg:my-0">
            <form use:form>
                <div class="my-2.5">
                    <InputDefault label="Nama" name="nama" placeholder="Nama" type="text" value={$data.nama}
                                  readonly={loading} isError={$errors.nama && $touched.nama}/>
                    {#if !!$errors.nama}
                        <ErrorMessage>
                            {$errors.nama}
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
                    <InputDefault label="No. Telepon" name="no_telp" placeholder="No. Telepon" type="number"
                                  value={$data.no_telp} readonly={loading}
                                  isError={$errors.no_telp && $touched.no_telp}/>
                    {#if !!$errors.no_telp}
                        <ErrorMessage>
                            {$errors.no_telp}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="my-2.5">
                    <InputDefault label="Alamat" name="alamat" placeholder="Alamat" type="text" value={$data.alamat}
                                  readonly={loading} isError={$errors.alamat && $touched.alamat}/>
                    {#if !!$errors.alamat}
                        <ErrorMessage>
                            {$errors.alamat}
                        </ErrorMessage>
                    {/if}
                </div>
                <div class="my-2.5">
                    <InputDefault label="Jabatan" name="jabatan" placeholder="Jabatan" type="text" value={$data.jabatan}
                                  readonly={loading} isError={$errors.jabatan && $touched.jabatan}/>
                    {#if !!$errors.jabatan}
                        <ErrorMessage>
                            {$errors.jabatan}
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