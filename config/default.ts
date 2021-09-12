import { config } from 'dotenv';

config();

const getStringValue = (input: string | undefined) => {
  return input === undefined ? '' : input;
};

export const Config = {
  currencySymbol: getStringValue(process.env.CURRENCY_SYMBOL),
  network: getStringValue(process.env.NETWORK_NAME),
  networkId: getStringValue(process.env.NETWORK_ID),
  url: getStringValue(process.env.HTTP_URL)
};
