package simplebdd

import(
    "fmt"
    "testing"
    "io/ioutil"
    "strings"
)

func Xunit(t *testing.T) {

    xml := ""
    group := ""
    failed := 0
    passed := 0
    skipped := 0

    for _, test := range testRun.tests {

        if group == "" {
            group = test.group
        }

        if test.skipped {
            skipped++
        } else if test.passed {
            passed++
        } else {
            failed++
        }

        xml += fmt.Sprintf("\t<testcase name=\"%s\" classname=\"%s\" time=\"%f\"/>\n", test.group, test.message, 0.0)
    }

    xml += "</testsuite>\n"
    xml = fmt.Sprintf("<testsuite name=\"%s\" time=\"%f\" tests=\"%d\" errors=\"%d\" skipped=\"%d\" failures=\"%d\">\n%s", group, 0.0, len(testRun.tests), 0, skipped, failed, xml)
    xml = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + xml

    ioutil.WriteFile("xunit_"+strings.Replace(group, " ", "_", -1)+".xml", []byte(xml), 0777)
}
