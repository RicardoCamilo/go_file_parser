CREATE TABLE IF NOT EXISTS phone_calls(
                                          id SERIAL PRIMARY KEY,
                                          call_date_time TIMESTAMP NOT NULL,
                                          disposition TEXT NOT NULL,
                                          phone_number TEXT NOT NULL,
                                          first_name VARCHAR(100),
                                          last_name VARCHAR(100),
                                          address1 TEXT,
                                          address2 TEXT,
                                          city TEXT,
                                          state TEXT,
                                          zip TEXT
);
