module lesson

go 1.22.0

replace (
	greeting/10 => github.com/automacon-gromoff/greeting v1.0.0
	greeting/11 => github.com/automacon-gromoff/greeting v1.1.0
)

require (
	greeting/10 v0.0.0
	greeting/11 v0.0.0
)
