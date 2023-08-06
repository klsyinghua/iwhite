```bash
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

```
```sql
CREATE TABLE servers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    hostname VARCHAR(255) NOT NULL,
    ip_address VARCHAR(50) NOT NULL,
    owner VARCHAR(100) NOT NULL,
    status ENUM('Running', 'Stopped') NOT NULL,
    expiration_date DATE NOT NULL,
    category VARCHAR(50) NOT NULL,
    uuid VARCHAR(36) NOT NULL,
    offline_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO servers (hostname, ip_address, owner, status,expiration_date,category,uuid)
VALUES ("ITSM-0001", "192.168.1.1", "ITSM", "Running","2022-12-31","infra","12345jdlaj");
```