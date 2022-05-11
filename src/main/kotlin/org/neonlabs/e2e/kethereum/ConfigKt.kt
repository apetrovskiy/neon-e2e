package org.neonlabs.e2e.kethereum

import io.github.cdimascio.dotenv.Dotenv

class ConfigKt {
    companion object {
        val config = Dotenv.configure().systemProperties().load()
        val proxyUrl = config["PROXY_URL"]
        val networkName = config["NETWORK_NAME"]
        val networkId = config["NETWORK_ID"]
        val addressFrom = config["ADDRESS_FROM"]
        val addressTo = config["ADDRESS_TO"]
        val privateKey = config["PRIVATE_KEY"]
        val faucetUrl = config["FAUCET_URL"]
        val useFaucet = if ("TRUE" == config["USE_FAUCET"].toString().toUpperCase()) true else false
        val requestAmount = config["REQUEST_AMOUNT"]
        val solanaExplorer = config["SOLANA_EXPLORER"]
        val solanaUrl = config["SOLANA_URL"]
        val usersNumber = config["USERS_NUMBER"]?.toInt()
    }
}
