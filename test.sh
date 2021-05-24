#!/bin/sh

start_time="$(date -u +%s)"

## Relational

arangoimp --file "./Dataset/Customer/person_0_0.csv" --type csv --translate "id=_key" --separator "|" --collection "Customer" --server.username root  --server.password ""  --create-collection true 
arangoimp --file "./Dataset/Feedback/feedback.csv" --type csv  --separator "|" --collection "Feedback" --server.username root  --server.password ""    --create-collection true 
arangoimp --file "./Dataset/Vendor/vendor.csv" --type csv  --translate "id=_key"  --collection "Vendor" --server.username root  --server.password ""    --create-collection true
arangoimp --file "./Dataset/Product/Product.csv" --type csv --translate "productId=_key" --collection "Product" --server.username root  --server.password ""    --create-collection true 
end_time1="$(date -u +%s)"
elapsed1="$(($end_time1-$start_time))"
echo "The importing time for Relation data is $elapsed1" >> arangoDBImportingTime

## JSON
arangoimp --file "./Dataset/Order/order.json" --type jsonl --translate "id=_key" --collection "Order" --server.username root  --server.password ""    --create-collection true
end_time2="$(date -u +%s)"
elapsed2="$(($end_time2-$end_time1))"
echo "The importing time for JSON data is $elapsed2" >> arangoDBImportingTime


## Graph

arangoimp --file "./Dataset/SocialNetwork/post_0_0.csv" --type csv --separator "|" --translate "id=_key" --collection "Post"  --server.username root  --server.password "" --create-collection true 
arangoimp --file "./Dataset/SocialNetwork/person_knows_person_0_0.csv" --type csv --separator "|" --translate "from=_from" --translate "to=_to" --collection "KnowsGraph" --from-collection-prefix Customer --to-collection-prefix Customer --server.username root  --server.password "" --create-collection true --create-collection-type edge 
arangoimp --file "./Dataset/SocialNetwork/post_hasCreator_person_0_0.csv" --type csv --separator "|" --translate "PersonId=_from" --translate "PostId=_to" --collection "PersonHasPost" --from-collection-prefix Customer --to-collection-prefix Post --server.username root  --server.password "" --create-collection true --create-collection-type edge 
arangoimp --file "./Dataset/SocialNetwork/post_hasTag_tag_0_0.csv" --type csv --separator "|" --translate "PostId=_from" --translate "TagId=_to" --collection "PostHasTag" --from-collection-prefix Post --to-collection-prefix Product --server.username root  --server.password "" --create-collection true --create-collection-type edge 

end_time3="$(date -u +%s)"
elapsed3="$(($end_time3-$end_time2))"
echo "The importing time for Graph data is $elapsed3" >> arangoDBImportingTime