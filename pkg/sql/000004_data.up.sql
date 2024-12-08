--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.0

-- Started on 2024-10-23 10:29:23

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3470 (class 0 OID 21964)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.users (id, email, phone_number, first_name, last_name, image) VALUES (1, 'sa@sa.com', '123', 'sa', 'sa', '');
INSERT INTO public.users (id, email, phone_number, first_name, last_name, image) VALUES (2, 'hung@gmail.com', '08767162531', 'Duc Hung', 'Ho', '');


--
-- TOC entry 3468 (class 0 OID 21949)
-- Dependencies: 217
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.accounts (id, username, role, email, password, type, created_at) VALUES (1, 'user#1', 'Admin', 'sa@sa.com', '$2a$12$x9jSMLNRSTX.kOmLThUFH.VanhMBTUJrFxUwNOkx7GOLTGrMWl5Kq', 'swc', '2024-07-31 07:51:50+00');


--
-- TOC entry 3474 (class 0 OID 21986)
-- Dependencies: 223
-- Data for Name: addresses; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.addresses (id, city, ward, district, street, user_id) VALUES (1, 'Hồ Chí Minh', 'Phường Thảo Điền', 'Thành Phố Thủ Đức', 'Xa lộ Hà Nội', 1);
INSERT INTO public.addresses (id, city, ward, district, street, user_id) VALUES (2, 'Hồ Chí Minh', 'Phường Thảo Điền', 'Thành Phố Thủ Đức', 'Xa lộ Hà Nội', 2);
INSERT INTO public.addresses (id, city, ward, district, street, user_id) VALUES (4, 'Hồ Chí Minh', 'Phường Thảo Điền', 'Thành Phố Thủ Đức', 'Xa lộ Hà Nội', 2);
INSERT INTO public.addresses (id, city, ward, district, street, user_id) VALUES (13, 'Hồ Chí Minh', 'Phường Thảo Điền', 'Thành Phố Thủ Đức', 'Xa lộ Hà Nội', 2);


--
-- TOC entry 3482 (class 0 OID 22026)
-- Dependencies: 231
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.categories (id, name, description) VALUES (1, 'phone', 'iPhone');
INSERT INTO public.categories (id, name, description) VALUES (2, 'accessories', 'accessories');
INSERT INTO public.categories (id, name, description) VALUES (3, 'tablet', 'iPad');
INSERT INTO public.categories (id, name, description) VALUES (4, 'watch', 'Watch');
INSERT INTO public.categories (id, name, description) VALUES (5, 'earphone', 'AirPod');
INSERT INTO public.categories (id, name, description) VALUES (6, 'desktop', 'Mac');
INSERT INTO public.categories (id, name, description) VALUES (7, 'display', 'Mac');
INSERT INTO public.categories (id, name, description) VALUES (8, 'laptop', 'Mac');


--
-- TOC entry 3480 (class 0 OID 22015)
-- Dependencies: 229
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.suppliers (id, name, email) VALUES (1, 'Apple', 'exam@example2.com');


--
-- TOC entry 3478 (class 0 OID 22005)
-- Dependencies: 227
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (7, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729393722/swc-storage/ycsfb8cdonhdz3czvj0j.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406776/swc-storage/precoqtpdgk7cacgfsjt.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406777/swc-storage/tbkun6bkgweeeqg47h2a.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406777/swc-storage/kwyl1ahevlgqzbphe5d4.webp', 'Từ 21.199.000đ', 'Trải nghiệm iPad cực đỉnh với công nghệ tiên tiến nhất.', 'iPad Pro', 1, 3, '2024-10-20 03:06:43.091122', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (8, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729393790/swc-storage/xigvxeoqyf2abbmiiyzw.png', '', 'Từ 15.599.000đ', 'Hiệu năng mạnh mẽ trong một thiết kế mỏng nhẹ.', 'iPad Air', 1, 3, '2024-10-20 03:09:35.887728', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (9, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729393862/swc-storage/q7evmia7ej933nzide4p.png', '', 'Từ 11.499.000đ', 'iPad màn hình toàn phần, đầy màu sắc. Cho mọi tác vụ hàng ngày.', 'iPad', 1, 3, '2024-10-20 03:10:51.206349', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (10, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729393915/swc-storage/l21ujs2j4rmxxm1wc5qu.png', '', 'Từ 8.399.000đ', 'Tất cả tính năng bạn cần gói gọn trong chiếc iPad vừa túi tiền nhất.', 'iPad', 1, 3, '2024-10-20 03:11:42.364928', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (11, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729393973/swc-storage/iz0kvlvmykbua3dxtja2.png', '', 'Từ 12.799.000đ', 'Trọn vẹn trải nghiệm iPad nằm gọn trên tay.', 'iPad mini', 1, 3, '2024-10-20 03:12:42.253262', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (12, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407136/swc-storage/fuj9iif1fmmbllqb9k8e.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407086/swc-storage/dcgmuvsu0djvbj7kelpy.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407087/swc-storage/mq94cg2azb0g4txezjov.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407088/swc-storage/hawnkybu826ogjgalds3.webp', 'Từ 12.799.000đ', 'Trọn vẹn trải nghiệm iPad nằm gọn trên tay.', 'iPad Pro 13 inch', 1, 3, '2024-10-20 06:48:34.985496', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (14, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524710/swc-storage/x72b4jpu9zhiyc8lhuwe.jpg', '', 'Từ 10.499.000đ hoặc 427.000đ/th. trong 24 tháng.', 'Cảm biến mạnh mẽ, tính năng sức khỏe tiên tiến.', 'Apple Watch Series 9', 1, 4, '2024-10-20 14:57:10.877648', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (16, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524881/swc-storage/lbtk7kyhwfkmijoop8k5.png', '', 'Từ 999đ hoặc 83.25đ/tháng trong 12 tháng.*', 'Chiếc laptop Mac giá cả phải chăng nhất để hoàn thành công việc khi di chuyển.', 'MacBook Air 13', 1, 8, '2024-10-21 15:04:01.007458', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (1, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571252/swc-storage/k4jskdsqnkwxbdabfne5.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571470/swc-storage/uasxqatjfu2kopdjljdl.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571472/swc-storage/uzdavi6lqcqhc3wkywbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571474/swc-storage/nl8qic2e0zm5jfz1mtx0.jpg', 'Từ 28.999.000đ', 'Một iPhone cực đỉnh.', 'iPhone 15 Pro', 1, 1, '2024-10-11 08:19:45.449374', '{"RAM": null, "SSD": null, "screen": "6.1 inch", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (19, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524805/swc-storage/ubkzcgmzgenw5cavpzj9.png', '', 'Từ $1299 hoặc $108.25/tháng trong 12 tháng.*', 'Một chiếc máy tính để bàn tất cả trong một tuyệt đẹp cho sự sáng tạo và năng suất.', 'iMac', 1, 6, '2024-10-21 15:06:52.683861', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (18, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524908/swc-storage/dsjsatk8dyjk0mne0cag.png', '', 'Từ 1599 hoặc 133.25/tháng trong 12 tháng.', 'Những chiếc laptop Mac tiên tiến nhất cho các quy trình làm việc đòi hỏi cao.', 'MacBook Pro 14 và 16', 1, 8, '2024-10-21 15:05:20.477541', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (30, '', '', '499.000đ', 'Gắn thêm một ốp lưng, một ví da từ tính hay cả hai. Để sạc không dây nhanh và hiệu quả.', 'MagSafe', 1, 2, '2024-10-21 15:21:38.277349', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (31, '', '', '199.000đ', 'Gắn một cái vào chìa khóa của bạn. Đặt một cái khác vào ba lô. Nhỡ có thất lạc, chỉ cần dùng ứng dụng Tìm.', 'AirTag', 1, 2, '2024-10-21 15:22:01.887987', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (26, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524451/swc-storage/v5czbatouvymcjmv1wzi.png', '', '3.499.000đ', 'Thế hệ thứ 2', 'AirPods', 1, 5, '2024-10-21 15:18:25.843997', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (29, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524481/swc-storage/ylqwe3iroh8ozd8efmhq.png', '', '4.499.000đ', 'Thế hệ thứ 3', 'AirPods', 1, 5, '2024-10-21 15:20:36.695189', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (28, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524529/swc-storage/qjbju5lcb6wgysedqkrq.png', '', '6.199.000đ', 'Thế hệ thứ 2', 'AirPods Pro', 1, 5, '2024-10-21 15:19:34.419569', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (27, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524543/swc-storage/fd9jmejtxajgv3uoyrgq.png', '', '13.199.000đ', 'Thế hệ thứ 3', 'AirPods Max', 1, 5, '2024-10-21 15:19:12.93202', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (13, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524649/swc-storage/vstqak2nqnwyqdkndnr7.jpg', '', 'Từ 6.399.000đ hoặc 261.000đ/th. trong 24 tháng.', 'Tất cả tính năng bạn cần. Giá nhẹ nhàng.', 'Apple Watch SE', 1, 4, '2024-10-20 14:56:30.774649', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (15, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524675/swc-storage/bfobkjysu4n10veerg9v.jpg', '', 'Từ 21.999.000đ hoặc 896.000đ/th. trong 24 tháng.', 'Ngầu và giàu năng lực nhất.', 'Apple Watch Ultra 2', 1, 4, '2024-10-20 14:57:32.775267', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (21, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524820/swc-storage/ujskhucqaesgri0q8jsw.png', '', 'Từ $599 hoặc $49.91/tháng trong 12 tháng.*', 'Chiếc máy tính để bàn Mac giá cả phải chăng nhất với hiệu suất vượt trội.', 'Mac mini', 1, 6, '2024-10-21 15:12:59.405665', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (23, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524842/swc-storage/rr2edmll8pfaxosusoou.png', '', 'Từ $6999 hoặc $583.25/tháng trong 12 tháng', 'Một máy trạm chuyên nghiệp với khả năng mở rộng PCIe cho các quy trình làm việc đòi hỏi cao.', 'Mac Pro', 1, 6, '2024-10-21 15:14:13.266605', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (22, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524864/swc-storage/ezu2pbr3jdbpokbv5kdv.png', '', 'Từ $1999 hoặc $166.58/tháng trong 12 tháng', 'Hiệu suất mạnh mẽ và kết nối mở rộng cho các quy trình làm việc chuyên nghiệp.', 'Mac Studio', 1, 6, '2024-10-21 15:13:23.059416', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (17, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524894/swc-storage/pjh9ujyflmcdlzp3ybbk.png', '', 'Từ 1099 hoặc 91.58/tháng trong 12 tháng', 'Mỏng và nhanh đến mức bạn có thể làm việc, chơi hoặc sáng tạo ở bất cứ đâu.', 'MacBook Air 13 và 15', 1, 8, '2024-10-21 15:04:47.187676', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (24, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524954/swc-storage/bhkvueujyphwwamjqat4.png', '', 'Từ $1599 hoặc $133.25/tháng trong 12 tháng.', 'Màn hình Retina 5K với camera và âm thanh tuyệt vời.', 'Studio Display', 1, 7, '2024-10-21 15:15:13.389073', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (25, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729524972/swc-storage/ozwunptltv47m2mov9uk.png', '', 'Từ $4999 hoặc $416.58/tháng trong 12 tháng.*', 'Màn hình XDR 6K tiên tiến cho các quy trình làm việc chuyên nghiệp.', 'Pro Display XDR', 1, 7, '2024-10-21 15:15:41.691094', '{"RAM": [], "SSD": [], "screen": "", "display": ""}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (2, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571269/swc-storage/g5xyddn9yslxzozdenmn.png', '', 'Từ 22.999.000đ', 'Siêu mạnh mẽ trên mọi mặt.', 'iPhone 15', 1, 1, '2024-10-10 14:37:26.298902', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (3, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571302/swc-storage/h8skd6ufmiy3x3ciedfi.png', '', 'Từ 19.999.000đ', 'Luôn tuyệt vời như thế.', 'iPhone 14', 1, 1, '2024-10-10 14:38:02.415319', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (4, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571323/swc-storage/fm6cretpjbzkgwbns9do.png', '', 'Từ 17.299.000đ', 'Hội tụ mọi điều tuyệt diệu.', 'iPhone 13', 1, 1, '2024-10-10 14:39:12.689638', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (5, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728571338/swc-storage/czmjjtvcgcqqc8csxsjx.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406138/swc-storage/shcwkpuynmqkqwpxgcea.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406138/swc-storage/bwrbvsbwrgzz8ljgvsic.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406496/swc-storage/htuygcmdownqvfpzyig7.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406497/swc-storage/j2lwwhfhwvntydoizwij.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406498/swc-storage/bkzxusjmufq9vnjro9na.webp', 'Từ 11.999.000đ', 'Thực sự mạnh mẽ. Thực sự giá trị.', 'iPhone SE', 1, 1, '2024-10-10 14:39:38.656298', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "6.71", "display": "Super Retina XDR"}', 'active');
INSERT INTO public.products (id, image, shop_image, price, description, name, supplier_id, category_id, created, specs, status) VALUES (6, 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729409168/swc-storage/miankeryqfd8d6orenf2.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635169/swc-storage/mreuudhtmybm0pjluspy.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635170/swc-storage/rqeksvzxoqy5nemujwme.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635171/swc-storage/anaby1wfox1zccuahpi0.jpg', 'Từ 28.999.000đ', 'Một iPhone cực đỉnh.', 'iPhone 15 Pro Max', 1, 1, '2024-10-11 08:22:32.021943', '{"RAM": [4, 8], "SSD": [128, 256, 512], "screen": "Super Retina XDR", "display": "6.71"}', 'active');


--
-- TOC entry 3494 (class 0 OID 22077)
-- Dependencies: 243
-- Data for Name: inventories; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (5, 1, 43999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615248/swc-storage/b5pipnb8pmkdmxweys0z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615249/swc-storage/jaysghjdmyace5fxxta8.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615250/swc-storage/utcd96odzu8dloe24085.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615251/swc-storage/k7vs30bg9b3hd4geguxr.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (4, 1, 37999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (1, 1, 28999000.0000, 'active', 'VND', 17, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (6, 1, 28999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (2, 1, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (3, 1, 31999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728613819/swc-storage/azovkmoamlblc12w9qcn.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612372/swc-storage/kqi4sd9vkepbq1g9cbkr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612373/swc-storage/oeqsjzeepwartkbpbxk5.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612374/swc-storage/mkszige0ic5xqqijasbz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728612375/swc-storage/yynzgungzxiccc93ycbg.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (7, 1, 31999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (8, 1, 37999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (24, 6, 37999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (25, 6, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (9, 1, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615654/swc-storage/durp3pzj2byfydsmoly1.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615734/swc-storage/ugpupsbeprvrmgqwqv4z.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615735/swc-storage/j9l2mnnwhk7doexv3uoz.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615736/swc-storage/uwl3nvt98vnvapokeiad.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728615737/swc-storage/sdlc32yyo2nxdk5llfqz.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (10, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (11, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (12, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (13, 6, 28999000.0000, 'active', 'VND', 1234, 'Nature Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635362/swc-storage/gxezo9nlkmmn92phb7lz.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635632/swc-storage/g6gc6f0lv87wzraxg3ln.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635633/swc-storage/bd31l9axivuampkoc7qh.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635634/swc-storage/ybd0ckqnutfl7pgblhcw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728635635/swc-storage/pb8s7rf8ak9udufrwlbj.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (18, 6, 28999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (19, 6, 31999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (20, 6, 37999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (14, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (15, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (16, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (17, 6, 28999000.0000, 'active', 'VND', 1234, 'Blue Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637270/swc-storage/xwvqintf96zjrxtznlzt.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637312/swc-storage/q07sqroceuxlhl6x9flc.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637313/swc-storage/xpod1mhem1suzwpklm8d.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637314/swc-storage/cpntkt5ofxermu54kigx.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728637315/swc-storage/kvqwmjfop9t4bvxgkvtr.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (21, 6, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696337/swc-storage/venyokfppmi2vq9thxgj.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696452/swc-storage/naz0mksrntlmwkqnfnrg.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696453/swc-storage/qycmsy2faakb8baylia6.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696454/swc-storage/xz8gkq3f6wfzsptfdegi.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696455/swc-storage/kogtl9jfuujlf786jsmp.jpg', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (22, 6, 28999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (23, 6, 31999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696770/swc-storage/l2er9ptyweeoxyx2zqmu.jpg', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696813/swc-storage/nzbfah33w2emvcqxzclr.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696814/swc-storage/xkdlsfuebn0v2o7banbw.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696815/swc-storage/c24zwxmhl3lxoqxqqkkd.jpg,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1728696817/swc-storage/cygibdl6tnjujb77mhbc.jpg', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (27, 2, 43999000.0000, 'active', 'VND', 1234, 'Nature Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (28, 2, 43999000.0000, 'active', 'VND', 1234, 'Blue Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (29, 2, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (30, 2, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (31, 3, 43999000.0000, 'active', 'VND', 1234, 'Natural Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (32, 3, 43999000.0000, 'active', 'VND', 1234, 'Blue Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (33, 3, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (34, 3, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (35, 4, 43999000.0000, 'active', 'VND', 1234, 'Natural Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (36, 4, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (37, 4, 43999000.0000, 'active', 'VND', 1234, 'Blue Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (38, 4, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', '', '', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (41, 7, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406333/swc-storage/jahmb7xvr1lmlrxu4mb0.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406300/swc-storage/xunxqmbkbgljqo3egkvl.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/aqcwxccfqgj9trkotlhx.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/xl65e1obpnfyiwhgv8cf.webp', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (42, 7, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406414/swc-storage/svnxrjvater7s66fhzhp.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406459/swc-storage/ibiac6n5xzcgytj5wbkg.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/rwz9yvsjvx0kfulz9vl8.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/fjvq65jov1qscrklowtq.webp', '{"ram": "8GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (51, 12, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407183/swc-storage/xkeszwa78sanhffb5qxn.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407302/swc-storage/oj80auuxnps84jjndtud.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407303/swc-storage/xn1k4c55mlud6ermd7j3.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407304/swc-storage/fnodmozhd7p1yz0cn4zq.webp', '{"ram": "8GB", "ssd": "126GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (52, 12, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407183/swc-storage/xkeszwa78sanhffb5qxn.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407302/swc-storage/oj80auuxnps84jjndtud.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407303/swc-storage/xn1k4c55mlud6ermd7j3.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407304/swc-storage/fnodmozhd7p1yz0cn4zq.webp', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (49, 12, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407212/swc-storage/dgm8drp45suziwhig4h9.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407275/swc-storage/eivxrcforrcmxdnjgtnd.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/pr47i07a1qk1ww1xdwpu.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/asytiwipjuejynjrylem.webp', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (55, 12, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407212/swc-storage/dgm8drp45suziwhig4h9.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407275/swc-storage/eivxrcforrcmxdnjgtnd.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/pr47i07a1qk1ww1xdwpu.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/asytiwipjuejynjrylem.webp', '{"ram": "8GB", "ssd": "126GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (43, 7, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406414/swc-storage/svnxrjvater7s66fhzhp.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406459/swc-storage/ibiac6n5xzcgytj5wbkg.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/rwz9yvsjvx0kfulz9vl8.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/fjvq65jov1qscrklowtq.webp', '{"ram": "16GB", "ssd": "128GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (44, 7, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406414/swc-storage/svnxrjvater7s66fhzhp.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406459/swc-storage/ibiac6n5xzcgytj5wbkg.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/rwz9yvsjvx0kfulz9vl8.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/fjvq65jov1qscrklowtq.webp', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (45, 7, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406414/swc-storage/svnxrjvater7s66fhzhp.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406459/swc-storage/ibiac6n5xzcgytj5wbkg.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/rwz9yvsjvx0kfulz9vl8.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406460/swc-storage/fjvq65jov1qscrklowtq.webp', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (46, 7, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406333/swc-storage/jahmb7xvr1lmlrxu4mb0.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406300/swc-storage/xunxqmbkbgljqo3egkvl.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/aqcwxccfqgj9trkotlhx.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/xl65e1obpnfyiwhgv8cf.webp', '{"ram": "8GB", "ssd": "256TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (48, 7, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406333/swc-storage/jahmb7xvr1lmlrxu4mb0.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406300/swc-storage/xunxqmbkbgljqo3egkvl.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/aqcwxccfqgj9trkotlhx.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/xl65e1obpnfyiwhgv8cf.webp', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (47, 7, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406333/swc-storage/jahmb7xvr1lmlrxu4mb0.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406300/swc-storage/xunxqmbkbgljqo3egkvl.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/aqcwxccfqgj9trkotlhx.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729406302/swc-storage/xl65e1obpnfyiwhgv8cf.webp', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (56, 12, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407212/swc-storage/dgm8drp45suziwhig4h9.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407275/swc-storage/eivxrcforrcmxdnjgtnd.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/pr47i07a1qk1ww1xdwpu.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/asytiwipjuejynjrylem.webp', '{"ram": "16GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (57, 12, 43999000.0000, 'active', 'VND', 1234, 'Black Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407212/swc-storage/dgm8drp45suziwhig4h9.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407275/swc-storage/eivxrcforrcmxdnjgtnd.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/pr47i07a1qk1ww1xdwpu.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407276/swc-storage/asytiwipjuejynjrylem.webp', '{"ram": "16GB", "ssd": "512GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (53, 12, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407183/swc-storage/xkeszwa78sanhffb5qxn.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407302/swc-storage/oj80auuxnps84jjndtud.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407303/swc-storage/xn1k4c55mlud6ermd7j3.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407304/swc-storage/fnodmozhd7p1yz0cn4zq.webp', '{"ram": "8GB", "ssd": "256GB", "desc": "", "connection": ""}');
INSERT INTO public.inventories (id, product_id, price, status, currency_code, available, color, color_img, image, specs) VALUES (54, 12, 43999000.0000, 'active', 'VND', 1234, 'White Titanium', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407183/swc-storage/xkeszwa78sanhffb5qxn.png', 'https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407302/swc-storage/oj80auuxnps84jjndtud.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407303/swc-storage/xn1k4c55mlud6ermd7j3.webp,https://res.cloudinary.com/dqsiqqz7q/image/upload/v1729407304/swc-storage/fnodmozhd7p1yz0cn4zq.webp', '{"ram": "16GB", "ssd": "1TB", "desc": "", "connection": ""}');


--
-- TOC entry 3484 (class 0 OID 22035)
-- Dependencies: 233
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.carts (id, user_id, inventory_id, quantity) VALUES (1, 1, 24, 1);
INSERT INTO public.carts (id, user_id, inventory_id, quantity) VALUES (2, 1, 20, 1);

--
-- TOC entry 3496 (class 0 OID 22086)
-- Dependencies: 245
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.news (id, created, "category", header, body) VALUES (1, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/1.jpg", "title": "You can do more with AI.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Artificial Intelligence"}');
INSERT INTO public.news (id, created, "category", header, body) VALUES (2, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/2.jpg", "title": "Enhance your productivity.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Productivity"}');
INSERT INTO public.news (id, created, "category", header, body) VALUES (3, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/3.jpg", "title": "Launching the new Apple Vision Pro.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Product"}');
INSERT INTO public.news (id, created, "category", header, body) VALUES (4, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/4.jpg", "title": "Maps for your iPhone 15 Pro Max.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Product"}');
INSERT INTO public.news (id, created, "category", header, body) VALUES (5, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/5.jpg", "title": "Photography just got better.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "iOS"}');
INSERT INTO public.news (id, created, "category", header, body) VALUES (6, '2024-08-03 17:26:15.088919', 'phone', 'Get to know your iPhone.', '{"src": "/img/posts/6.jpg", "title": "Hiring for a Staff Software Engineer.", "content": [{"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}, {"src": "/img/posts/8.jpg", "content": "Keep a journal, quickly jot down a grocery list, and take amazing class notes. Want to convert those notes to text? No problem. Langotiya jeetu ka mara hua yaar is ready to capture every thought."}], "category": "Hiring"}');

SELECT pg_catalog.setval('public.news_id_seq', 6, true);


--
-- TOC entry 3490 (class 0 OID 22061)
-- Dependencies: 239
-- Data for Name: deliveries; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.deliveries (id, user_id, address_id, sent_date, status, method, note) VALUES (1, 1, 1, '0001-01-01 00:00:00+00', 'active', 'COD', '');
INSERT INTO public.deliveries (id, user_id, address_id, sent_date, status, method, note) VALUES (2, 2, 2, '2024-10-21 03:28:37.489141+00', 'active', 'COD', '');
INSERT INTO public.deliveries (id, user_id, address_id, sent_date, status, method, note) VALUES (4, 2, 4, '2024-10-21 03:42:07.28624+00', 'active', 'COD', '');
INSERT INTO public.deliveries (id, user_id, address_id, sent_date, status, method, note) VALUES (13, 2, 13, '2024-10-21 05:43:30.692936+00', 'active', 'COD', '');


--
-- TOC entry 3492 (class 0 OID 22070)
-- Dependencies: 241
-- Data for Name: favorite; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- TOC entry 3486 (class 0 OID 22044)
-- Dependencies: 235
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.orders (id, uuid, "time", user_id, delivery_id, total_amount, status, payment_method) VALUES (1, 'Q9DVXAG7O1JY3WDJ', '2024-10-21 03:28:37.479607', 2, 2, 28999000.0000, 'active', 'COD');
INSERT INTO public.orders (id, uuid, "time", user_id, delivery_id, total_amount, status, payment_method) VALUES (2, 'JBI81XHSWAUSY4MB', '2024-10-21 03:42:07.284393', 2, 4, 43999000.0000, 'active', 'COD');
INSERT INTO public.orders (id, uuid, "time", user_id, delivery_id, total_amount, status, payment_method) VALUES (3, 'Y8KTLRT64A2HXT0R', '2024-10-21 05:43:30.685709', 2, 13, 87998000.0000, 'active', 'COD');


--
-- TOC entry 3472 (class 0 OID 21977)
-- Dependencies: 221
-- Data for Name: payments; Type: TABLE DATA; Schema: public; Owner: admin
--



--
-- TOC entry 3488 (class 0 OID 22054)
-- Dependencies: 237
-- Data for Name: product_in_order; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.product_in_order (id, order_id, inventory_id, quantity, total_amount, currency_code) VALUES (1, 1, 17, 1, 28999000.0000, 'VND');
INSERT INTO public.product_in_order (id, order_id, inventory_id, quantity, total_amount, currency_code) VALUES (2, 2, 56, 1, 43999000.0000, 'VND');
INSERT INTO public.product_in_order (id, order_id, inventory_id, quantity, total_amount, currency_code) VALUES (3, 3, 56, 1, 43999000.0000, 'VND');
INSERT INTO public.product_in_order (id, order_id, inventory_id, quantity, total_amount, currency_code) VALUES (4, 3, 25, 1, 43999000.0000, 'VND');


--
-- TOC entry 3502 (class 0 OID 0)
-- Dependencies: 216
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.accounts_id_seq', 1, true);


--
-- TOC entry 3503 (class 0 OID 0)
-- Dependencies: 222
-- Name: addresses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.addresses_id_seq', 13, true);


--
-- TOC entry 3504 (class 0 OID 0)
-- Dependencies: 232
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.carts_id_seq', 2, true);


--
-- TOC entry 3505 (class 0 OID 0)
-- Dependencies: 230
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.categories_id_seq', 8, true);



--
-- TOC entry 3508 (class 0 OID 0)
-- Dependencies: 238
-- Name: deliveries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.deliveries_id_seq', 13, true);


--
-- TOC entry 3509 (class 0 OID 0)
-- Dependencies: 240
-- Name: favorite_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.favorite_id_seq', 1, false);


--
-- TOC entry 3510 (class 0 OID 0)
-- Dependencies: 242
-- Name: inventories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.inventories_id_seq', 57, true);


--
-- TOC entry 3511 (class 0 OID 0)
-- Dependencies: 234
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.orders_id_seq', 3, true);


--
-- TOC entry 3513 (class 0 OID 0)
-- Dependencies: 236
-- Name: product_in_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.product_in_order_id_seq', 4, true);


--
-- TOC entry 3514 (class 0 OID 0)
-- Dependencies: 226
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.products_id_seq', 31, true);


--
-- TOC entry 3515 (class 0 OID 0)
-- Dependencies: 228
-- Name: suppliers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.suppliers_id_seq', 1, true);


--
-- TOC entry 3516 (class 0 OID 0)
-- Dependencies: 218
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.users_id_seq', 2, true);


-- Completed on 2024-10-23 10:29:24

--
-- PostgreSQL database dump complete
--

