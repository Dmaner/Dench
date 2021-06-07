from pyArango import connection
from pyArango import database
from pprint import pprint
from functools import wraps
from time import time
import pyorient


def timerec(funcname: str):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kw):
            start = time()
            res = func(*args, **kw)
            end = time()
            print("{} cost {:.5f} sec".format(funcname, end - start))
            return res
        return wrapper
    return decorator


class WorkLoadOrientDB:
    def __init__(self, db: database.DBHandle):
        self.db = db

    def AQLShow(self, AQL:str, show):
        result = db.AQLQuery(AQL, rawResults=True)
        if show and len(result) != 0:
            pprint(result[0])    
    
    @timerec("OrientDB Q1")
    def Q1(self, num, show=True):
        AQL = 'RETURN DOCUMENT("Customer", "{:d}")'.format(num)
        self.AQLShow(AQL, show)

    @timerec("OrientDB Q2")
    def Q2(self, num, show=True):
        AQL = 'let persons = (for o in Order for so in o.Suborders filter '
        AQL += 'so.Product.ProductId == {:d} return o.CustomerId) '.format(num)
        AQL += 'return {{ "product": {:d}, "person": Unique(persons) }}'.format(
            num)
        self.AQLShow(AQL, show)

    @timerec("OrientDB Q3")
    def Q3(self, num, show=True):
        AQL = 'let persons = (for f in Feedback '
        AQL += 'filter f.ProductId == {:d} filter f.Star < 5 return '.format(
            num)
        AQL += '{"person": f.CustomerId, "feedback": f.Comment}) '
        AQL += 'return {{ "productId": {:d}, "persons": Unique(persons) }} '.format(
            num)
        self.AQLShow(AQL, show)

    @timerec("OrientDB Q4")
    def Q4(self, show=True):
        AQL = 'let persons = (for o in Order sort o.Cost desc limit 2'
        AQL += 'return {Customer : o.CustomerId, Money:o.Cost}) '
        AQL += 'let set1 = (for vertex in 1..3 outbound concat("Customer/", persons[1].Customer) KnowsGraph return vertex) '
        AQL += 'let set2 = (for vertex in 1..3 outbound concat("Customer/", persons[1].Customer) KnowsGraph return vertex) '
        AQL += 'return count(intersection(set1, set2))'
        self.AQLShow(AQL, show)

    @timerec("OrientDB Q5")
    def Q5(self, num, vender, show=True):
        AQL = 'let persons = (for friend in 1..1 outbound "Customer/{:d}" KnowsGraph '.format(
            num)
        AQL += 'for o in Order filter o.CustomerId == +friend._key and {:d} '.format(vender)
        AQL += 'in o.Suborders[*].Product.Vender.VenderId return friend) return persons'
        self.AQLShow(AQL, show)
        
    @timerec("OrientDB Q6")
    def Q6(self, a, b, show=True):
        AQL = 'let shortpath = (for vertex, edge in outbound shortest_path "Customer/{:d}" to "Customer/{:d}" '.format(a, b)
        AQL += 'KnowsGraph return vertex) '
        AQL += 'let plist = flatten( for item in shortpath for o in Order filter +item._key == o.CustomerId '
        AQL += 'return o.Suborders) '
        AQL += 'For item in plist collect vender=item.Product.Vender.VenderId with count into cnt Sort cnt desc limit 5 '
        AQL += "Return {vender,cnt}"
        self.AQLShow(AQL, show)

    @timerec("OrientDB Q7")
    def Q7(self, country, start, end, show=True):
        AQL = 'let brands = (for b in Vender filter b.Country == "China" return +b._key) '.format(country)
        AQL += 'let orderlines=flatten(For o in Order for so in o.Suborders Filter '
        AQL += 'so.CreationDate > "{}" and so.CreationDate < "{}" and '.format(start, end)
        AQL += 'so.Product.Vender.VenderId in brands return so) '
        AQL += 'let popularity=Count(Unique(orderlines)) '
        AQL += 'return popularity'
        self.AQLShow(AQL, show)


if __name__ == "__main__":
    conn = connection.Connection(username="dman", password="test")
    db = conn["mydb"]
    w = WorkLoadOrientDB(db)
    show = False
    w.Q1(10, show)
    w.Q2(9908, show)
    w.Q3(9908, show)
    w.Q4(show)
    w.Q5(10, 290, show)
    w.Q6(1, 10, show)
    w.Q7('China', '1900', '2020', show)
