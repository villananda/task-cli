# Task Tracker CLI

Task Tracker adalah aplikasi Command Line Interface (CLI) sederhana berbasis Golang untuk melacak dan mengelola daftar tugas Anda. Aplikasi ini dirancang agar berjalan di terminal, memproses setiap argumen baris perintah (CLI), dan secara otomatis menyimpan seluruh data tugas ke dalam berkas `tasks.json` di dalam direktori aplikasi.

## Fitur

- Menambah, mengubah, dan menghapus tugas.
- Menandai tugas dengan status `in-progress` (sedang dikerjakan) atau `done` (selesai).
- Menampilkan seluruh tugas yang ada.
- Memfilter tampilan tugas berdasarkan status (`todo`, `in-progress`, atau `done`).
- Menggunakan standar hierarki *folder* proyek Go modern sehingga mudah untuk di-maintance atau dikembangkan kembali.

## Struktur Proyek

```text
task-cli/
├── cmd/
│   └── task-cli/
│       └── main.go       (Titik masuk utama aplikasi / Entry point)
├── internal/
│   ├── cli/
│   │   └── handler.go    (Penyambung logika baris perintah dan aksi)
│   ├── storage/
│   │   └── storage.go    (Operasi baca-tulis pada berkas JSON)
│   └── task/
│       └── task.go       (Model data dan konstanta untuk Tugas)
├── go.mod                (Definisi modul Go)
└── tasks.json            (Berkas data yang dibuat secara otomatis oleh aplikasi)
```

## Persyaratan

Pastikan komputer atau sistem operasi Anda sudah menginstal [Go](https://go.dev/dl/) versi terbaru.

## Instalasi

1. Masuk ke dalam direktori aplikasi.
2. Anda bisa langsung menjalankan dengan `go run`, atau akan jauh lebih cepat dan disarankan apabila Anda membangun *binary executable*-nya terlebih dahulu.

Untuk kompilasi program, jalankan perintah berikut:

```bash
go build -o task-cli cmd/task-cli/main.go
```

Hal ini akan menghasilkan berkas yang dapat langsung dieksekusi bernama `task-cli`.

## Panduan Penggunaan

Aplikasi dirancang dengan metode argumen posisi (argumen pertama adalah perintah, sisanya adalah input).

### Menambahkan Tugas

Digunakan untuk menambahkan tugas baru (status otomatis tersetting pada `todo`):
```bash
./task-cli add "Beli perlengkapan harian"
```

### Mengubah Tugas

Memperbarui deskripsi tugas (membutuhkan ID tugas):
```bash
./task-cli update 1 "Beli perlengkapan harian dan susu"
```

### Menghapus Tugas

Menghapus sebuah tugas secara permanen dari daftar:
```bash
./task-cli delete 1
```

### Mengubah Status Tugas

Tandai sebuah tugas sebagai sedang dikerjakan:
```bash
./task-cli mark-in-progress 1
```

Tandai sebuah tugas sebagai telah selesai dikerjakan:
```bash
./task-cli mark-done 1
```

### Menampilkan Daftar Tugas

Tampilkan seluruh daftar tugas dengan format tabel yang rapi:
```bash
./task-cli list
```

Tampilkan tugas yang masih harus dikerjakan:
```bash
./task-cli list todo
```

Tampilkan tugas yang sedang dalam proses:
```bash
./task-cli list in-progress
```

Tampilkan tugas yang sudah diselesaikan:
```bash
./task-cli list done
```

referensi: [Task Tracker](https://roadmap.sh/projects/task-tracker)
