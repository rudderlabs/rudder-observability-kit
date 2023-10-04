from typing import Self
from ..labels.type import Label
class Logger:
    def __init__(self, *labels: Label) -> None:
        self.labels: list[Label] = list(labels)
    def get_labels(self) -> list[Label]:
        return self.labels
    def create_logger(self, *labels: Label) -> Self:
        allLabels = self.get_labels() + list(labels)
        return Logger(*allLabels)
    
    # Child classes should override these methods
    def debug(self, msg: str, *labels: Label):
        pass
    def info(self, msg: str, *labels: Label):
        pass
    def warn(self, msg: str, *labels: Label):
        pass
    def error(self, msg: str, *labels: Label):
        pass
    def fatal(self, msg: str, *labels: Label):
        pass
