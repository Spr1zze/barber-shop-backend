-- =========================
-- SALONS
-- (assignment data + old schema's extra columns with sensible defaults)
-- =========================
INSERT INTO salons (slug, name, address, postal_code, city, region, latitude, longitude, description, hero_image_url, phone, email) VALUES
(
    'runde-taarn',
    'Salon Runde Tårn',
    'Frederiksgade 1',
    '8000', 'Aarhus', 'Jylland',
    56.154, 10.203,
    'En klassisk barbersalon i hjertet af Aarhus.',
    'https://images.unsplash.com/photo-1621605815971-fbc98d665033?w=1200&q=80',
    '+45 86 00 00 01',
    'info@rundetaarn.dk'
),
(
    'hc-andersens-skaeve-hus',
    'Salon H.C. Andersens Skæve Hus',
    'Overgade 24',
    '5000', 'Odense', 'Fyn',
    55.399, 10.383,
    'Kreativ salon med sjæl midt i Odense.',
    'https://images.unsplash.com/photo-1517832606299-7ae9b720a186?w=1200&q=80',
    '+45 66 00 00 02',
    'info@skaevehus.dk'
),
(
    'lille-klippetorv',
    'Salon Lille Klippetorv',
    'Bispensgade 34',
    '9000', 'Aalborg', 'Jylland',
    57.048, 9.920,
    'Hyggelig salon med fokus på det gode klip i Aalborg.',
    'https://images.unsplash.com/photo-1503951914875-452162b0f3f1?w=1200&q=80',
    '+45 98 00 00 03',
    'info@lilleklippetorv.dk'
),
(
    'klippekrogen',
    'Salon Klippekrogen',
    'Kongensgade 52',
    '6700', 'Esbjerg', 'Jylland',
    55.476, 8.451,
    'Lokal favorit i Esbjerg med venlig betjening.',
    'https://images.unsplash.com/photo-1622288432450-277d0fef5ed6?w=1200&q=80',
    '+45 75 00 00 04',
    'info@klippekrogen.dk'
),
(
    'lokkehuset',
    'Salon Lokkehuset',
    'Sct. Mathias Gade 45',
    '8800', 'Viborg', 'Jylland',
    56.451, 9.402,
    'Moderne salon i rolige Viborg-omgivelser.',
    'https://images.unsplash.com/photo-1559599101-f09722fb4948?w=1200&q=80',
    '+45 86 00 00 05',
    'info@lokkehuset.dk'
),
(
    'frisurevej',
    'Salon Frisurevej',
    'Algade 15',
    '4000', 'Roskilde', 'Sjælland',
    55.641, 12.081,
    'Professionel salon tæt på Roskilde Domkirke.',
    'https://images.unsplash.com/photo-1596728325488-58c87691e9af?w=1200&q=80',
    '+45 46 00 00 06',
    'info@frisurevej.dk'
);


-- =========================
-- OPENING HOURS (all salons, Mon–Sun)
-- =========================
INSERT INTO salon_opening_hours (salon_id, day_name, day_order, open_time, close_time, is_closed)
SELECT s.id, h.day_name, h.day_order, h.open_time::TIME, h.close_time::TIME, h.is_closed
FROM salons s
JOIN (VALUES
    ('Mandag',  1, '09:00', '18:00', FALSE),
    ('Tirsdag', 2, '09:00', '18:00', FALSE),
    ('Onsdag',  3, '09:00', '18:00', FALSE),
    ('Torsdag', 4, '09:00', '19:00', FALSE),
    ('Fredag',  5, '09:00', '18:00', FALSE),
    ('Lørdag',  6, '10:00', '15:00', FALSE),
    ('Søndag',  7, NULL,    NULL,    TRUE)
) AS h(day_name, day_order, open_time, close_time, is_closed) ON TRUE;


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


-- =========================
-- BARBER <-> SALON (M:N)
-- =========================
INSERT INTO barber_salons (barber_id, salon_id)
SELECT b.id, s.id
FROM barbers b
JOIN salons s ON (
    (b.name = 'Klippet Kirsten'       AND s.slug = 'runde-taarn') OR
    (b.name = 'Hårdtarbejdende Hans'  AND s.slug = 'lille-klippetorv') OR
    (b.name = 'Hårdtarbejdende Hans'  AND s.slug = 'lokkehuset') OR
    (b.name = 'Sakse-Søren'           AND s.slug = 'hc-andersens-skaeve-hus') OR
    (b.name = 'Frisure-Freja'         AND s.slug = 'klippekrogen') OR
    (b.name = 'Lokke-Lars'            AND s.slug = 'runde-taarn') OR
    (b.name = 'Krølle-Karen'          AND s.slug = 'hc-andersens-skaeve-hus') OR
    (b.name = 'Hårde-Hanne'           AND s.slug = 'frisurevej') OR
    (b.name = 'Bølle-Bob'             AND s.slug = 'lokkehuset') OR
    (b.name = 'Bølle-Bob'             AND s.slug = 'klippekrogen') OR
    (b.name = 'Sakse-Sally'           AND s.slug = 'runde-taarn') OR
    (b.name = 'Klippe-Kasper'         AND s.slug = 'lille-klippetorv') OR
    (b.name = 'Lokke-Louise'          AND s.slug = 'hc-andersens-skaeve-hus') OR
    (b.name = 'Krøllet-Kurt'          AND s.slug = 'klippekrogen') OR
    (b.name = 'Krøllet-Kurt'          AND s.slug = 'runde-taarn') OR
    (b.name = 'Hår-Hilde'             AND s.slug = 'hc-andersens-skaeve-hus') OR
    (b.name = 'Stylet-Stine'          AND s.slug = 'frisurevej') OR
    (b.name = 'Bølget-Benny'          AND s.slug = 'runde-taarn') OR
    (b.name = 'Sakse-Simone'          AND s.slug = 'lille-klippetorv') OR
    (b.name = 'Sakse-Simone'          AND s.slug = 'lokkehuset') OR
    (b.name = 'Klip-Kim'              AND s.slug = 'klippekrogen') OR
    (b.name = 'Lokke-Lea'             AND s.slug = 'hc-andersens-skaeve-hus') OR
    (b.name = 'Krøl-Kalle'            AND s.slug = 'runde-taarn') OR
    (b.name = 'Hår-Henriette'         AND s.slug = 'frisurevej') OR
    (b.name = 'Stylet-Stefan'         AND s.slug = 'lokkehuset')
);


-- =========================
-- SERVICES
-- (assignment names/prices + old schema's extra columns)
-- Services are shared across the chain — assigned to a representative salon
-- or duplicated per salon as needed. Here they're global via a shared salon.
-- Since salon_id is now required, we spread the assignment services across
-- relevant salons based on the barber assignments above.
-- =========================
INSERT INTO services (salon_id, slug, name, duration_minutes, price, display_order) VALUES
-- Runde Tårn
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'striber',          'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'farvet-kort',      'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'farvet-langt',     'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'herre-klip',       'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'dame-klip',        'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'skaeg-trim',       'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'runde-taarn'),       'vask-og-styling',  'Vask & styling',      30, 200, 7),
-- H.C. Andersens Skæve Hus
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'striber',         'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'farvet-kort',     'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'farvet-langt',    'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'herre-klip',      'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'dame-klip',       'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'skaeg-trim',      'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'), 'vask-og-styling', 'Vask & styling',      30, 200, 7),
-- Lille Klippetorv
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'striber',          'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'farvet-kort',      'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'farvet-langt',     'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'herre-klip',       'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'dame-klip',        'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'skaeg-trim',       'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'lille-klippetorv'), 'vask-og-styling',  'Vask & styling',      30, 200, 7),
-- Klippekrogen
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'striber',          'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'farvet-kort',      'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'farvet-langt',     'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'herre-klip',       'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'dame-klip',        'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'skaeg-trim',       'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'klippekrogen'), 'vask-og-styling',  'Vask & styling',      30, 200, 7),
-- Lokkehuset
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'striber',          'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'farvet-kort',      'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'farvet-langt',     'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'herre-klip',       'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'dame-klip',        'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'skaeg-trim',       'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'lokkehuset'), 'vask-og-styling',  'Vask & styling',      30, 200, 7),
-- Frisurevej
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'striber',          'Striber',             90, 350, 1),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'farvet-kort',      'Farvet hår (kort)',   60, 450, 2),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'farvet-langt',     'Farvet hår (langt)', 90, 650, 3),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'herre-klip',       'Herre klip',          30, 250, 4),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'dame-klip',        'Dame klip',           45, 300, 5),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'skaeg-trim',       'Skæg trim',           20, 150, 6),
((SELECT id FROM salons WHERE slug = 'frisurevej'), 'vask-og-styling',  'Vask & styling',      30, 200, 7);


-- =========================
-- BOOKINGS
-- (exact assignment data — barber/service/salon/date_time/customer/phone)
-- =========================
INSERT INTO bookings (barber_id, service_id, salon_id, date_time, customer_name, phone) VALUES
(
    (SELECT id FROM barbers WHERE name = 'Klippet Kirsten'),
    (SELECT id FROM services WHERE slug = 'herre-klip' AND salon_id = (SELECT id FROM salons WHERE slug = 'runde-taarn')),
    (SELECT id FROM salons WHERE slug = 'runde-taarn'),
    '2026-03-25 10:00:00',
    'Mads Jensen',
    '+4522334455'
),
(
    (SELECT id FROM barbers WHERE name = 'Hårdtarbejdende Hans'),
    (SELECT id FROM services WHERE slug = 'farvet-langt' AND salon_id = (SELECT id FROM salons WHERE slug = 'lokkehuset')),
    (SELECT id FROM salons WHERE slug = 'lokkehuset'),
    '2026-03-25 12:30:00',
    'Emma Nielsen',
    '+4511223344'
),
(
    (SELECT id FROM barbers WHERE name = 'Sakse-Søren'),
    (SELECT id FROM services WHERE slug = 'striber' AND salon_id = (SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus')),
    (SELECT id FROM salons WHERE slug = 'hc-andersens-skaeve-hus'),
    '2026-03-26 09:15:00',
    'Lucas Pedersen',
    '+4544556677'
),
(
    (SELECT id FROM barbers WHERE name = 'Frisure-Freja'),
    (SELECT id FROM services WHERE slug = 'dame-klip' AND salon_id = (SELECT id FROM salons WHERE slug = 'klippekrogen')),
    (SELECT id FROM salons WHERE slug = 'klippekrogen'),
    '2026-03-26 14:00:00',
    'Sofie Larsen',
    '+4599887766'
),
(
    (SELECT id FROM barbers WHERE name = 'Bølle-Bob'),
    (SELECT id FROM services WHERE slug = 'skaeg-trim' AND salon_id = (SELECT id FROM salons WHERE slug = 'lokkehuset')),
    (SELECT id FROM salons WHERE slug = 'lokkehuset'),
    '2026-03-27 11:00:00',
    'Jonas Mikkelsen',
    '+4533445566'
);