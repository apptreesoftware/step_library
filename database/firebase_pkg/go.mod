module firebase_pkg

go 1.12

require (
	cloud.google.com/go v0.34.0
	firebase.google.com/go v3.7.0+incompatible
	github.com/apptreesoftware/go-workflow v0.0.0-20190422155832-a31ef04a817f
	github.com/golang/mock v1.2.0 // indirect
	github.com/google/martian v2.1.0+incompatible // indirect
	github.com/googleapis/gax-go v2.0.2+incompatible // indirect
	github.com/kr/pretty v0.1.0 // indirect
	go.opencensus.io v0.20.0 // indirect
	golang.org/x/xerrors v0.0.0-20190513163551-3ee3066db522
	google.golang.org/api v0.3.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
