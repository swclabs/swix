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

