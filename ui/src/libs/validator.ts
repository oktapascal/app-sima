function required(value: any): boolean {
    return value === "" || value == null;
}

function isEmailValid(value: any): boolean {
    return (/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/).test(value);
}

function min(value: any, standard: number): boolean {
    if (typeof value === "string") {
        return standard > value.length;
    }

    if (typeof value === "number") {
        return standard > value;
    }
}

function max(value: any, standard: number): boolean {
    if (typeof value === "string") {
        return standard < value.length;
    }

    if (typeof value === "number") {
        return standard < value;
    }
}

export const Validators = {
    required,
    isEmailValid,
    min,
    max,
};