package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureExternallyOwnedAccounts;
import static org.neonlabs.e2e.web3j.Constants.StoryWallet;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;

/**
 * Externally Owned Accounts tests.
 */
@Epic(Epic)
@Feature(FeatureExternallyOwnedAccounts)
@Slf4j
@SuppressWarnings({ "PMD.JUnitTestContainsTooManyAsserts", "PMD.DataflowAnomalyAnalysis" })
class WalletTest {
  @Test
  @Story(StoryWallet)
  @Description("Creating a new wallet")
  void test() {
    final var wallet = Wallet.create();

    assertAll(
        () -> assertNotNull(wallet, "Wallet is null"),
        () -> assertNotNull(wallet.getAddress(), "Wallet address is null"),
        () -> assertNotNull(wallet.getPrivateKey(), "Wallet private key is null"));
    log.info("Wallet has been created");
  }
}
