echo "Import dataset to Arangoddb"

if [ $# -eq 3 ] 
then 
    user=$1
    passwd=$2
    database=$3
    echo "User: $user "
    echo "Passwd: ***** "
    echo "Database: $database" 
else 
    echo "Error: Usage sh importdn/ArangodbDataImport.sh <username> <passwd> <database>"
    exit 1
fi

# relation data
start="$(date -u +%s)"

arangoimp --file "data/venders.csv" \
          --type csv \
          --translate "VenderId=_key" \
          --collection "Vender" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --overwrite true 
arangoimp --file "data/customers.csv" \
          --type csv \
          --translate "CustomerId=_key" \
          --collection "Customer" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --overwrite true 
arangoimp --file "data/products.csv" \
          --type csv \
          --translate "ProductId=_key" \
          --collection "Product" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --overwrite true 
arangoimp --file "data/feedbacks.csv" \
          --type csv \
          --collection "Feedback" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --overwrite true 

end="$(date -u +%s)"
cost="$(($end-$start))"
echo "Data importing cost $cost sec"

