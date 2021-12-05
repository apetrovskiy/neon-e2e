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
