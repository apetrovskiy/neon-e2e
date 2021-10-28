namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using Nethereum.RPC.Eth;
  using Xunit;

  [AllureSuite("Nethereum")]
  public class BlockNumberTest
  {
    [AllureXunit(DisplayName = "Latest block number test")]
    public async void ShouldGetLatestBlockNumber()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
