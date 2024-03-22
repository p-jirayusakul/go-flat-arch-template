# Golang Echo and SQLC
 1. ก่อนรัน โปรคทุกครั้งต้องมีไฟล์ .env เสมอ
 2. ไฟล์ .env ห้าม upload ขึ้น git เพราะว่ามี config ที่ห้าม public เช่น IP database, token ต่าง ๆ
 3. หากรับช่วงต่อโปรเจคหรือมาดูโครงเพื่อต่อยอด ต้องศึกษาตัว SQLC ด้วย [อ่านวิธีใช้ SQLC เพิ่มเติม](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html)
 4. โปรเจคนี้ใช้ SQLC ฉนั้นจะมีข้อสังเกตุหลายอย่าง [ตามนี้](#ข้อสังเกตุการใช้งาน-sqlc)

## Setup local development

### Install tools

- [Golang](https://golang.org/)

- [Homebrew](https://brew.sh/)

- [Make](https://makefiletutorial.com/) (ถ้าใช้ Mac OS จะมีมาให้อยู่แล้วตรวจสอบด้วยคำสั่งนี้)

	```bash
	make --version
	```

- [Sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)

	```bash
	brew install sqlc
	```

- [Gomock (for unitesting)](https://github.com/uber-go/mock)

	``` bash
	go install go.uber.org/mock/mockgen@latest
	```

### How to generate code

- Generate SQL CRUD with sqlc:

	```bash
	make sqlc
	```


### How to run

- .env (ดูค่าต่าง ๆ ในโฟลเดอร์ helm ตาม envelopment):

	```bash
    DATABASE_USER=postgres
    DATABASE_HOST=localhost
    DATABASE_PASSWORD=1234
    DATABASE_PORT=5432
    DATABASE_NAME=homework1
    API_PORT=3000
    JWT_SECRET=XyvnrmjDFkCLaUwYZ0zyiPapYSdyVMD8
    SECRET=3nSSLymRXuUnDNXzM50BCaSKgjbcKAK8
	EXTERNAL_API_URL=https://jsonplaceholder.typicode.com
	```

- Run server:
	```bash
	make server
	```
	or
	```bash
	go run main.go
	```
- generate mockup: (สร้างไฟล์ mockup สำหรับทำ unitest)
	```bash
	make mock
	```
- Run test:
	```bash
	make test
	```