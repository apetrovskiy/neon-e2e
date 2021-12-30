from requests import post
from config.config import FAUCET_URL, USE_FAUCET
from src.helpers.faucet.faucet_request import FaucetRequest


async def request_faucet(wallet: str, amount: int):
    if not USE_FAUCET or FAUCET_URL == '':
        print("Skipping faucet request")
        return
    data: FaucetRequest = FaucetRequest(amount=amount, wallet=wallet)
    print(FAUCET_URL)
    print(data)
    try:
        response = post(FAUCET_URL, data)
    except Exception as e:
        print(f"Failed to obtain token from faucet {FAUCET_URL}")
        print(e)
    assert response.status_code == 200
