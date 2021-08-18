import { GlobalCounter } from "../../src/samples/GlobalCounter";

describe("test setup and teardown", () => {
    let globalCounter: GlobalCounter;
    beforeAll(() => {
        globalCounter = new GlobalCounter();
    })

    beforeEach(() => { globalCounter.count = 0; });

    afterEach(() => { console.log(`globalCounter.count = ${globalCounter.count}`) });

    it("should increment", () => {
        globalCounter.increment();
        expect(globalCounter.count).toEqual(1);
    });

    it("should increment twice", () => {
        globalCounter.increment();
        globalCounter.increment();
        expect(globalCounter.count).toEqual(2);
    });
});