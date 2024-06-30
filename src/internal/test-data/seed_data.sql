INSERT INTO member (id, first_name, last_name, title, dob, email, profile_url, created_at)
VALUES (1,'Jack','Goodby','Mr.','1977-03-08','jackgoodby@me.com','feds','2024-06-21 16:36:00');

INSERT INTO member (id, first_name, last_name, title, dob, email, profile_url, created_at)
VALUES (2,'Annabel','Goodby','Miss.','2000-08-26','agoodby@me.com','feds','2024-06-21 16:36:00');

INSERT INTO court_slot_set (name, uuid, slot_times)
VALUES ('Court 1', '456991b5-58d2-4b42-81fd-9bc5fbe102e4',
        '[
          {
            "start_time": "9:00",
            "duration": 90
          },
          {
            "start_time": "10:30",
            "duration": 90
          },
          {
            "start_time": "12:00",
            "duration": 90
          },
          {
            "start_time": "13:30",
            "duration": 90
          },
          {
            "start_time": "15:00",
            "duration": 120
          },
          {
            "start_time": "17:00",
            "duration": 120
          },
          {
            "start_time": "19:00",
            "duration": 120
          }
        ]');

INSERT INTO court_slot_set (name, uuid, slot_times)
VALUES ('Plexi 5,6,7 and Tiger 10,11,12', '83e4854d-ba01-4e0a-96a4-9fc7466dca17',
        '[
          {
            "start_time": "9:30",
            "duration": 90
          },
          {
            "start_time": "11:00",
            "duration": 90
          },
          {
            "start_time": "12:30",
            "duration": 60
          },
          {
            "start_time": "13:30",
            "duration": 60
          },
          {
            "start_time": "14:30",
            "duration": 90
          },
          {
            "start_time": "16:00",
            "duration": 90
          },
          {
            "start_time": "17:30",
            "duration": 90
          },
          {
            "start_time": "19:00",
            "duration": 90
          }
        ]');

INSERT INTO court_slot_set (name, uuid, slot_times)
VALUES ('Tiger 8,9', '98379934-be70-483b-b45a-2d216c460cb9',
        '[
          {
            "start_time": "9:00",
            "duration": 90
          },
          {
            "start_time": "10:30",
            "duration": 90
          },
          {
            "start_time": "12:00",
            "duration": 90
          },
          {
            "start_time": "13:30",
            "duration": 90
          },
          {
            "start_time": "15:00",
            "duration": 90
          },
          {
            "start_time": "16:30",
            "duration": 90
          },
          {
            "start_time": "18:00",
            "duration": 90
          },
          {
            "start_time": "19:30",
            "duration": 90
          }
        ]');

INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (1, 1, 'Centre Court', 'Plexi', '456991b5-58d2-4b42-81fd-9bc5fbe102e4');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (2, 2, NULL, 'Clay', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (3, 3, NULL, 'Clay', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (4, 4, NULL, 'Clay', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (5, 5, NULL, 'Plexi', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (6, 6, NULL, 'Plexi', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (7, 7, NULL, 'Plexi', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (8, 8, NULL, 'Tiger Turf', '98379934-be70-483b-b45a-2d216c460cb9');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (9, 9, NULL, 'Tiger Turf', '98379934-be70-483b-b45a-2d216c460cb9');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (10, 10, NULL, 'Tiger Turf', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (11, 11, NULL, 'Tiger Turf', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (12, 12, NULL, 'Tiger Turf', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (13, 13, NULL, 'Tarmac', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');
INSERT INTO court (id, court_number, alt_name, surface, slot_set_id) VALUES (14, 14, NULL, 'Tarmac', '83e4854d-ba01-4e0a-96a4-9fc7466dca17');

INSERT INTO public.competition (id, uuid, comp_type, name, identifier, is_internal, start_date, end_date) VALUES (1, 'ed480663-af77-48f6-a84f-1eb2de862fa9', 'competition', 'SLTC Annual Club Championships', '2024', true, '2024-06-17 00:00Z', '2024-07-21 23:59Z');
