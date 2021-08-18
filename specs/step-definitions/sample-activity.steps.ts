import { defineFeature, loadFeature } from "jest-cucumber";
import { AccountCreator } from "../../src/helpers/create_account";
const feature = loadFeature("specs/features/sample-activity.feature");
defineFeature(feature, (test) => {
    let user01: object | null;
    let user02: object | null;

    test("a new account's activity", ({ given, when, then }) => {
        beforeAll(() => {
            console.log("BEFORE ALL");
        });
        given("there is a new user in Ethereum", () => {
            user01 = new AccountCreator().create();
            console.log(user01);
        });
        given("there is another new user in Ethereum", () => {
            user02 = new AccountCreator().create();
            console.log(user02);
        });
        when("user sends Eth", () => {
            console.log('when');
        });
        then("the recipient has balance increased", () => {
            console.log('then 1');
        });
        then("the sender has balance decreased", () => {
            console.log('then 2');
        });
    });
});