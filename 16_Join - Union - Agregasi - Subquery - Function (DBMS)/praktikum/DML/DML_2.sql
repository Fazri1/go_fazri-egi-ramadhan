-- SELECT
-- A
SELECT name FROM users WHERE gender='L';

-- B
SELECT * FROM products WHERE id=3;

-- C
SELECT * FROM users WHERE name LIKE '%a%' AND created_at BETWEEN '2023-03-11 12:00:00' AND '2023-03-18 12:00:00';

-- D
SELECT COUNT(gender) AS total_women FROM users WHERE gender='P';

-- E
SELECT * FROM users ORDER BY name; 

-- F
SELECT * FROM products LIMIT 5;