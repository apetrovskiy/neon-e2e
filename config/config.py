from dotenv import load_dotenv
import os

load_dotenv()

BASE_URL = os.environ.get('PROXY_URL').replace('/solana', ''),
PROXY_URL = os.environ.get('PROXY_URL')
NETWORK_ID = os.environ.get('NETWORK_ID')
NETWORK_NAME = os.environ.get('NETWORK_NAME')
REQUEST_AMOUNT = os.environ.get('REQUEST_AMOUNT')
FAUCET_URL = os.environ.get('FAUCET_URL')
USE_FAUCET = os.environ.get('USE_FAUCET')
ADDRESS_FROM = os.environ.get('ADDRESS_FROM')
ADDRESS_TO = os.environ.get('ADDRESS_TO')
PRIVATE_KEY = os.environ.get('PRIVATE_KEY')
SOLANA_EXPLORER = os.environ.get('SOLANA_EXPLORER')
SOLANA_URL = os.environ.get('SOLANA_URL')
USERS_NUMBER = os.environ.get('USERS_NUMBER')
