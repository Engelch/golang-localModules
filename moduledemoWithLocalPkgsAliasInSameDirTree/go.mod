module moduledemo

go 1.19

replace mypackage => ./mypackage

require mypackage v0.0.0-00010101000000-000000000000 // indirect
