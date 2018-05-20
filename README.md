# handle-toml-with-go-assets

## Requirement

- Docker

## Usage

```bash
$ docker build -t localhost/handle-toml-with-go-assets .
$ docker run --rm localhost/handle-toml-with-go-assets
```

## Output

```console
$ docker run --rm localhost/handle-toml-with-go-assets
type: string, value: "Hello Toml!"
type: int, value: 0
type: float32, value: 1
type: bool, value: true
type: []string, value: []string{"pen", "pineapple", "apple", "pen"}
```