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
  public class BalanceTest
  {
    [AllureStory(new string[] { StoryBalance })]
    [AllureXunit(DisplayName = "GetBalance Async")]
    public async void ShouldGetBalanceAsync()
    {
      var balance = await Balance.GetBalance();
      Assert.NotNull(balance);
    }
    [AllureStory(new string[] { StoryBalance })]
    [AllureXunit(DisplayName = "Initial balance")]
    public void ShouldGetInitialBalanceAsync()
    {
      var account = AccountFactory.CreateAccount();
      // TODO: logging
      Console.WriteLine(account.ChainId);
      // Console.WriteLine(account.TransactionManager.Client.GetType().Name);
      // Assert.NotNull(balance);
    }
    [AllureStory(new string[] { StoryQueryBlock })]
    // [AllureXunit(DisplayName = "Get the latest block balance")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetLatestBlockBalance()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { StoryQueryBlock })]
    // [AllureXunit(DisplayName = "Get a specific block balance")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetSpecifigBlockBalance()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
    [AllureStory(new string[] { StoryQueryBlock })]
    // [AllureXunit(DisplayName = "Get pending balance")]
    [Fact(Skip = "not yet done")]
    public async void ShouldGetPendingBalance()
    {
      var latestBlockNumber = await Block.GetLatestBlockNumber();
      Assert.NotNull(latestBlockNumber);
    }
  }
}
