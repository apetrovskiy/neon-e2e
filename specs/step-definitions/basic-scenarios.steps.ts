import { loadFeature, defineFeature } from 'jest-cucumber';
// import { PasswordValidator } from '../../src/password-validator';

const feature = loadFeature('specs/features/basic-scenarios.feature');

defineFeature(feature, (test) => {
    // let passwordValidator = new PasswordValidator();
    let accessGranted = false;

    beforeEach(() => {
        // passwordValidator = new PasswordValidator();
        console.log('BEFORE EACH');
    });

    test.skip('Entering a correct password', ({ given, when, then }) => {
        given('I have previously created a password', () => {
            // passwordValidator.setPassword('1234');
            console.log('GIVEN');
        });

        when('I enter my password correctly', () => {
            // accessGranted = passwordValidator.validatePassword('1234');
            console.log('WHEN');
        });

        then('I should be granted access', () => {
            console.log('THEN');
            // expect(accessGranted).toBe(true);
        });
    });
});
