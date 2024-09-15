INSERT INTO users (id, email, phone_number, first_name, last_name, image) 
VALUES (1, 'sa@sa.com', '123', 'sa', 'sa', '')
ON CONFLICT (email) DO NOTHING;

INSERT INTO accounts (username, role, email, password, type, created_at)
VALUES ('user#1', 'Admin', 'sa@sa.com', '$2a$12$x9jSMLNRSTX.kOmLThUFH.VanhMBTUJrFxUwNOkx7GOLTGrMWl5Kq', 'swc', '2024-07-31 07:51:50+00')
ON CONFLICT (username) DO NOTHING;
--
-- TOC entry 3416 (class 0 OID 16468)
-- Dependencies: 230
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO categories (id, name, description) VALUES (1, 'phone', 'iPhone');
INSERT INTO categories (id, name, description) VALUES (2, 'accessories', 'accessories');

--
-- TOC entry 3418 (class 0 OID 16457)
-- Dependencies: 228
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO suppliers (id, name, email) VALUES (1, 'Apple', 'exam@example2.com');

--
-- TOC entry 3418 (class 0 OID 16447)
-- Dependencies: 226
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (1, '', '1.000.000 - 2.000.000', 'iPhone 12', 'iPhone 12', 1, 1, '2024-08-03 16:43:30.048345', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.1 inch", "display": "Super Retina XDR display"}', 'active');
INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (2, '', '2.000.000 - 3.000.000', 'iPhone 12 Pro', 'iPhone 12 Pro', 1, 1, '2024-08-03 16:44:35.806808', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.7 inch", "display": "Super Retina XDR display"}', 'active');
INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (3, '', '1.000.000 - 2.000.000', 'iPhone 12 Mini', 'iPhone 12 Mini', 1, 1, '2024-08-03 16:45:45.19697', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "5.4 inch", "display": "Super Retina XDR display"}', 'active');
INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (4, '', '500.000', 'Apple iPhone Adapter 20W', 'Apple iPhone Adapter', 1, 2, '2024-08-03 16:54:21.488746', '{}', 'active');
INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (5, '', '500.000', 'Apple iPhone Case', 'Apple iPhone Case', 1, 2, '2024-08-03 16:54:48.684444', '{}', 'active');
INSERT INTO products (id, image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (6, '', '500.000', 'Apple iPhone Screen Protector', 'Apple iPhone Screen Protector', 1, 2, '2024-08-03 16:55:02.747345', '{}', 'active');

--
-- TOC entry 3416 (class 0 OID 16547)
-- Dependencies: 248
-- Data for Name: collections; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO collections (id, created, "position", headline, body) VALUES (1, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/1.jpg", "title": "You can do more with AI.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Artificial Intelligence"}');
INSERT INTO collections (id, created, "position", headline, body) VALUES (2, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/2.jpg", "title": "Enhance your productivity.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Productivity"}');
INSERT INTO collections (id, created, "position", headline, body) VALUES (3, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/3.jpg", "title": "Launching the new Apple Vision Pro.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Product"}');
INSERT INTO collections (id, created, "position", headline, body) VALUES (4, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/4.jpg", "title": "Maps for your iPhone 15 Pro Max.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Product"}');
INSERT INTO collections (id, created, "position", headline, body) VALUES (5, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/5.jpg", "title": "Photography just got better.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "iOS"}');
INSERT INTO collections (id, created, "position", headline, body) VALUES (6, '2024-08-03 17:26:15.088919', 'phone_1', 'Get to know your iPhone.', '{"src": "/img/posts/6.jpg", "title": "Hiring for a Staff Software Engineer.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Hiring"}');

--
-- TOC entry 3421 (class 0 OID 16538)
-- Dependencies: 246
-- Data for Name: inventories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO inventories (id, product_id, price, status, currency_code, available, color, color_img, image) VALUES (2, 1, 10000.0000, 'active', 'VND', 1000, 'Blue Titanium', 'https://example.com/grey-titanium.jpg', 'https://example.com/iphone-12.jpg,https://example.com/iphone-12.jpg');
INSERT INTO inventories (id, product_id, price, status, currency_code, available, color, color_img, image) VALUES (3, 1, 10000.0000, 'active', 'VND', 1000, 'Red Titanium', 'https://example.com/grey-titanium.jpg', 'https://example.com/iphone-12.jpg,https://example.com/iphone-12.jpg');
INSERT INTO inventories (id, product_id, price, status, currency_code, available, color, color_img, image) VALUES (1, 1, 10000.0000, 'active', 'VND', 5000, 'Grey Titanium', 'https://example.com/grey-titanium.jpg', 'https://example.com/iphone-12.jpg,https://example.com/iphone-12.jpg');
INSERT INTO inventories (id, product_id, price, status, currency_code, available, color, color_img, image) VALUES (4, 1, 10000000.0000, 'active', 'VND', 70, 'Black Titanium', 'https://example.com/black-titanium.jpg', 'https://example.com/iphone-12.jpg,https://example.com/iphone-12-2.jpg');

--
-- TOC entry 3421 (class 0 OID 16557)
-- Dependencies: 250
-- Data for Name: specifications; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO specifications (id, inventory_id, content) VALUES (1, 1, '{"ram": "16", "ssd": "512"}');
INSERT INTO specifications (id, inventory_id, content) VALUES (2, 1, '{"ram": "32", "ssd": "512"}');
INSERT INTO specifications (id, inventory_id, content) VALUES (3, 2, '{"ram": "16", "ssd": "512"}');
INSERT INTO specifications (id, inventory_id, content) VALUES (4, 3, '{"ram": "32", "ssd": "512"}');
INSERT INTO specifications (id, inventory_id, content) VALUES (5, 4, '{"ram": "8GB", "ssd": "256GB"}');
INSERT INTO specifications (id, inventory_id, content) VALUES (6, 4, '{"ram": "16GB", "ssd": "512"}');

--
-- TOC entry 3421 (class 0 OID 16557)
-- Dependencies: 250
-- Data for Name: specifications; Type: TABLE DATA; Schema: public; Owner: admin
--
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (1, 0, 0, 1, 1, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (2, 0, 0, 1, 1, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (3, 0, 0, 1, 2, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (4, 0, 0, 1, 3, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (5, 0, 0, 1, 4, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (6, 0, 0, 1, 5, 5, 0, 0, 'This is a great product');