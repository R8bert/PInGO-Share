-- Fix timestamp columns to use TIMESTAMP WITH TIME ZONE
-- This is needed because the Go backend created columns as TIMESTAMP without timezone

-- Fix users table
ALTER TABLE users 
ALTER COLUMN created_at TYPE TIMESTAMP WITH TIME ZONE;

-- Fix settings table (if exists)
ALTER TABLE settings 
ALTER COLUMN updated_at TYPE TIMESTAMP WITH TIME ZONE;

-- Fix uploads table
ALTER TABLE uploads 
ALTER COLUMN created_at TYPE TIMESTAMP WITH TIME ZONE;

-- Fix reverse_share_tokens table
ALTER TABLE reverse_share_tokens 
ALTER COLUMN created_at TYPE TIMESTAMP WITH TIME ZONE;

-- Fix deletion_logs table
ALTER TABLE deletion_logs 
ALTER COLUMN uploaded_at TYPE TIMESTAMP WITH TIME ZONE,
ALTER COLUMN deleted_at TYPE TIMESTAMP WITH TIME ZONE;
