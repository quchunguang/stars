# stars
++++------ your data in your shell, and more. Inspired by holman/spark.

## Install

```sh
go get github.com/quchunguang/stars
```

## Run

```sh
stars 1.2 4.4 3.1 0.5
echo 1.2 4.4 3.1 0.5 | stars
stars --spark 1.2 4.4 3.1 0.5
```

## Cooler Usage

```sh
stars -H --spark "Average scores" 1.2 4.4 3.1 0.5
stars -H --spark "Changed lines in stars.go" 10 5
```

## Contributing

Contributions welcome! Like seriously, I think contributions are real nifty.

Make your changes and be sure the tests all pass:

```sh
go test
```

That also means you should probably be adding your own tests as well as changing the code. Wouldn't want to lose all your good work down the line, after all!

Once everything looks good, open a pull request.
