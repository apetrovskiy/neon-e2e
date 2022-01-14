package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
// import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureExternallyOwnedAccounts;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;
import static org.neonlabs.e2e.web3j.Constants.StoryTransfer;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
// import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/**
 * Externally-owned Accounts tests.
 */
@Epic(FeatureExternallyOwnedAccounts)
@Feature(StoryTransfer)
@DisplayName(FeatureExternallyOwnedAccounts)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class TransferTokenTest {
  @Test
  // @Story(StoryTransfer)
  @DisplayName(FeatureExternallyOwnedAccounts)
  @Description("Transfer tokens")
  @Disabled(NotYetDone)
  void test() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
