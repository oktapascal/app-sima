function required(value: any): boolean {
    return value === "" || value == null;
}

export const Validators = {
    required,
};