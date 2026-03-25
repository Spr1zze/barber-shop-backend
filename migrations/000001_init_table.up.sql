CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- =========================
-- SALONS
-- =========================
CREATE TABLE IF NOT EXISTS salons (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug          VARCHAR(255) NOT NULL UNIQUE,
    name          VARCHAR(255) NOT NULL,
    address       VARCHAR(255) NOT NULL,
    postal_code   VARCHAR(10)  NOT NULL,
    city          VARCHAR(100) NOT NULL,
    region        VARCHAR(100) NOT NULL,
    latitude      DECIMAL(9,6),
    longitude     DECIMAL(9,6),
    description   TEXT,
    hero_image_url TEXT,
    phone         VARCHAR(30),
    email         VARCHAR(255),
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_salons_city ON salons(city);


-- =========================
-- SALON OPENING HOURS
-- =========================
CREATE TABLE IF NOT EXISTS salon_opening_hours (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    salon_id   UUID NOT NULL REFERENCES salons(id) ON DELETE CASCADE,
    day_name   VARCHAR(20) NOT NULL,
    day_order  SMALLINT    NOT NULL CHECK (day_order BETWEEN 1 AND 7),
    open_time  TIME,
    close_time TIME,
    is_closed  BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE (salon_id, day_order)
);


-- =========================
-- BARBERS
-- =========================
CREATE TABLE IF NOT EXISTS barbers (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- =========================
-- BARBER <-> SALON (M:N)
-- =========================
CREATE TABLE IF NOT EXISTS barber_salons (
    barber_id UUID NOT NULL REFERENCES barbers(id) ON DELETE CASCADE,
    salon_id  UUID NOT NULL REFERENCES salons(id)  ON DELETE CASCADE,
    PRIMARY KEY (barber_id, salon_id)
);

CREATE INDEX IF NOT EXISTS idx_barber_salons_barber ON barber_salons(barber_id);
CREATE INDEX IF NOT EXISTS idx_barber_salons_salon  ON barber_salons(salon_id);


-- =========================
-- SERVICES
-- =========================
CREATE TABLE IF NOT EXISTS services (
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    salon_id         UUID NOT NULL REFERENCES salons(id) ON DELETE CASCADE,
    slug             VARCHAR(255) NOT NULL,
    name             VARCHAR(255) NOT NULL,
    duration_minutes INTEGER NOT NULL,
    price            INTEGER NOT NULL CHECK (price >= 0),
    display_order    INTEGER NOT NULL DEFAULT 0,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (salon_id, slug)
);


-- =========================
-- BOOKINGS
-- =========================
CREATE TABLE IF NOT EXISTS bookings (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    barber_id     UUID NOT NULL REFERENCES barbers(id)  ON DELETE CASCADE,
    service_id    UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    salon_id      UUID NOT NULL REFERENCES salons(id)   ON DELETE CASCADE,
    date_time     TIMESTAMP NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    phone         VARCHAR(30)  NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_bookings_barber   ON bookings(barber_id);
CREATE INDEX IF NOT EXISTS idx_bookings_salon    ON bookings(salon_id);
CREATE INDEX IF NOT EXISTS idx_bookings_datetime ON bookings(date_time);