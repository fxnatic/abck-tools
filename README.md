# ABCK Tools (abck-tools)

A collection of functions necessary for generating Akamai sensor data. The functions in this package are designed to work together to help you in generating valid sensor_data.

## Functions

```go
func Jrs(t int) ([]int, error)

func CalcDis(floatArr []float64) (int, error)

func Ab(input string) int

func GenerateSeparator(inputStrings []string, number int) string

func Od(t string, a string) string

func Rir(t int, a int, e int, n int) int

func Encrypt(sensor string, keys []int) string

func ShuffleString(input string, seed int) string

func EncryptString(input string, key int) string

func MN_S(byteArr []byte) []byte

func ATS(byteArr []byte) string
```

## Installation 

To install this package, you can use the following command in your terminal:

```bash
go get github.com/fxnatic/abck-tools
```

## Usage 

Once the package is installed, you can import it in your Go programs like this:

```go
import "github.com/fxnatic/abck-tools"
```

## License

This project is licensed under the [MIT License](/LICENSE).