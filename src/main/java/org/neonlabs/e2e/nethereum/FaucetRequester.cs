namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System.Net;
  using System.Net.Http;
  using System.Text;
  using System.Text.Json;
  using System.Threading.Tasks;
  using Allure.Xunit;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum.model;

  public class FaucetRequester
  {
    public static async Task<HttpStatusCode> RequestFaucetAsync(FaucetRequest requestDto)
    {
      return await Steps.Step("Requesting faucet", async () => {
        Config.Init();
        if (!Config.UseFaucet || string.IsNullOrEmpty(Config.FaucetUrl))
        {
          return HttpStatusCode.NotFound;
        }
        var client = new HttpClient();
        // client.BaseAddress = new Uri(Config.FaucetUrl);
        // client.DefaultRequestHeaders.Accept.Clear();
        // client.DefaultRequestHeaders.Accept.Add(
        //     new MediaTypeWithQualityHeaderValue("application/json"));
        var requestData = JsonSerializer.Serialize(requestDto);

        var requestContent = new StringContent(requestData, Encoding.UTF8, "application/json");

        var response = await client.PostAsync(Config.FaucetUrl, requestContent);
        response.EnsureSuccessStatusCode();

        // var content = await response.Content.ReadAsStringAsync();
        // var createdCompany = JsonSerializer.Deserialize<CompanyDto>(content, _options);

        // HttpResponseMessage response = await client.PostAsync(Config.FaucetUrl, requestDto as HttpContent);

        return response.StatusCode;
      });
    }
  }
}
