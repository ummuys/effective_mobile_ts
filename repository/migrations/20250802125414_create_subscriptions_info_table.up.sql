CREATE TABLE subscriptions.info (
    service_name TEXT NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID UNIQUE NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL 
);
