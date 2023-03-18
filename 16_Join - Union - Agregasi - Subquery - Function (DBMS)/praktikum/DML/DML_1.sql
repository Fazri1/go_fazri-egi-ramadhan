-- INSERT
-- A
INSERT INTO operators (name) VALUES
    ('abadi jaya'),
    ('furnibest'),
    ('msiofficial'),
    ('VansOfficialStore'),
    ('chelsStore');

-- B
INSERT INTO product_types (name) VALUES
    ('Komputer'),
    ('Elektronik'),
    ('Sepatu'); 

-- C
INSERT INTO products (code, name, product_type_id, operator_id, price, stock) VALUES
    ('komp0001', 'MSI GF63 THIN', 1, 3, 11499000, 100),
    ('komp0002', 'MSI Geforce RTX 3060', 1, 3, 6499000, 100); 

-- D
INSERT INTO products (code, name, product_type_id, operator_id, price, stock) VALUES
    ('elt0001', 'SHARP LED TV 24 inch', 2, 1, 1550000, 120),
    ('elt0002', 'TOSHIBA LED TV 43 inch', 2, 1, 4499000, 80),
    ('elt0003', 'Kipas Angin COSMOS', 2, 1, 239000, 150);

-- E 
INSERT INTO products (code, name, product_type_id, operator_id, price, stock) VALUES
    ('sep0001', 'Vans Old Skool Blackwhite', 3, 4, 449000, 200),
    ('sep0002', 'Vans Slip On', 3, 4, 599000, 200),
    ('sep0003', 'Vans Old Skool Navy', 3, 4, 499000, 200); 

-- F
INSERT INTO product_descriptions (description) VALUES
    ('The GeForce RTX™ 3060 lets you take on the latest games using the power of Ampere-NVIDIA 2nd generation RTX architecture. Get incredible performance with enhanced Ray Tracing Cores and Tensor Cores, new streaming multiprocessors, and high-speed G6 memory.'),
    ('MSI GF63 Thin 10SC-860ID, Spesifikasi : Display : 15.6" FHD (1920*1080), IPS-Level, Thin Bezel, Warna : Black, Camera : HD type (30fps@720p), VGA, V-RAM : GTX1650 GDDR6 4GB, CPU : Comet lake i5-10500H+HM470, Keyboard : Single backlight KB(Red), Memory : DDR IV 8GB (3200MHz), Memory Slot, Max Capacity : 2 slots, Max 64GB, Storage : 256GB NVMe PCIe SSD, Storage Slot: 1x M.2 SSD slot (NVMe PCIe Gen3)'),
    ('SHARP LED TV 24 Inch HD Digital - 2T-C24DC1i hadir dengan display berukuran  24 Inch yang memanjakan Anda. Dengan resolusi 1366 x 768 HD Ready menampilkan warna yang tajam dan terang. Anda dapat menyambungkan dengan perangkat lainnya menggunakan  HDMI atau USB. Sangat cocok menjadi pilihan Anda. '),
    ('TOSHIBA LED TV - FHD SMART VIDAA 43" adalah smart TV yang dilengkapi prosesor tercanggih Vidaa U4.2 dan berkemampuan mengolah data dengan begitu cepat sehingga Anda bisa lebih mudah dan nyaman menikmati berbagai hiburan. Audio visual yang dihasilkan 43E31KP juga sangat berkualitas. Anda dan keluarga akan disajikan banyak hiburan mewah di rumah setiap saat.'),
    ('Warnanya yang memberikan ketenangan serta dengan diameter baling-balingnya yang besar dapat memberikan kesejukan kesetiap penjuru ruangan anda, membuat Cosmos Stand Fan 16 inch 16-SDB sangat cocok berada di ruangan rumah anda.'),
    ('Vans Old Skool Black White Classic ORIGINAL 100% Global Release'),
    ('VANS SLIP ON CHECKERBOARD ORIGINAL 100% Global Release'),
    ('VANS OLD SKOOL CLASSIC NAVY ORIGINAL 100% Global Release'); 

UPDATE products SET product_description_id=2 WHERE id=1;
UPDATE products SET product_description_id=1 WHERE id=2;
UPDATE products SET product_description_id=3 WHERE id=3;
UPDATE products SET product_description_id=4 WHERE id=4;
UPDATE products SET product_description_id=5 WHERE id=5;
UPDATE products SET product_description_id=6 WHERE id=6;
UPDATE products SET product_description_id=7 WHERE id=7;
UPDATE products SET product_description_id=8 WHERE id=8;

-- G
INSERT INTO payment_methods (name) VALUES
    ('Bank'),
    ('Dana'),
    ('Kartu Kredit');

-- H
INSERT INTO users (name, address, birth, gender) VALUES
    ('Bambang', 'Kota Bogor', '2001-08-01', 'L'),
    ('Handoko', 'Depok', '1990-01-01', 'L'),
    ('Siti', 'Karawang', '2005-03-21', 'P'),
    ('Putri', 'Cimahi', '2000-12-12', 'P'),
    ('Bayu', 'Kalibata', '1987-07-15', 'L');

-- I 
INSERT INTO transactions (code, payment_method_id, user_id) VALUES
    ('tr03230001', 1, 1),
    ('tr03230002', 1, 1),
    ('tr03230003', 1, 1),
    ('tr03230004', 1, 2),
    ('tr03230005', 1, 2),
    ('tr03230006', 1, 2),
    ('tr03230007', 3, 3),
    ('tr03230008', 3, 3),
    ('tr03230009', 3, 3),
    ('tr03230010', 2, 4),
    ('tr03230011', 2, 4),
    ('tr03230012', 2, 4),
    ('tr03230013', 3, 5),
    ('tr03230014', 3, 5),
    ('tr03230015', 3, 5);

-- J
INSERT INTO transaction_details (transaction_id, product_id, price, qty) VALUES
    (1, 2, (SELECT price FROM products WHERE id=2), 1),
    (1, 1, (SELECT price FROM products WHERE id=1), 1),
    (1, 4, (SELECT price FROM products WHERE id=4), 1),
    (2, 7, (SELECT price FROM products WHERE id=7), 1),
    (2, 5, (SELECT price FROM products WHERE id=5), 1),
    (2, 2, (SELECT price FROM products WHERE id=2), 1),
    (3, 3, (SELECT price FROM products WHERE id=3), 1),
    (3, 6, (SELECT price FROM products WHERE id=6), 1),
    (3, 1, (SELECT price FROM products WHERE id=1), 1),
    (4, 6, (SELECT price FROM products WHERE id=6), 1),
    (4, 7, (SELECT price FROM products WHERE id=7), 1),
    (4, 8, (SELECT price FROM products WHERE id=8), 1),
    (5, 5, (SELECT price FROM products WHERE id=5), 2),
    (5, 3, (SELECT price FROM products WHERE id=3), 3),
    (5, 1, (SELECT price FROM products WHERE id=1), 1),
    (6, 2, (SELECT price FROM products WHERE id=2), 6),
    (6, 5, (SELECT price FROM products WHERE id=5), 1),
    (6, 7, (SELECT price FROM products WHERE id=7), 1),
    (7, 1, (SELECT price FROM products WHERE id=1), 2),
    (7, 3, (SELECT price FROM products WHERE id=3), 1),
    (7, 5, (SELECT price FROM products WHERE id=5), 1),
    (8, 7, (SELECT price FROM products WHERE id=7), 1),
    (8, 2, (SELECT price FROM products WHERE id=2), 1),
    (8, 4, (SELECT price FROM products WHERE id=4), 1),
    (9, 7, (SELECT price FROM products WHERE id=7), 1),
    (9, 6, (SELECT price FROM products WHERE id=6), 1),
    (9, 8, (SELECT price FROM products WHERE id=8), 1),
    (10, 2, (SELECT price FROM products WHERE id=2), 3),
    (10, 1, (SELECT price FROM products WHERE id=1), 1),
    (10, 3, (SELECT price FROM products WHERE id=3), 1),
    (11, 7, (SELECT price FROM products WHERE id=7), 1),
    (11, 5, (SELECT price FROM products WHERE id=5), 1),
    (11, 4, (SELECT price FROM products WHERE id=4), 1),
    (12, 2, (SELECT price FROM products WHERE id=2), 2),
    (12, 5, (SELECT price FROM products WHERE id=5), 2),
    (12, 7, (SELECT price FROM products WHERE id=7), 1),
    (13, 5, (SELECT price FROM products WHERE id=5), 1),
    (13, 3, (SELECT price FROM products WHERE id=3), 3),
    (13, 7, (SELECT price FROM products WHERE id=7), 1),
    (14, 2, (SELECT price FROM products WHERE id=2), 1),
    (14, 1, (SELECT price FROM products WHERE id=1), 1),
    (14, 6, (SELECT price FROM products WHERE id=6), 1),
    (15, 1, (SELECT price FROM products WHERE id=1), 1),
    (15, 5, (SELECT price FROM products WHERE id=5), 1),
    (15, 2, (SELECT price FROM products WHERE id=2), 1);

-- UPDATE tabel transaction
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=1), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=1) WHERE id=1;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=2), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=2) WHERE id=2;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=3), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=3) WHERE id=3;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=4), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=4) WHERE id=4;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=5), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=5) WHERE id=5;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=6), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=6) WHERE id=6;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=7), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=7) WHERE id=7;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=8), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=8) WHERE id=8;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=9), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=9) WHERE id=9;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=10), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=10) WHERE id=10;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=11), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=11) WHERE id=11;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=12), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=12) WHERE id=12;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=13), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=13) WHERE id=13;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=14), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=14) WHERE id=14;
UPDATE transactions SET total_price=(SELECT SUM(price*qty) FROM transaction_details WHERE transaction_id=15), total_qty=(SELECT SUM(qty) FROM transaction_details WHERE transaction_id=15) WHERE id=15;