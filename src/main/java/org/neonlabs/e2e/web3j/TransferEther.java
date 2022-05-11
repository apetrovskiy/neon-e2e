package org.neonlabs.e2e.web3j;

import java.math.BigDecimal;
import lombok.SneakyThrows;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Transfer;
import org.web3j.utils.Convert;

/** Contains transfer operations. */
public class TransferEther {
    /** Transfers Ether between accounts. */
    @SneakyThrows
    public static TransactionReceipt send() {
        final var web3 = new Connection().createConnection();
        Credentials credentials = WalletUtils.loadCredentials("password", "/path/to/walletfile");
        TransactionReceipt transactionReceipt =
                Transfer.sendFunds(
                                web3,
                                credentials,
                                "0x<address>|<ensName>",
                                BigDecimal.valueOf(1.0),
                                Convert.Unit.ETHER)
                        .send();
        return transactionReceipt;
    }
}
