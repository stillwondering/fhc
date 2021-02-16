CREATE TABLE IF NOT EXISTS hose (
    id INTEGER NOT NULL PRIMARY KEY,
    code INTEGER NOT NULL,
    type CHAR(1) NOT NULL,
    length INTEGER NOT NULL,

    created_at TEXT,
    updated_at TEXT
);

CREATE TRIGGER hose_update_timestamps_after_insert AFTER INSERT ON hose
BEGIN
    UPDATE
        hose
    SET
        created_at = DATETIME('now'),
        updated_at = DATETIME('now')
    WHERE
        id = NEW.id;
END;

CREATE TRIGGER hose_update_timestamps_after_update AFTER UPDATE ON hose
BEGIN
    UPDATE
        hose
    SET
        updated_at = DATETIME('now')
    WHERE
        id = OLD.id;
END;
