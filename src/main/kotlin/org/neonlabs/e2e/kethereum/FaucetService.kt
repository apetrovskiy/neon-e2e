package org.neonlabs.e2e.kethereum

import org.neonlabs.e2e.kethereum.model.Model
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.POST

public interface FaucetService {
    @POST("/") fun requestFaucet(@Body faucetRequest: Model.FaucetRequest): Call<Model.FaucetRequest>
}
