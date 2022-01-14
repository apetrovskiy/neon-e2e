package org.neonlabs.e2e.kethereum

import io.qameta.allure.Description
import io.qameta.allure.Epic
import io.qameta.allure.Feature
// import io.qameta.allure.Story
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Assertions.assertNotNull
import org.junit.jupiter.api.Disabled
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.Test

// import org.kethereum.rpc

@Epic("Common")
@Feature("Connection")
@DisplayName("Common")
class ConnectionKtTest {
  @Test
  @DisplayName("Connection Kethereum")
  @Description("Connection Kethereum")
  // @Story("Connection")
  @Disabled("not yet done")
  fun shouldWork() {
    assertEquals(1, ConnectionKt().getId(), "kotlin")
  }

  @Test
  @DisplayName("Connection Kethereum 2")
  @Description("Connection Kethereum 2")
  // @Story("Connection")
  @Disabled("not yet done")
  fun shouldConnectToNetwork() {
    assertNotNull(ConfigKt.proxyUrl)
    assertNotNull(ConfigKt.faucetUrl)
    // val connection = HttpProvider() // RPCTransport(ConfigKt.proxyUrl) //
    //     HttpTransport(ConfigKt.proxyUrl)
    // assertNotNull(connection)
    // println(connection)
    // connection.
  }
}
