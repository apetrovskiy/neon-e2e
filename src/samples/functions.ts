export function hasValueNoWhiteSpace(value: string): boolean {
    if (value && value.length > 0 && value.trim().length > 0) {
        return true;
    }
    return false;
}

export function testUsing<T>(values: T[], func: Function) {
    for (let value of values) {
        func.apply(Object, [value]);
    }
}