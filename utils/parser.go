package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const standardURL = "http://nvd.nist.gov/static/feeds/xml/sp80053/rev4/800-53-controls.xml"

type XMLStandard struct {
	Controls []XMLControl `xml:"control"`
	XMLName  xml.Name     `xml:"http://scap.nist.gov/schema/sp800-53/feed/2.0 controls"`
}

type Standard struct {
	Name     string             `yaml:"name"`
	Controls map[string]Control `yaml:",inline"`
}

type XMLControl struct {
	Family               string                    `xml:"family"`
	Number               string                    `xml:"number"`
	Title                string                    `yaml:"name" xml:"title"`
	Priority             string                    `xml:"priority"`
	BaselineImpact       []string                  `xml:"baseline-impact"`
	Statements           []XMLStatement            `xml:"statement"`
	SupplementalGuidance []XMLSupplementalGuidance `xml:"supplemental-guidance"`
	ControlEnhancements  []XMLControlEnhancement   `xml:"control-enhancements>control-enhancement"`
	References           []XMLReference            `xml:"references"`
}

type Control struct {
	Family      string `yaml:"family"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type XMLStatement struct {
	Number      string         `xml:"number"`
	Description string         `xml:"description"`
	Statements  []XMLStatement `xml:"statement"`
}

type XMLSupplementalGuidance struct {
	Description string   `xml:"description"`
	Related     []string `xml:"related"`
}

type XMLControlEnhancement struct {
	Number         string   `xml:"number"`
	Title          string   `xml:"title"`
	BaselineImpact []string `xml:"baseline-impact"`
	Withdrawn      struct {
		IncorporatedInto string `xml:"incorporated-into"`
	} `xml:"withdrawn"`
	Statements []XMLStatement `xml:"statement"`
}

type XMLReference struct {
	Item string `xml:"item"`
	Link string `xml:"href,attr"`
}

func main() {
	resp, err := http.Get(standardURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var xmlStandard XMLStandard
	if err := xml.Unmarshal(data, &xmlStandard); err != nil {
		panic(err)
	}

	controls := map[string]Control{}

	for _, xmlControl := range xmlStandard.Controls {
		// if !parseBaseline("HIGH", xmlControl.BaselineImpact) {
		//	continue
		//}

		family := parseFamily(xmlControl.Family)
		name := strings.Title(strings.ToLower(xmlControl.Title))
		description := "\"" + strings.TrimSuffix(parseDescription(family, xmlControl.Statements), "\n") + "\""

		control := Control{
			Family:      family,
			Name:        name,
			Description: description,
		}

		controls[xmlControl.Number] = control

		for _, controlEnhancement := range xmlControl.ControlEnhancements {
			//if !parseBaseline("HIGH", controlEnhancement.BaselineImpact) {
			//	continue
			//}

			name := strings.Title(strings.ToLower(controlEnhancement.Title))
			description := "\"" + strings.TrimSuffix(parseDescription(family, controlEnhancement.Statements), "\n") + "\""

			control := Control{
				Family:      family,
				Name:        name,
				Description: description,
			}

			controls[controlEnhancement.Number] = control
		}
	}

	standard := Standard{Name: "NIST-800-53", Controls: controls}

	y, err := yaml.Marshal(&standard)
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(os.Args[1], y, 0644); err != nil {
		panic(err)
	}
}

func parseFamily(family string) string {
	switch family {
	case "ACCESS CONTROL":
		return "AC"
	case "AUDIT AND ACCOUNTABILITY":
		return "AU"
	case "AWARENESS AND TRAINING":
		return "AT"
	case "CONFIGURATION MANAGEMENT":
		return "CM"
	case "CONTINGENCY PLANNING":
		return "CP"
	case "IDENTIFICATION AND AUTHENTICATION":
		return "IA"
	case "INCIDENT RESPONSE":
		return "IR"
	case "MAINTENANCE":
		return "MA"
	case "MEDIA PROTECTION":
		return "MP"
	case "PERSONNEL SECURITY":
		return "PS"
	case "PHYSICAL AND ENVIRONMENTAL PROTECTION":
		return "PE"
	case "PLANNING":
		return "PL"
	case "PROGRAM MANAGEMENT":
		return "PM"
	case "RISK ASSESSMENT":
		return "RA"
	case "SECURITY ASSESSMENT AND AUTHORIZATION":
		return "CA"
	case "SYSTEM AND COMMUNICATIONS PROTECTION":
		return "SC"
	case "SYSTEM AND INFORMATION INTEGRITY":
		return "SI"
	case "SYSTEM AND SERVICES ACQUISITION":
		return "SA"
	default:
		return ""
	}
}

func parseBaseline(baseline string, baselines []string) bool {
	for _, b := range baselines {
		if b == baseline {
			return true
		}
	}
	return false
}

func parseDescription(family string, statements []XMLStatement) string {
	var description string
	for _, s := range statements {
		var number string
		if s.Number != "" {
			re := regexp.MustCompile(fmt.Sprintf("%s-[0-9]+", family))
			lenNumber := len(re.FindStringSubmatch(s.Number)[0])
			number = strings.TrimSuffix(s.Number[lenNumber:len(s.Number)], ".")
			chars := strings.Split(number, ".")

			if len(chars) > 1 {
				for _, char := range chars[1:] {
					description += fmt.Sprintf("    %s.  %s\n", char, s.Description)

				}
			} else {
				description += fmt.Sprintf("  %s.  %s\n", chars[0], s.Description)
			}

		} else {
			description += fmt.Sprintf("%s\n", s.Description)
		}
		description += parseDescription(family, s.Statements)

	}

	return description
}
