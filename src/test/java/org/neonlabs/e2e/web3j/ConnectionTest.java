package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertNotNull;
// import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureCommon;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
// import io.qameta.allure.Story;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

@Epic(FeatureCommon)
@Feature("Connection")
@DisplayName(FeatureCommon)
class ConnectionTest {
  @Test
  // @Story("Connection")
  @DisplayName(FeatureCommon)
  @Description("Connection web3j")
  void shouldConnectToNetwork() {
    final var web3 = new Connection().createConnection();
    assertNotNull(web3, "Connecton to the network should not be null");
  }
}
