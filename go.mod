module github.com/karnalab/Karna

go 1.13

require (
	github.com/aws/aws-sdk-go v1.25.19
	github.com/aws/aws-sdk-go-v2 v0.15.0
	github.com/gorilla/mux v1.7.3
	github.com/graphql-go/graphql v0.7.8
	github.com/johnnadratowski/golang-neo4j-bolt-driver v0.0.0-20181101021923-6b24c0085aae
	github.com/karnalab/karna v1.1.0

	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/spf13/cobra v0.0.5
)

replace github.com/karnalab/karna v1.1.0 => ./../karna
