package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import java.util.concurrent.ExecutionException;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.methods.response.EthAccounts;

/**
 * Contains services for accounts.
 */
@SuppressWarnings("RequireEmptyLineBeforeBlockTagGroupCheck")
public class Accounts {
  /**
   * retrieves list of accounts.
   */
  @Step
  public EthAccounts getEthAccounts(Web3j web3j)
      throws InterruptedException, ExecutionException {
    EthAccounts result = web3j.ethAccounts()
        .sendAsync()
        .get();
    return result;
  }
}
