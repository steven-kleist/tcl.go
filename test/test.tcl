# test.tcl
set val 11
set user "Steven"
# list add platform windows

proc ::square {x} {
	return [* $x $x]
}

append user " Kleist" " (stevenkl)"

puts -nonewline [::square [set val]]
puts " Finish"

puts "Hello, $user"
