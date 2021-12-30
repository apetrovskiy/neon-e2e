package org.neonlabs.e2e.kethereum

import io.github.cdimascio.dotenv.Dotenv

class ConfigKt {
  val config = Dotenv.configure().systemProperties().load()
  val proxyUrl = config["PROXY_URL"]
  val faucetUrl = config["FAUCET_URL"]
  val networkName = config["NETWORK_NAME"]
  val networkId = config["NETWORK_ID"]
  val addressFrom = config["ADDRESS_FROM"]
  val addressTo = config["ADDRESS_TO"]
  val privateKey = config["PRIVATE_KEY"]
  val requestAmount = config["REQUEST_AMOUNT"]
}
