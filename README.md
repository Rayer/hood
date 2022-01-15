# Hood

Hood conceals confidential data in struct by adding tag

## Install
`go get github.com/Rayer/hood`

## Example

Given a struct like this

```
type ConfigWithTag struct {
    Host     string
    User     string `confidential:"1,1"`
    Password string `confidential:"2,2"`
}

func (c ConfigWithTag) String() string {
    ret, _ := hood.PrintConfidentialData(c)
    return ret
}
```
Adding `func (c ConfigWithTag) String() string` will override `fmt.Println`, `fmt.Printf("%v")` and `fmt.Printf("%+v")` result, it will return with concealed data. Please refer to example folder for more information.

## Parameters
A standard tag will like :
`confidential:"1,1"`. The first parameter means "how many words keep not be concealed in head", and the second one is for tail.

Given input as `ABCDE12345`
- `confidential:"1,1"` will be `A********5`
- `confidential:"0,5"` will be `*****12345`
