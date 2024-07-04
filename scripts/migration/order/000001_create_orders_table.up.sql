-- migrate_up.sql

-- Create orders table
CREATE TABLE orders
(
    id            VARCHAR(255) PRIMARY KEY,
    user_id       VARCHAR(255)   NOT NULL,
    restaurant_id VARCHAR(255)   NOT NULL,
    status        VARCHAR(20)    NOT NULL,
    total_amount  DECIMAL(10, 2) NOT NULL,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delivery_id   VARCHAR(255)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='Order table';

-- Create order_items table
CREATE TABLE order_items
(
    order_id VARCHAR(255),
    item_id  VARCHAR(255),
    quantity INT            NOT NULL,
    price    DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (order_id, item_id),
    FOREIGN KEY (order_id) REFERENCES orders (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci COMMENT ='Order items table';
