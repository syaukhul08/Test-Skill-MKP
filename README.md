B.	System Design Test 
Anda diminta untuk merancang sebuah sistem E-Ticketing Transportasi Publik yang beroperasi 24 jam, dimana tarif ditentukan berdasarkan titik terminal ketika checkin dan titik terminal ketika checkout. Untuk titik terminal ada 5 dan mempunyai gate/gerbang validasi yang terhubung pada server (bisa 1 atau lebih). Pembayaran dilakukan menggunakan kartu prepaid (seperti pada tol) yang mampu berjalan saat offline.

1.	Gambarkan desain rancangan anda
   
   ![Capture](https://github.com/syaukhul08/Test-Skill-MKP/assets/61621568/0595dd1b-9f24-4605-b353-fdf376eead05)

3. 	Ceritakan rancangan anda dengan jelas saat ada jaringan internet

    Ketika ada jaringan internet, sistem akan berfungsi sebagai berikut:
   - Pengguna melakukan login menggunakan aplikasi mereka dengan mengirimkan kredensial (nama pengguna dan kata sandi) ke server.
   - Setelah login berhasil, server akan menghasilkan token JWT yang akan digunakan untuk otentikasi di permintaan selanjutnya.
   - Pengguna dapat melakukan top-up saldo kartu prepaid mereka melalui aplikasi dengan menghubungkan ke server. Saldo kartu akan diperbarui di server dan dikirimkan ke kartu yang sesuai.
   - Ketika pengguna ingin naik transportasi publik, mereka akan melakukan check-in dengan menggesekkan kartu mereka pada gate terminal. Informasi check-in dan token JWT akan dikirimkan ke server.
   - Saat pengguna melakukan check-out, mereka akan melakukan tindakan yang sama. Server akan menghitung tarif perjalanan berdasarkan titik check-in dan check-out, kemudian mengurangkan saldo kartu pengguna.

5. 	Ceritakan solusi anda dengan jelas (apabila memungkinkan) saat tidak ada jaringan internet

   	Ketika tidak ada jaringan internet, sistem masih dapat berfungsi, tetapi dalam mode terbatas:
   	- Pengguna masih dapat melakukan check-in dan check-out menggunakan kartu mereka pada gate terminal. Data transaksi akan disimpan secara lokal di terminal.
    - Saldo kartu dan data transaksi yang tersimpan di kartu akan diperbarui ketika pengguna melakukan top-up di mesin pengisian yang kompatibel. Mesin ini akan menghubungkan ke server ketika kembali ada jaringan internet untuk memperbarui saldo dan data transaksi.
    - Ketika ada koneksi internet kembali, terminal akan mengirimkan data transaksi yang tertunda ke server. Server akan melakukan perhitungan tarif dan mengurangkan saldo di akun pengguna.

C.	Database Design Test
Gambarkan rancangan database anda (lebih detail lebih baik) sesuai dengan analisa 	sistem pada point test B. Mohon dapat lakukan sesuai instruksi.

Link bddiagram.io
https://dbdiagram.io/d/65096a4202bd1c4a5ed86463
![Untitled](https://github.com/syaukhul08/Test-Skill-MKP/assets/61621568/22a759e8-7456-4d36-b5b2-879bfd1c5a44)

D.	Skill Test
Buat API sederhana menggunakan Golang dengan database point test C. [Uploading MKP.postman_collection.json…]()

API tersebut adalah :
1.	API Login dengan output JSON beserta token JWT
2.	Api Create Terminal dengan Authorization Token dari JWT Login [Uploading MKP.postman_collection.json…]()


Export Postman on the Expost Postman Folder
