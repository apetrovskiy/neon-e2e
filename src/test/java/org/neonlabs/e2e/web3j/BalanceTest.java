package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureExternallyOwnedAccounts;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;
import static org.neonlabs.e2e.web3j.Constants.StoryBalance;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

/**
 * Balance tests.
 */
@Epic(Epic)
@Feature(FeatureExternallyOwnedAccounts)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class BalanceTest {
  @Test
  @Story(StoryBalance)
  @Description("Get the latest block balance")
  @Disabled(NotYetDone)
  void test01() {
    final var web3 = new Connection().createConnection();
    assertNotNull(web3, "Connection should not be null");

    /*
     * client, err := connect()
     * assert.Nil(t, err, fmt.Sprintf(FaileToConnectTo, GetConfig().ProxyURL, err))
     * if err != nil {
     * t.Errorf(FaileToConnectTo, GetConfig().ProxyURL, err)
     * }
     * 
     * account := createWallet()
     * assert.NotEqual(t, 0, len(account.Address.Hash()), FailedToCreateWallet)
     * if len(account.Address) == 0 {
     * t.Error(FailedToCreateWallet)
     * }
     * 
     * balance := getLastBlockBalance(client, account.Address.Hex())
     * log.Println(balance)
     * assert.Equal(t, GetConfig().InitialBalance, balance, InitialBalanceIsWrong)
     */
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story(StoryBalance)
  @Description("Get a specific block balance")
  @Disabled(NotYetDone)
  void test02() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story(StoryBalance)
  @Description("Get pending balance")
  @Disabled(NotYetDone)
  void test03() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
