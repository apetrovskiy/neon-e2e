namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit;
  using dotenv.net.Utilities;
  using Nethereum.Web3;

  public class Connection
  {
    public static Web3 Connect()
    {
      return Steps.Step("Loading config", () => {
        Config.Init();

        // TODO: logging
        Console.WriteLine("===================");
        Console.WriteLine(Config.ProxyUrl);

        var web3 = new Web3(Config.ProxyUrl);
        // TODO: logging
        Console.WriteLine(web3);
        return web3;
      });
    }
  }
}
