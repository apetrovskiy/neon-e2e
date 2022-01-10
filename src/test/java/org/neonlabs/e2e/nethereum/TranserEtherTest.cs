namespace NeonEndToEnd.Tests.org.neonlabs.e2e.nethereum
{
  using System;
  using Allure.Xunit.Attributes;
  using NeonEndToEnd.org.neonlabs.e2e.nethereum;
  using Nethereum.RPC.Eth;
  using Nethereum.Web3;
  using Nethereum.Web3.Accounts;
  using Nethereum.Web3.Accounts.Managed;
  using Nethereum.Hex.HexTypes;
  using Xunit;
  using static NeonEndToEnd.org.neonlabs.e2e.nethereum.Constants;

  [AllureSuite(Suite)]
  [AllureEpic(Epic)]
  [AllureFeature(new string[] { FeatureExternallyOwnedAccounts })]
  public class TransferEtherTest
  {
    [AllureStory(new string[] { StoryTransfer })]
    [AllureXunit(DisplayName = "Transfer with the default gas price")]
    [AllureDescription("Sending Ether using the EtherTransferService with the default gas price and gas amount")]
    public async void ShouldTransferWithDefaultGasPrice()
    {
      // TODO: write code
      // var latestBlockNumber = await Block.GetLatestBlockNumber();
      // Assert.NotNull(latestBlockNumber);

      // var transaction = await web3.Eth.GetEtherTransferService().TransferEtherAndWaitForReceiptAsync(toAddress, 1.11m);
    }
  }
}
