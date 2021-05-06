.open pad.db


DROP TABLE categories;
DROP TABLE products;
DROP TABLE product_types;
DROP TABLE customers;
DROP TABLE admins;
DROP TABLE orders;
DROP TABLE ordered_products;


CREATE TABLE categories (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  image_url TEXT NOT NULL
);

CREATE TABLE products (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  price DECIMAL NOT NULL,
  description TEXT NOT NULL,
  delivery_date TIMESTAMP NOT NULL,
  image_url TEXT NOT NULL,
  product_type_id TEXT NOT NULL,
  FOREIGN KEY (product_type_id) REFERENCES product_types (id) ON DELETE CASCADE
);

CREATE TABLE product_types (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  category_id TEXT NOT NULL,
  FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);

CREATE TABLE customers (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  phone TEXT NOT NULL,
  address TEXT NOT NULL,
  city_region TEXT NOT NULL,
  cc_number TEXT NOT NULL
);

CREATE TABLE admins (
  id TEXT PRIMARY KEY,
  login TEXT NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE orders (
  id TEXT PRIMARY KEY,
  amount DECIMAL NOT NULL,
  created_at TIMESTAMP NOT NULL,
  customer_id TEXT NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
);

CREATE TABLE ordered_products (
    product_id TEXT NOT NULL,
    customer_order_id TEXT NOT NULL,
    FOREIGN KEY (customer_order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
    PRIMARY KEY (product_id, customer_order_id)
);


