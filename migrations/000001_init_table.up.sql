-- Enable extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =========================
-- SALONS
-- =========================
CREATE TABLE IF NOT EXISTS salons (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    postal_code VARCHAR(10) NOT NULL,
    city VARCHAR(100) NOT NULL,
    region VARCHAR(100) NOT NULL,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_salons_city ON salons(city);


-- =========================
-- BARBERS
-- =========================
CREATE TABLE IF NOT EXISTS barbers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- =========================
-- BARBER <-> SALON (M:N)
-- =========================
CREATE TABLE IF NOT EXISTS barber_salons (
    barber_id UUID NOT NULL,
    salon_id UUID NOT NULL,
    PRIMARY KEY (barber_id, salon_id),
    FOREIGN KEY (barber_id) REFERENCES barbers(id) ON DELETE CASCADE,
    FOREIGN KEY (salon_id) REFERENCES salons(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_barber_salons_barber ON barber_salons(barber_id);
CREATE INDEX IF NOT EXISTS idx_barber_salons_salon ON barber_salons(salon_id);


-- =========================
-- SERVICES
-- =========================
CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL CHECK (price >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- =========================
-- BOOKINGS
-- =========================
CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    barber_id UUID NOT NULL,
    service_id UUID NOT NULL,
    salon_id UUID NOT NULL,

    date_time TIMESTAMP NOT NULL,

    customer_name VARCHAR(255) NOT NULL,
    phone VARCHAR(30) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (barber_id) REFERENCES barbers(id) ON DELETE CASCADE,
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE,
    FOREIGN KEY (salon_id) REFERENCES salons(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_bookings_barber ON bookings(barber_id);
CREATE INDEX IF NOT EXISTS idx_bookings_salon ON bookings(salon_id);
CREATE INDEX IF NOT EXISTS idx_bookings_datetime ON bookings(date_time);