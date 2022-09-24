# 轻量化minio客户端

## 配置
在./config目录下建立config.yml文件夹  
### 示例：
```yaml
AppDebug: true   # 调试模式

HttpServer:
  Addr: "http://localhost"
  Port: ":9002" # 网站运行端口号

Logger:
  LogFile: "minio-client.log"
  Console: "enable" # 开关命令行输出 enable or disable

Jwt:
  Secret: "{Secret}"
  Issuer: "{Issuer}"
  ExpireDays: 14

Minio: # 此处位Minio储存库的配置
  Endpoint: "minio:9000"
  AccessKeyID: "{AccessKeyID}"
  SecretAccessKey: "{AccessKeyID}"
  Bucket: "{BucketName}"

Minio-Client: # 此处为访问Minio客户端的配置
  AccessKeyID: "{AccessKeyID}"
  SecretAccessKey: "{SecretAccessKey}"

Document:
  Url: "https://github.com/RockRockWhite/minio-client"
```

## 接口文档
[使用postman共享接口文档](https://www.getpostman.com/collections/9109091476a689689c08)

## 运行
```bash
  git clone https://github.com/RockRockWhite/minio-clinet.git
  cd ./minio-client
  docker compose up
```