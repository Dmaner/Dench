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

# Json
arangoimp --file "data/orders.json" \
          --type json \
          --translate "OrderId=_key" \
          --collection "Order" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --overwrite true 

#Graph
arangoimp --file "data/cinps.csv" \
          --type csv \
          --translate "CustomerId=_from" \
          --from-collection-prefix Customer \
          --translate "ProductId=_to" \
          --to-collection-prefix Product \
          --collection "CustomerInterestGraph" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --create-collection-type edge \
          --overwrite true 

arangoimp --file "data/pknowps.csv" \
          --type csv \
          --translate "CustomerFromId=_from" \
          --from-collection-prefix Customer \
          --translate "CustomerToId=_to" \
          --to-collection-prefix Customer \
          --collection "KnowsGraph" \
          --server.username $user \
          --server.password $passwd \
          --server.database $database \
          --create-collection true \
          --create-collection-type edge \
          --overwrite true 

end="$(date -u +%s)"
cost="$(($end-$start))"

echo "Data importing cost $cost sec"

