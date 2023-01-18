export interface IAlert {
    text: string;
    type: "success" | "error" | "warning" | "info";
    show: boolean;
}