export class MockAsync {
    executeSlowFunction(complete: (value: string) => void) {
        setTimeout(() => {
            complete(`completed`);
        }, 1000);
    }
}