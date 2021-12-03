namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using dotenv.net.Utilities;
  using Nethereum.Web3;

  public class Connection
  {
    public static Web3 Connect()
    {
      // new Config();
      // TODO: improve it
      Config.Init();

      // TODO: logging
      Console.WriteLine("===================");
      // Console.WriteLine(EnvReader.GetStringValue("PROXY_URL"));
      Console.WriteLine(Config.ProxyUrl);

      // var web3 = new Web3(EnvReader.GetStringValue("PROXY_URL"));
      var web3 = new Web3(Config.ProxyUrl);
      // TODO: logging
      Console.WriteLine(web3);
      return web3;
    }
  }
}
