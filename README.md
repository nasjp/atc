[![Actions Status](https://github.com/NasSilverBullet/atc/workflows/CI/badge.svg)](https://github.com/NasSilverBullet/atc/actions)

# [WIP] atc

## atc is light test tools for atcoder

## Getting Started

```sh
$ go get -u github.com/NasSilverBullet/atc/cmd/atc

$ atc init
Config file generated successfully!!
Please edit config file
```

## Run

```sh
$ atc test
=== RUN   abc051_b.py
    --- PASS   abc051_b.py  sample1
    --- PASS   abc051_b.py  sample2
=== RUN   abc081_a.py
    --- PASS   abc081_a.py  sample1
    --- PASS   abc081_a.py  sample2
=== RUN   abc081_b.py
    --- FAIL   abc081_b.py  sample1
        Runtime error: exit status 1
    --- FAIL   abc081_b.py  sample2
        Runtime error: exit status 1
    --- FAIL   abc081_b.py  sample3
        Runtime error: exit status 1
=== RUN   abc083_b.py
    --- PASS   abc083_b.py  sample1
    --- PASS   abc083_b.py  sample2
    --- PASS   abc083_b.py  sample3
=== RUN   abc085_b.py
    --- PASS   abc085_b.py  sample1
    --- PASS   abc085_b.py  sample2
    --- PASS   abc085_b.py  sample3
=== RUN   abc085_c.py
    --- FAIL   abc085_c.py  sample1
        EXPECT: 0 9 0
        ACTUAL: 4 0 5
    --- PASS   abc085_c.py  sample2
    --- FAIL   abc085_c.py  sample3
        EXPECT: 2 54 944
        ACTUAL: 14 27 959
    --- PASS   abc085_c.py  sample4
=== RUN   abc086_a.py
    --- PASS   abc086_a.py  sample1
    --- PASS   abc086_a.py  sample2
=== RUN   abc087_b.py
    --- PASS   abc087_b.py  sample1
    --- PASS   abc087_b.py  sample2
    --- PASS   abc087_b.py  sample3
=== RUN   abc088_b.py
    --- PASS   abc088_b.py  sample1
    --- PASS   abc088_b.py  sample2
    --- PASS   abc088_b.py  sample3
=== RUN   arc065_a.py
    --- PASS   arc065_a.py  sample1
    --- PASS   arc065_a.py  sample2
    --- PASS   arc065_a.py  sample3
=== RUN   arc089_a.py
    --- PASS   arc089_a.py  sample1
    --- PASS   arc089_a.py  sample2
    --- PASS   arc089_a.py  sample3
```

## License

MIT License. See LICENSE.txt for more information.
