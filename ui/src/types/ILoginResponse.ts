import type {Response} from "./IResponse";

export interface ILoginResponse extends Response<{
    nik: string
    id_location: string
    role: string
    photo: string | null
}> {
}