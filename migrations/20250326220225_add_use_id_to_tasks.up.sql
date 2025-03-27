BEGIN;

-- Add the user_id column
ALTER TABLE tasks ADD COLUMN user_id INTEGER;

-- Create an index
CREATE INDEX idx_tasks_user_id ON tasks(user_id);

-- Add foreign key constraint
ALTER TABLE tasks
    ADD CONSTRAINT fk_tasks_user
        FOREIGN KEY (user_id) REFERENCES users(id)
            ON DELETE CASCADE;

COMMIT;
