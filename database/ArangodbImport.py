from pyArango import connection
from pyArango import database

def Q1(db: database.DBHandle):
    AQL = 'RETURN DOCUMENT("Customer", "10")'
    result = db.AQLQuery(AQL)
    print(result)


if __name__ == "__main__":
    conn = connection.Connection(username="dman", password="test")
    db = conn["mydb"]
    Q1(db)
