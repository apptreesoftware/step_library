module oracle_pkg

go 1.12

require (
	github.com/apptreesoftware/go-workflow v0.0.0-20190613131947-e477173659c8
	github.com/apptreesoftware/step_library/database/db_common v0.0.0
	golang.org/x/xerrors v0.0.0-20190513163551-3ee3066db522
	gopkg.in/goracle.v2 v2.12.3
)

replace github.com/apptreesoftware/step_library/database/db_common => ../db_common
