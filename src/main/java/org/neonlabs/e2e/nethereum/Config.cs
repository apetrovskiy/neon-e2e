namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using dotenv.net;
  using dotenv.net.Utilities;

  public class Config
  {
    public string NetworkName { get; set; }
    public string ProxyUrl { get; set; }
    public string NetworkId { get; set; }
    public string CurrencySymbol { get; set; }
    public int FaucetQuotient { get; set; }
    public string FaucetUrl { get; set; }
    public string AddressFrom { get; set; }
    public string AddressTo { get; set; }

    public bool DisableConfirmation { get; set; }
    public string PrivateKey { get; set; }

    public Config()
    {
      DotEnv.Load(options: new DotEnvOptions(envFilePaths: new[] { "../../../../../.env" }));
      NetworkName = dotenv.net.Utilities.EnvReader.GetStringValue("NETWORK_NAME");
      ProxyUrl = EnvReader.GetStringValue("PROXY_URL");
      NetworkId = EnvReader.GetStringValue("NETWORK_ID");
      CurrencySymbol = EnvReader.GetStringValue("CURRENCY_SYMBOL");
      FaucetQuotient = Convert.ToInt32(EnvReader.GetStringValue("FAUCET_QUOTIENT"));
      FaucetUrl = EnvReader.GetStringValue("FAUCET_URL");
      AddressFrom = EnvReader.GetStringValue("ADDRESS_FROM");
      AddressTo = EnvReader.GetStringValue("ADDRESS_TO");
      DisableConfirmation = Convert.ToBoolean(EnvReader.GetStringValue("DISABLE_CONFIRMATION"));
      PrivateKey = EnvReader.GetStringValue("PPRIVATE_KEY");
    }
  }
}
