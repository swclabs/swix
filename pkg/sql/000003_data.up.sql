INSERT INTO users (id, email, phone_number, first_name, last_name, image) 
VALUES (1, 'sa@sa.com', '123', 'sa', 'sa', '')
ON CONFLICT (email) DO NOTHING;


-- setval for users_id_seq
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));

INSERT INTO accounts (id, username, role, email, password, type, created_at)
VALUES (1, 'user#1', 'Admin', 'sa@sa.com', '$2a$12$x9jSMLNRSTX.kOmLThUFH.VanhMBTUJrFxUwNOkx7GOLTGrMWl5Kq', 'swc', '2024-07-31 07:51:50+00')
ON CONFLICT (username) DO NOTHING;

-- setval for accounts_id_seq
SELECT setval('accounts_id_seq', (SELECT MAX(id) FROM accounts));   

INSERT INTO public.categories VALUES (1, 'phone', 'iPhone');
INSERT INTO public.categories VALUES (2, 'accessories', 'accessories');
INSERT INTO public.categories VALUES (3, 'tablet', 'iPad');
INSERT INTO public.categories VALUES (4, 'watch', 'Watch');
INSERT INTO public.categories VALUES (5, 'earphone', 'AirPod');
INSERT INTO public.categories VALUES (6, 'computer', 'Mac');


--
-- TOC entry 3413 (class 0 OID 0)
-- Dependencies: 230
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.categories_id_seq', 6, true);

--
-- TOC entry 3418 (class 0 OID 16457)
-- Dependencies: 228
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO suppliers (id, name, email) VALUES (1, 'Apple', 'exam@example2.com');
-- setval for suppliers_id_seq
SELECT setval('suppliers_id_seq', (SELECT MAX(id) FROM suppliers));

--
-- TOC entry 3409 (class 0 OID 20173)
-- Dependencies: 227
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.products VALUES (2, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571269/swc-storage/g5xyddn9yslxzozdenmn.png', '', 'Từ 22.999.000đ hoặc 936.000đ/th. trong 24 tháng*', 'Siêu mạnh mẽ trên mọi mặt.', 'iPhone 15', 1, 1, '2024-10-10 14:37:26.298902', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products VALUES (3, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571302/swc-storage/h8skd6ufmiy3x3ciedfi.png', '', 'Từ 19.999.000đ hoặc 814.000đ/th. trong 24 tháng*', 'Luôn tuyệt vời như thế.', 'iPhone 14', 1, 1, '2024-10-10 14:38:02.415319', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products VALUES (4, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571323/swc-storage/fm6cretpjbzkgwbns9do.png', '', 'Từ 17.299.000đ hoặc 704.000đ/th. trong 24 tháng*', 'Hội tụ mọi điều tuyệt diệu.', 'iPhone 13', 1, 1, '2024-10-10 14:39:12.689638', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products VALUES (5, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571338/swc-storage/czmjjtvcgcqqc8csxsjx.png', '', 'Từ 11.999.000đ hoặc 489.000đ/th. trong 24 tháng*', 'Thực sự mạnh mẽ. Thực sự giá trị.', 'iPhone SE', 1, 1, '2024-10-10 14:39:38.656298', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products VALUES (6, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635080/swc-storage/fce1xe5tpgchkpoe2jgu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635169/swc-storage/mreuudhtmybm0pjluspy.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635170/swc-storage/rqeksvzxoqy5nemujwme.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635171/swc-storage/anaby1wfox1zccuahpi0.jpg', 'Từ 28.999.000đ hoặc 1.181.000đ/th. trong 24 tháng*', 'Một iPhone cực đỉnh.', 'iPhone 15 Pro Max', 1, 1, '2024-10-11 08:22:32.021943', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products VALUES (1, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571252/swc-storage/k4jskdsqnkwxbdabfne5.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571470/swc-storage/uasxqatjfu2kopdjljdl.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571472/swc-storage/uzdavi6lqcqhc3wkywbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571474/swc-storage/nl8qic2e0zm5jfz1mtx0.jpg', 'Từ 28.999.000đ hoặc 1.181.000đ/th. trong 24 tháng*', 'Một iPhone cực đỉnh.', 'iPhone 15 Pro', 1, 1, '2024-10-11 08:19:45.449374', '{"RAM": null, "SSD": null, "screen": "6.1 inch", "display": "Super Retina XDR"}', 'active');


--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 226
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.products_id_seq', 6, true);

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


SELECT setval('collections_id_seq', (SELECT MAX(id) FROM collections));

--
-- TOC entry 3408 (class 0 OID 20245)
-- Dependencies: 243
-- Data for Name: inventories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO inventories VALUES (5, 1, 43999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615248/swc-storage/b5pipnb8pmkdmxweys0z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615249/swc-storage/jaysghjdmyace5fxxta8.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615250/swc-storage/utcd96odzu8dloe24085.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615251/swc-storage/k7vs30bg9b3hd4geguxr.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (4, 1, 37999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (1, 1, 28999000.0000, 'active', 'VND', 17, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (6, 1, 28999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (2, 1, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (3, 1, 31999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (7, 1, 31999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (8, 1, 37999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (24, 6, 37999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (25, 6, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (9, 1, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (10, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (11, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (12, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (13, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (18, 6, 28999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (19, 6, 31999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (20, 6, 37999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (14, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (15, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (16, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (17, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (21, 6, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (22, 6, 28999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO inventories VALUES (23, 6, 31999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');


--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 242
-- Name: inventories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('inventories_id_seq', 25, true);


INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (1, 0, 0, 1, 1, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (2, 0, 0, 1, 1, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (3, 0, 0, 1, 2, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (4, 0, 0, 1, 3, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (5, 0, 0, 1, 4, 5, 0, 0, 'This is a great product');
INSERT INTO comments (id, level, parent_id, user_id, product_id, rating, liked, disliked, content) VALUES (6, 0, 0, 1, 5, 5, 0, 0, 'This is a great product');

SELECT setval('comments_id_seq', (SELECT MAX(id) FROM comments));   

-- Completed on 2024-09-21 20:03:14

--
-- PostgreSQL database dump complete
--

--
-- TOC entry 3411 (class 0 OID 16706)
-- Dependencies: 222
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.addresses VALUES (1, 'Hồ Chí Minh', 'Phường Thảo Điền', 'Thành Phố Thủ Đức', 'Xa lộ Hà Nội', 1);


--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 221
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.addresses_id_seq', 1, true);


-- Completed on 2024-09-21 20:00:17

--
-- PostgreSQL database dump complete
--

--
-- TOC entry 3412 (class 0 OID 16781)
-- Dependencies: 238
-- Data for Name: deliveries; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.deliveries (id, user_id, address_id, sent_date, received_date, status, method, note) VALUES (1, 1, 1, '0001-01-01 00:00:00+00', '0001-01-01 00:00:00+00', 'active', 'COD', '');


--
-- TOC entry 3418 (class 0 OID 0)
-- Dependencies: 237
-- Name: deliveries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.deliveries_id_seq', 1, true);


-- Completed on 2024-09-22 13:09:17

--
-- PostgreSQL database dump complete
--