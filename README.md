# Перша лаба ПКПЗ

## How to use

```shell
go test
```

```shell
# For more verbose output
go test -v
```

## What to expect

`TestToRoman_TwentyOne` should fail because of the incorrect expected value.

`TestToRoman_Forty` should fail because the number is out of the range (>3999)

All other tests should show a `PASS` state.
