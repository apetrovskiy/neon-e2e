namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System.Threading.Tasks;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum.model;
  using Nethereum.Hex.HexConvertors.Extensions;
  using Nethereum.Web3.Accounts;
  using static Nethereum.Signer.EthECKey;

  public class AccountFactory
  {
    public static async Task<Account> CreateAccount(int amount)
    {
      // TODO: improve it
      Config.Init();
      var ecKey = GenerateKey();
      var privateKey = ecKey.GetPrivateKeyAsBytes().ToHex();
      var account = new Account(privateKey);
      if (amount > 0)
      {
        await FaucetRequester.RequestFaucetAsync(new FaucetRequest
        {
          Wallet = account.Address,
          Amount = amount
        });
      }
      return account;
    }
  }
}
