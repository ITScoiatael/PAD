# PAD 

EShop project

### To run

```
```


## Backend

Installation and launch:
```bash
$ cd backend
$ go get github.com/99designs/gqlgen
$ go get github.com/google/uuid
$ go run ./server.go
```

GraphQL playground:
http://localhost:8080/

### IDs
The database is based on the use of UUID.

UUID is Universally Unique Identifiers, or UUIDS, are 128 bit numbers, composed of 16 octets and represented as 32 base-16 characters, that can be used to identify information across a computer system. 

Example:
> 123e4567-e89b-12d3-a456-426614174000

### Queries
- Categories - Returns all categories
- Category (ID) - Returns the category by ID
- Customers - Returns all customers
- Customer (ID) - Returns the customer by ID
- Products - Returns all products
- Product (ID) - Returns the product by ID
- SubProducts (productID) - Returns all subproducts by product ID
- SubProduct (ID) - Returns the subproduct by ID
- Orders (customerID) - Returns all orders by customer ID
- Order (ID) - Returns the order by ID
- OrderedSubProducts (orderID) - Returns all products in the order by its ID
- OrderedSubProduct (ID) - Returns the order product by ID
- Admins - Returns all admins
- Admin (ID) - Returns the admin by ID

### Mutations
- CreateCategory (name, imageURL) - Creates a category
- DeleteCategory (ID) - Deletes a category by ID
- AddProduct (categoryID) - Adds a product to a category by its ID
- DeleteProduct (ID) - Deletes a product by ID
- CreateSubProduct (productID, price, size, color, amount) - Creates a subproduct into a product by its ID
- DeleteSubProduct (ID) - Deletes a subproduct by ID
- CreateCustomer (name, email, phone, address, region, ccNumber) - Creates a customer
- DeleteCustomer (ID) - Deletes the customer by ID
- AddOrder (customerID, amount, createdAt) - Adds an order to the customer by his ID
- RemoveOrder (ID) - Removes an order by ID
- AddSubProduct (orderID, price, size, color, amount) - Adds a product to the order by its ID
- RemoveSubProduct (ID) - Removes a product from the order by ID
- CreateAdmin (login, password) - Creates an admin
- DeleteAdmin (ID) - Deletes the admin by id

> Data editing is not yet implemented

### Creators
###### Backend:
- [Slava](https://github.com/Wedyarit)
- [Denis](https://github.com/FaneNohman)

###### Nuxt GraphQL Apollo
- [Zhan](https://github.com/Vafailis)
- [Aldiyar](https://github.com/AldiyarSergazy)

###### Frontend
- [Artyom](https://github.com/ket02jfu)
- [Daniyar](https://github.com/DanikBruh)

###### Design
- [Asima](https://github.com/AsiyaBl)
- [Tima](https://github.com/Hicfok)

###### Other
- [Vlad](https://github.com/MiyRon-Code)

Penetration | Ass we can | Dungeon Master
