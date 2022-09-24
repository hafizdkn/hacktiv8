## Assigment 2
### List Endpoint 
- Create Order
  - http://localhost:8080/order/
  -  Data
      ```json
      {
      "orderedAt": "2019-11-09T21:21:46+00:00",
      "customerName": "Tom Jerry",
      "items": [
        {
          "itemCode": "123",
          "description": "IPhone 10X",
          "quantity": 1
        }
        ]
      } 
      ```
  ![create_order](https://user-images.githubusercontent.com/40784871/192100533-e225de60-587c-488e-b2dc-f16028238b84.png)

- Get Order
  - http://localhost:8080/orders
  
  ![get_orders](https://user-images.githubusercontent.com/40784871/192100651-53111dc0-9b78-400f-a7b3-cc8f5827d04c.png)
