DROP TABLE Admin;
DROP TABLE Category;
DROP TABLE Product;
DROP TABLE Sub_Product;
DROP TABLE Customer;
DROP TABLE 'Order';
DROP TABLE OrderedProduct;
-- CREATION
CREATE TABLE Admin (
  id TEXT PRIMARY KEY,
  login TEXT NOT NULL,
  password TEXT NOT NULL
);
CREATE TABLE Category (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  image_url TEXT NOT NULL
);
CREATE TABLE Product (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  manufacturer TEXT NOT NULL,
  fabric TEXT NOT NULL,
  category_id TEXT REFERENCES Category(id)
);
CREATE TABLE Sub_Product (
  id TEXT PRIMARY KEY,
  price REAL NOT NULL,
  size TEXT NOT NULL,
  color TEXT NOT NULL,
  amount INTEGER NOT NULL,
  product_id TEXT REFERENCES Product(id)
);
CREATE TABLE Customer (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  phone TEXT NOT NULL,
  address TEXT NOT NULL,
  region TEXT NOT NULL,
  cc_number TEXT NOT NULL
);
CREATE TABLE 'Order' (
  id TEXT PRIMARY KEY,
  amount INTEGER NOT NULL,
  created_at TEXT NOT NULL,
  customer_id TEXT REFERENCES Customer(id)
);
CREATE TABLE OrderedProduct (
  id TEXT PRIMARY KEY,
  amount INTEGER NOT NULL,
  order_id TEXT REFERENCES 'Order'(id),
  sub_product_id TEXT REFERENCES Sub_Product(id)
);
CREATE TABLE Images (
  id TEXT PRIMARY KEY,
  image_url TEXT NOT NULL,
  sub_product_id TEXT REFERENCES Sub_Product(id)
)