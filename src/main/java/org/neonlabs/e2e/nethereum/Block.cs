namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using System;
  using System.Threading.Tasks;
  using Nethereum.Hex.HexTypes;
  using Nethereum.Web3;

  public class Block
  {
    public static async Task<HexBigInteger> GetLatestBlockNumber()
    {
      var web3 = Connection.Connect();
      var latestBlockNumber = await web3.Eth.Blocks.GetBlockNumber.SendRequestAsync();
      Console.WriteLine($"Latest block number: {latestBlockNumber}");
      return latestBlockNumber;
    }
  }
}
