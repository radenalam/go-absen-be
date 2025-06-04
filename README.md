# Dokumentasi Migrasi Database

README ini menjelaskan cara mengelola migrasi database menggunakan **golang-migrate** dalam proyek Go “go-absen-be”. Di sini Anda akan menemukan langkah‐langkah untuk membuat, menjalankan, dan men‐rollback migrasi.

---

## Prasyarat

1. **Go** (versi minimal 1.18)

2. **PostgreSQL** (versi minimal 12)

3. **golang-migrate (migrate CLI)**

   - ◦ macOS/Linux (Homebrew):

     ```bash
     brew install golang-migrate
     ```

   - ◦ Atau unduh binary sesuai OS dari [Release Page golang-migrate](https://github.com/golang-migrate/migrate/releases).

4. **psql** (PostgreSQL CLI)

5. Pastikan extension `uuid-ossp` sudah aktif di database.

---

## Struktur Folder Migrasi

Dalam folder proyek, struktur direktori migrasi terletak di:

```
go-absen-be/
└── db/
    └── migrations/
        ├── 20250603065000_create_users_table.up.sql
        ├── 20250603065000_create_users_table.down.sql
        ├── 20250603072026_create_refresh_tokens_table.up.sql
        ├── 20250603072026_create_refresh_tokens_table.down.sql
        ├── 20250603100000_create_roles_table.up.sql
        ├── 20250603100000_create_roles_table.down.sql
        ├── 20250603110000_create_user_roles_table.up.sql
        ├── 20250603110000_create_user_roles_table.down.sql
        └── … (migrasi lain diurutkan berdasarkan timestamp/nomor prefix)
```

- File dengan ekstensi `.up.sql` berisi skrip untuk **membuat/memodifikasi** tabel.
- File `.down.sql` berisi skrip untuk **rollback** (menghapus atau mengembalikan perubahan).
- Penamaan file mengikuti pola `<timestamp>_<deskripsi>.up.sql` dan `<timestamp>_<deskripsi>.down.sql`, di mana `timestamp` berupa timerpanjang (misalnya `20250603110000`).

---

## Konfigurasi Koneksi Database

Sebelum menjalankan migrasi, atur variabel environment `DATABASE_URL` (format PostgreSQL):

```bash
DATABASE_URL="postgres://<db_user>:<db_pass>@<host>:<port>/<db_name>?sslmode=disable"
```

Contoh:

```bash
DATABASE_URL="postgres://alam@localhost:5432/go-absen?sslmode=disable"
```

> **Catatan:**
>
> - `sslmode=disable` biasanya digunakan di lingkungan lokal.

---

## Contoh Penggunaan Makefile (Alias Pendek)

Untuk menghindari mengetik ulang flag, Anda bisa memanfaatkan **Makefile** yang sudah disiapkan:

#### Contoh Perintah via Makefile

- **Jalankan semua migrasi**

  ```bash
  make up
  ```

- **Rollback satu level**

  ```bash
  make down
  ```

- **Buat migrasi baru**

  ```bash
  make create add_index_to_users_email
  ```

- **Force versi (misal versi bersih: 20250603110000)**

  ```bash
  make force 20250603110000
  ```

- **Drop semua dan reset**

  ```bash
  make drop
  ```

---

## Contoh Isi File Migrasi

### 1. Contoh `20250603070506_create_users_table.up.sql`

```sql
-- Pastikan extension uuid-ossp sudah aktif
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    fcm_token VARCHAR(255) DEFAULT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
```

### 2. Contoh `20250603070506_create_users_table.down.sql`

```sql
DROP TABLE IF EXISTS users;
```

---

## Tips dan Best Practices

1. **Urutkan file migrasi secara kronologis**
   Gunakan timestamp (YYYYMMDDHHMMSS) agar nama file sesuai urutan eksekusi.

2. **Jangan ubah file migrasi yang sudah di‐push ke repository**
   Bila ada koreksi setelah migrasi dijalankan pada environment lain, buat migrasi baru untuk mengubah skema, bukan memodifikasi file lama.

3. **Gunakan `ON DELETE CASCADE` untuk relasi penting**
   Saat membuat foreign key di SQL, tambahkan `ON DELETE CASCADE` jika ingin otomatis menghapus data anak ketika data induk dihapus.

4. **Selalu cek versinya sebelum menjalankan migrasi**
   Jika muncul `dirty`, jangan langsung `up`—pertama gunakan `force` atau rollback secara manual, lalu jalankan kembali.

5. **Backup Data**
   Sebelum menjalankan perintah `drop` atau force downgrade, selalu pastikan backup database tersedia—karena perintah tersebut bisa menghapus data secara permanen.

---

## Troubleshooting

- **Error “Dirty database version …”**
  → Gunakan `make force <versi_bersih>`, atau rollback manual (`make down`) sampai versi bersih, lalu `make up`.

- **Syntax error pada file `.up.sql`**
  → Pastikan sintaks SQL benar (`CREATE TABLE IF NOT EXISTS ...`). Cek ada/tidaknya koma berlebih sebelum `)` dan penggunaan `IF NOT EXISTS` pada posisi yang tepat.

- **Extension `uuid-ossp` tidak ditemukan**
  → Jalankan di psql:

  ```sql
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  ```

- **Tabel sudah ada saat migrasi up**
  → Jika memang ingin memulai baru, gunakan `make drop` lalu `make up`. Jika bukan, cek skrip `.down.sql` untuk memastikan migrasi rollback berjalan sempurna sebelum ulang.

---
