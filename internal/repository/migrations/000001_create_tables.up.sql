-- Типы перечислений
CREATE TYPE status AS ENUM ('active', 'inactive', 'deleted');
CREATE TYPE subscription_type AS ENUM ('individual', 'group');
CREATE TYPE payment_status AS ENUM ('paid', 'not paid');


CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    hash_password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    date_birthday DATE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    payment_date DATE NOT NULL,
    payment_interval INTEGER NOT NULL,
    status status NOT NULL DEFAULT 'active',
    type subscription_type NOT NULL DEFAULT 'individual',
    days_notify_before_payment INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    image BYTEA,
    color VARCHAR(50),
    invate_token VARCHAR(255)
);

CREATE TABLE  IF NOT EXISTS subscription_members (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    subscription_id INTEGER NOT NULL REFERENCES subscriptions(id) ON DELETE CASCADE,
    amount NUMERIC(10, 2) NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW(),
    days_notify_before_payment INTEGER NOT NULL DEFAULT 0,
    color VARCHAR(50),
    payment_status payment_status NOT NULL DEFAULT 'not paid'
);
