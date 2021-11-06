namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using Nethereum.RPC.Eth;
  using Xunit;

  [AllureSuite("Nethereum")]
  [AllureEpic("Nethereum")]
  public class ConnectionTest
  {
    [AllureXunit(DisplayName = "Connection test")]
    public void ShouldConnect()
    {
      var web3 = Connection.Connect();
      Assert.NotNull(web3);
    }

    [AllureXunit(DisplayName = "Subprojects are filled in")]
    public void ShouldHaveData()
    {
      var web3 = Connection.Connect();
      Assert.NotNull(web3);
      // Assert.All<IEthAccounts>(web3.Eth.Accounts.,w=>
      // {
      //     Assert.NotNull(w);
      //     // Assert.Collection(web3.Eth.Accounts);
      // });
      // web3.Eth.Accounts
      // web3.Eth.Accounts
      // web
      // Assert.All(web3, w=>w);
    }

    [AllureXunit(DisplayName = "GetBalance Async")]
    public async void ShouldGetBalanceAsync()
    {
      var balance = await Balance.GetBalance();
      Assert.NotNull(balance);
    }

    [AllureXunit(DisplayName = "Initial balance")]
    public void ShouldGetInitialBalanceAsync()
    {
      var account = AccountFactory.CreateAccount();
      Console.WriteLine(account.ChainId);
      // Console.WriteLine(account.TransactionManager.Client.GetType().Name);
      // Assert.NotNull(balance);
    }
  }
}
