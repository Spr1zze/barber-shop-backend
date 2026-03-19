-- What do I need?
-- barber (id)
-- service (id)
-- salon (id)
-- timestamp
-- address (salon address)
-- barber (name)


SELECT 
    bookings.id,
    bookings.dateTime,
    salons.address,

FROM bookings
JOIN barbers ON bookings.barber_id = barbers.id
JOIN salons ON bookings.salon_id = salons.id
JOIN services ON bookings.service_id = services.id