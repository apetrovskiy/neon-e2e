package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
// import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureBlocks;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;
import static org.neonlabs.e2e.web3j.Constants.StoryQueryBlock;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
// import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/**
 * Blocks tests.
 */
@Epic(FeatureBlocks)
@Feature(StoryQueryBlock)
@DisplayName(FeatureBlocks)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class BlockTest {
  @Test
  // @Story(StoryQueryBlock)
  @Description("Block Header")
  @Disabled(NotYetDone)
  void test01() {

    /*
     * client, err := connect()
     * assert.Nil(t, err, fmt.Sprintf(FaileToConnectTo, GetConfig().ProxyURL, err))
     * header, err := client.HeaderByNumber(context.Background(), nil)
     * assert.Nil(t, err, "Failed to get block header")
     * if err != nil {
     * log.Fatal(err)
     * }
     * assert.Greater(t, 0, header.Number, "Block header number greater than 0")
     */

    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }

  @Test
  // @Story(StoryQueryBlock)
  @DisplayName(FeatureBlocks)
  @Description("Full Block")
  @Disabled(NotYetDone)
  void test02() {
    assertAll(
        () -> assertNotNull(1, ""),
        () -> assertEquals(1, 1, ""));
  }
}
