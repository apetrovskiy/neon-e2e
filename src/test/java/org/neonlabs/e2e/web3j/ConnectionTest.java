package org.neonlabs.e2e.web3j;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.neonlabs.e2e.web3j.Constants.Epic;
import static org.neonlabs.e2e.web3j.Constants.FeatureCommon;

import io.qameta.allure.Description;
import io.qameta.allure.Epic;
import io.qameta.allure.Feature;
import io.qameta.allure.Story;
import org.junit.jupiter.api.Test;

@Epic(Epic)
@Feature(FeatureCommon)
class ConnectionTest {
  @Test
  @Description("allure description")
  @Story("allure story")
  void probe() {
    assertEquals(1, 1, "sample assertion");
  }

  @Test
  @Description("Connection web3j")
  @Story("Connection")
  void shouldConnectToNetwork() {
    // var config = new Config();
    var web3 = new Connection().createConnection();
    assertNotNull(web3, "Connecton to the network should not be null");
  }
}
