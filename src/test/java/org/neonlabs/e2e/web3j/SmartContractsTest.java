package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.FeatureSmartContracts;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/** Smart Contracts tests. */
@Epic(FeatureSmartContracts)
// @Feature(FeatureSmartContracts)
@DisplayName(FeatureSmartContracts)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class SmartContractsTest {
    @Test
    @Story("Deploy smart contracts")
    // @DisplayName(FeatureSmartContracts)
    @Description("Deploy smart contracts")
    @Disabled(NotYetDone)
    void test01() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    @Story("Load smart contracts")
    // @DisplayName(FeatureSmartContracts)
    @Description("Load smart contracts")
    @Disabled(NotYetDone)
    void test02() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    @Story("Query smart contracts")
    // @DisplayName(FeatureSmartContracts)
    @Description("Query smart contracts")
    @Disabled(NotYetDone)
    void test03() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    @Story("Write to smart contracts")
    // @DisplayName(FeatureSmartContracts)
    @Description("Write to smart contracts")
    @Disabled(NotYetDone)
    void test04() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    @Story("Read smart contract bytecode")
    @Description("Read smart contract bytecode")
    @Disabled(NotYetDone)
    void test05() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    @Story("Query smart contracts")
    // @DisplayName(FeatureSmartContracts)
    @Description("Query ERC20 token smart contracts")
    @Disabled(NotYetDone)
    void test06() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }
}
