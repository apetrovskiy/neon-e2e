package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureBlocks;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;
import static org.neonlabs.e2e.web3j.Constants.StoryQueryBlock;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

/**
 * Blocks tests.
 */
@Epic(Epic)
@Feature(FeatureBlocks)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class BlockTest {
  @Test
  @Story(StoryQueryBlock)
  @Description("Block Header")
  @Disabled(NotYetDone)
  void test01() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  @Story(StoryQueryBlock)
  @Description("Full Block")
  @Disabled(NotYetDone)
  void test02() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
