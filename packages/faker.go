package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "syreclabs.com/go/faker"
)

func init() {
  env.Packages["faker"] = map[string]reflect.Value{
    "Address":                reflect.ValueOf(faker.Address),
    "App":                    reflect.ValueOf(faker.App),
    "Avatar":                 reflect.ValueOf(faker.Avatar),
    "Business":               reflect.ValueOf(faker.Business),
    "Code":                   reflect.ValueOf(faker.Code),
    "Commerce":               reflect.ValueOf(faker.Commerce),
    "Company":                reflect.ValueOf(faker.Company),
    "Date":                   reflect.ValueOf(faker.Date),
    "Internet":               reflect.ValueOf(faker.Internet),
    "Lorem":                  reflect.ValueOf(faker.Lorem),
    "Name":                   reflect.ValueOf(faker.Name),
    "Number":                 reflect.ValueOf(faker.Number),
    "PhoneNumber":            reflect.ValueOf(faker.PhoneNumber),
    "Time":                   reflect.ValueOf(faker.Time),
    "Finance":                reflect.ValueOf(faker.Finance),
  }
  env.PackageTypes["faker"] = map[string]reflect.Type{

  }
}
