package db

const createTables = `
CREATE TABLE IF NOT EXISTS sequences (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    open_tracking_enabled BOOLEAN NOT NULL DEFAULT false,
    click_tracking_enabled BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS steps (
    id BIGSERIAL PRIMARY KEY,
    sequence_id BIGINT NOT NULL REFERENCES sequences(id),
    subject VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    step_order INTEGER NOT NULL CHECK (step_order >= 0),
    wait_days INTEGER NOT NULL CHECK (wait_days >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);
`

// MigrateDB performs all necessary database migrations
func (db *DB) MigrateDB() error {
	// Create sequences table first
	if _, err := db.Conn.Exec(createTables); err != nil {
		return err
	}

	return nil
}
