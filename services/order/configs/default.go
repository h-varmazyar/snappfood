package configs

var (
	DefaultConfig = []byte(`
service_name: order
version: v1.0.1
grpc_port: 6060
http_port: 11000
db:
  mysql_dsn: root:password@tcp(localhost:3306)/orderDB?charset=utf8mb4&parseTime=True&loc=Local
  redis_address: localhost:6379
  redis_password:
  redis_lock_db: 1
manager_app:
  order_queue: "defaultOrders"
  service_configs:
  controller_configs:
reader_app:
  worker_configs:
    order_worker_interval: 1s
    order_queue_key: "defaultOrders"
`)
)
