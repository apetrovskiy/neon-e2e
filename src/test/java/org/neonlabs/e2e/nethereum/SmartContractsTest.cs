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
  [AllureFeature(new string[] { FeatureSmartContracts })]
  public class SmartContractsTest
  {
    [AllureStory(new string[] { "Deploy smart contracts" })]
    // [AllureXunit(DisplayName = "Deploy smart contracts")]
    [Fact(Skip = "not yet done")]
    public async void ShouldDeploySmartContracts()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { "Load smart contracts" })]
    // [AllureXunit(DisplayName = "Load smart contracts")]
    [Fact(Skip = "not yet done")]
    public async void ShouldLoadSmartContracts()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { "Query smart contracts" })]
    // [AllureXunit(DisplayName = "Query smart contracts")]
    [Fact(Skip = "not yet done")]
    public async void ShouldQuerySmartContracts()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { "Write to smart contracts" })]
    // [AllureXunit(DisplayName = "Write to smart contracts")]
    [Fact(Skip = "not yet done")]
    public async void ShouldWriteToSmartContracts()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { "Read smart contract bytecode" })]
    // [AllureXunit(DisplayName = "Read smart contract bytecode")]
    [Fact(Skip = "not yet done")]
    public async void ShouldReadSmartContractBytecode()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { "Query smart contracts" })]
    // [AllureXunit(DisplayName = "Query ERC20 token smart contracts")]
    [Fact(Skip = "not yet done")]
    public async void ShouldQuestERC20TokenSmartContracts()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
