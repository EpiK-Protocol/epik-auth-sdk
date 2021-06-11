module github.com/EpiK-Protocol/epik-auth-sdk

go 1.16

require (
	github.com/EpiK-Protocol/go-epik v0.4.2-0.20210530114614-689ada0d7b94
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-state-types v0.1.0
	github.com/shirou/gopsutil v2.20.5+incompatible // indirect
	github.com/supranational/blst v0.2.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
)

replace github.com/filecoin-project/specs-actors/v2 => github.com/EpiK-Protocol/go-epik-actors/v2 v2.4.0-alpha.0.20210517033919-7cac385c0096
