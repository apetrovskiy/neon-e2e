package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import java.util.concurrent.ExecutionException;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.methods.response.EthBlockNumber;

/** Blocks methods. */
@SuppressWarnings("RequireEmptyLineBeforeBlockTagGroupCheck")
public class Blocks {
    /** Retrieves tha latest block's number. */
    @Step
    public EthBlockNumber getBlockNumber(Web3j web3j)
            throws InterruptedException, ExecutionException {
        EthBlockNumber result = web3j.ethBlockNumber().sendAsync().get();
        return result;
    }
}
