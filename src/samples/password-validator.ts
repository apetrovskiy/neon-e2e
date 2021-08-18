export class PasswordValidator {
    private password: string = "";
    setPassword(password: string): void {
        this.password = password;
    }
    validatePassword(password: string): boolean {
        return this.password === password;
    }
}