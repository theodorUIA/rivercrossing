module github.com/theodoruia/rivercrossing

go 1.17

replace github.com/theodoruia/rivercrossing/state => C:\Users\Theodor\go\src\github.com\theodoruia\rivercrossing\state

replace github.com/theodoruia/rivercrossing/addremove => C:\Users\Theodor\go\src\github.com\theodoruia\rivercrossing\addremove

replace github.com/theodoruia/myqoute => C:\Users\Theodor\go\src\github.com\theodoruia\myqoute

require (
	github.com/theodoruia/myqoute v1.2.0
	github.com/theodoruia/rivercrossing/addremove v0.0.0-20220128133946-688123256d9e
	github.com/theodoruia/rivercrossing/state v0.0.0-20220128133946-688123256d9e
)

require (
	golang.org/x/text v0.3.7 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)
