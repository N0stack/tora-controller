from enum import Enum


class MessageType(Enum):
    SPEC = 1
    NOTIFICATION = 2


class Message:
    """
    DO NOT REUSE INSTANCE.
    """

    def __init__(self, spec_id, type):
        # type: (str, Enum) -> (None)
        self.__spec_id = spec_id
        self.__type = type

    @property
    def type(self):
        # type: () -> Enum
        return self.__type

    @property
    def spec_id(self):
        # type: () -> str
        return self.__spec_id