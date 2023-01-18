import type {Response} from "./IResponse"

export interface IUserResponse extends Response<{
    nama: string
    foto: string
}>{}