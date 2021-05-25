from pyArango import connection
from pyArango import database
from pprint import pprint


class WorkLoadArangodb:
    def __init__(self, db: database.DBHandle):
        self.db = db

    def Q1(self, num, show=True):
        AQL = 'RETURN DOCUMENT("Customer", "{:d}")'.format(num)
        result = db.AQLQuery(AQL, rawResults=True)
        if show:
            pprint(result[0])

    def Q2(self, num, show=True):
        AQL = 'let persons = (for o in Order for so in o.Suborders filter '
        AQL += 'so.Product.ProductId == {:d} return o.CustomerId) '.format(num)
        AQL += 'return {{ "product": {:d}, "person": Unique(persons) }}'.format(num)
        result = self.db.AQLQuery(AQL, rawResults=True)
        if show:
            pprint(result[0])
    
    

if __name__ == "__main__":
    conn = connection.Connection(username="dman", password="test")
    db = conn["mydb"]
    w = WorkLoadArangodb(db)
    w.Q2(9908)
