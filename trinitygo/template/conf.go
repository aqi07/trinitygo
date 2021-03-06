package template

func init() {
	_templates["/conf/conf.toml"] = genConf()
}

func genConf() string {
	return `
[project]
name = "{{.PackageName}}"
version = "v1.0"
tags = ["primary"]

[runtime]
debug = true

[security]
    [security.jwt]
    secret_key = "xxx"
    issuer = "isser"
    expire_hour = 24
    header_prefix = "Bearer"
    verify_issuer = false
    verify_expire_hour = false
    [security.cors]
    enable = true
    allow_origins = ["http://localhost:9000"]
    allow_methods = ["GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"]
    allow_headers = ["Origin", "Content-Length", "Content-Type", "authorization", "x-api-key"]
    expose_Headers = ["Content-Length"]
    allow_credentials = true
    max_age_hour =  12
    [security.tls]
    enable = false
    ca_pem_file =  "/Users/daniel/Documents/workspace/SolutionDelivery/conf/ca.pem"
    server_pem_file = "/Users/daniel/Documents/workspace/SolutionDelivery/conf/server/server.pem"
    server_key_file = "/Users/daniel/Documents/workspace/SolutionDelivery/conf/server/server.key"
    client_pem_file = "xx"
    client_key_file = "xx"

[app]
type = "HTTP" # HTTP GRPC
address = "127.0.0.1"
port = 8088
read_timeout = 30 #second
read_header_timeout = 30 #second
writer_timeout=  30 #second
idle_timeout = 30 #second
max_header_bytes = 1048576
template_path = "templates/*"
media_url = "/api/media"
media_path =  "./media"
static_url =  "/api/static"
static_path =  "./static"
migration_path =  "./static/migrations"
page_size =  20
max_body_size =  5242880
atomic_request = true
base_url = "/{{.PackageName}}"

[log]
log_root_path = "/var/log/trinity"
log_name = "app.log"

[cache]
    [cache.redis]
      host = "127.0.0.1"
      port =  6379
      password =  "123456"
      max_idle =  30
      max_active =  30
      idle_timeout =  200
    [cache.gcache] 
      cache_size =  50
      timeout =  24

[database]
db_type = "postgres" #mysql  postgres
server = "host=127.0.0.1 port=60901 user={{.PackageName}} password= dbname={{.PackageName}} sslmode=disable" #mysql option =  charset=utf8&parseTime=True&loc=Local
table_prefix =  "{{.PackageName}}_"
max_idle_conn =  10
max_open_conn =  100

[service_discovery]
type =  "etcd"  # consult , etcd
address = "127.0.0.1"
port =  2379
timeout = 5 
deregister_after_critical =  60
health_check_interval =  10
auto_register =  true
`
}
