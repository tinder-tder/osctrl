module admin/sessions

go 1.17

replace github.com/jmpsec/osctrl/nodes => ../../nodes

replace github.com/jmpsec/osctrl/queries => ../../queries

replace github.com/jmpsec/osctrl/types => ../../types

replace github.com/jmpsec/osctrl/users => ../../users

require (
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/jmpsec/osctrl/nodes v0.3.1 // indirect
	github.com/jmpsec/osctrl/queries v0.3.1 // indirect
	github.com/jmpsec/osctrl/types v0.3.1 // indirect
	github.com/jmpsec/osctrl/users v0.3.1
)

require (
	github.com/golang-jwt/jwt/v4 v4.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce // indirect
)
