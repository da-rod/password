package password

import "testing"

func TestGenerate(t *testing.T) {
	pwdList := Generate(nil)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, true, true, false})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, true, false, false})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, true, false, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, false, false, false})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, false, true, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, false, true, false})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, true, false, false, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, false, false, false, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, false, true, false, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, false, true, true, true})...)
	pwdList = append(pwdList, Generate(&Config{10, 21, false, false, false, false})...)
	pwdList = append(pwdList, Generate(&Config{0, 21, true, true, true, true})...)
	// ...
	for _, pwd := range pwdList {
		t.Log(pwd)
	}
}
