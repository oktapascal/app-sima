import type {Response} from "./IResponse";

export interface IProfileResponse extends Response<{
    address: string | null
    email: string | null
    id_location: string | null
    name: string | null
    photo: string | null
    nik: string | null
    no_telp: string | null
    username: string | null
    role: string | null
}> {
}