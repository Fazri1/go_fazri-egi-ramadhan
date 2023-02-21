# (6) Version Control and Branch Management (Git)

Apa itu versioning?
Versioning adalah mengatur versi dari source code program.

## Version Control System
VCS adalah kategori alat perangkat lunak yang membantu merekam perubahan yang dilakukan pada file dengan melacak modifikasi yang dilakukan dalam kode. Git adalah salah satu version control system populer yang digunakan para developer untuk mengembangkan software secara bersama-bersama. Git merupakan distributed VCS yang dibuat oleh Linus Torvalds pada tahun 2005. Git melacak setiap perubahan pada file. Cukup rumit untuk mengatur server git, maka dari itu kita dapat menggunakan github sebagai layanan hosting.

## Branch Management
### Perintah Git
Beberapa perintah yang dapat digunakan pada git ke github :

#### Konfigurasi
`git config --global user.name "<nama>"`
Mengatur nama yang ingin ditautkan pada transaksi commit.

`git config --global user.email "<email>"`
Mengatur alamat email yang ingin ditautkan pada transaksi commit.

`git init`
Membuat repositori lokal baru dari folder.

`git clone <repo_url>`
Unduh sebuah proyek dan seluruh riwayat revisinya.


#### Menyimpan Perubahan
`git status`
Melihat perubahan yang siap didaftarkan dalam commit.

```
git add <directory> # sebuah direktori
git add hello.py # satu file
git add . # semua file dan direktori
```
Rekam direktori yang akan didaftarkan ke dalam commit.

`git commit -m "<message>"`
Daftarkan perubahan direktori secara permanen di riwayat revisi.
Dalam melakukan commit, pesannya harus jelas dan mudah dimengerti.

`git diff --staged`
Menunjukkan perbedaan direktori-direktori yang belum didaftarkan dalam commit.

`git stash`
Menyimpan semua perubahan direktori yang terlacak untuk sementara.

`git stash apply`
Menerapkan kembali perubahan simpanan Anda.


#### Memeriksa Repositori
`git log --oneline`
Melihat revisi lama.

`git checkout <branch_name>`
Berpindah ke cabang tertentu dan perbarui direktori yang sedang dikerjakan.

`git reset <commit_id> <mode>`
Membatalkan semua commit (setelah id commit). Terdapat dua mode, yaitu :
- `--soft` -> uncommit perubahan, perubahan dibiarkan bertahap (indeks).
- `--hard` -> uncommit + unstage + hapus perubahan, tidak ada yang tersisa.

#### Sinkronisasi
`git remote add <remote_name> <remote_repo_url>`
Menyinkronisasi repositori yang disimpan di server lain.

`git fetch`
Unduh semua riwayat dari repositori remote

`git pull <remote_name> <branch_name>`
Unduh riwayat marka dan gabungkan perubahan

`git push <remote_name> <local_branch_name>`
Unggah semua commit dari cabang lokal ke GitHub.

#### Percabangan
```
# tampilkan semua branch
git branch --list

# membuat branch baru
git branch <branch>

# paksa hapus sebuah branch
git branch -D <branch>

# list remote branch
git branch -a
```

`git merge <branch>`
Menggabungkan cabang remote ke dalam cabang lokal saat ini


### Workflow Collaboration
Untuk mengoptimalkan kolaborasi dalam development, kita tidak bisa hanya bekerja dalam satu branch, perlu dibuat beberapa branch agar kolaborasi dapat berjalan dengan optimal. Best practice :
1. BUAT BRANCH MASTER DARI BRANCH DEVELOPMENT
2. HINDARI DIRECT EDIT KE BRANCH DEVELOPMENT
3. MERGE BRANCH FEATURE HANYA KE BRANCH DEVELOPMENT
4. MERGE  BRANCH DEVELOPMENT KE BRANCH MASTER JIKA DEVELOPMENT SELESAI

Apa yang membuat git conflict?
Ketika ada dua orang yang membuat perubahan di file dan baris yang sama

Bagaimana cara resolve git conflict?
Resolve dengan satu computer, gunakan plugin untuk membandingkan perubahan

Jika terlalu banyak branch, apakah branch boleh dihapus?
Ya tentu boleh, hapus saja branch yang sudah lama.
