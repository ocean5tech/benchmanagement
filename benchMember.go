package main

type benchMember struct {
	id         string
	name       string
	email      string
	pemName    string
	pemEmail   string
	benchstart string
	benchEnd   string
}

func NewBenchMember(id string, name string, email string, pemName string, pemEmail string, benchStart string, benchEnd string) *benchMember {
	return &benchMember{
		id, name, email, pemName, pemEmail, benchStart, benchEnd,
	}
}

func (bm *benchMember) Id() string {
	return bm.id
}
func (bm *benchMember) Name() string {
	return bm.name
}
func (bm *benchMember) Email() string {
	return bm.email
}
func (bm *benchMember) PemName() string {
	return bm.pemName
}
func (bm *benchMember) PemEmail() string {
	return bm.pemEmail
}
func (bm *benchMember) BenchStart() string {
	return bm.benchstart
}
func (bm *benchMember) BenchEnd() string {
	return bm.benchEnd
}
