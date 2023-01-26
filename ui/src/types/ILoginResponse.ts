import type {Response} from "./IResponse";

export interface ILoginResponse extends Response<{
    nik: string
    kode_lokasi: string
    role: string
}> {
}