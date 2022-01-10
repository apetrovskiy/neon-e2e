namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using Nethereum.RPC.Eth;
  using Xunit;
  using static NeonEndToEnd.org.neonlabs.e2e.nethereum.Constants;

  [AllureSuite(Suite)]
  [AllureEpic(Epic)]
  [AllureFeature(new string[] { FeatureBlocks })]
  public class BlockNumberTest
  {
    [AllureSubSuite(StoryQueryBlock)]
    [AllureStory(new string[] { StoryQueryBlock })]
    [AllureXunit(DisplayName = "Latest block number test")]
    public async void ShouldGetLatestBlockNumber()
    {
      await Steps.Step("Getting the latest block number", async () => { 
        var latestBlockNumber = await Block.GetLatestBlockNumber();
        Assert.NotNull(latestBlockNumber);
      });
    }
    [AllureSubSuite(StoryQueryBlock)]
    [AllureStory(new string[] { StoryQueryBlock })]
    // [AllureXunit(DisplayName = "Block header")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetBlockHeader()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureSubSuite(StoryQueryBlock)]
    [AllureStory(new string[] { StoryQueryBlock })]
    // [AllureXunit(DisplayName = "Full block")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetFullBlock()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
