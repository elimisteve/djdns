package model

import (
	"github.com/miekg/dns"
	"reflect"
	"testing"
)

type ToDnsTest struct {
	DjRecord        Record
	ExtRecordString string
}

func (test *ToDnsTest) ExtRecord() (dns.RR, error) {
	return dns.NewRR(test.ExtRecordString)
}

func Test_Record_ToDns(t *testing.T) {
	// TODO: Add more tests
	tests := []ToDnsTest{
		ToDnsTest{
			Record{
				DomainName: "host.name.",
				Rttl:       4000,
				Rtype:      "MX",
				Rdata:      "10 9.9.9.9",
			},
			"host.name. 4000 IN MX 10 9.9.9.9",
		},
		ToDnsTest{
			Record{
				DomainName: "another.host.",
				Rdata:      "10.10.10.10",
			},
			"another.host. 3600 IN A 10.10.10.10",
		},
		ToDnsTest{
			Record{
				DomainName: "ipv6.example.",
				Rtype:      "AAAA",
				Rdata:      "fcd5:7d07:2146:f18f:f937:d46e:77c9:80e7",
			},
			"ipv6.example. 3600 IN AAAA fcd5:7d07:2146:f18f:f937:d46e:77c9:80e7",
		},
		ToDnsTest{
			Record{
				DomainName: "ipv6.short.",
				Rtype:      "AAAA",
				Rdata:      "fcd5::77c9:80e7",
			},
			"ipv6.short. 3600 IN AAAA fcd5::77c9:80e7",
		},
	}
	for _, test := range tests {
		expected, err := test.ExtRecord()
		if err != nil {
			t.Fatal(err)
		}
		test.DjRecord.Normalize()
		record, err := test.DjRecord.ToDns()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(record, expected) {
			t.Log(record)
			t.Log(expected)
			t.Fatal("record != expected")
		}
	}
}
