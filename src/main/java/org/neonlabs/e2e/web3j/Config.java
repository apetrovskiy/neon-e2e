package org.neonlabs.e2e.web3j;

import io.github.cdimascio.dotenv.Dotenv;

/**
 * Loads configuration.
 */
public class Config {

  public Config() {
    Dotenv.configure().systemProperties().load();
  }

  public String getNetworkName() {
    return System.getProperty("NETWORK_NAME");
  }

  public String getProxyUrl() {
    return System.getProperty("PROXY_URL");
  }

  public String getNetworkId() {
    return System.getProperty("NETWORK_ID");
  }

  public String getAddressFrom() {
    return System.getProperty("ADDRESS_FROM");
  }

  public String getAddressTo() {
    return System.getProperty("ADDRESS_TO");
  }

  public String getPrivateKey() {
    return System.getProperty("PRIVATE_KEY");
  }

  public String getFaucetQuotient() {
    return System.getProperty("FAUCET_QUOTIENT");
  }

  public String getFaucetUrl() {
    return System.getProperty("FAUCET_URL");
  }

  public Boolean getUseFaucet() {
    return Boolean.parseBoolean(System.getProperty("USE_FAUCET"));
  }

  public Boolean getSolanaExplorer() {
    return Boolean.parseBoolean(System.getProperty("SOLANA_EXPLORER"));
  }

  public Boolean getSolanaUrl() {
    return Boolean.parseBoolean(System.getProperty("SOLANA_URL"));
  }

  public Integer getUsersNumber() {
    return Integer.parseInt(System.getProperty("USERS_NUMBER"));
  }

  public Boolean getFaucetAmount() {
    return Boolean.parseBoolean(System.getProperty("FAUCET_AMOUNT"));
  }
}
