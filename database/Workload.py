from abc import abstractmethod, ABCMeta

class WorkLoad(metaclass=ABCMeta):
    # show given person's all information
    @abstractmethod
    def Q1(self, customerId):
        pass

    # find all person who bought given product by search orders
    @abstractmethod
    def Q2(self, productId):
        pass
    
    # find all person whoo bought given product and star score less than 5(means bad)
    @abstractmethod
    def Q3(self, productId):
        pass

    # find who spend the highest amount of money in orders.
    # Then foreach person, traverse her knows-graph with 3-hop to find the friends, 
    # and finally return the common friends of these two persons 
    @abstractmethod
    def Q4(self):
        pass


    # Given a start customer and a product category, find persons who are this 
    # customer’s friends within 3-hop friendships in knows-graph, and they have 
    # bought products in the given category. Finally, return feedback with the 5-rating 
    # review of those bought products
    @abstractmethod
    def Q5(self, customer1, customer2):
        pass

    # During given years find the popularity of given coutry vender 
    @abstractmethod
    def Q6(self):
        pass

    # find the most popularity vender of given country 
    @abstractmethod
    def Q7(self, country):
        pass

    