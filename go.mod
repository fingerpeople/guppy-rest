module github.com/fingerpeople/guppy-rest

go 1.14

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/Luzifer/go-openssl v2.0.0+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/fingerpeople/guppy v0.0.0-20200610111803-7181468a8ff3
	github.com/fingerpeople/guppy-cli v0.0.0-20200610112002-8bd9ed670302
	github.com/gin-gonic/gin v1.6.3
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/urfave/cli v1.22.4
	gopkg.in/yaml.v2 v2.3.0
)
