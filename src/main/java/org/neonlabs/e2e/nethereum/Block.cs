namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using System.Threading.Tasks;
  using Allure.Xunit;
  using Nethereum.Hex.HexTypes;
  using Nethereum.Web3;

  public class Block
  {
    public static async Task<HexBigInteger> GetLatestBlockNumber()
    {
      return await Steps.Step("Getting the latest block number", async () => {
      var web3 = Connection.Connect();
        var latestBlockNumber = await web3.Eth.Blocks.GetBlockNumber.SendRequestAsync();
        // TODO: logging
        Console.WriteLine($"Latest block number: {latestBlockNumber}");
        return latestBlockNumber;
      });
    }
  }
}
