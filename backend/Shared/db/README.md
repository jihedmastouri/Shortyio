

## Recommendations:

```sql
-- Create a user called shorty with permissions to read and write to the database for Commands
CREATE USER shorty WITH PASSWORD 'shorty';
GRANT CONNECT ON DATABASE shorty TO shorty;
GRANT USAGE ON SCHEMA blocks TO shorty;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA blocks TO shorty;
GRANT SELECT, USAGE ON ALL SEQUENCES IN SCHEMA blocks TO shorty;

-- Create a user with only read permissions for Aggregations
CREATE USER shorty_read WITH PASSWORD 'shorty_read';
GRANT CONNECT ON DATABASE shorty TO shorty_read;
GRANT USAGE ON SCHEMA blocks TO shorty_read;
GRANT SELECT ON ALL TABLES IN SCHEMA blocks TO shorty_read;
GRANT SELECT, USAGE ON ALL SEQUENCES IN SCHEMA blocks TO shorty_read;
```
