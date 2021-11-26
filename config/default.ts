import { config } from 'dotenv';

config();

const getStringValue = (input: string | undefined) => {
  return input === undefined ? '' : input;
};

export const Config = {
  baseUrl: getStringValue(process.env.PROXY_URL).replace('/solana', ''),
  faucetQuotient: Number.parseInt(getStringValue(process.env.FAUCET_QUOTIENT)),
  faucetUrl: getStringValue(process.env.FAUCET_URL),
  network: getStringValue(process.env.NETWORK_NAME),
  networkId: getStringValue(process.env.NETWORK_ID),
  url: getStringValue(process.env.PROXY_URL)
};
