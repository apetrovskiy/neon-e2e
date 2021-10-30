// import the actual Api class
const { Api } = require('@parity/api');

// do the setup
describe('Light.js', () => {
  it('Connection test', () => {
    const provider = new Api.Provider.Http(process.env.PROXY_URL);
    const api = new Api(provider);
  });
});
