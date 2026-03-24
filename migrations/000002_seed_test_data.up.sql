-- =========================
-- SALONS
-- =========================
INSERT INTO salons (name, address, postal_code, city, region, latitude, longitude) VALUES
('Salon Runde Tårn', 'Frederiksgade 1', '8000', 'Aarhus', 'Jylland', 56.154, 10.203),
('Salon H.C. Andersens Skæve Hus', 'Overgade 24', '5000', 'Odense', 'Fyn', 55.399, 10.383),
('Salon Lille Klippetorv', 'Bispensgade 34', '9000', 'Aalborg', 'Jylland', 57.048, 9.920),
('Salon Klippekrogen', 'Kongensgade 52', '6700', 'Esbjerg', 'Jylland', 55.476, 8.451),
('Salon Lokkehuset', 'Sct. Mathias Gade 45', '8800', 'Viborg', 'Jylland', 56.451, 9.402),
('Salon Frisurevej', 'Algade 15', '4000', 'Roskilde', 'Sjælland', 55.641, 12.081);


-- =========================
-- BARBERS
-- =========================
INSERT INTO barbers (name) VALUES
('Klippet Kirsten'),
('Hårdtarbejdende Hans'),
('Sakse-Søren'),
('Frisure-Freja'),
('Lokke-Lars'),
('Krølle-Karen'),
('Hårde-Hanne'),
('Bølle-Bob'),
('Sakse-Sally'),
('Klippe-Kasper'),
('Lokke-Louise'),
('Krøllet-Kurt'),
('Hår-Hilde'),
('Stylet-Stine'),
('Bølget-Benny'),
('Sakse-Simone'),
('Klip-Kim'),
('Lokke-Lea'),
('Krøl-Kalle'),
('Hår-Henriette'),
('Stylet-Stefan');



INSERT INTO services (name, price) VALUES
('Striber', 350),
('Farvet hår (kort)', 450),
('Farvet hår (langt)', 650),
('Herre klip', 250),
('Dame klip', 300),
('Skæg trim', 150),
('Vask & styling', 200);


INSERT INTO bookings (
    barber_id,
    service_id,
    salon_id,
    date_time,
    customer_name,
    phone
)
VALUES
(
    (SELECT id FROM barbers WHERE name = 'Klippet Kirsten'),
    (SELECT id FROM services WHERE name = 'Herre klip'),
    (SELECT id FROM salons WHERE name = 'Salon Runde Tårn'),
    '2026-03-25 10:00:00',
    'Mads Jensen',
    '+4522334455'
),
(
    (SELECT id FROM barbers WHERE name = 'Hårdtarbejdende Hans'),
    (SELECT id FROM services WHERE name = 'Farvet hår (langt)'),
    (SELECT id FROM salons WHERE name = 'Salon Lokkehuset'),
    '2026-03-25 12:30:00',
    'Emma Nielsen',
    '+4511223344'
),
(
    (SELECT id FROM barbers WHERE name = 'Sakse-Søren'),
    (SELECT id FROM services WHERE name = 'Striber'),
    (SELECT id FROM salons WHERE name = 'Salon H.C. Andersens Skæve Hus'),
    '2026-03-26 09:15:00',
    'Lucas Pedersen',
    '+4544556677'
),
(
    (SELECT id FROM barbers WHERE name = 'Frisure-Freja'),
    (SELECT id FROM services WHERE name = 'Dame klip'),
    (SELECT id FROM salons WHERE name = 'Salon Klippekrogen'),
    '2026-03-26 14:00:00',
    'Sofie Larsen',
    '+4599887766'
),
(
    (SELECT id FROM barbers WHERE name = 'Bølle-Bob'),
    (SELECT id FROM services WHERE name = 'Skæg trim'),
    (SELECT id FROM salons WHERE name = 'Salon Lokkehuset'),
    '2026-03-27 11:00:00',
    'Jonas Mikkelsen',
    '+4533445566'
);

-- =========================
-- BARBER <-> SALON RELATION
-- =========================
INSERT INTO barber_salons (barber_id, salon_id)
SELECT b.id, s.id
FROM barbers b
JOIN salons s ON (
    (b.name = 'Klippet Kirsten' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Hårdtarbejdende Hans' AND s.name = 'Salon Lille Klippetorv') OR
    (b.name = 'Hårdtarbejdende Hans' AND s.name = 'Salon Lokkehuset') OR
    (b.name = 'Sakse-Søren' AND s.name = 'Salon H.C. Andersens Skæve Hus') OR
    (b.name = 'Frisure-Freja' AND s.name = 'Salon Klippekrogen') OR
    (b.name = 'Lokke-Lars' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Krølle-Karen' AND s.name = 'Salon H.C. Andersens Skæve Hus') OR
    (b.name = 'Hårde-Hanne' AND s.name = 'Salon Frisurevej') OR
    (b.name = 'Bølle-Bob' AND s.name = 'Salon Lokkehuset') OR
    (b.name = 'Bølle-Bob' AND s.name = 'Salon Klippekrogen') OR
    (b.name = 'Sakse-Sally' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Klippe-Kasper' AND s.name = 'Salon Lille Klippetorv') OR
    (b.name = 'Lokke-Louise' AND s.name = 'Salon H.C. Andersens Skæve Hus') OR
    (b.name = 'Krøllet-Kurt' AND s.name = 'Salon Klippekrogen') OR
    (b.name = 'Krøllet-Kurt' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Hår-Hilde' AND s.name = 'Salon H.C. Andersens Skæve Hus') OR
    (b.name = 'Stylet-Stine' AND s.name = 'Salon Frisurevej') OR
    (b.name = 'Bølget-Benny' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Sakse-Simone' AND s.name = 'Salon Lille Klippetorv') OR
    (b.name = 'Sakse-Simone' AND s.name = 'Salon Lokkehuset') OR
    (b.name = 'Klip-Kim' AND s.name = 'Salon Klippekrogen') OR
    (b.name = 'Lokke-Lea' AND s.name = 'Salon H.C. Andersens Skæve Hus') OR
    (b.name = 'Krøl-Kalle' AND s.name = 'Salon Runde Tårn') OR
    (b.name = 'Hår-Henriette' AND s.name = 'Salon Frisurevej') OR
    (b.name = 'Stylet-Stefan' AND s.name = 'Salon Lokkehuset')
);