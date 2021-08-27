from dotenv import load_dotenv
import os
from web3 import Web3

load_dotenv()

[print(f"{k}={v}") for k, v in sorted(os.environ.items())]

w3 = Web3(os.getenv("RPC_URL"))

[print(account) for account in w3.eth.accounts]
