package org.neonlabs.e2e.kethereum.model
object Model {
  data class FaucetRequest(val wallet: String, val amount: Int)
}
