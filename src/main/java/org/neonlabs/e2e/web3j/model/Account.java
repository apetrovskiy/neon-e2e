package org.neonlabs.e2e.web3j.model;

import java.math.BigInteger;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

/** Is a model for accounts. */
@AllArgsConstructor
@Builder
@Data
@NoArgsConstructor
public class Account {
    private String address;
    private String privateKey;
    private BigInteger privateKeyDec;
}
