module github.com/ozoncp/ocp-roadmap-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.10.2
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.23.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api => ./pkg/ocp-roadmap-api
