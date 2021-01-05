# gorm-imp

## Database
Aplikasi dibuat bukan untuk production, pengaturan koneksi database ditulis hardcoded pada file `database/postgree.go` 
Sesuaikan dengan database postgreesql yang tersedia, program akan menjalankan auto migrasi dan melakukan relasi dengan sendirinya.
```go
func InitDB(){
	dsn := "host=localhost user=postgres password=Password! dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  ...
}
```

## Logika bisnis
Logika bisnis dikumpulkan pada package service

## Endpoint
### Pasien
1. `GET` `{{url}}/api/pasien?gender=`
Menampilkan pasien dengan urutan jumlah terapi terbanyak paling atas. 
Query gender akan memfilter jenis kelamin dari pasien. gender hanya bisa di isi huruf L atau P atau boleh diabaikan untuk menampilkan keduanya.
2. `POST` `{{url}}/api/pasien`  
Body : 
```json
{
    "nama" : "Muchlis",
    "no_wa" : "0990099090",
    "no_hp" : "0990099090",
    "alamat" : "jl . belitung",
    "jk" : 0
}
```
jenis kelamin disederhanakan menjadi 0 untuk perempuan dan 1 untuk laki-laki.
Data yang tersimpan dari pasien pada field `NoPasien` akan memiliki urutan berbeda antar gender. Sedangkan `IDPasien` akan sesuai urutan primaryKey semestinya.

### Pegawai
1. `GET` `{{url}}/api/pegawai` menampilkan list pegawai
2. `POST` `{{url}}/api/pegawai`  
Body :
```json
{
    "nama" : "MuchlisAdmin",
    "kontak" : "081231847434",
    "username" : "Muchlis",
    "password": "Password",
    "level": 2
}
```
Password akan di hash menggunakan Bcrypt (hashing yang paling banyak digunakan dan populer saat ini). Password tidak akan ditampilkan pada response.

### Terapi
1. `GET` `{{url}}/api/terapi` menampilkan list terapi
2. `POST` `{{url}}/api/terapi-range` menampilkan list terapi pada tanggal tertentu  
Body :
```json
{
    "start_date" : "2021-01-02T19:49:58.828478+08:00",
    "end_date" : "2021-01-02T19:49:58.828478+08:00"
}
```
Menggunakan method post karena lebih mudah menulis tanggalnya di body daripada di query Url.  

3. `POST` `{{url}}/api/terapi`  
Body :
```json
{
    "nama" : "XXX",
    "biaya" : 500000,
    "layanan" : "Massage",
    "pasien_id" : 1
}
```

tanggal terapi akan menyesuaikan waktu saat ini. Pegawai yang di tunjuk untuk terapi di pilih secara random oleh sistem.
dan untuk upah akan dihitung secara otomatis.

### Pengeluaran
1. `GET` `{{url}}/api/pengeluaran` menampilkan list pengeluaran pegawai
2. `POST` `{{url}}/api/pengeluaran`  
Body :
```json
{
    "pegawai_id" : 1,
    "deskripsi" : "Beli Jarum suntik gajah",
    "biaya_satuan" : "500000",
    "qty": "2",
    "lampiran": ""
}
```
Karena belum ada auth maka pegawai_id di isi manual. Lampiran dapat diisi dengan Encoded file base64
