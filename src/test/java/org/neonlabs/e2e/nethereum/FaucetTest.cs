namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum.model;
  using Xunit;
  using static NeonEndToEnd.org.neonlabs.e2e.nethereum.Constants;

  [AllureSuite(Suite)]
  [AllureEpic(Epic)]
  [AllureFeature(new string[] { FeatureFaucet })]
  public class FaucetTest
  {
    [AllureSubSuite(StoryFaucet)]
    [AllureStory(new string[] { StoryFaucet })]
    [AllureXunit(DisplayName = "Single request to faucet")]
    public async void ShouldProvideTokensRequested()
    {
      int amountRequested = 3, expectedBalance = 3;
      Console.WriteLine("first param = " + amountRequested);
      Console.WriteLine("second param = " + expectedBalance);
      var account = await AccountFactory.CreateAccount(amountRequested);
      var balance = await Balance.GetBalanceInEther(account);
      Assert.Equal(expectedBalance, balance);
    }

    [AllureSubSuite(StoryFaucet)]
    [AllureStory(new string[] { StoryFaucet })]
    [AllureXunit(DisplayName = "Consequent requests to faucet")]
    public async void ShouldProvideTokensRequestedInTwoRequests()
    {
      int firstAmountRequested = 5, secondAmountRequested = 3, expectedBalance = 8;
      Console.WriteLine("first param = " + firstAmountRequested);
      Console.WriteLine("second param = " + expectedBalance);
      var account = await AccountFactory.CreateAccount(0);
      await FaucetRequester.RequestFaucetAsync(new FaucetRequest
      {
        Wallet = account.Address,
        Amount = firstAmountRequested
      });
      await FaucetRequester.RequestFaucetAsync(new FaucetRequest
      {
        Wallet = account.Address,
        Amount = secondAmountRequested
      });
      var balance = await Balance.GetBalanceInEther(account);
      Assert.Equal(expectedBalance, balance);
    }
  }
}


