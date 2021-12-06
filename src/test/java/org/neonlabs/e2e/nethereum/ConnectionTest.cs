namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using Nethereum.RPC.Eth;
  using Xunit;
  using static NeonEndToEnd.org.neonlabs.e2e.nethereum.Constants;

  [AllureSuite(Suite)]
  [AllureEpic(Epic)]
  [AllureFeature(new string[] { FeatureCommon })]
  public class ConnectionTest
  {
    [AllureStory(Story)]
    [AllureXunit(DisplayName = "Connection test")]
    public void ShouldConnect()
    {
      var web3 = Connection.Connect();
      Assert.NotNull(web3);
    }
[AllureStory(Story)]
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

  }
}
