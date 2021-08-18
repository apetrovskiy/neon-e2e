export class AsyncWithPromise {
    delayedPromise(): Promise<string> {
        return new Promise<string>((resolve: (str: string) => void, reject: (str: string) => void) => {
            setTimeout(() => {
                console.log(`2. rturning success`);
                resolve("success");
            }, 1000);
        })
    }
}