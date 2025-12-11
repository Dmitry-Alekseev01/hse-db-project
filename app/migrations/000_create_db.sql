SELECT 'CREATE DATABASE football_db'
WHERE NOT EXISTS (
    SELECT FROM pg_database WHERE datname = 'football_db'
)\gexec