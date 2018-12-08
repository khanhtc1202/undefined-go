package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	//convertBase64ToImage()
	b := fakeImageBase64()
	printBase64Image(b)
}

func convertBase64ToImage() {
	b := fakeImageBase64()
	unbased, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		panic("Cannot decode b64")
	}

	r := bytes.NewReader(unbased)
	im, err := png.Decode(r)
	if err != nil {
		panic("Bad png")
	}

	f, err := os.OpenFile("image/example.png", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic("Cannot open file")
	}

	png.Encode(f, im)
}

func fakeImageBase64() string {
	return "iVBORw0KGgoAAAANSUhEUgAAAIAAAAB+CAYAAADsphmiAAAEDWlDQ1BJQ0MgUHJvZmlsZQAAOI2NVV1oHFUUPrtzZyMkzlNsNIV0qD8NJQ2TVjShtLp/3d02bpZJNtoi6GT27s6Yyc44M7v9oU9FUHwx6psUxL+3gCAo9Q/bPrQvlQol2tQgKD60+INQ6Ium65k7M5lpurHeZe58853vnnvuuWfvBei5qliWkRQBFpquLRcy4nOHj4g9K5CEh6AXBqFXUR0rXalMAjZPC3e1W99Dwntf2dXd/p+tt0YdFSBxH2Kz5qgLiI8B8KdVy3YBevqRHz/qWh72Yui3MUDEL3q44WPXw3M+fo1pZuQs4tOIBVVTaoiXEI/MxfhGDPsxsNZfoE1q66ro5aJim3XdoLFw72H+n23BaIXzbcOnz5mfPoTvYVz7KzUl5+FRxEuqkp9G/Ajia219thzg25abkRE/BpDc3pqvphHvRFys2weqvp+krbWKIX7nhDbzLOItiM8358pTwdirqpPFnMF2xLc1WvLyOwTAibpbmvHHcvttU57y5+XqNZrLe3lE/Pq8eUj2fXKfOe3pfOjzhJYtB/yll5SDFcSDiH+hRkH25+L+sdxKEAMZahrlSX8ukqMOWy/jXW2m6M9LDBc31B9LFuv6gVKg/0Szi3KAr1kGq1GMjU/aLbnq6/lRxc4XfJ98hTargX++DbMJBSiYMIe9Ck1YAxFkKEAG3xbYaKmDDgYyFK0UGYpfoWYXG+fAPPI6tJnNwb7ClP7IyF+D+bjOtCpkhz6CFrIa/I6sFtNl8auFXGMTP34sNwI/JhkgEtmDz14ySfaRcTIBInmKPE32kxyyE2Tv+thKbEVePDfW/byMM1Kmm0XdObS7oGD/MypMXFPXrCwOtoYjyyn7BV29/MZfsVzpLDdRtuIZnbpXzvlf+ev8MvYr/Gqk4H/kV/G3csdazLuyTMPsbFhzd1UabQbjFvDRmcWJxR3zcfHkVw9GfpbJmeev9F08WW8uDkaslwX6avlWGU6NRKz0g/SHtCy9J30o/ca9zX3Kfc19zn3BXQKRO8ud477hLnAfc1/G9mrzGlrfexZ5GLdn6ZZrrEohI2wVHhZywjbhUWEy8icMCGNCUdiBlq3r+xafL549HQ5jH+an+1y+LlYBifuxAvRN/lVVVOlwlCkdVm9NOL5BE4wkQ2SMlDZU97hX86EilU/lUmkQUztTE6mx1EEPh7OmdqBtAvv8HdWpbrJS6tJj3n0CWdM6busNzRV3S9KTYhqvNiqWmuroiKgYhshMjmhTh9ptWhsF7970j/SbMrsPE1suR5z7DMC+P/Hs+y7ijrQAlhyAgccjbhjPygfeBTjzhNqy28EdkUh8C+DU9+z2v/oyeH791OncxHOs5y2AtTc7nb/f73TWPkD/qwBnjX8BoJ98VVBg/m8AAAAJcEhZcwAACxMAAAsTAQCanBgAAAQJaVRYdFhNTDpjb20uYWRvYmUueG1wAAAAAAA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJYTVAgQ29yZSA1LjQuMCI+CiAgIDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+CiAgICAgIDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiCiAgICAgICAgICAgIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIgogICAgICAgICAgICB4bWxuczpzdFJlZj0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL3NUeXBlL1Jlc291cmNlUmVmIyIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyI+CiAgICAgICAgIDx4bXBNTTpEZXJpdmVkRnJvbSByZGY6cGFyc2VUeXBlPSJSZXNvdXJjZSI+CiAgICAgICAgICAgIDxzdFJlZjppbnN0YW5jZUlEPnhtcC5paWQ6NDBCRjlEOEE5NkQxMTFFMEEzMjFDOEE4QTZGNTk1M0Q8L3N0UmVmOmluc3RhbmNlSUQ+CiAgICAgICAgICAgIDxzdFJlZjpkb2N1bWVudElEPnhtcC5kaWQ6NDBCRjlEOEI5NkQxMTFFMEEzMjFDOEE4QTZGNTk1M0Q8L3N0UmVmOmRvY3VtZW50SUQ+CiAgICAgICAgIDwveG1wTU06RGVyaXZlZEZy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Ypdddil10vPmgrBdGDkUMF9R44On8LY6LCxbAERECoc80iYycIApllJKwrV7G2kUcCuATaXddt890Nd0cTxpBSTKac649ZhR09x/eGoXazzx6/4CghQj9mb+rb/eemnlVVZJUh4J62sUSAv6rwsAgbDDx5gRmDBhQjwdOT5m4RuCPqsCPKR82F3ad599opgaEr5eXATABdf/P8mgQdpl113T4kssEc6ziwCYSLOiIFA2/8o33njj4OtNN90UT34D8Qvgi5k/gKZiVE8vcXYQADOdsloQghCz0M18xIrZSjJxd/kVV4SxTYo4pz3//Oc/R5EnT54cTxMmPmbRGzWD8muLdtJWt7KUs1rZzWPsLW633XbpbllswzYioI6i5o71L5A1q3Do9P+ftIcrQY0XxiN33HHHpL13s2QrgAAwJwCsJXuGAKZ4gegDYBzxBnUApWEali/Dp0M3MueitnGBiN/NgA5l1XYyrhmOiCAAQ42X8RmI5wpvfnMkK2XSeIYAYKCZZuHt664bRpbxqRInQrd5g5AYWARoabii49HzjjsCYaEg/GARn3xcRs9/YCwasHt8zGI3DGcD7vON4sP/g5VXXpnPfjPBTISQWNU86+yz0/g3vSm9WdK35JJLRiuDFU0MLyMQVaCFMJDGQDAEHBFIqU0lWecIU7X8I4GByM/lGaqny4R1dEA6g2FHOQTA/4Nll102PE2c+OjDjcxIg4tJpXdvs02sT+dJrbXWWon1BjqdkyR4NgyNiXcYkwMCQVqk64L4mYdr5500yYfFr9/+9rfpxz/+cZict3s7aY2ksKwNYIgbnstGUQoBcHOAeXQAYvcXciHYcMMN0z333BM2gQ899NDEf+hZ/W/552qvQbpCw5IHhNB8aoIRine84x0hGEsvvXRadNFFW2olwLdVoeCXw1Q3lsRtu3irrbbqb5GHPL4rSS1ETBvOVVhSfMY8P7/nUUS69777Ig6zRYADx0c/bqQDsWnmbWeflSmPNJw0tU574JO2OiXptKe99947vGRvN3qt0m9Lk2VcWurOoaFEM1arlSAdgHx9hUN2Mz44nX766TH5xQQYZwIQv1Nlz7IctNdWcEfw+dU98fjjMfubsN2LZU81w4VO4pA8/P/qYHz08/aKdqgA2MATJYp/acURQDlBBI/36o1Vya232SbCEye/cP/iF79YsFn1tttv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"
}

func printBase64Image(imgBase64Str string) {
	fmt.Printf("\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}

func printImage(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("\x1b]1337;File=inline=1:%s\a\n", imgBase64Str)
}
