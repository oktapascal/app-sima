import type {Response} from "./IResponse";

export interface IUploadResponse extends Response<{
    file_name: string | null
}> {
}