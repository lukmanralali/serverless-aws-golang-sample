# AWS Serverless Sample With DynamoDB

pastikan serverless terinstall di environtment kerja kita.<br>
jika belum silahkan install terlebih dahulu [how to](https://serverless.com/framework/docs/getting-started)

__Generator__<br>
untuk membuat project baru, bisa menggunakan basic comand serverless [doc](https://serverless.com/framework/docs/providers/aws/)
contoh:
```bash
$ serverless create -t aws-go-dep -p service_name
```

__Manager__<br>
project ini dimanage menggunakan [dep](https://golang.github.io/dep/docs/introduction.html)

__Deployment__<br>
jika AWS CLI kita sudah ter setup dengan benar, bisa langsung lakukan panggil perintah dari makefile.
jika belum ter-setup bisa ke [sini](https://docs.aws.amazon.com/comprehend/latest/dg/setup-awscli.html) atau ke [sini](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)
