package pegawai

type Pegawai struct {
	ID       uint
	Nama     string
	Kontak   string
	Username string
	Password string
	Level    int // 0 admin , 1 superAdmin
}
