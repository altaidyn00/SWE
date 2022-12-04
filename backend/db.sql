CREATE DATABASE Hospital;

USE Hospital;

  create TABLE department (
    id int,
    descript varchar(30),
    Primary key(id)
  );

  create TABLE specialization (
    id int,
    descript varchar(30),
    Primary key(id)
  );



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
    department_id int,
    foreign key (department_id) REFERENCES department(id),
    specialization_id int,
    foreign key (specialization_id) REFERENCES specialization(id),

    experience_in_years INTEGER,
    appointment_price INTEGER,
    education_degree VARCHAR(30),
    schedule_detail VARCHAR(60),
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
  doctor_id INT,
foreign key (doctor_id) REFERENCES doctor(government_id),
  preferred_date VARCHAR(20),
  preferred_time VARCHAR(20),
  name VARCHAR(20),
  surname varchar(20),
  email varchar(30),
  specialization_id int,
Primary key (email, doctor_id)
);


INSERT into department value (1, "main");
INSERT into department value (2, "second");
INSERT into department value (3, "third");


INSERT into users value(0, 'admin', 'pass1', 'user1','','admin@med.com');
INSERT into specialization value (1, "cardiology");
INSERT into specialization value (2, "brain surgeon");
INSERT into specialization value (3, "dermatiologist");
INSERT into specialization value (4, "Endocrinologists");
INSERT into specialization value (5, "Gastroenterologists");
INSERT into specialization value (6, "Infectious Disease Specialists");
INSERT into specialization value (7, "Oncologists");
INSERT into specialization value (8, "Rheumatologists");
INSERT into specialization value (9, "Urologists");
INSERT into specialization value (10, "Physiatrists");
