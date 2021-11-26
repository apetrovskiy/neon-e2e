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
    public int FaucetQuotient { get; set; }
    public string FaucetUrl { get; set; }
    public string AddressFrom { get; set; }
    public string AddressTo { get; set; }
    public bool DisableConfirmation { get; set; }
    public string PrivateKey { get; set; }
    public string SolanaExplorer { get; set; }
    public string SolanaUrl { get; set; }
    public int UsersNumber { get; set; }

    public Config()
    {
      string networkName;
      string proxyUrl;
      string networkId;
      string faucetQuotient;
      string faucetUrl;
      string addressFrom;
      string addressTo;
      string disableConfirmation;
      string privateKey;
      string solanaExplorer;
      string solanaUrl;
      string usersNumber;

      DotEnv.Load(options: new DotEnvOptions(envFilePaths: new[] { "../../../../../.env" }));
      EnvReader.TryGetStringValue("NETWORK_NAME", out networkName);
      EnvReader.TryGetStringValue("PROXY_URL", out proxyUrl);
      EnvReader.TryGetStringValue("NETWORK_ID", out networkId);
      EnvReader.TryGetStringValue("FAUCET_QUOTIENT", out faucetQuotient);
      EnvReader.TryGetStringValue("FAUCET_URL", out faucetUrl);
      EnvReader.TryGetStringValue("ADDRESS_FROM", out addressFrom);
      EnvReader.TryGetStringValue("ADDRESS_TO", out addressTo);
      EnvReader.TryGetStringValue("DISABLE_CONFIRMATION", out disableConfirmation);
      EnvReader.TryGetStringValue("PRIVATE_KEY", out privateKey);
      EnvReader.TryGetStringValue("SOLANA_EXPLORER", out solanaExplorer);
      EnvReader.TryGetStringValue("SOLANA_URL", out solanaUrl);
      Convert.ToInt32(EnvReader.TryGetStringValue("USERS_NUMBER", out usersNumber));

      NetworkName = networkId;
      ProxyUrl = proxyUrl;
      NetworkId = networkId;
      FaucetQuotient = Convert.ToInt32(faucetQuotient);
      FaucetUrl = faucetUrl;
      AddressFrom = addressFrom;
      AddressTo = addressTo;
      DisableConfirmation = Convert.ToBoolean(disableConfirmation);
      PrivateKey = privateKey;
      SolanaExplorer = solanaExplorer;
      SolanaUrl = solanaUrl;
      UsersNumber = Convert.ToInt32(usersNumber);
    }
  }
}
