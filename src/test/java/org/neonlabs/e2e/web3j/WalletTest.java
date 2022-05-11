package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.FeatureExternallyOwnedAccounts;
import static org.neonlabs.e2e.web3j.Constants.StoryWallet;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/** Externally-owned Accounts tests. */
@Epic(FeatureExternallyOwnedAccounts)
@Feature(StoryWallet)
@DisplayName(FeatureExternallyOwnedAccounts)
@Slf4j
@SuppressWarnings({"PMD.JUnitTestContainsTooManyAsserts", "PMD.DataflowAnomalyAnalysis"})
class WalletTest {
    @Test
    // @Story(StoryWallet)
    // @DisplayName(FeatureExternallyOwnedAccounts)
    @Description("Creating a new wallet")
    void shouldCreateWallet() {
        final var wallet = Wallet.create();

        assertAll(
                () -> assertNotNull(wallet, "Wallet is null"),
                () -> assertNotNull(wallet.getAddress(), "Wallet address is null"),
                () -> assertNotNull(wallet.getPrivateKey(), "Wallet private key is null"));
        log.info("Wallet has been created");
    }
}
