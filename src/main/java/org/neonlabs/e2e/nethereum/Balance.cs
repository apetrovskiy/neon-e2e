namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using System.Threading.Tasks;
  using Nethereum.Hex.HexTypes;
  using Nethereum.Web3;

  public class Balance
  {
    public static async Task<HexBigInteger> GetBalance()
    {
      var web3 = Connection.Connect();
      var account = AccountFactory.CreateAccount();
      var balance = await web3.Eth.GetBalance.SendRequestAsync(account.Address);
      // TODO: logging
      Console.WriteLine($"Balance in Wei: {balance.Value}");

      var etherAmount = Web3.Convert.FromWei(balance.Value);
      // TODO: logging
      Console.WriteLine($"Balance in Ether: {etherAmount}");

      return balance;
    }
  }
}
