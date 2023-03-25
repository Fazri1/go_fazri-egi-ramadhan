-- PERINTAH YANG DIJALANKAN UNTUK MENGERJAKAN SOAL 2

-- 1
CREATE DATABASE alta_online_shop;
USE alta_online_shop;

-- 2
CREATE TABLE users(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    address TEXT,
    birth DATE,
    status SMALLINT,
    gender VARCHAR(1),
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE payment_methods(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE transactions(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(50) UNIQUE,
    date TIMESTAMP,
    total_price DOUBLE,
    total_qty INT,
    payment_method_id INT(11),
    user_id INT(11),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id),
    FOREIGN KEY (user_id) REFERENCES users(id));

CREATE TABLE product_descriptions(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    description TEXT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE product_types(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP); 

CREATE TABLE operators(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE products(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    code VARCHAR(50) UNIQUE,
    name VARCHAR(100),
    product_description_id INT(11),
    product_type_id INT(11),
    operator_id INT(11),
    price DOUBLE,
    stock INT,
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_description_id) REFERENCES product_descriptions(id),
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id));

CREATE TABLE transaction_details(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    transaction_id INT(11),
    product_id INT(11),
    price DOUBLE,
    qty INT,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)); 


-- 3
CREATE TABLE kurir(
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    updated_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);

-- 4
ALTER TABLE kurir ADD ongkos_dasar DOUBLE;

-- 5
ALTER TABLE kurir RENAME TO shipping;

-- 6
DROP TABLE shipping;