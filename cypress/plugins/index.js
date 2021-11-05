/* eslint-disable @typescript-eslint/no-unsafe-assignment */
/// <reference types="cypress" />
// ***********************************************************
// This example plugins/index.js can be used to load plugins
//
// You can change the location of this file or turn off loading
// the plugins file with the 'pluginsFile' configuration option.
//
// You can read more here:
// https://on.cypress.io/plugins-guide
// ***********************************************************

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

// import { load } from 'cypress-browser-extension-plugin/loader';
const loadExtension = require('cypress-browser-extension-plugin').load;

/**
 * @type {Cypress.PluginConfig}
 */
// eslint-disable-next-line no-unused-vars
// export default (on) => {
module.exports = (on, config) => {
  // `on` is used to hook into various events Cypress emits
  // `config` is the resolved Cypress config
  on('before:browser:launch', (browser, launchOptons) => {
    // loadExtension("./cypress/extensions/metamask-10.2.2-an+fx.xpi"));
    launchOptons.extensions.push('./cypress/extensions/metamask-10.2.2-an+fx.xpi');
    loadExtension('./cypress/extensions/metamask-10.2.2-an+fx.xpi');

    return launchOptons;
  });
};

/*
on('before:browser:launch', loadExtension({
  source: 'cypress/extensions/metamask-10.2.2-an+fx.xpi',
  skipHooks: true,
}));
*/
