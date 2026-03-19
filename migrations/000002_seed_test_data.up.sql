-- Insert test salons
INSERT INTO salons (salonName, address) VALUES 
    ('Cool Cuts', '123 Main St'),
    ('Sharp Styles', '456 Oak Ave');

-- Insert test barbers
INSERT INTO barbers (barberName, salon_id) VALUES 
    ('John Smith', (SELECT id FROM salons WHERE salonName = 'Cool Cuts')),
    ('Jane Doe', (SELECT id FROM salons WHERE salonName = 'Sharp Styles'));

-- Insert test services
INSERT INTO services (name, price) VALUES 
    ('Haircut', 30),
    ('Beard Trim', 15),
    ('Hair Color', 60);

-- Insert test bookings
INSERT INTO bookings (barber_id, service_id, salon_id, dateTime, customer_name, phone) VALUES 
    ((SELECT id FROM barbers WHERE barberName = 'John Smith'),
     (SELECT id FROM services WHERE name = 'Haircut'),
     (SELECT id FROM salons WHERE salonName = 'Cool Cuts'),
     '2026-03-25 14:00',
     'Alice Johnson',
     '555-1234');