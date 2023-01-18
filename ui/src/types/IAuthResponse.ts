import type {Response} from "./IResponse"
export interface IAuthResponse extends Response<{
    id_user: string
    kode_lokasi: string
    role: string
    isAuthenticated: boolean
}> {}