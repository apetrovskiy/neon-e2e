package org.neonlabs.e2e.kethereum

import org.neonlabs.e2e.kethereum.model.Model
import retrofit2.Retrofit

fun requestFaucet(wallet: String, amount: Int) {
    val faucetService = buildRequester().create(FaucetService::class.java)
    faucetService.requestFaucet(Model.FaucetRequest(wallet, amount))
}

fun buildRequester(): Retrofit {
    val retrofit =
        Retrofit.Builder()
            .baseUrl(ConfigKt.faucetUrl)
            // .addConverterFactory(GsonConverterFactory.create(gson))
            .build()
    return retrofit
}
