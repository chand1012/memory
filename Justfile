default:
  just --list --unsorted

# test the module
test:
  go test .

# benchmark the module
bench: 
  go test -bench .
