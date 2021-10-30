// load_extension.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test

const loadExtension = require('cypress-browser-extension-plugin').load;
describe('Test 0001', () => {
  it('yeees', () => {
    // Commands.clearExtensionStorage;
    loadExtension('./cypress/extensions/metamask-10.2.2-an+fx.xpi');
    cy.visit('http://www.yandex.ru');
  });
});
