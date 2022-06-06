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
	AnzahlBrennerStarts    AttrID = 111    // anzahl_brennerstart_r
	InternerPumpenStatus   AttrID = 245    // zustand_interne_pumpe_r
	HeizPumpenStatusM1     AttrID = 729    // zustand_heizkreispumpe_r
	HeizPumpenStatusM2     AttrID = 730    // zustand_heizkreispumpe_r
	ZirkPumpenStatus       AttrID = 7181   // zustand_zirkulationspumpe_r
	PartyModusM1           AttrID = 7855   // konf_partybetrieb_rw
	PartyModusM2           AttrID = 7856   // konf_partybetrieb_rw
	EnergieSparmodusM1     AttrID = 7852   // konf_sparbetrieb_rw
	EnergieSparmodusM2     AttrID = 7853   // konf_sparbetrieb_rw
	DatumUhrzeit           AttrID = 5385   // konf_uhrzeit_rw
	AktuellerFehler        AttrID = 7184   // aktuelle_fehler_r
	FerienStartM1          AttrID = 306    // konf_ferien_start_rw
	FerienStartM2          AttrID = 307    // konf_ferien_start_rw
	FerienEndeM1           AttrID = 309    // konf_ferien_ende_rw
	FerienEndeM2           AttrID = 310    // konf_ferien_ende_rw
	ZustandFerienProgM1    AttrID = 714    // zustand_ferienprogramm_r
	ZustandFerienProgM2    AttrID = 715    // zustand_ferienprogramm_r
	AktuelleBetriebsartM1  AttrID = 708    // aktuelle_betriebsart_r
	AktuelleBetriebsartM2  AttrID = 709    // aktuelle_betriebsart_r
	ZustandFrostgefahrM1   AttrID = 717    // zustand_frostgefahr_r
	ZustandFrostgefahrM2   AttrID = 718    // zustand_frostgefahr_r
	BetriebsartM1          AttrID = 92     // konf_betriebsart_rw
	BetriebsartM2          AttrID = 94     // konf_betriebsart_rw
	NeigungM1              AttrID = 2869   // konf_neigung_rw
	NeigungM2              AttrID = 2871   // konf_neigung_rw
	NiveauM1               AttrID = 2875   // konf_niveau_rw
	NiveauM2               AttrID = 2877   // konf_niveau_rw
	Heizungsschema         AttrID = 801    // konf_heizungsschema_r
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
	AnzahlBrennerStarts: {
		Type:   TypeDouble,
		Access: ReadWrite,
		Doc:    "Anzahl von Brennerstarts",
		Name:   "AnzahlBrennerStarts",
	},
	InternerPumpenStatus: {
		Type: NewEnum([]string{ // 0 -> 3
			"Aus",
			"Ein",
			"Aus2",
			"Ein2",
		}),
		Access: ReadOnly,
		Doc:    "Interner Pumpen Status",
		Name:   "InternerPumpenStatus",
	},
	HeizPumpenStatusM1: {
		Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
		Access: ReadOnly,
		Doc:    "Zustand Heizkreispumpe Heizkörper",
		Name:   "HeizPumpenStatusM1",
	},
	HeizPumpenStatusM2: {
                Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
                Access: ReadOnly,
                Doc:    "Zustand Heizkreispumpe Hußbodenheizung",
                Name:   "HeizPumpenStatusM2",
        },
	ZirkPumpenStatus: {
		Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
		Access: ReadOnly,
		Doc:    "Zustand Zirkulationspumpe",
		Name:   "ZirkPumpenStatus",
	},
	PartyModusM1: {
		Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
		Access: ReadWrite,
		Doc:    "Partymodus Heizkörper",
		Name:   "PartyModusM1",
	},
	PartyModusM2: {
                Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
                Access: ReadWrite,
                Doc:    "Partymodus Fußbodenheizung",
                Name:   "PartyModusM2",
        },
	EnergieSparmodusM1: {
		Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
		Access: ReadWrite,
		Doc:    "Energiesparmodus Heizkörper",
		Name:   "EnergieSparmodusM1",
	},
	EnergieSparmodusM2: {
                Type: NewEnum([]string{ // 0 -> 1
                        "Aus",
                        "Ein",
                }),
                Access: ReadWrite,
                Doc:    "Energiesparmodus Fußbodenheizung",
                Name:   "EnergieSparmodusM2",
        },
	DatumUhrzeit: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Aktuelles Datum mir Uhrzeit",
		Name:   "DatumUhrzeit",
	},
	AktuellerFehler: {
		Type:   TypeString,
		Access: ReadOnly,
		Doc:    "Fehlermeldung",
		Name:   "AktuellerFehler",
	},
	FerienStartM1: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Start der Ferienzeit für Heizkörper",
		Name:   "FerienStartM1",
	},
	FerienStartM2: {
                Type:   TypeDate,
                Access: ReadWrite,
                Doc:    "Start der Ferienzeit für Fußbodenheizung",
                Name:   "FerienStartM2",
        },
	FerienEndeM1: {
		Type:   TypeDate,
		Access: ReadWrite,
		Doc:    "Ende der Ferienzeit für Heizkörper",
		Name:   "FerienEndeM1",
	},
	FerienEndeM2: {
                Type:   TypeDate,
                Access: ReadWrite,
                Doc:    "Ende der Ferienzeit für Fußbodenheizung",
                Name:   "FerienEndeM2",
        },
	ZustandFerienProgM1: {
		Type: NewEnum([]string{ // 0 -> 1
                        "inaktiv",
                        "aktiv",
                }),
		Access: ReadOnly,
		Doc:    "Zustand Ferienprogramm Heizkörper",
		Name:   "ZustandFerienProgM1",
	},
	ZustandFerienProgM2: {
                Type: NewEnum([]string{ // 0 -> 1
                        "inaktiv",
                        "aktiv",
                }),
                Access: ReadOnly,
                Doc:    "Zustand Ferienprogramm Fußbodenheizung",
                Name:   "ZustandFerienProgM2",
        },
	AktuelleBetriebsartM1: {
		Type: NewEnum([]string{ // 0 -> 3
			"Abschaltbetrieb",
			"Reduzierter Betrieb",
			"Normalbetrieb",
			"Dauernd Normalbetrieb",
		}),
		Access: ReadOnly,
		Doc:    "Betriebsart Heizkörper",
		Name:   "AktuelleBetriebsartM1",
	},
	AktuelleBetriebsartM2: {
                Type: NewEnum([]string{ // 0 -> 3
                        "Abschaltbetrieb",
                        "Reduzierter Betrieb",
                        "Normalbetrieb",
                        "Dauernd Normalbetrieb",
                }),
                Access: ReadOnly,
                Doc:    "Betriebsart Fußbodenheizung",
                Name:   "AktuelleBetriebsartM2",
        },
	ZustandFrostgefahrM1: {
		Type: NewEnum([]string{ // 0 -> 1
                        "inaktiv",
                        "aktiv",
                }),
		Access: ReadOnly,
		Doc:    "Zustand Frostgefahr Heizkörper",
		Name:   "ZustandFrostgefahrM1",
	},
	ZustandFrostgefahrM2: {
                Type: NewEnum([]string{ // 0 -> 1
                        "inaktiv",
                        "aktiv",
                }),
                Access: ReadOnly,
                Doc:    "Zustand Frostgefahr Fußbodenheizung",
                Name:   "ZustandFrostgefahrM2",
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
	Heizungsschema: {
                Type: NewEnum([]string{ // 0 -> 10
                        "",
                        "1 A1",
                        "2 A1 + WW",
                        "3 M2",
                        "4 M2 + WW",
                        "5 A1 + M2",
                        "6 A1 + M2 + WW",
                        "7 M2 + M3",
                        "8 M2 + M3 + WW",
                        "9 A1 + M2 + M3",
                        "10 A1 + M2 + M3 + WW",
                }),
                Access: ReadOnly,
                Doc:    "Heizungsschema für Anlage",
                Name:   "Heizungsschema",
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
