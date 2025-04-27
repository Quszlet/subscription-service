DELETE FROM subscription_members;
DELETE FROM subscriptions;
DELETE FROM users;


ALTER SEQUENCE subscription_members_id_seq RESTART WITH 1;
ALTER SEQUENCE subscriptions_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
