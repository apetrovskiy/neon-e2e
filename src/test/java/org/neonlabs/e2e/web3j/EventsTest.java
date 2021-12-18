package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureEvents;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

/**
 * Events tests.
 */
@Epic(Epic)
@Feature(FeatureEvents)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class EventsTest {
  @Test
  @Story("Subscription to new blocks")
  @Description("Subscription to new blocks")
  @Disabled(NotYetDone)
  void test01() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story("Subscription to event logs")
  @Description("Subscription to event logs")
  @Disabled(NotYetDone)
  void test02() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story("Read event logs")
  @Description("Read event logs")
  @Disabled(NotYetDone)
  void test03() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story("Read ERC20 token event logs")
  @Description("Read ERC20 token event logs")
  @Disabled(NotYetDone)
  void test04() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story("Read 0x protocol event logs")
  @Description("Read 0x protocol event logs")
  @Disabled(NotYetDone)
  void test05() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
