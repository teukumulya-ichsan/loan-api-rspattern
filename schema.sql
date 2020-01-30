-- postgresSQL dump db

CREATE DATABASE db_loan;

CREATE TABLE loans
(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    date_loan TIMESTAMP
    WITH TIME ZONE, gender VARCHAR
    (10) NOT NULL, KTP VARCHAR
    (16) NOT NULL, birthdate DATE NOT NULL DEFAULT CURRENT_DATE, amount NUMERIC NOT NULL, period integer);


    -------------------------------------------------------------------------------------------
    INSERT INTO loans
        (date_loan, name, gender, ktp, birthdate, amount, period)
    values('2018-09-01', 'Wiro', 'Male', '3522582509010002', '2001-09-25', 2000000, 12);

    INSERT INTO loans
        (date_loan, name, gender, ktp, birthdate, amount, period)
    values('2020-01-01', 'Ichsan', 'Male', '3522582509010004', '1995-06-02', 20000000, 12);

    -------------------------------------------------------------------------------------------


    -- getting last 7 days data from current date
    select *
    from loans
    where date_loan >= NOW() - interval
    '2' day;

    select *
    from loans
    where date_loan >= current_date - interval
    '2' day;

    -- getting last 7 days data from choice date
    select *
    from loans
    where date_loan > '2020-01-01' - interval
    '7' day;


    select * from loans where date_loan between '2020-01-29'::date - integer '6' AND '2020-01-29'::date;
    


    