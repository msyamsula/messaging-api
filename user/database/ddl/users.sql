create table users (
id serial primary key,
username varchar(100) not null unique constraint namecheck check(char_length(username) >= 5),
"password" varchar(100) not null constraint passwordcheck check(char_length("password") >= 5),
"created_at" timestamp default now() not null,
"updated_at" timestamp default now() not null,
"is_active" bool default false not null
);


CREATE OR REPLACE FUNCTION update_timestamp()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';

create trigger update_users_updated_at before update on users for each row execute procedure update_timestamp();