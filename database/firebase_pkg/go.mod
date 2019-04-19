module firebase_pkg

go 1.12

require (
	cloud.google.com/go v0.34.0
	firebase.google.com/go v3.7.0+incompatible
	github.com/apptreesoftware/go-workflow v0.0.0-20190417191803-3da7584c654b
	github.com/golang/mock v1.2.0 // indirect
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/kr/pretty v0.1.0 // indirect
	go.opencensus.io v0.20.0 // indirect
	google.golang.org/api v0.3.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
