package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureTransaction;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

/**
 * Transaction tests.
 */
@Epic(Epic)
@Feature(FeatureTransaction)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class TransactionTest {
  @Test
  @Story("Create raw transaction")
  @Description("Create raw transaction")
  @Disabled(NotYetDone)
  void test01() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story("Send raw transaction")
  @Description("Send raw transaction")
  @Disabled(NotYetDone)
  void test02() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
