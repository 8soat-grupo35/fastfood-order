CREATE TABLE IF NOT EXISTS customers(
    id serial primary key,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    cpf varchar(11) NOT NULL,
    created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL
);

CREATE TABLE IF NOT EXISTS items(
    id serial primary key,
    name varchar(255) NOT NULL,
    category varchar(30) NOT NULL,
    price numeric NOT NULL,
    image_url varchar(255) NOT NULL,
    created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL
);

CREATE TABLE IF NOT EXISTS orders(
    id serial primary key,
    status varchar(50) NOT NULL,
    customer_id int NOT NULL,
    created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,

    CONSTRAINT fk_customer_orders
      FOREIGN KEY(customer_id) 
      REFERENCES customers(id)
      ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS order_items(
    id serial primary key,
    order_id int NOT NULL,
    item_id int NOT NULL,
    quantity int NOT NULL,
    created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,

    CONSTRAINT fk_order_orders_items
      FOREIGN KEY(order_id) 
      REFERENCES orders(id)
      ON DELETE SET NULL,

    CONSTRAINT fk_item_orders_items
      FOREIGN KEY(item_id) 
      REFERENCES items(id)
      ON DELETE SET NULL
);

INSERT INTO items (name, category, price, image_url, created_at, updated_at, deleted_at) VALUES ('X-Burguer', 'LANCHE', 28, 'https://fastly.picsum.photos/id/8/200/200.jpg?hmac=7z37E8o2M_U09oSFIN5CdqKXlYXuLeWxTHJVlT9UUlY', 'NOW'::timestamptz, 'NOW'::timestamptz, null);

INSERT INTO items (name, category, price, image_url, created_at, updated_at, deleted_at) VALUES ('X-Bacon', 'LANCHE', 35, 'https://fastly.picsum.photos/id/8/200/200.jpg?hmac=7z37E8o2M_U09oSFIN5CdqKXlYXuLeWxTHJVlT9UUlY', 'NOW'::timestamptz, 'NOW'::timestamptz, null);

INSERT INTO customers (name, email, cpf, created_at, updated_at, deleted_at) VALUES ('John Doe', 'john@gmail.com', '12345678911', 'NOW'::timestamptz, 'NOW'::timestamptz, null);