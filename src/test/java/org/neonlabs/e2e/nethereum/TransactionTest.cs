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
  [AllureFeature(new string[] { FeatureTransaction })]
  public class TransactionTest
  {
    [AllureStory(new string[] { "Transaction" })]
    // [AllureXunit(DisplayName = "TBD")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetLatestBlockNumber()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
