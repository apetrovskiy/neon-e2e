import httpx
from config.config import FAUCET_URL, USE_FAUCET


async def request_faucet(wallet: str, amount: int):
    if not USE_FAUCET or FAUCET_URL == '':
        print("Skipping faucet request")
        return
    data = {'amount': int(amount), 'wallet': wallet}
    try:
        async with httpx.AsyncClient() as client:
            response = await client.post(FAUCET_URL, json=data, timeout=20)
            assert response.status_code == 200, "The response status is not OK"
    except Exception as e:
        print(e)
        print(f"Failed to obtain tokens from faucet {FAUCET_URL}")
