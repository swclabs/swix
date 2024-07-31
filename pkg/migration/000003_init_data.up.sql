INSERT INTO users (id, email, phone_number, first_name, last_name, image) 
VALUES (1, 'sa@sa.com', '123', 'sa', 'sa', '')
ON CONFLICT (email) DO NOTHING;


INSERT INTO accounts (username, role, email, password, type, created_at)
VALUES ('user#1', 'Admin', 'sa@sa.com', '$2a$12$x9jSMLNRSTX.kOmLThUFH.VanhMBTUJrFxUwNOkx7GOLTGrMWl5Kq', 'swc', '2024-07-31 07:51:50+00')
ON CONFLICT (username) DO NOTHING;

