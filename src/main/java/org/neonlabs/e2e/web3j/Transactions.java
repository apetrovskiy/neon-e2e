package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import java.util.concurrent.ExecutionException;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.methods.response.EthGetBalance;

/** Transaction methods. */
@SuppressWarnings("RequireEmptyLineBeforeBlockTagGroupCheck")
public class Transactions {
    /** Retrieves transaction number by account address. */
    @Step
    public EthGetBalance getEthBalance(Web3j web3j, String address)
            throws InterruptedException, ExecutionException {
        EthGetBalance result = new EthGetBalance();
        web3j.ethGetBalance(address, DefaultBlockParameter.valueOf("latest")).sendAsync().get();
        return result;
    }
}
