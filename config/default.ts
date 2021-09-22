import { config } from 'dotenv';

config();

const getStringValue = (input: string | undefined) => {
  return input === undefined ? '' : input;
};

export const Config = {
  currencySymbol: getStringValue(process.env.CURRENCY_SYMBOL),
  network: getStringValue(process.env.NETWORK_NAME),
  networkId: getStringValue(process.env.NETWORK_ID),
  tokenAddressA: getStringValue(process.env.TOKEN_A_ADDRESS),
  tokenAddressB: getStringValue(process.env.TOKEN_B_ADDRESS),
  tokenNameA: getStringValue(process.env.TOKEN_A_NAME),
  tokenNameB: getStringValue(process.env.TOKEN_B_NAME),
  url: getStringValue(process.env.HTTP_URL)
};
