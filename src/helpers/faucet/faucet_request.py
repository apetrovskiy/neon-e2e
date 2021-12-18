from dataclasses import dataclass


@dataclass
class FaucetRequest:
    amount: int
    wallet: str
