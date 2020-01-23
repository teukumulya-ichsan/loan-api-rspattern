CREATE TABLE pinjaman
(
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    date_loan TIMESTAMP
    WITH TIME ZONE, gender VARCHAR
    (10) NOT NULL, KTP VARCHAR
    (16) NOT NULL, birthdate DATE NOT NULL DEFAULT CURRENT_DATE, amount NUMERIC NOT NULL, period integer);


    -------------------------------------------------------------------------------------------
    INSERT INTO pinjaman
        (date_loan, name, gender, ktp, birthdate, amount, period)
    values('2018-09-01', 'Wiro', 'Male', '3522582509010002', '2001-09-25', 2000000, 12);

    INSERT INTO pinjaman
        (date_loan, name, gender, ktp, birthdate, amount, period)
    values('2020-01-01', 'Ichsan', 'Male', '3522582509010004', '1995-06-02', 20000000, 12);