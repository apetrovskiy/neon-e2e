package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.http.HttpService;

/**
 * Returns a connection to blockchain.
 */
public class Connection {
  /**
   * Creates a connection to blockchain.
   */
  @Step
  public Web3j createConnection() {
    var config = new Config();
    var web3 = Web3j.build(new HttpService(config.getProxyUrl()));
    return web3;
  }
}
