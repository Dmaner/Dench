from abc import abstractmethod, ABCMeta

class WorkLoad(metaclass=ABCMeta):
    # show given person's all information
    @abstractmethod
    def Q1(self):
        pass

    # find all person who bought given product
    @abstractmethod
    def Q2(self):
        pass
    