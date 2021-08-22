module github.com/ozoncp/ocp-roadmap-api

go 1.16

require (
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.14.0
	github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api v0.0.0-00010101000000-000000000000
	github.com/pashagolub/pgxmock v1.4.0 // indirect
	github.com/pressly/goose/v3 v3.1.0 // indirect
	github.com/rs/zerolog v1.23.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
)

replace github.com/ozoncp/ocp-roadmap-api/pkg/ocp-roadmap-api => ./pkg/ocp-roadmap-api
