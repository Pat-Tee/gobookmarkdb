package db

const initDB = `
    CREATE TABLE IF NOT EXISTS bookmark(
      url TEXT NOT NULL UNIQUE,
      description TEXT NOT NULL,
      created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
      updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
    ); 
    CREATE TRIGGER IF NOT EXISTS update_Timestamp_Trigger AFTER UPDATE ON bookmark
    BEGIN
    UPDATE bookmark SET updated_at = CURRENT_TIMESTAMP WHERE rowid = old.rowid;
    END;
    `
