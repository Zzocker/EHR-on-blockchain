package main

// DocTyp of blockchain data
// used for doing rich query
const (
	REPORT    = "REPORT"
	DRUGS     = "DRUGS"
	TESTS     = "TESTS"
	TREATMENT = "TREATMENT"
	CONSENT   = "CONSENT"
)

// Report of patient
type Report struct {
	DocTyp    string `json:"docTyp"`
	ID        string `json:"report_id"`
	PatientID string `json:"patient_id"`
	// DrugsID []string `json:"drugs_id"`      			///
	// TreatmentID []string `json:"treatment_id"`		/// these will be stored
	// TestID []string `json:"test_id"`					///
	Status      string            `json:"status"`
	RefDoctorID string            `json:"doctor_id"`
	Comments    map[string]string `json:"comments"`
	CreateTime  int64             `json:"create_time"`
	UpdateTime  int64             `json:"updated_time"`
}

// Drugs model
type Drugs struct {
	DocTyp     string            `json:"docTyp"`
	ReportID   string            `json:"report_id"`
	ID         string            `json:"drugs_id"`
	For        string            `json:"patient__id"`
	RefDoctor  string            `json:"ref_doctor"`
	Drug       map[string]string `josn:"drug"`   // name of drug mapped to doses
	Status     int               `json:"status"` // 0 - requested 1-  given
	Pending    map[string]string // name of ignored drugs mapped to when will that be 	available
	CreateTime int64             `json:"create_time"`
	UpdateTime int64             `josn:"updated_time"`
}

// Test model file
type Test struct {
	DocTyp            string   `json:"docTyp"`
	ReportID          string   `json:"report_id"`
	ID                string   `json:"test_id"`
	PatientID         string   `json:"patient_id"`
	MediaFileLocation []string `json:"media_file_location"`
	Name              string   `json:"test_name"`
	Supervisor        string   `json:"supervisor_details"` // this will name of supervisor, aadress , path Lab
	RefDoctor         string   `json:"ref_doctor"`
	Result            string   `json:"test_result"`
	Status            int      `json:"status"`       // status of test 0 - not done 1 - done
	TypeOfT           int      `json:"type_of_test"` // 0- normal 1-abnormal
	CreateTime        int64    `json:"create_time"`
	UpdateTime        int64    `josn:"updated_time"`
}

// Treatment model
type Treatment struct {
	DocTyp            string            `json:"docTyp"`
	PatientID         string            `json:"patient_id"`
	ReportID          string            `json:"report_id"`
	ID                string            `json:"treatment_id"`
	Supervisor        string            `json:"supervisor_details"` // deatils of nurses , doctor
	RefDoctor         string            `json:"ref_doctor"`
	Name              string            `josn:"treatment_name"`
	MediaFileLocation []string          `josn:"media_file_location"`
	Comments          map[string]string `json:"comments"`
	Status            int               // 0 not done 1 started 2  done 3 failed
	CreateTime        int64             `json:"create_time"`
	UpdateTime        int64             `josn:"updated_time"`
}

// Consent model file
type Consent struct {
	DocTyp              string           `json:"docTyp"`
	ID                  string           `json:"patient_aadhaar"`
	PermanentConsenters map[string]bool  `json:"parma_consenters"` // list of permanent consenter
	TemporaryConsenters map[string]int64 `josn:"temp_consenters"`  // id of consenters mapped to expiry time unix
// 	Status              string           `json:"status"`           // defined status crises status
// 	Track               []string         `json:"track"`            // track to who he/she meet (id)
}
