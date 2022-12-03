CREATE DATABASE Hospital;

USE Hospital;

CREATE TABLE users(
  government_id INT,
  role VARCHAR(20),
  password VARCHAR(40),
  first_name varchar(50),
  last_name VARCHAR(50),
  email VARCHAR(50),
  PRIMARY KEY (government_id)
);
CREATE TABLE patient(
    iin VARCHAR(20),
  date_of_birth VARCHAR(20),
  government_id INT,
FOREIGN KEY (government_id) REFERENCES users(government_id), 
PRIMARY KEY (government_id),
  blood_group VARCHAR(10),
  emergency_contact_number VARCHAR(20),
  contact_number VARCHAR(20),
  address VARCHAR(60),
  marital_status VARCHAR(20),
  registration_date VARCHAR(20)
);
CREATE TABLE doctor(
    date_of_birth VARCHAR(20),
    iin VARCHAR(20),
    government_id INT,
Foreign key(government_id) REFERENCES users(government_id),
PRIMARY KEY (government_id),
    contact_number VARCHAR(15),
    number_of_patients INTEGER,
    category VARCHAR(20),
    department_id VARCHAR(20),
    specialization_details_id VARCHAR(100),
    experience_in_years INTEGER,
    appointment_price INTEGER,
    education_degree VARCHAR(30),
    schedule_detaile VARCHAR(60),
    rating INTEGER,
    address VARCHAR(60),
    photolocation varchar(140)
);
CREATE TABLE administration(
  government_id INT,
FOREIGN KEY (government_id) REFERENCES users(government_id),
 PRIMARY KEY (government_id)
);

CREATE TABLE appointment(
  patient_id INTEGER ,
Foreign key (patient_id) REFERENCES patient(government_id),
  doctor_id INT,
foreign key (doctor_id) REFERENCES doctor(government_id),
  preferred_date VARCHAR(20),
  prefered_time VARCHAR(20),
Primary key (patient_id, doctor_id)
);

INSERT into users value(0, 'admin', 'pass1', 'user1','','admin@med.com');