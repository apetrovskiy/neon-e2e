package org.neonlabs.e2e.web3j;

import io.qameta.allure.Step;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.Sign;
import org.web3j.utils.Numeric;

// @UtilityClass
/** Contains Ethereum-related utils. */
public class EthereumUtils {
    /** Signs linking code. */
    @Step
    public static String signLinkingCode(String code, String privateKey) {
        var credentials = Credentials.create(privateKey);
        byte[] data = code.getBytes();
        Sign.SignatureData signature = Sign.signPrefixedMessage(data, credentials.getEcKeyPair());
        return Numeric.toHexString(signature.getR())
                + Numeric.toHexString(signature.getS()).substring(2)
                + Numeric.toHexString(signature.getV()).substring(2);
    }
}
