// import the actual Api class
const Api = require('@parity/api');
const makeContract = require('@parity/light.js');
// do the setup
describe('Light.js', () => {
  it('Connection test', () => {
    const provider = new Api.Provider.Http(process.env.PROXY_URL);
    const api = new Api(provider);
  });

  it('Contract', () => {
    // TODO: use a real ABI
    const ABI = '';
    makeContract(ABI, '0x00..ff').balanceOf$('0xabc');
    // .subscribe(console.log); // Logs the result when balance changes
  });
});
