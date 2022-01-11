namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using System.Threading.Tasks;
  using Allure.Xunit;
  using Nethereum.Hex.HexTypes;
  using Nethereum.Web3;
  using Nethereum.Web3.Accounts;

  public class Balance
  {
    public static async Task<HexBigInteger> GetBalance(Account account)
    {
      return await Steps.Step("Getting balance", async () =>
      {
        var web3 = Connection.Connect();
        var balance = await web3.Eth.GetBalance.SendRequestAsync(account.Address);
        // TODO: logging
        Console.WriteLine($"Balance in Wei: {balance.Value}");

        var etherAmount = Web3.Convert.FromWei(balance.Value);
        // TODO: logging
        Console.WriteLine($"Balance in Ether: {etherAmount}");

        return balance;
      });
    }

    public static async Task<decimal> GetBalanceInEther(Account account)
    {
      return await Steps.Step("Getting balance", async () =>
      {
        var web3 = Connection.Connect();
        var balance = await web3.Eth.GetBalance.SendRequestAsync(account.Address);
        // TODO: logging
        Console.WriteLine($"Balance in Wei: {balance.Value}");

        var etherAmount = Web3.Convert.FromWei(balance.Value);
        return etherAmount;
      });
    }
  }
}
