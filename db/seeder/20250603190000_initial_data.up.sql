-- ==============================================
-- Seeder: initial_data (admin user + roles + permissions)
-- ==============================================

-- 1. Buat role "admin"
INSERT INTO roles (
    id,
    name,
    created_at,
    updated_at
) VALUES 
    (uuid_generate_v4(), 'Super Admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (uuid_generate_v4(), 'Staff', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- -- 3. Buat daftar permission
-- --    Anda bisa menambahkan atau mengurangi yang diperlukan sesuai bisnis logic aplikasi
-- INSERT INTO permissions (
--     id,
--     name,
--     description,
--     created_at,
--     updated_at
-- ) VALUES
--     (uuid_generate_v4(), 'manage_users',       'Mengelola data user (create, update, delete)', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
--     (uuid_generate_v4(), 'manage_roles',       'Mengelola data role (create, update, delete)', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
--     (uuid_generate_v4(), 'manage_permissions', 'Mengelola data permission (create, update, delete)', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
--     (uuid_generate_v4(), 'view_reports',       'Melihat laporan dan statistik',                  CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- -- 4. Kaitkan semua permission ke role "admin"
-- INSERT INTO role_permissions (
--     id,
--     role_id,
--     permission_id,
--     created_at,
--     updated_at
-- )
-- SELECT
--     uuid_generate_v4(),
--     r.id,
--     p.id,
--     CURRENT_TIMESTAMP,
--     CURRENT_TIMESTAMP
-- FROM roles r
-- JOIN permissions p ON p.name IN ('manage_users','manage_roles','manage_permissions','view_reports')
-- WHERE r.name = 'admin';

-- -- 5. Buat user “admin” 
-- --    Catatan: kolom password di sini wajib berupa hash. 
-- --    Ganti 'SHA256_HASH_DI_SINI' dengan hash yang sesuai (misal dari bcrypt atau fungsi hash lain). 
-- --    Contoh di Linux: echo -n "rahasia123" | sha256sum  → hasilnya diganti di bawah.
-- INSERT INTO users (
--     id,
--     name,
--     username,
--     password,
--     email,
--     created_at,
--     updated_at
-- ) VALUES (
--     uuid_generate_v4(),
--     'Administrator',
--     'admin',
--     'SHA256_HASH_DI_SINI',
--     'admin@example.com',
--     CURRENT_TIMESTAMP,
--     CURRENT_TIMESTAMP
-- );

-- -- 6. Kaitkan user "admin" dengan role "admin"
-- INSERT INTO user_roles (
--     id,
--     user_id,
--     role_id,
--     created_at,
--     updated_at
-- )
-- SELECT
--     uuid_generate_v4(),
--     u.id,
--     r.id,
--     CURRENT_TIMESTAMP,
--     CURRENT_TIMESTAMP
-- FROM users u
-- JOIN roles r ON r.name = 'admin'
-- WHERE u.username = 'admin';

-- 7. (Opsional) Jika Anda ingin memberi role “user” pada akun demo lain misalnya:
-- INSERT INTO users (id, name, username, password, email, created_at, updated_at)
-- VALUES (
--     uuid_generate_v4(),
--     'Demo User',
--     'demo',
--     'HASH_DEMO', 
--     'demo@example.com',
--     CURRENT_TIMESTAMP,
--     CURRENT_TIMESTAMP
-- );
-- INSERT INTO user_roles (id, user_id, role_id, created_at, updated_at)
-- SELECT
--     uuid_generate_v4(),
--     u.id,
--     r.id,
--     CURRENT_TIMESTAMP,
--     CURRENT_TIMESTAMP
-- FROM users u
-- JOIN roles r ON r.name = 'user'
-- WHERE u.username = 'demo';

-- ==============================================
-- Akhir seeder
-- ==============================================
