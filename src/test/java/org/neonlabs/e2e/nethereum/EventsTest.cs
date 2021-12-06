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
  [AllureFeature(new string[] { FeatureEvents })]
  public class EventsTest
  {
    [AllureStory("Subscription to new blocks")]
    [AllureXunit(DisplayName = "Subscription to new blocks")]
    [Fact(Skip = "not yet done")]
    public async void ShouldSubscribeToNewBlocks()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory("Subscription to event logs")]
    [AllureXunit(DisplayName = "Subscription to event logs")]
    [Fact(Skip = "not yet done")]
    public async void ShouldSubscribeToEventLogs()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory("Read event logs")]
    [AllureXunit(DisplayName = "Read event logs")]
    [Fact(Skip = "not yet done")]
    public async void ShouldReadEventLogs()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory("Read ERC20 token event logs")]
    [AllureXunit(DisplayName = "Read ERC20 token event logs")]
    [Fact(Skip = "not yet done")]
    public async void ShouldReadERC20TokenEventLogs()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory("Read 0x protocol event logs")]
    [AllureXunit(DisplayName = "Read 0x protocol event logs")]
    [Fact(Skip = "not yet done")]
    public async void ShouldRead0xProtocalEventLogs()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
