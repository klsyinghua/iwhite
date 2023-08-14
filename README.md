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
```yml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"
    metrics_path: "/api/metrics"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["10.66.207.133:8080"]


```

```bash
podman run -d -p 3306:3306 --name mysql --env="MYSQL_ROOT_PASSWORD=password" --network=host docker.io/mysql 
podman run -d -p 9090:9090 --name prometheus --network=host -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml docker.io/prom/prometheus
podman run -d -p 3000:3000 --name=granfa --network=host docker.io/grafana/grafana
```

```sql
INSERT INTO servers (id, name, v_cpus, ram, disk, ipv4_address, ipv6_address, created_at, updated_at, terminate_at, power_state, host_name, host_status, owner, environment, features)
VALUES
    ('9f819649-1f7c-44ea-a315-7aa6706de50b', 'ecs-test01', 4, 4096, 0, '192.168.0.116', '', '2023-08-12 15:00:00', '2023-08-12 15:00:00', '2023-08-12 18:00:00', 1, 'instance-000ffcfa', 'ACTIVE', 'John Doe', 'Production', 'Feature1,Feature2');
    ('d6c4f518-28c1-43a2-8d24-8cb66b335320', 'ecs-test02', 8, 8192, 100, '192.168.0.117', '', '2023-08-12 15:30:00', '2023-08-12 15:30:00', '2023-08-12 19:00:00', 2, 'instance-00abcd12', 'ACTIVE', 'Jane Smith', 'Development', 'Feature2,Feature3');
    ('d6c4f518-28c1-43a2-8d24-8cb66b335322', 'ecs-test03', 8, 8192, 100, '192.168.0.117', '', '2023-08-12 15:30:00', '2023-08-12 15:30:00', '2023-08-12 19:00:00', 2, 'instance-00abcd12', 'DOWN', 'Jane Smith', 'Development', 'Feature2,Feature3');
    ('d6c4f518-28c1-43a2-8d24-8cb66b335323', 'ecs-test05', 8, 8192, 100, '192.168.0.112', '', '2023-08-12 15:30:00', '2023-08-12 15:30:00', '2023-08-12 19:00:00', 2, 'instance-00abcd13', 'DOWN', 'Jane Smith', 'Development', 'Feature2,Feature3');
    ('d6c4f518-28c1-43a2-8d24-8cb66b335324', 'ecs-test06', 8, 8192, 100, '192.168.0.110', '', '2023-08-12 15:30:00', '2023-08-12 15:30:00', '2023-08-12 19:00:00', 2, 'instance-00abcd14', 'DOWN', 'Jane Smith', 'Development', 'Feature2,Feature3');
    ('d6c4f518-28c1-43a2-8d24-8cb66b335325', 'ecs-test03', 8, 8192, 100, '192.168.0.119', '', '2023-08-12 15:30:00', '2023-08-12 15:30:00', '2023-08-12 19:00:00', 2, 'instance-00abcd15', 'DOWN', 'Jane Smith', 'Development', 'Feature2,Feature3');

-- Add more test data as needed

```

```bash
curl -X PUT -u username:password -H "Content-Type: application/json" -d '{ 
    "ID": "9f819649-1231231231231231233333312312312asdasdsadasdasdasdasd1f7c-44ea-a315-12312",
    "Name": "New Server Nam22222321312312312312322e22222213123123123123222222",
    "VCPUs": 4,
    "RAM": 8192,
    "Disk": 500,
    "IPv4Address": "192.168.0.134",
    "IPv6Address": "2001:db8::1",
    "CreatedAt": "2023-08-12T15:00:00Z",
    "UpdatedAt": "2023-08-12T16:00:00Z",
    "TerminateAt": "2023-08-12T19:00:00Z",
    "PowerState": 1,
    "HostName": "ITSM-0008",
    "HostStatus": "ACTIVE",
    "Owner": "New Owner",
    "Environment": "Production",
    "Features": "Feature1,Feature2"
}' http://localhost:8080/api/servers/d6c4f518-28c1-43a2-8d24-8cb66b335325
{"message":"Server updated successfully"}

```