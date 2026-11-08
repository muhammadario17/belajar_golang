Modul Berikutnya: Error Handling Ala Go (if err != nil)

Jika kamu pernah belajar Java, Python, atau JavaScript, error biasanya ditangani dengan blok try - catch.

Namun, perancang bahasa Go sengaja MEMBUANG konsep try - catch. Mengapa?

Di bahasa lain, try-catch sering kali menyembunyikan alur jalannya program (control flow) dan membuat programmer malas mengecek kemungkinan error.
Di Go, error dianggap sebagai data biasa (First-Class Citizen).
Setiap fungsi yang berisiko gagal akan mengembalikan 2 nilai sekaligus: (Hasil, Error).

Kalimat paling terkenal dan paling sering diketik oleh semua programmer Go di seluruh dunia adalah:

go
if err != nil {
    // Tangani masalahnya di sini!
}
Mari Kita Coba Studi Kasus Nyata: Mesin ATM (Tarik Tunai)

Bayangkan sebuah fungsi penarikan uang di ATM:

Jika saldo cukup 
→
→ Transaksi berhasil, saldo berkurang, error-nya nil (kosong/tidak ada error).
Jika saldo kurang 
→
→ Transaksi gagal, saldo tetap, dan fungsi mengembalikan pesan error.