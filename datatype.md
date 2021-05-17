# 数据类型

声明数据生成类型

## 目录

- [数据类型](#数据类型)
  - [目录](#目录)
  - [关系型数据](#关系型数据)
    - [Customer](#customer)
    - [Vender](#vender)
  - [KV型数据](#kv型数据)
    - [Feedback](#feedback)
  - [XML数据](#xml数据)
    - [Invoice](#invoice)
  - [Json类型数据](#json类型数据)
    - [Order](#order)
    - [Product](#product)
  - [图数据](#图数据)
    - [PersonInterestProduct](#personinterestproduct)
    - [PersonKnowPerson](#personknowperson)

## 关系型数据

### Customer

```
id, 
firstName, 
lastName, 
gender(性别),
birthday, 
creationDate, 
locationIP, 
browserUsed(使用浏览器), 
place
```

### Vender

```
id,
Country,
Industry
```

## KV型数据

### Feedback

```
asin(亚马逊商品id),
PersonId,
feedback(评价)
```

## XML数据

### Invoice 

```xml
<OrderId>
<PersonId>
<OrderDate>
<TotalPrice>
<Orderline>
  <productId>
  <asin>
  <title>
  <price>
  <brand>
</Orderline>
```
## Json类型数据

### Order

```json
{
    "OrderId": "016f6a4a-ec18-4885-b1c7-9bf2306c76d6",
    "PersonId": "10995116278711",
    "OrderDate": "2022-09-01",
    "TotalPrice": 723.88,
    "Orderline": [
        {
            "productId": "6465",
            "asin": "B000FIE4WC",
            "title": "Topeak Dual Touch Bike Storage Stand",
            "price": 199.95,
            "brand": "MYLAPS_Sports_Timing"
        },
    ]
}
```

### Product 

```
asin,
title,
price,
imgUrl,
productId,
brand
```

## 图数据

```
Customer -- PersonKnowPerson -- > Customer
```

### PersonInterestProduct

```
Person.id,
Product.id
IntersetValue
```

### PersonKnowPerson

```
from,
to,
creationDate
```