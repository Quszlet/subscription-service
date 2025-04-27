INSERT INTO users (login, hash_password, email, date_birthday, created_at) VALUES
('alice', 'hashed_password_1', 'alice@example.com', '1990-01-15', NOW()),
('bob', 'hashed_password_2', 'bob@example.com', '1988-06-20', NOW()),
('charlie', 'hashed_password_3', 'charlie@example.com', '1995-03-10', NOW()),
('diana', 'hashed_password_4', 'diana@example.com', '1992-12-01', NOW()),
('eve', 'hashed_password_5', 'eve@example.com', '1985-09-05', NOW());


INSERT INTO subscriptions (owner_id, name, amount, payment_date, payment_interval, status, type, days_notify_before_payment, created_at, updated_at, image, color, invate_token) VALUES
(1, 'Netflix', 9.99, '2025-05-01', 30, 'active', 'individual', 5, NOW(), NOW(), NULL, 'red', 'token1'),
(2, 'Spotify', 4.99, '2025-05-05', 30, 'active', 'individual', 3, NOW(), NOW(), NULL, 'green', 'token2'),
(3, 'Gym Membership', 29.99, '2025-05-10', 30, 'inactive', 'individual', 7, NOW(), NOW(), NULL, 'blue', 'token3'),
(4, 'Family Amazon Prime', 12.99, '2025-05-03', 30, 'active', 'group', 2, NOW(), NOW(), NULL, 'yellow', 'token4'),
(5, 'Disney+', 7.99, '2025-05-15', 30, 'deleted', 'individual', 1, NOW(), NOW(), NULL, 'purple', 'token5');

INSERT INTO subscription_members (user_id, subscription_id, amount, joined_at, days_notify_before_payment, color, payment_status) VALUES
(1, 1, 9.99, NOW(), 5, 'red', 'paid'),
(2, 2, 4.99, NOW(), 3, 'green', 'not paid'),
(3, 3, 29.99, NOW(), 7, 'blue', 'paid'),
(4, 4, 6.50, NOW(), 2, 'yellow', 'paid'), -- Делим Family подписку на 2 участников
(5, 4, 6.49, NOW(), 2, 'yellow', 'not paid');
