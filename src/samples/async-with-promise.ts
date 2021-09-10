export class AsyncWithPromise {
    delayedPromise(): Promise<string> {
        return new Promise<string>((resolve: (str: string) => void) => {
            setTimeout(() => {
                console.log(`2. returning success`);
                resolve("success");
            }, 1000);
        })
    }
}