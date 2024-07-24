# module-graph-search
Tool to search for full dependency chains for a specific dependency. Helps with fast removal/update of dependencies with security issues.

## Installation

```bash
git clone git@github.com:amikhailau-ms/module-graph-search.git
go install module-graph-search/module-graph-search.go
```

## Usage

```bash
go mod graph | module-graph-search --deps=module1[,module2,...]
```

## Usage example

```bash
go mod graph | module-graph-search --deps=github.com/golang/protobuf@v1.5.3,github.com/dgrijalva/jwt-go@v3.2.0+incompatible
```

## Output example

```bash
Dependency chains for github.com/golang/protobuf@v1.5.3

github.com/golang/protobuf@v1.5.3
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go@v0.112.0
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go/bigquery@v1.58.0
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go/compute@v1.23.3
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go/iam@v1.1.5
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go/pubsub@v1.34.0
github.com/golang/protobuf@v1.5.3<--cloud.google.com/go/storage@v1.36.0
github.com/golang/protobuf@v1.5.3<--github.com/google/s2a-go@v0.1.7
github.com/golang/protobuf@v1.5.3<--github.com/googleapis/gax-go/v2@v2.12.0
github.com/golang/protobuf@v1.5.3<--github.com/spf13/afero@v1.11.0
github.com/golang/protobuf@v1.5.3<--github.com/spf13/viper@v1.18.2
github.com/golang/protobuf@v1.5.3<--go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc@v0.46.1
github.com/golang/protobuf@v1.5.3<--golang.org/x/oauth2@v0.15.0
github.com/golang/protobuf@v1.5.3<--google.golang.org/api@v0.155.0
github.com/golang/protobuf@v1.5.3<--google.golang.org/genproto@v0.0.0-20240125205218-1f4bbc51befe
github.com/golang/protobuf@v1.5.3<--google.golang.org/genproto/googleapis/api@v0.0.0-20240116215550-a9fa1716bcac
github.com/golang/protobuf@v1.5.3<--google.golang.org/grpc@v1.61.0
github.com/golang/protobuf@v1.5.3<--github.com/spf13/afero@v1.11.0<--github.com/sagikazarmark/locafero@v0.4.0
github.com/golang/protobuf@v1.5.3<--google.golang.org/genproto@v0.0.0-20240125205218-1f4bbc51befe<--google.golang.org/genproto/googleapis/rpc@v0.0.0-20240205150955-31a09d347014


Dependency chains for github.com/dgrijalva/jwt-go@v3.2.0+incompatible

github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/markbates/goth@v1.67.1
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.7.0<--github.com/gobuffalo/buffalo@v0.16.14
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.4.0<--github.com/spf13/cobra@v0.0.6
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.4.0<--github.com/spf13/cobra@v0.0.6<--github.com/gobuffalo/packr/v2@v2.8.0
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.4.0<--github.com/spf13/cobra@v0.0.6<--github.com/markbates/refresh@v1.11.1
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.6.2<--github.com/gobuffalo/buffalo@v0.15.4<--github.com/gobuffalo/mw-csrf@v1.0.0
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.6.2<--github.com/gobuffalo/buffalo@v0.15.4<--github.com/gobuffalo/mw-forcessl@v0.0.0-20200131175327-94b2bd771862
github.com/dgrijalva/jwt-go@v3.2.0+incompatible<--github.com/spf13/viper@v1.6.2<--github.com/gobuffalo/buffalo@v0.15.4<--github.com/gobuffalo/mw-paramlogger@v1.0.0
```
