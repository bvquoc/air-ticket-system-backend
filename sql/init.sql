-- Users table
CREATE TABLE IF NOT EXISTS users (
                                     id BIGINT AUTO_INCREMENT PRIMARY KEY,
                                     email VARCHAR(255),
    name VARCHAR(255),
    password VARCHAR(255),
    user_role VARCHAR(255),
    username VARCHAR(255) UNIQUE,
    CHECK (user_role IN ('USER', 'ADMIN'))
    );

-- Indexes on users
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_username ON users (username);

-- Bookings table
CREATE TABLE IF NOT EXISTS bookings (
                                        booking_id INT AUTO_INCREMENT PRIMARY KEY,
                                        user_id BIGINT,
                                        booking_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        total_amount DECIMAL(10, 2),
    status VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

-- Payments table
CREATE TABLE IF NOT EXISTS payments (
                                        payment_id INT AUTO_INCREMENT PRIMARY KEY,
                                        booking_id INT,
                                        payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        amount DECIMAL(10, 2),
    payment_method VARCHAR(50),
    status VARCHAR(50),
    FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE
    );

-- Index on payments
CREATE INDEX idx_booking_id ON payments (booking_id);

-- Direction Bookings table
CREATE TABLE IF NOT EXISTS direction_bookings (
                                                  direction_booking_id INT AUTO_INCREMENT PRIMARY KEY,
                                                  booking_id INT,
                                                  origin VARCHAR(100),
    destination VARCHAR(100),
    departure_date TIMESTAMP,
    arrival_date TIMESTAMP,
    FOREIGN KEY (booking_id) REFERENCES bookings(booking_id) ON DELETE CASCADE
    );

-- Passengers table
CREATE TABLE IF NOT EXISTS passengers (
                                          passenger_id INT AUTO_INCREMENT PRIMARY KEY,
                                          first_name VARCHAR(100),
    last_name VARCHAR(100),
    date_of_birth DATE,
    passport_number VARCHAR(50),
    nationality VARCHAR(50)
    );

-- Aircrafts table
CREATE TABLE IF NOT EXISTS aircrafts (
                                         aircraft_id INT AUTO_INCREMENT PRIMARY KEY,
                                         aircraft_model VARCHAR(100),
    manufacturer VARCHAR(100),
    seating_capacity INT
    );

-- Flights table
CREATE TABLE IF NOT EXISTS flights (
                                       flight_id INT AUTO_INCREMENT PRIMARY KEY,
                                       flight_number VARCHAR(50) UNIQUE,
    departure_time TIMESTAMP,
    arrival_time TIMESTAMP,
    origin VARCHAR(100),
    destination VARCHAR(100),
    aircraft_id INT,
    FOREIGN KEY (aircraft_id) REFERENCES aircrafts(aircraft_id) ON DELETE CASCADE
    );

-- Seats table
CREATE TABLE IF NOT EXISTS seats (
                                     seat_id INT AUTO_INCREMENT PRIMARY KEY,
                                     aircraft_id INT,
                                     seat_number VARCHAR(10),
    seat_type VARCHAR(50),
    is_available BOOLEAN DEFAULT TRUE,
    FOREIGN KEY (aircraft_id) REFERENCES aircrafts(aircraft_id) ON DELETE CASCADE
    );

-- Tickets table
CREATE TABLE IF NOT EXISTS tickets (
                                       ticket_id INT AUTO_INCREMENT PRIMARY KEY,
                                       direction_booking_id INT,
                                       flight_id INT,
                                       seat_id INT,
                                       passenger_id INT,
                                       ticket_status VARCHAR(50),
    FOREIGN KEY (direction_booking_id) REFERENCES direction_bookings(direction_booking_id) ON DELETE CASCADE,
    FOREIGN KEY (flight_id) REFERENCES flights(flight_id) ON DELETE CASCADE,
    FOREIGN KEY (seat_id) REFERENCES seats(seat_id) ON DELETE CASCADE,
    FOREIGN KEY (passenger_id) REFERENCES passengers(passenger_id) ON DELETE CASCADE
    );

-- Indexes on tickets
CREATE INDEX idx_flight_id ON tickets (flight_id);
CREATE INDEX idx_seat_id ON tickets (seat_id);
CREATE INDEX idx_passenger_id ON tickets (passenger_id);
