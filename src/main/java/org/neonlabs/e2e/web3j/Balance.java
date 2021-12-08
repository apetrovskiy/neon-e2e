package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import java.util.concurrent.ExecutionException;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.methods.response.EthGetBalance;

/**
 * Balance information.
 */
@SuppressWarnings("RequireEmptyLineBeforeBlockTagGroupCheck")
public class Balance {
  /**
   * retrieves balance by account address.
   */
  @Step
  public EthGetBalance getEthBalance(Web3j web3j, String address)
      throws InterruptedException, ExecutionException {
    EthGetBalance result = new EthGetBalance();
    web3j.ethGetBalance(address,
        DefaultBlockParameter.valueOf("latest"))
        .sendAsync()
        .get();
    return result;
    /*
     * EthGetBalance ethGetBalance = web3
     * .ethGetBalance(“0xcF8B652b0173FBABE734f5F388C2da24a2359993”,
     * DefaultBlockParameterName.LATEST)
     * .sendAsync().get();
     */
  }
}
