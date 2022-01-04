namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using dotenv.net;
  using dotenv.net.Utilities;

  public static class Config
  {
    public static string NetworkName { get; set; }
    public static string ProxyUrl { get; set; }
    public static string NetworkId { get; set; }
    public static int RequestAmount { get; set; }
    public static string FaucetUrl { get; set; }
    public static bool UseFaucet { get; set; }
    public static string AddressFrom { get; set; }
    public static string AddressTo { get; set; }
    public static string PrivateKey { get; set; }
    public static string SolanaExplorer { get; set; }
    public static string SolanaUrl { get; set; }
    public static int UsersNumber { get; set; }
    private static bool isInitialized;

    public static void Init()
    {
      if (isInitialized) return;
      string networkName;
      string proxyUrl;
      string networkId;
      string requestAmount;
      string faucetUrl;
      bool useFaucet;
      string addressFrom;
      string addressTo;
      string privateKey;
      string solanaExplorer;
      string solanaUrl;
      string usersNumber;

      DotEnv.Load(options: new DotEnvOptions(envFilePaths: new[] { "../../../../../.env" }));
      EnvReader.TryGetStringValue("NETWORK_NAME", out networkName);
      EnvReader.TryGetStringValue("PROXY_URL", out proxyUrl);
      EnvReader.TryGetStringValue("NETWORK_ID", out networkId);
      EnvReader.TryGetStringValue("REQUEST_AMOUNT", out requestAmount);
      EnvReader.TryGetStringValue("FAUCET_URL", out faucetUrl);
      EnvReader.TryGetBooleanValue("USE_FAUCET", out useFaucet);
      EnvReader.TryGetStringValue("ADDRESS_FROM", out addressFrom);
      EnvReader.TryGetStringValue("ADDRESS_TO", out addressTo);
      EnvReader.TryGetStringValue("PRIVATE_KEY", out privateKey);
      EnvReader.TryGetStringValue("SOLANA_EXPLORER", out solanaExplorer);
      EnvReader.TryGetStringValue("SOLANA_URL", out solanaUrl);
      Convert.ToInt32(EnvReader.TryGetStringValue("USERS_NUMBER", out usersNumber));

      NetworkName = networkId;
      ProxyUrl = proxyUrl;
      NetworkId = networkId;
      RequestAmount = Convert.ToInt32(requestAmount);
      FaucetUrl = faucetUrl;
      UseFaucet = useFaucet;
      AddressFrom = addressFrom;
      AddressTo = addressTo;
      PrivateKey = privateKey;
      SolanaExplorer = solanaExplorer;
      SolanaUrl = solanaUrl;
      UsersNumber = Convert.ToInt32(usersNumber);
    }
  }
}
