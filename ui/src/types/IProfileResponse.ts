import type {Response} from "./IResponse";

export interface IProfileResponse extends Response<{
    alamat: string | null
    email: string | null
    kode_lokasi: string | null
    nama: string | null
    foto: string | null
    nik: string | null
    no_telp: string | null
    username: string | null
    role: string | null
}> {
}