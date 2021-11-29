import requests
from config.config import FAUCET_URL
from src.helpers.faucet.faucet_request import FaucetRequest


async def request_faucet(wallet: str, amount: int):
    data: FaucetRequest = FaucetRequest(amount=amount, wallet=wallet)
    print(FAUCET_URL)
    print(data)
    requests.post(FAUCET_URL, data)
