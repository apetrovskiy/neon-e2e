import { hasValueNoWhiteSpace, testUsing } from "../../src/samples/functions";

[1, 2, 3, 4, 5].forEach((value: number) => {
    it.skip(`${value} should be less than 5`, () => {
        expect(value).toBeLessThan(5);
    });
});

testUsing(
    [
        [undefined, false],
        [null, false],
        [" ", false],
        ["  ", false],
        [" a ", true],
    ], ([value, isValid]: [string, boolean]) => {
        it(`${value} hasValueNoWhiteSpace ? ${isValid}`, () => {
            isValid ? expect(hasValueNoWhiteSpace(value)).toBeTruthy() :
                expect(hasValueNoWhiteSpace(value)).toBeFalsy();
        })
    }
);