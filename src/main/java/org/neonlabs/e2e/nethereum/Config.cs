namespace NeonEndToEnd.org.neonlabs.e2e.nethereum
{
  using dotenv.net;

  public class Config
  {
    public Config()
    {
      DotEnv.Load(options: new DotEnvOptions(envFilePaths: new[] { "../../../../../.env" }));
    }
  }
}
