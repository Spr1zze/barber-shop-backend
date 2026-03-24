-- Insert test salons
INSERT INTO salons (slug, name, address, description, hero_image_url, phone, email) VALUES
    (
        'downtown-hair',
        'Downtown Hair',
        'Nørregade 14, 1165 København K',
        'Downtown Hair er en moderne barbersalon i indre by med fokus på præcise klip, rolige omgivelser og en enkel oplevelse fra booking til færdigt resultat.',
        'https://images.unsplash.com/photo-1621605815971-fbc98d665033?w=1200&q=80',
        '+45 31 23 45 67',
        'hej@downtownhair.dk'
    ),
    (
        'sharp-styles',
        'Sharp Styles',
        'Østerbrogade 42, 2100 København Ø',
        'Sharp Styles er en lokal salon med fokus på klassiske klip, skægtrim og hurtig booking.',
        'https://images.unsplash.com/photo-1517832606299-7ae9b720a186?w=1200&q=80',
        '+45 20 30 40 50',
        'hej@sharpstyles.dk'
    );

-- Insert opening hours
INSERT INTO salon_opening_hours (salon_id, day_name, day_order, open_time, close_time, is_closed) VALUES
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Mandag', 1, '09:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Tirsdag', 2, '09:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Onsdag', 3, '09:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Torsdag', 4, '09:00', '19:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Fredag', 5, '09:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Lørdag', 6, '10:00', '15:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'Søndag', 7, NULL, NULL, TRUE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Mandag', 1, '10:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Tirsdag', 2, '10:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Onsdag', 3, '10:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Torsdag', 4, '10:00', '19:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Fredag', 5, '10:00', '18:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Lørdag', 6, '10:00', '14:00', FALSE),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'Søndag', 7, NULL, NULL, TRUE);

-- Insert test barbers
INSERT INTO barbers (name, salon_id) VALUES
    ('Mads Nielsen', (SELECT id FROM salons WHERE slug = 'downtown-hair')),
    ('Jonas Holm', (SELECT id FROM salons WHERE slug = 'downtown-hair')),
    ('Sara Jensen', (SELECT id FROM salons WHERE slug = 'sharp-styles'));

-- Insert test services
INSERT INTO services (salon_id, slug, name, duration_minutes, price_from, display_order) VALUES
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'herreklip', 'HerreKlip', 30, 220, 1),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'skin-fade', 'Skin Fade', 45, 300, 2),
    ((SELECT id FROM salons WHERE slug = 'downtown-hair'), 'skaeg-trim', 'Skæg Trim', 20, 160, 3),
    ((SELECT id FROM salons WHERE slug = 'sharp-styles'), 'classic-cut', 'Classic Cut', 30, 210, 1);

-- Insert test bookings
INSERT INTO bookings (barber_id, service_id, salon_id, start_time, customer_name, phone) VALUES
    ((SELECT id FROM barbers WHERE name = 'Mads Nielsen'),
     (SELECT id FROM services WHERE slug = 'herreklip' AND salon_id = (SELECT id FROM salons WHERE slug = 'downtown-hair')),
     (SELECT id FROM salons WHERE slug = 'downtown-hair'),
     '2026-03-25 14:00',
     'Alice Johnson',
     '555-1234');
