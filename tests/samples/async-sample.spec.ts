import { AsyncWithPromise } from "../../src/samples/async-with-promise";
import { MockAsync } from "../../src/samples/MockAsync";

describe.skip("failing asyng tests", () => {
    it("should wqit for callback to complete", () => {
        let mockAsync = new MockAsync();
        console.log(`1. calling executeSlowFunction`);
        let returnedValue !: string;
        mockAsync.executeSlowFunction((value: string) => {
            console.log(`2. complate called`);
            returnedValue = value;
        });
        console.log(`3. checking return value`);
        expect(returnedValue).toBe("completed");
    });
});

describe.skip("async test with done", () => {
    let returnedValue!: string;

    beforeEach((done: jest.DoneCallback) => {
        let mockAsync = new MockAsync();
        console.log(`1. calling executeSlowFunction`);
        mockAsync.executeSlowFunction((value: string) => {
            console.log(`2. executeSlowFunction returned`);
            returnedValue = value;
            done();
        });
    });

    it("should return value after 1 second", () => {
        console.log(`3. checking returnedValue`);
        expect(returnedValue).toBe("completed");
    });
});

describe.skip("async test", () => {
    it("should wait 1 second for promise to resolve", async () => {
        let asyncWithPromise = new AsyncWithPromise();
        console.log(`1. calling delayedPromise`);
        let returnValue = await asyncWithPromise.delayedPromise();
        console.log(`3. after await`);
        expect(returnValue).toEqual("success");
    });
});