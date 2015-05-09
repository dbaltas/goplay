@test "build hello" {
  run go build hello.go
  [ "$status" -eq 0 ]
  [[ "$output" = "" ]]
}

@test "run hello" {
  run ./hello
  [ "$status" -eq 0 ]
  [[ "$output" =~ "hello" ]]
}

@test "remove hello" {
  run rm hello
  [ "$status" -eq 0 ]
  [[ "$output" = "" ]]
}

