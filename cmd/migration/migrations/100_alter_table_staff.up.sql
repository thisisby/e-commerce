ALTER TABLE staff
    DROP COLUMN IF EXISTS start_time,
    DROP COLUMN IF EXISTS end_time;

ALTER TABLE staff
    ADD COLUMN IF NOT EXISTS time_slot TEXT NOT NULL DEFAULT '[
    {"id": 1, "time": "08:00", "isAvailable": true},
    {"id": 2, "time": "08:15", "isAvailable": true},
    {"id": 3, "time": "08:30", "isAvailable": true},
    {"id": 4, "time": "08:45", "isAvailable": true},
    {"id": 5, "time": "09:00", "isAvailable": true},
    {"id": 6, "time": "09:15", "isAvailable": true},
    {"id": 7, "time": "09:30", "isAvailable": true},
    {"id": 8, "time": "09:45", "isAvailable": true},
    {"id": 9, "time": "10:00", "isAvailable": true},
    {"id": 10, "time": "10:15", "isAvailable": true},
    {"id": 11, "time": "10:30", "isAvailable": true},
    {"id": 12, "time": "10:45", "isAvailable": true},
    {"id": 13, "time": "11:00", "isAvailable": true},
    {"id": 14, "time": "11:15", "isAvailable": true},
    {"id": 15, "time": "11:30", "isAvailable": true},
    {"id": 16, "time": "11:45", "isAvailable": true},
    {"id": 17, "time": "12:00", "isAvailable": false},
    {"id": 18, "time": "12:15", "isAvailable": false},
    {"id": 19, "time": "12:30", "isAvailable": false},
    {"id": 20, "time": "12:45", "isAvailable": false},
    {"id": 21, "time": "13:00", "isAvailable": false},
    {"id": 22, "time": "13:15", "isAvailable": true},
    {"id": 23, "time": "13:30", "isAvailable": true},
    {"id": 24, "time": "13:45", "isAvailable": true},
    {"id": 25, "time": "14:00", "isAvailable": true},
    {"id": 26, "time": "14:15", "isAvailable": true},
    {"id": 27, "time": "14:30", "isAvailable": true},
    {"id": 28, "time": "14:45", "isAvailable": true},
    {"id": 29, "time": "15:00", "isAvailable": true},
    {"id": 30, "time": "15:15", "isAvailable": true},
    {"id": 31, "time": "15:30", "isAvailable": true},
    {"id": 32, "time": "15:45", "isAvailable": true},
    {"id": 33, "time": "16:00", "isAvailable": true},
    {"id": 34, "time": "16:15", "isAvailable": true},
    {"id": 35, "time": "16:30", "isAvailable": true},
    {"id": 36, "time": "16:45", "isAvailable": true},
    {"id": 37, "time": "17:00", "isAvailable": true},
    {"id": 38, "time": "17:15", "isAvailable": true},
    {"id": 39, "time": "17:30", "isAvailable": true},
    {"id": 40, "time": "17:45", "isAvailable": true},
    {"id": 41, "time": "18:00", "isAvailable": true},
    {"id": 42, "time": "18:15", "isAvailable": true},
    {"id": 43, "time": "18:30", "isAvailable": true},
    {"id": 44, "time": "18:45", "isAvailable": true},
    {"id": 45, "time": "19:00", "isAvailable": true},
    {"id": 46, "time": "19:15", "isAvailable": true},
    {"id": 47, "time": "19:30", "isAvailable": true},
    {"id": 48, "time": "19:45", "isAvailable": true},
    {"id": 49, "time": "20:00", "isAvailable": true},
    {"id": 50, "time": "20:15", "isAvailable": true},
    {"id": 51, "time": "20:30", "isAvailable": true},
    {"id": 52, "time": "20:45", "isAvailable": true},
    {"id": 53, "time": "21:00", "isAvailable": true},
    {"id": 54, "time": "21:15", "isAvailable": true},
    {"id": 55, "time": "21:30", "isAvailable": true},
    {"id": 56, "time": "21:45", "isAvailable": true},
    {"id": 57, "time": "22:00", "isAvailable": true},
    {"id": 58, "time": "22:15", "isAvailable": true},
    {"id": 59, "time": "22:30", "isAvailable": true},
    {"id": 60, "time": "22:45", "isAvailable": true},
    {"id": 61, "time": "23:00", "isAvailable": true}
]';

ALTER TABLE staff
    ADD COLUMN IF NOT EXISTS working_days TEXT[] NOT NULL DEFAULT '{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}';