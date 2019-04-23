module oracle_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190422155832-a31ef04a817f
	github.com/apptreesoftware/step_library/database/db_common v0.0.0
	gopkg.in/goracle.v2 v2.12.3
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
