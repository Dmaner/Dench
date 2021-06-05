import pymongo

myclient = pymongo.MongoClient("mongodb://localhost:27017/")
dbs = myclient.list_database_names()
print(dbs)