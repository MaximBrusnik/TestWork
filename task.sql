CREATE TABLE task(
                     id SERIAL PRIMARY KEY,
                     title VARCHAR(255) NOT NULL,
                     description TEXT,
                     due_date TIMESTAMP NOT NULL,
                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO task (title, description, due_date) VALUES
('Task 1', 'Description for Task 1', '2023-10-15 10:00:00'),
('Task 2', 'Description for Task 2', '2023-10-20 14:30:00'),
('Task 3', 'Description for Task 3', '2023-10-25 09:45:00'),
('Task 4', 'Description for Task 4', '2023-11-01 16:20:00'),
('Task 5', 'Description for Task 5', '2023-11-05 11:15:00');