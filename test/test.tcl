# test.tcl
set val 11
list add platform windows

proc ::square {x} {
	return [* $x $x]
}

puts -nonewline [::square $val]
puts " Finish"
