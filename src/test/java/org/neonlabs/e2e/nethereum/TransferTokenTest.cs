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
  [AllureFeature(new string[] { FeatureExternallyOwnedAccounts })]
  public class TransferTokenTest
  {
    [AllureStory(new string[] { StoryTransfer })]
    // [AllureXunit(DisplayName = "TBD")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetLatestBlockNumber()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
