import pyorient

client = pyorient.OrientDB("localhost", 2480) 
session_id = client.connect( "admin", "admin" )
print(session_id)