-- Erstelle das Schema 'test_schema'
CREATE SCHEMA IF NOT EXISTS test_schema;

-- Erstelle den Benutzer 'read_user' mit Lesezugriff
CREATE USER read_user WITH PASSWORD 'password';
GRANT CONNECT ON DATABASE test TO read_user;
GRANT USAGE ON SCHEMA test_schema TO read_user;
GRANT SELECT ON ALL TABLES IN SCHEMA test_schema TO read_user;

-- Setze den Benutzer 'admin' mit allen Rechten auf die Datenbank und das Schema
GRANT ALL PRIVILEGES ON DATABASE test TO admin;
GRANT ALL PRIVILEGES ON SCHEMA test_schema TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA test_schema TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA test_schema TO admin;

-- Optional: Setze Standardrechte für alle zukünftigen Tabellen und Sequenzen
ALTER DEFAULT PRIVILEGES IN SCHEMA test_schema
GRANT ALL ON TABLES TO admin;
ALTER DEFAULT PRIVILEGES IN SCHEMA test_schema
GRANT SELECT ON TABLES TO read_user;
