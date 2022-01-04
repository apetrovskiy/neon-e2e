using System.Net.Http;
using System.Text.Json.Serialization;

namespace NeonEndToEnd.org.neonlabs.e2e.nethereum.model
{
  public class FaucetRequest // : HttpContent
  {
    [JsonPropertyName("amount")]
    public int Amount { get; set; }
    [JsonPropertyName("wallet")]
    public string Wallet { get; set; }
  }
}
