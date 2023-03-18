-- JOIN, UNION, SUB QUERY, FUNCTION
-- 1
SELECT * FROM transactions WHERE user_id=1
UNION                                                          
SELECT * FROM transactions WHERE user_id=2;

-- 2
SELECT SUM(t.total_price) FROM users u INNER JOIN transactions t ON u.id = t.user_id WHERE u.id = 1;

-- 3
SELECT COUNT(t.id) AS total_transaction_product_type_2 FROM transactions t LEFT JOIN transaction_details td ON td.transaction_id=t.id LEFT JOIN products p ON td.product_id=p.id WHERE p.product_type_id=2; 

-- 4
SELECT p.*, pt.name AS product_type FROM products p LEFT JOIN product_types pt ON p.product_type_id=pt.id; 

-- 5
SELECT t.*, p.name AS product_name, u.name AS user_name FROM transactions t LEFT JOIN users u ON t.user_id=u.id LEFT JOIN transaction_details td ON t.id=td.transaction_id LEFT JOIN products p ON td.product_id=p.id;

-- 6
DELIMITER $$
CREATE TRIGGER delete_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN                                     
DECLARE v_transaction_id INT;             
SET v_transaction_id=OLD.id;              
DELETE FROM transaction_details WHERE transaction_id=v_transaction_id;
END$$
DELIMITER ;

-- 7
DELIMITER $$
CREATE TRIGGER update_total_qty
BEFORE DELETE ON transaction_details FOR EACH ROW   
BEGIN                                               
UPDATE transactions SET total_qty=total_qty-(SELECT qty FROM transaction_details WHERE id=OLD.id) WHERE id=OLD.transaction_id;                    
END $$
DELIMITER ;

-- 8
SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details);