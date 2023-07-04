CREATE TABLE ms_fertilizers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	stock int, 
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
);

CREATE TABLE ms_farmers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	address VARCHAR(50),
	phone_number VARCHAR(20),
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
);

CREATE TABLE ms_plants (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50),
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
);

CREATE TABLE ms_fertilizer_prices (
	id SERIAL PRIMARY KEY,
	fertilizer_id int,
	price int,
	is_active BOOL,
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
);

CREATE TABLE tr_bills (
	id SERIAL PRIMARY KEY,
	farmer_id int,
	"date" DATE,
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
	
);

CREATE TABLE tr_bill_details (
	id SERIAL PRIMARY KEY,
	bill_id int,
	fertilizer_price_id int,
	qty int, 
	create_at DATE,
	update_at DATE,
	create_by VARCHAR(50),
	update_by VARCHAR(50)
);

CREATE TABLE farmers_plants (
	plant_id int,
	farmer_id int
);