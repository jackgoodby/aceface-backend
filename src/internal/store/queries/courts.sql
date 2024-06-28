-- name: GetCourts :many
SELECT
    c.uuid,
    c.court_number,
    c.alt_name,
    c.surface,
    css.slot_times
FROM court c
JOIN court_slot_set css
    ON css.uuid = c.slot_set_id;

-- name: GetCourt :one
SELECT
    c.uuid,
    c.court_number,
    c.alt_name,
    c.surface,
    css.slot_times
FROM court c
         JOIN court_slot_set css
              ON css.uuid = c.slot_set_id
WHERE c.uuid = $1;