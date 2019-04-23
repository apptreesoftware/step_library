module postgres_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190422155832-a31ef04a817f
	github.com/apptreesoftware/step_library/database/db_common v0.0.0
	github.com/lib/pq v1.0.0
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
