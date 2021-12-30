namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using Nethereum.Hex.HexConvertors.Extensions;
  using Nethereum.Web3.Accounts;
  using static Nethereum.Signer.EthECKey;

  public class AccountFactory
  {
    public static Account CreateAccount()
    {
      // TODO: improve it
      Config.Init();
      var ecKey = GenerateKey();
      var privateKey = ecKey.GetPrivateKeyAsBytes().ToHex();
      return new Account(privateKey);
    }
  }
}
