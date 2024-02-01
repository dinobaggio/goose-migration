# **Membuat migrations dengan Goose pada Golang**

Migration adalah salah satu aspek yang penting dalam pengembangan perangkat lunak. Ini memungkinkan developer untuk mengelola perubahan struktural dalam database, seperti menambahkan kolom baru, mengubah nama tabel, atau membuat indeks baru dengan cara yang terkelola dan dapat direplikasi dengan mudah.

Goose adalah sebuah alat yang membantu dalam mengelola migration database dengan mudah dan efisien, terutama dalam proyek yang menggunakan bahasa pemrograman Go. Dalam artikel ini, kita akan membahas penggunaan Goose untuk melakukan migration database dalam proyek Go.

artikel penuh disini: [dev.to/dinobaggio](https://dev.to/dinobaggio/cara-membuat-migrations-dengan-goose-pada-golang-268c)

## **Install Goose**

Pertama-tama, Anda perlu menginstal Goose. Anda dapat melakukannya dengan perintah:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Hal ini akan menginstall binary `goose` kedalam directory `$GOPATH/bin`

Untuk pengguna macOS Anda bisa lakukannya dengan perintah:

```bash
brew install goose
```

Lebih lengkap bisa baca dokumentasi [installation instructions](https://pressly.github.io/goose/installation/)

## **Menjalankan Migration**

Selanjutnya kita bisa menjalankan migrations dengan meng-eksekusi file `main.go`.

```bash
go run .
```

Maka jika berhasil outputnya akan seperti ini.

```
2024/01/31 06:55:12 OK    20240131052901_create_users_table.sql
2024/01/31 06:55:12 goose: no migrations to run. current version: 20240131052901
```
