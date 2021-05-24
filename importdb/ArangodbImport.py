from pyArango import connection

conn = connection.Connection(username="dman", password="test")
db = conn["mydb"]
print(db)