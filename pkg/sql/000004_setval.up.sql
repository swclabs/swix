SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));
SELECT setval('suppliers_id_seq', (SELECT MAX(id) FROM suppliers));
SELECT setval('inventories_id_seq', (SELECT MAX(id) FROM inventories));
SELECT setval('collections_id_seq', (SELECT MAX(id) FROM collections));
SELECT setval('products_id_seq', (SELECT MAX(id) FROM products));
SELECT setval('accounts_id_seq', (SELECT MAX(id) FROM accounts));   