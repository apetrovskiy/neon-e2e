package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertAll;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.FeatureExternallyOwnedAccounts;
import static org.neonlabs.e2e.web3j.Constants.NotYetDone;
import static org.neonlabs.e2e.web3j.Constants.StoryBalance;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import lombok.SneakyThrows;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

/** Balance tests. */
@Epic(FeatureExternallyOwnedAccounts)
@Feature(StoryBalance)
@DisplayName(FeatureExternallyOwnedAccounts)
@SuppressWarnings("PMD.JUnitTestContainsTooManyAsserts")
class BalanceTest {
    @Test
    // @Story(StoryBalance)
    // @DisplayName(FeatureExternallyOwnedAccounts)
    @Description("Get the latest block balance")
    @SneakyThrows
    void shouldGetLatestBlockBalance() {
        final var web3 = new Connection().createConnection();
        assertNotNull(web3, "Connection should not be null");
        final var wallet = Wallet.create();
        assertNotNull(wallet, "Wallet is null");
        final var balance = new Balance().getEthBalance(web3, wallet.getAddress());

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
         * assert.Equal(t, GetConfig().RequestAmount, balance, InitialBalanceIsWrong)
         */
        assertEquals(
                new Config().getRequestAmount(),
                balance,
                "Balance is not equal to the initial balance");
    }

    @Test
    // @Story(StoryBalance)
    // @DisplayName(FeatureExternallyOwnedAccounts)
    @Description("Get a specific block balance")
    void test02() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }

    @Test
    // @Story(StoryBalance)
    // @DisplayName(FeatureExternallyOwnedAccounts)
    @Description("Get pending balance")
    @Disabled(NotYetDone)
    void test03() {
        assertAll(() -> assertNotNull(1, ""), () -> assertEquals(1, 1, ""));
    }
}

/*
 * func TestGetLatestBalance(t *testing.T) {
 *
 * allure.Test(t,
 * allure.Epic(Epic),
 * allure.Feature(FeatureExternallyOwnedAccounts),
 * allure.Story(StoryBalance),
 * allure.Description("Get the latest block balance"),
 * allure.Action(func() {
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
 * assert.Equal(t, GetConfig().RequestAmount, balance, InitialBalanceIsWrong)
 * }))
 * }
 */
