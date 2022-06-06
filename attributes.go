package vitotrol

import (
	"fmt"
	"strconv"
)

// An AttrID defines an attribute ID.
type AttrID uint16

// Attribute IDs currently supported by the library. For each, the
// Vitotrol™ name.
const (
	AussenTemp             AttrID = 5373   // temp_ats_r
	AbgasTemp              AttrID = 5372   // temp_agt_r
	BoilerTemp             AttrID = 5374   // temp_kts_r
	HeisswasserTemp        AttrID = 5381   // temp_ww_r
	HeisswasserAusgangTemp AttrID = 5382   // temp_auslauf_r
	HeizwasserAusgangTemp  AttrID = 6053   // temp_vts_r
	HeizNormalTempM1       AttrID = 82     // konf_raumsolltemp_rw
	HeizNormalTempM2       AttrID = 83     // konf_raumsolltemp_rw
	HeizPartyTempM1        AttrID = 79     // konf_partysolltemp_rw
	HeizPartyTempM2        AttrID = 80     // konf_partysolltemp_rw
	HeizReduziertTempM1    AttrID = 85     // konf_raumsolltemp_reduziert_rw
	HeizReduziertTempM2    AttrID = 86     // konf_raumsolltemp_reduziert_rw
	HeisswasserSollTemp    AttrID = 51     // konf_ww_solltemp_rw
	AnzahlBrennerstunden   AttrID = 104    // anzahl_brennerstunden_r
	BrennerStatus          AttrID = 600    // zustand_brenner_r
	BurnerStarts           AttrID = 111    // anzahl_brennerstart_r
	InternalPumpStatus     AttrID = 245    // zustand_interne_pumpe_r
	HeatingPumpStatus      AttrID = 729    // zustand_heizkreispumpe_r
	CirculationPumpState   AttrID = 7181   // zustand_zirkulationspumpe_r
	PartyMode              AttrID = 7855   // konf_partybetrieb_rw
	EnergySavingMode       AttrID = 7852   // konf_sparbetrieb_rw
	DateTime               AttrID = 5385   // konf_uhrzeit_rw
	CurrentError           AttrID = 7184   // aktuelle_fehler_r
	HolidaysStart          AttrID = 306    // konf_ferien_start_rw
	HolidaysEnd            AttrID = 309    // konf_ferien_ende_rw
	HolidaysStatus         AttrID = 714    // zustand_ferienprogramm_r
	Way3ValveStatus        AttrID = 5389   // info_status_umschaltventil_r
	OperatingModeCurrent   AttrID = 708    // aktuelle_betriebsart_r
	FrostProtectionStatus  AttrID = 717    // zustand_frostgefahr_r
	BetriebsartM1          AttrID = 92     // konf_betriebsart_rw
	BetriebsartM2          AttrID = 94     // konf_betriebsart_rw
	NeigungM1              AttrID = 2869   // konf_neigung_rw
	NeigungM2              AttrID = 2871   // konf_neigung_rw
	NiveauM1               AttrID = 2875   // konf_niveau_rw
	NiveauM2               AttrID = 2877   // konf_niveau_rw
	NoAttr                 AttrID = 0xffff // Used in error cases
)

// An AttrAccess defines attributes access rights.
type AttrAccess uint8

// Availables access rights.
const (
	ReadOnly AttrAccess = 1 << iota
	WriteOnly
	ReadWrite AttrAccess = ReadOnly | WriteOnly
)

// AccessToStr map allows to translate AttrAccess values to strings.
var AccessToStr = map[AttrAccess]string{
	ReadOnly:  "read-only",
	WriteOnly: "write-only",
	ReadWrite: "read/write",
}

// An AttrRef describes an attribute reference: its type, access and name.
type AttrRef struct {
	Type   VitodataType
	Access AttrAccess
	Name   string
	Doc    string
	Custom bool
}

// String returns all information contained in this attribute reference.
func (r *AttrRef) String() string {
	return fmt.Sprintf("%s: %s (%s - %s)",
		r.Name, r.Doc, r.Type.Type(), AccessToStr[r.Access])
}

// AttributesRef lists the reference for each attribute ID.
var AttributesRef = map[AttrID]*AttrRef{
	AussenTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Außen Temperatur",
		Name:   "AussenTemp",
	},
	AbgasTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Abgas Temperatur",
		Name:   "AbgasTemp",
	},
	BoilerTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Boiler Temperatur",
		Name:   "BoilerTemp",
	},
	HeisswasserTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Heißwasser Temperatur",
		Name:   "HeisswasserTemp",
	},
	HeisswasserAusgangTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Heißwasser Ausgangstemperatur",
		Name:   "HeisswasserAusgangTemp",
	},
	HeizwasserAusgangTemp: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Heizwasser Ausgangstemperatur",
		Name:   "HeizwasserAusgangTemp",
	},
	HeizNormalTempM1: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Normale Raumsolltemperatur Heizkörper",
		Name:   "HeizNormalTempM1",
	},
	HeizNormalTempM2: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Normale Raumsolltemperatur Fußbodenheizung",
                Name:   "HeizNormalTempM2",
        },
	HeizPartyTempM1: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Party Raumtemperatur Heizkörper",
		Name:   "HeizPartyTempM1",
	},
        HeizPartyTempM2: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Party Raumtemperatur Fußbodenheizung",
                Name:   "HeizPartyTempM2",
        },
	HeizReduziertTempM1: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Reduzierte Raumtemperatur Heizkörper",
		Name:   "HeizReduziertTempM1",
	},
        HeizReduziertTempM2: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Reduzierte Raumtemperatur Fußbodenheizung",
                Name:   "HeizReduziertTempM2",
        },
	HeisswasserSollTemp: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Solltemperatur Warmwasser",
		Name:   "HeisswasserSollTemp",
	},
	AnzahlBrennerstunden: {
		Type:   TypeDouble,
		Access: ReadOnly,
		Doc:    "Brennerstundenanzahl",
		Name:   "AnzahlBrennerstunden",
	},
	BrennerStatus: {
		Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
		Access: ReadOnly,
		Doc:    "Brenner Status",
		Name:   "BrennerStatus",
	},
	BurnerStarts: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Burner starts",
		Name:   "BurnerStarts",
	},
	InternalPumpStatus: {
		Type: NewEnum([]string{ // 0 -> 3
			"off",
			"on",
			"off2",
			"on2",
		}),
		Access: ReadOnly,
		Doc:    "Internal pump status",
		Name:   "InternalPumpStatus",
	},
	HeatingPumpStatus: {
		Type:   TypeOnOffEnum,
		Access: ReadOnly,
		Doc:    "Heating pump status",
		Name:   "HeatingPumpStatus",
	},
	CirculationPumpState: {
		Type:   TypeOnOffEnum,
		Access: ReadOnly,
		Doc:    "Statut pompe circulation",
		Name:   "CirculationPumpState",
	},
	PartyMode: {
		Type:   TypeEnabledEnum,
		Access: ReadWrite,
		Doc:    "Party mode status",
		Name:   "PartyMode",
	},
	EnergySavingMode: {
		Type:   TypeEnabledEnum,
		Access: ReadWrite,
		Doc:    "Energy saving mode status",
		Name:   "EnergySavingMode",
	},
	DateTime: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Current date and time",
		Name:   "DateTime",
	},
	CurrentError: {
		Type:   TypeString,
		Access: ReadOnly,
		Doc:    "Current error",
		Name:   "CurrentError",
	},
	HolidaysStart: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Holidays begin date",
		Name:   "HolidaysStart",
	},
	HolidaysEnd: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Holidays end date",
		Name:   "HolidaysEnd",
	},
	HolidaysStatus: {
		Type:   TypeEnabledEnum,
		Access: ReadOnly,
		Doc:    "Holidays program status",
		Name:   "HolidaysStatus",
	},
	Way3ValveStatus: {
		Type: NewEnum([]string{ // 0 -> 3
			"undefined",
			"heating",
			"middle position",
			"hot water",
		}),
		Access: ReadOnly,
		Doc:    "3-way valve status",
		Name:   "3WayValveStatus",
	},
	OperatingModeCurrent: {
		Type: NewEnum([]string{ // 0 -> 3
			"stand-by",
			"reduced",
			"normal",
			"continuous normal",
		}),
		Access: ReadOnly,
		Doc:    "Operating mode",
		Name:   "OperatingModeCurrent",
	},
	FrostProtectionStatus: {
		Type:   TypeEnabledEnum,
		Access: ReadOnly,
		Doc:    "Frost protection status",
		Name:   "FrostProtectionStatus",
	},
	BetriebsartM1: {
                Type: NewEnum([]string{ // 0 -> 4
                        "Abschalt",
                        "Nur WW",
                        "Heizen + WW",
                        "Dauernd Reduziert",
                        "Dauernd Normal",
                }),
                Access: ReadWrite,
                Doc:    "Betriebsart Heizkörper",
                Name:   "BetriebsartM1",
        },
	BetriebsartM2: {
		Type: NewEnum([]string{ // 0 -> 4
			"Abschalt",
			"Nur WW",
			"Heizen + WW",
			"Dauernd Reduziert",
			"Dauernd Normal",
		}),
		Access: ReadWrite,
		Doc:    "Betriebsart Fussbodenheizung",
		Name:   "konf_betriebsart_rw-0x005e",
	},
	NeigungM1: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Neigung Heizkörper",
                Name:   "NeigungM1",
	},
        NeigungM2: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Neigung Fussbodenheizung",
                Name:   "NeigungM2",
        },
	NiveauM1: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Niveau Heizkörper",
                Name:   "NiveauM1",
        },
	NiveauM2: {
                Type:   TypeDouble,
                Access: ReadWrite,
                Doc:    "Niveau Fussbodenheizung",
                Name:   "NiveauM2",
        },

}

// AddAttributeRef adds a new attribute to the "official" list. This
// new attribute will only differ from others by its Custom field set
// to true.
//
// No check is done to avoid overriding existing attributes.
func AddAttributeRef(attrID AttrID, ref AttrRef) {
	ref.Custom = true
	AttributesRef[attrID] = &ref

	AttributesNames2IDs = computeNames2IDs()
	Attributes = computeAttributes()
}

func computeNames2IDs() map[string]AttrID {
	ret := make(map[string]AttrID, len(AttributesRef))
	for attrID, pAttrRef := range AttributesRef {
		ret[pAttrRef.Name] = attrID
	}
	return ret
}

func computeAttributes() []AttrID {
	ret := make([]AttrID, 0, len(AttributesRef))
	for attrID := range AttributesRef {
		ret = append(ret, attrID)
	}
	return ret
}

// AttributesNames2IDs maps the attributes names to their AttrID
// counterpart.
var AttributesNames2IDs = computeNames2IDs()

// Attributes lists the AttrIDs for all available attributes.
var Attributes = computeAttributes()

// Value is the timestamped value of an attribute.
type Value struct {
	Value string
	Time  Time
}

// Num returns the numerical value of this value. If the value is not
// a numerical one, 0 is returned.
func (v *Value) Num() (ret float64) {
	ret, _ = strconv.ParseFloat(v.Value, 64)
	return
}
