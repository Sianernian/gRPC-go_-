### openssl 生成证书

```
1. 生成私钥
openssl genrsa -out server.key 2048
					XXX.key
2. 生成证书 
openssl req -new -x509 -key server.key -out server.crt -days 36500
							XXX.key      XXX.crt
3. 生成csr
openssl req -new -key server.key -out server.csr
						XXX.key 		server.csr
						
```

``` 
更改 openssl.cnf (linux 是openssl.cfg)
1) 复制一份安装 openssl  bin目录下 openssl.cnf文件
2) 找到 [ CA_default] 代开 copy_extensions = copy
3) 找到 [ req ] 打开 req_extensions =v3_req 
4) 找到 [v3_req] 添加 subjectAltName = @alt_name
5) 添加新的 标签 [ alt_name ] 和标签字段名
 	DNS。1=* (你需要的 域名 www.baidu.com)
 	...
```

```
#生成证书私钥
openssl genpkey -algorithm RSA -out an.key 
									XXX.key
#通过私钥 生成证书请求文件
openssl req -new -nodes -key -an.key -out an.csr -days 3650 -config ./openssl.cnf -extensions v3_req

#生成 证书 pem
openssl x509 -req -days 3650 -in an.csr -out an.pem -CA server.crt  -CAKye server.key  -CAcreateserial -extfile ./openssl.cnf -extensions v3_req
```

