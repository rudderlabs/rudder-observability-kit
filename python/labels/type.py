from typing import Any

class Label():
    def __init__(self, name: str, value: Any):
        self._name = name
        self._value = value

    @property
    def name(self) -> str:
        return self._name

    @property
    def value(self) -> Any:
        return self._value 
    
    @staticmethod
    def name(name: str):
        return lambda value: Label(name, value)
    
