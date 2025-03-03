package goft

import "time"

// This code is generated and then modified as need. Be careful using it as I'm fixing the structs as I use an need them

type _ struct {
	Accreditations        Accreditation       `json:"accreditations"`
	Achievements          Achievement         `json:"achievements"`
	AchievementsUsers     AchievementUser     `json:"achievements_users"`
	Announcements         Announcement        `json:"announcements"`
	Apps                  App                 `json:"apps"`
	Attachments           Attachment          `json:"attachments"`
	Balances              Balance             `json:"balances"`
	BlocDeadlines         BlocDeadlines       `json:"bloc_deadlines"`
	Blocs                 Blocs               `json:"blocs"`
	Campus                Campus              `json:"campus"`
	CampusUsers           CampusUser          `json:"campus_users"`
	Closes                Close               `json:"closes"`
	Coalitions            Coalition           `json:"coalitions"`
	CoalitionsUsers       CoalitionsUser      `json:"coalitions_users"`
	CommunityServices     CommunityService    `json:"community_services"`
	Cursus                Cursus              `json:"cursus"`
	CursusUsers           CursusUser          `json:"cursus_users"`
	DashesUsers           DasheUser           `json:"dashes_users"`
	Endpoints             Endpoint            `json:"endpoints"`
	Events                Event               `json:"events"`
	Exams                 Exam                `json:"exams"`
	Expertises            Expertise           `json:"expertises"`
	ExpertisesUsers       ExpertiseUser       `json:"expertises_users"`
	FlashUsers            FlashUser           `json:"flash_users"`
	Flashes               Flashe              `json:"flashes"`
	Groups                Group               `json:"groups"`
	GroupsUsers           GroupUser           `json:"groups_users"`
	Internships           Internship          `json:"internships"`
	Languages             Language            `json:"languages"`
	LanguagesUsers        LanguageUser        `json:"languages_users"`
	Locations             Location            `json:"locations"`
	Mailings              Mailing             `json:"mailings"`
	Notes                 Note                `json:"notes"`
	Notions               Notion              `json:"notions"`
	Partnerships          Partnership         `json:"partnerships"`
	PartnershipsUsers     PartnershipUser     `json:"partnerships_users"`
	Patronages            Patronage           `json:"patronages"`
	PatronagesReports     PatronageReport     `json:"patronages_reports"`
	Pools                 Pool                `json:"pools"`
	Products              Product             `json:"products"`
	ProjectData           ProjectData         `json:"project_data"`
	ProjectSessions       ProjectSession      `json:"project_sessions"`
	ProjectSessionsSkills ProjectSessionSkill `json:"project_sessions_skills"`
	Projects              Project             `json:"projects"`
	ProjectsUsers         ProjectUser         `json:"projects_users"`
	Quests                Quest               `json:"quests"`
	QuestsUsers           QuestUser           `json:"quests_users"`
	Roles                 Role                `json:"roles"`
	RolesEntities         RoleEntity          `json:"roles_entities"`
	Rules                 Rule                `json:"rules"`
	ScaleTeams            ScaleTeam           `json:"scale_teams"`
	Scales                Scale               `json:"scales"`
	Scores                Score               `json:"scores"`
	Subnotions            Subnotion           `json:"subnotions"`
	Tags                  Tag                 `json:"tags"`
	TeamsUploads          TeamUpload          `json:"teams_uploads"`
	TeamsUsers            TeamUser            `json:"teams_users"`
	Transactions          Transaction         `json:"transactions"`
	Translations          Translation         `json:"translations"`
	UserCandidatures      UserCandidature     `json:"user_candidatures"`
	Users                 User                `json:"users"`
	Waitlists             Waitlist            `json:"waitlists"`
}

type User struct {
	Achievements    []Achievement   `json:"achievements"`
	Active          bool            `json:"active?"`
	Alumni          bool            `json:"alumni?"`
	AnonymizeDate   string          `json:"anonymize_date"`
	Campus          []UserCampus    `json:"campus"`
	CampusUsers     []CampusUser    `json:"campus_users"`
	CorrectionPoint int64           `json:"correction_point"`
	CursusUsers     []CursusUser    `json:"cursus_users"`
	DataErasureDate any             `json:"data_erasure_date"`
	Displayname     string          `json:"displayname"`
	Email           string          `json:"email"`
	ExpertisesUsers []ExpertiseUser `json:"expertises_users"`
	FirstName       string          `json:"first_name"`
	Groups          []any           `json:"groups"`
	ID              int64           `json:"id"`
	Image           Image           `json:"image"`
	Kind            string          `json:"kind"`
	LanguagesUsers  []LanguageUser  `json:"languages_users"`
	LastName        string          `json:"last_name"`
	Location        any             `json:"location"`
	Login           string          `json:"login"`
	Partnerships    []Partnership   `json:"partnerships"`
	Patroned        []Patronage     `json:"patroned"`
	Patroning       []Patronage     `json:"patroning"`
	Phone           any             `json:"phone"`
	PoolMonth       string          `json:"pool_month"`
	PoolYear        string          `json:"pool_year"`
	ProjectsUsers   []ProjectUser   `json:"projects_users"`
	Roles           []Role          `json:"roles"`
	Staff           bool            `json:"staff?"`
	Titles          []Title         `json:"titles"`
	TitlesUsers     []TitleUser     `json:"titles_users"`
	URL             string          `json:"url"`
	UsualFirstName  string          `json:"usual_first_name"`
	UsualFullName   string          `json:"usual_full_name"`
	Wallet          int64           `json:"wallet"`
}

type Title struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TitleUser struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	TitleId   int       `json:"title_id"`
	Selected  bool      `json:"selected"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Achievement struct {
	Achievements []any  `json:"achievements"`
	Description  string `json:"description"`
	ID           int64  `json:"id"`
	Image        string `json:"image"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	NbrOfSuccess int64  `json:"nbr_of_success"`
	Parent       any    `json:"parent"`
	Tier         string `json:"tier"`
	Title        any    `json:"title"`
	UsersURL     string `json:"users_url"`
	Visible      bool   `json:"visible"`
}

type Campus struct {
	Active             bool     `json:"active"`
	Address            string   `json:"address"`
	City               string   `json:"city"`
	Country            string   `json:"country"`
	DefaultHiddenPhone bool     `json:"default_hidden_phone"`
	EmailExtension     string   `json:"email_extension"`
	Endpoint           Endpoint `json:"endpoint"`
	Facebook           string   `json:"facebook"`
	ID                 int64    `json:"id"`
	Language           Language `json:"language"`
	Name               string   `json:"name"`
	Public             bool     `json:"public"`
	TimeZone           string   `json:"time_zone"`
	Twitter            string   `json:"twitter"`
	UsersCount         int64    `json:"users_count"`
	VogsphereID        int64    `json:"vogsphere_id"`
	Website            string   `json:"website"`
	Zip                string   `json:"zip"`
}

type Internship struct {
	AdministrationID         int64           `json:"administration_id"`
	BreachAt                 any             `json:"breach_at"`
	CompanyAddress           string          `json:"company_address"`
	CompanyBossUserEmail     string          `json:"company_boss_user_email"`
	CompanyBossUserFirstName string          `json:"company_boss_user_first_name"`
	CompanyBossUserLastName  string          `json:"company_boss_user_last_name"`
	CompanyBossUserPhone     string          `json:"company_boss_user_phone"`
	CompanyCity              string          `json:"company_city"`
	CompanyCountry           string          `json:"company_country"`
	CompanyName              string          `json:"company_name"`
	CompanyPostal            string          `json:"company_postal"`
	CompanySiret             string          `json:"company_siret"`
	CompanyUserEmail         string          `json:"company_user_email"`
	CompanyUserFirstName     string          `json:"company_user_first_name"`
	CompanyUserLastName      string          `json:"company_user_last_name"`
	CompanyUserPhone         string          `json:"company_user_phone"`
	CompanyUserPost          string          `json:"company_user_post"`
	ContractType             string          `json:"contract_type"`
	Convention               Convention      `json:"convention"`
	Currency                 string          `json:"currency"`
	Days                     string          `json:"days"`
	Duration                 int64           `json:"duration"`
	EndAt                    time.Time       `json:"end_at"`
	ID                       int64           `json:"id"`
	InternshipAddress        string          `json:"internship_address"`
	InternshipCity           string          `json:"internship_city"`
	InternshipCountry        string          `json:"internship_country"`
	InternshipPostal         string          `json:"internship_postal"`
	LanguageID               int64           `json:"language_id"`
	Movement                 any             `json:"movement"`
	NbDays                   int64           `json:"nb_days"`
	NbHours                  int64           `json:"nb_hours"`
	OfferID                  any             `json:"offer_id"`
	Salary                   int64           `json:"salary"`
	StartAt                  time.Time       `json:"start_at"`
	State                    string          `json:"state"`
	Subject                  string          `json:"subject"`
	User                     TransactionUser `json:"user"`
	UserAddress              string          `json:"user_address"`
	UserCity                 string          `json:"user_city"`
	UserCountry              string          `json:"user_country"`
	UserPostal               string          `json:"user_postal"`
}

type QuestUser struct {
	Advancement any             `json:"advancement"`
	CreatedAt   time.Time       `json:"created_at"`
	EndAt       time.Time       `json:"end_at"`
	ID          int64           `json:"id"`
	Prct        any             `json:"prct"`
	Quest       Quest           `json:"quest"`
	QuestID     int64           `json:"quest_id"`
	UpdatedAt   time.Time       `json:"updated_at"`
	User        TransactionUser `json:"user"`
	ValidatedAt time.Time       `json:"validated_at"`
}

type PatronageReport struct {
	Answers     []any           `json:"answers"`
	BeginAt     time.Time       `json:"begin_at"`
	CreatedAt   time.Time       `json:"created_at"`
	ID          int64           `json:"id"`
	Patronage   Patronage       `json:"patronage"`
	PatronageID int64           `json:"patronage_id"`
	Report      Report          `json:"report"`
	ReportID    int64           `json:"report_id"`
	UpdatedAt   time.Time       `json:"updated_at"`
	User        TransactionUser `json:"user"`
	UserID      int64           `json:"user_id"`
	ValidatedAt time.Time       `json:"validated_at"`
}

type Mailing struct {
	Attachment  any       `json:"attachment"`
	Attachments any       `json:"attachments"`
	Bcc         []any     `json:"bcc"`
	Cc          []any     `json:"cc"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	From        string    `json:"from"`
	HTMLContent string    `json:"html_content"`
	ID          int64     `json:"id"`
	Identifier  string    `json:"identifier"`
	Meta        Meta      `json:"meta"`
	Subject     string    `json:"subject"`
	Subtitle    any       `json:"subtitle"`
	Title       string    `json:"title"`
	To          []string  `json:"to"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Project struct {
	Attachments     []any            `json:"attachments"`
	Campus          []UserCampus     `json:"campus"`
	Children        []any            `json:"children"`
	CreatedAt       time.Time        `json:"created_at"`
	Cursus          []Cursus         `json:"cursus"`
	Description     string           `json:"description"`
	Difficulty      int64            `json:"difficulty"`
	Exam            bool             `json:"exam"`
	ID              int64            `json:"id"`
	Name            string           `json:"name"`
	Objectives      []string         `json:"objectives"`
	Parent          any              `json:"parent"`
	ProjectSessions []ProjectSession `json:"project_sessions"`
	Skills          []Skill          `json:"skills"`
	Slug            string           `json:"slug"`
	Tags            []Tag            `json:"tags"`
	UpdatedAt       time.Time        `json:"updated_at"`
	Videos          []any            `json:"videos"`
}

type Subnotion struct {
	Attachments []any     `json:"attachments"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Notepad     any       `json:"notepad"`
	Notion      any       `json:"notion"`
	Slug        string    `json:"slug"`
}

type Announcement struct {
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
	ID        int64     `json:"id"`
	Kind      string    `json:"kind"`
	Link      any       `json:"link"`
	Text      string    `json:"text"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Attachment struct {
	BaseID        int64         `json:"base_id"`
	CreatedAt     time.Time     `json:"created_at"`
	ID            int64         `json:"id"`
	Language      Language      `json:"language"`
	Name          string        `json:"name"`
	PageCount     int64         `json:"page_count"`
	Pdf           AttachmentPdf `json:"pdf"`
	PdfProcessing bool          `json:"pdf_processing"`
	Slug          string        `json:"slug"`
	ThumbURL      any           `json:"thumb_url"`
	Type          string        `json:"type"`
	URL           any           `json:"url"`
}

type ProjectSession struct {
	BeginAt          time.Time    `json:"begin_at"`
	Campus           any          `json:"campus"`
	CampusID         any          `json:"campus_id"`
	CreatedAt        time.Time    `json:"created_at"`
	Cursus           any          `json:"cursus"`
	CursusID         any          `json:"cursus_id"`
	DurationDays     any          `json:"duration_days"`
	EndAt            time.Time    `json:"end_at"`
	EstimateTime     int64        `json:"estimate_time"`
	Evaluations      []Evaluation `json:"evaluations"`
	ID               int64        `json:"id"`
	IsSubscriptable  bool         `json:"is_subscriptable"`
	MaxPeople        any          `json:"max_people"`
	Project          Project      `json:"project"`
	ProjectID        int64        `json:"project_id"`
	Scales           []Scale      `json:"scales"`
	Solo             bool         `json:"solo"`
	TeamBehaviour    string       `json:"team_behaviour"`
	TerminatingAfter any          `json:"terminating_after"`
	UpdatedAt        time.Time    `json:"updated_at"`
	Uploads          []any        `json:"uploads"`
}

type Product struct {
	BeginAt         time.Time    `json:"begin_at"`
	CategoryID      int64        `json:"category_id"`
	CreatedAt       time.Time    `json:"created_at"`
	Description     string       `json:"description"`
	EndAt           time.Time    `json:"end_at"`
	ID              int64        `json:"id"`
	Image           ProductImage `json:"image"`
	IsUniq          bool         `json:"is_uniq"`
	Kind            string       `json:"kind"`
	Name            string       `json:"name"`
	OneTimePurchase bool         `json:"one_time_purchase"`
	Price           int64        `json:"price"`
	Quantity        int64        `json:"quantity"`
	Slug            string       `json:"slug"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type BlocDeadlines struct {
	BeginAt     time.Time `json:"begin_at"`
	BlocID      int64     `json:"bloc_id"`
	CoalitionID int64     `json:"coalition_id"`
	CreatedAt   time.Time `json:"created_at"`
	EndAt       time.Time `json:"end_at"`
	ID          int64     `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Exam struct {
	BeginAt        time.Time `json:"begin_at"`
	Campus         Campus    `json:"campus"`
	CreatedAt      time.Time `json:"created_at"`
	Cursus         []Cursus  `json:"cursus"`
	EndAt          time.Time `json:"end_at"`
	ID             int64     `json:"id"`
	IPRange        string    `json:"ip_range"`
	Location       string    `json:"location"`
	MaxPeople      int64     `json:"max_people"`
	Name           string    `json:"name"`
	NbrSubscribers int64     `json:"nbr_subscribers"`
	Projects       []Project `json:"projects"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Location struct {
	BeginAt  time.Time       `json:"begin_at"`
	CampusID int64           `json:"campus_id"`
	EndAt    time.Time       `json:"end_at"`
	Floor    any             `json:"floor"`
	Host     string          `json:"host"`
	ID       int64           `json:"id"`
	Post     any             `json:"post"`
	Primary  bool            `json:"primary"`
	Row      any             `json:"row"`
	User     TransactionUser `json:"user"`
}

type Event struct {
	ID                        int       `json:"id"`
	Name                      string    `json:"name"`
	Description               string    `json:"description"`
	Location                  string    `json:"location"`
	Kind                      string    `json:"kind"`
	MaxPeople                 int       `json:"max_people"`
	NbrSubscribers            int       `json:"nbr_subscribers"`
	BeginAt                   time.Time `json:"begin_at"`
	EndAt                     time.Time `json:"end_at"`
	CampusIds                 []int     `json:"campus_ids"`
	CursusIds                 []int     `json:"cursus_ids"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	ProhibitionOfCancellation int       `json:"prohibition_of_cancellation"`
	Waitlist                  Waitlist  `json:"waitlist,omitempty"`
}

type ScaleTeam struct {
	BeginAt        time.Time `json:"begin_at"`
	Comment        any       `json:"comment"`
	Correcteds     string    `json:"correcteds"`
	Corrector      string    `json:"corrector"`
	CreatedAt      time.Time `json:"created_at"`
	Feedback       any       `json:"feedback"`
	FeedbackRating any       `json:"feedback_rating"`
	Feedbacks      []any     `json:"feedbacks"`
	FilledAt       time.Time `json:"filled_at"`
	FinalMark      any       `json:"final_mark"`
	Flag           Flag      `json:"flag"`
	ID             int64     `json:"id"`
	Scale          Scale     `json:"scale"`
	ScaleID        int64     `json:"scale_id"`
	Truant         any       `json:"truant"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Resources_sub105 struct {
	BeginAt      time.Time       `json:"begin_at"`
	Cursus       Cursus          `json:"cursus"`
	CursusID     int64           `json:"cursus_id"`
	EndAt        time.Time       `json:"end_at"`
	Grade        any             `json:"grade"`
	HasCoalition bool            `json:"has_coalition"`
	ID           int64           `json:"id"`
	Level        float32         `json:"level"`
	Skills       []any           `json:"skills"`
	User         TransactionUser `json:"user"`
}

type CursusUser struct {
	BeginAt      time.Time       `json:"begin_at"`
	Cursus       Cursus          `json:"cursus"`
	CursusID     int64           `json:"cursus_id"`
	EndAt        time.Time       `json:"end_at"`
	Grade        string          `json:"grade"`
	HasCoalition bool            `json:"has_coalition"`
	ID           int64           `json:"id"`
	Level        float32         `json:"level"`
	Skills       []any           `json:"skills"`
	User         TransactionUser `json:"user"`
}

type Balance struct {
	BeginAt time.Time `json:"begin_at"`
	EndAt   time.Time `json:"end_at"`
	ID      int64     `json:"id"`
	PoolID  int64     `json:"pool_id"`
}

type UserCandidature struct {
	BirthCity          string    `json:"birth_city"`
	BirthCountry       string    `json:"birth_country"`
	BirthDate          string    `json:"birth_date"`
	ContactAffiliation any       `json:"contact_affiliation"`
	ContactFirstName   any       `json:"contact_first_name"`
	ContactLastName    any       `json:"contact_last_name"`
	ContactPhone1      any       `json:"contact_phone1"`
	ContactPhone2      any       `json:"contact_phone2"`
	Country            string    `json:"country"`
	CreatedAt          time.Time `json:"created_at"`
	Email              string    `json:"email"`
	Gender             string    `json:"gender"`
	HiddenPhone        any       `json:"hidden_phone"`
	ID                 int64     `json:"id"`
	Language           any       `json:"language"`
	MaxLevelLogic      any       `json:"max_level_logic"`
	MaxLevelMemory     any       `json:"max_level_memory"`
	MeetingDate        any       `json:"meeting_date"`
	OtherInformation   any       `json:"other_information"`
	Phone              any       `json:"phone"`
	PhoneCountryCode   any       `json:"phone_country_code"`
	Pin                any       `json:"pin"`
	PiscineDate        any       `json:"piscine_date"`
	PostalCity         string    `json:"postal_city"`
	PostalComplement   any       `json:"postal_complement"`
	PostalCountry      any       `json:"postal_country"`
	PostalStreet       string    `json:"postal_street"`
	PostalZipCode      any       `json:"postal_zip_code"`
	UpdatedAt          time.Time `json:"updated_at"`
	UserID             int64     `json:"user_id"`
	ZipCode            string    `json:"zip_code"`
}

type ProjectData struct {
	By               []any   `json:"by"`
	Coordinates      []int64 `json:"coordinates"`
	ID               int64   `json:"id"`
	Kind             string  `json:"kind"`
	ProjectSessionID int64   `json:"project_session_id"`
}

type Score struct {
	CalculationID    int64     `json:"calculation_id"`
	CoalitionID      int64     `json:"coalition_id"`
	CoalitionsUserID int64     `json:"coalitions_user_id"`
	CreatedAt        time.Time `json:"created_at"`
	ID               int64     `json:"id"`
	Reason           string    `json:"reason"`
	ScoreableID      int64     `json:"scoreable_id"`
	ScoreableType    string    `json:"scoreable_type"`
	UpdatedAt        time.Time `json:"updated_at"`
	Value            int64     `json:"value"`
}

type Endpoint struct {
	Campus      []Campus  `json:"campus"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int64     `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	URL         string    `json:"url"`
}

type Quest struct {
	Campus       any       `json:"campus"`
	CampusID     any       `json:"campus_id"`
	CreatedAt    time.Time `json:"created_at"`
	Cursus       Cursus    `json:"cursus"`
	CursusID     int64     `json:"cursus_id"`
	Description  string    `json:"description"`
	Grade        any       `json:"grade"`
	GradeID      any       `json:"grade_id"`
	ID           int64     `json:"id"`
	InternalName any       `json:"internal_name"`
	Kind         string    `json:"kind"`
	Name         string    `json:"name"`
	Position     int64     `json:"position"`
	Slug         string    `json:"slug"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Blocs struct {
	CampusID   int64       `json:"campus_id"`
	Coalitions []Coalition `json:"coalitions"`
	CreatedAt  time.Time   `json:"created_at"`
	CursusID   int64       `json:"cursus_id"`
	ID         int64       `json:"id"`
	SquadSize  int64       `json:"squad_size"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type Pool struct {
	CampusID      int64 `json:"campus_id"`
	CurrentPoints int64 `json:"current_points"`
	CursusID      int64 `json:"cursus_id"`
	ID            int64 `json:"id"`
	MaxPoints     int64 `json:"max_points"`
}

type CampusUser struct {
	CampusID  int64 `json:"campus_id"`
	ID        int64 `json:"id"`
	IsPrimary bool  `json:"is_primary"`
	UserID    int64 `json:"user_id"`
}

type CommunityService struct {
	Close      Close     `json:"close"`
	CreatedAt  time.Time `json:"created_at"`
	Duration   int64     `json:"duration"`
	ID         int64     `json:"id"`
	Occupation string    `json:"occupation"`
	ScheduleAt time.Time `json:"schedule_at"`
	State      string    `json:"state"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Team struct {
	Closed           bool      `json:"closed?"`
	ClosedAt         time.Time `json:"closed_at"`
	CreatedAt        time.Time `json:"created_at"`
	FinalMark        any       `json:"final_mark"`
	ID               int64     `json:"id"`
	Locked           bool      `json:"locked?"`
	LockedAt         time.Time `json:"locked_at"`
	Name             string    `json:"name"`
	ProjectID        int64     `json:"project_id"`
	ProjectSessionID int64     `json:"project_session_id"`
	RepoURL          any       `json:"repo_url"`
	RepoUUID         string    `json:"repo_uuid"`
	Status           string    `json:"status"`
	TerminatingAt    time.Time `json:"terminating_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	URL              string    `json:"url"`
	Users            []User    `json:"users"`
	Validated        any       `json:"validated?"`
}

type Close struct {
	Closer            TransactionUser    `json:"closer"`
	CommunityServices []CommunityService `json:"community_services"`
	CreatedAt         time.Time          `json:"created_at"`
	EndAt             time.Time          `json:"end_at"`
	ID                int64              `json:"id"`
	Reason            string             `json:"reason"`
	State             string             `json:"state"`
	UpdatedAt         time.Time          `json:"updated_at"`
	User              TransactionUser    `json:"user"`
}

type CoalitionsUser struct {
	CoalitionID int64     `json:"coalition_id"`
	CreatedAt   time.Time `json:"created_at"`
	ID          int64     `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int64     `json:"user_id"`
}

type Coalition struct {
	Color    string `json:"color"`
	ID       int64  `json:"id"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Score    int64  `json:"score"`
	Slug     string `json:"slug"`
	UserID   int64  `json:"user_id"`
}

type Scale struct {
	Comment            string     `json:"comment"`
	CorrectionNumber   int64      `json:"correction_number"`
	CreatedAt          time.Time  `json:"created_at"`
	DisclaimerMd       string     `json:"disclaimer_md"`
	Duration           int64      `json:"duration"`
	Evaluation         Evaluation `json:"evaluation"`
	EvaluationID       int64      `json:"evaluation_id"`
	GuidelinesMd       string     `json:"guidelines_md"`
	ID                 int64      `json:"id"`
	IntroductionMd     string     `json:"introduction_md"`
	IsPrimary          bool       `json:"is_primary"`
	Languages          []Language `json:"languages"`
	ManualSubscription bool       `json:"manual_subscription"`
	Name               string     `json:"name"`
	Sections           []Section  `json:"sections"`
}

type Report struct {
	Comment        string    `json:"comment"`
	CreatedAt      time.Time `json:"created_at"`
	DelayDays      int64     `json:"delay_days"`
	DisclaimerMd   string    `json:"disclaimer_md"`
	GuidelinesMd   string    `json:"guidelines_md"`
	ID             int64     `json:"id"`
	IntroductionMd string    `json:"introduction_md"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TeamUpload struct {
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	FinalMark int64     `json:"final_mark"`
	ID        int64     `json:"id"`
	Upload    Upload    `json:"upload"`
	UploadID  int64     `json:"upload_id"`
}

type ExpertiseUser struct {
	ContactMe   bool            `json:"contact_me"`
	CreatedAt   time.Time       `json:"created_at"`
	Expertise   Expertise       `json:"expertise"`
	ExpertiseID int64           `json:"expertise_id"`
	ID          int64           `json:"id"`
	Interested  bool            `json:"interested"`
	User        TransactionUser `json:"user"`
	UserID      int64           `json:"user_id"`
	Value       int64           `json:"value"`
}

type Resources_sub106 struct {
	ContactMe   bool      `json:"contact_me"`
	CreatedAt   time.Time `json:"created_at"`
	ExpertiseID int64     `json:"expertise_id"`
	ID          int64     `json:"id"`
	Interested  bool      `json:"interested"`
	UserID      int64     `json:"user_id"`
	Value       int64     `json:"value"`
}

type Note struct {
	Content   string          `json:"content"`
	CreatedAt time.Time       `json:"created_at"`
	FromUser  TransactionUser `json:"from_user"`
	ID        int64           `json:"id"`
	Subject   string          `json:"subject"`
	User      TransactionUser `json:"user"`
}

type Flashe struct {
	Content    string `json:"content"`
	Duration   int64  `json:"duration"`
	ID         int64  `json:"id"`
	Identifier string `json:"identifier"`
	Selector   string `json:"selector"`
	Title      string `json:"title"`
}

type Convention struct {
	Convention Thumb `json:"convention"`
}

type Notion struct {
	CreatedAt  time.Time `json:"created_at"`
	Cursus     []Cursus  `json:"cursus"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Subnotions []any     `json:"subnotions"`
	Tags       []Tag     `json:"tags"`
}

type Param struct {
	CreatedAt    time.Time `json:"created_at"`
	DataType     string    `json:"data_type"`
	DefaultValue string    `json:"default_value"`
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	RuleID       int64     `json:"rule_id"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Translation struct {
	CreatedAt               time.Time        `json:"created_at"`
	Default                 bool             `json:"default"`
	Fields                  Resources_sub100 `json:"fields"`
	ID                      int64            `json:"id"`
	LanguageID              int64            `json:"language_id"`
	TranslatableID          int64            `json:"translatable_id"`
	TranslatableType        string           `json:"translatable_type"`
	TranslationsStructure   Resources_sub102 `json:"translations_structure"`
	TranslationsStructureID int64            `json:"translations_structure_id"`
	UpToDate                bool             `json:"up_to_date"`
	UpdatedAt               time.Time        `json:"updated_at"`
	UserID                  any              `json:"user_id"`
}

type App struct {
	CreatedAt   time.Time       `json:"created_at"`
	Description any             `json:"description"`
	ID          int64           `json:"id"`
	Image       any             `json:"image"`
	Name        string          `json:"name"`
	Owner       TransactionUser `json:"owner"`
	Public      bool            `json:"public"`
	RateLimit   int64           `json:"rate_limit"`
	Roles       []Role          `json:"roles"`
	Scopes      []any           `json:"scopes"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Website     any             `json:"website"`
}

type Entity struct {
	CreatedAt   time.Time `json:"created_at"`
	Description any       `json:"description"`
	ID          int64     `json:"id"`
	Image       any       `json:"image"`
	Name        string    `json:"name"`
	Owner       any       `json:"owner"`
	Public      bool      `json:"public"`
	RateLimit   int64     `json:"rate_limit"`
	Scopes      []any     `json:"scopes"`
	UpdatedAt   time.Time `json:"updated_at"`
	Website     any       `json:"website"`
}

type Upload struct {
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	EvaluationID int64     `json:"evaluation_id"`
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Rule struct {
	CreatedAt            time.Time `json:"created_at"`
	Description          string    `json:"description"`
	ID                   int64     `json:"id"`
	InternalName         string    `json:"internal_name"`
	Kind                 string    `json:"kind"`
	Name                 string    `json:"name"`
	Params               []Param   `json:"params"`
	ProjectSessionsRules []any     `json:"project_sessions_rules"`
	Slug                 string    `json:"slug"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type RoleEntity struct {
	CreatedAt  time.Time `json:"created_at"`
	Entity     Entity    `json:"entity"`
	EntityID   int64     `json:"entity_id"`
	EntityType string    `json:"entity_type"`
	ExpiresAt  time.Time `json:"expires_at"`
	ID         int64     `json:"id"`
	Role       Role      `json:"role"`
}

type Expertise struct {
	CreatedAt          time.Time `json:"created_at"`
	ExpertisesUsersURL string    `json:"expertises_users_url"`
	ID                 int64     `json:"id"`
	Kind               string    `json:"kind"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	URL                string    `json:"url"`
}

type Resources_sub102 struct {
	CreatedAt          time.Time        `json:"created_at"`
	FieldsOrganisation Resources_sub101 `json:"fields_organisation"`
	ID                 int64            `json:"id"`
	IdentifiedBy       []string         `json:"identified_by"`
	SearchableBy       []string         `json:"searchable_by"`
	StructuresKind     string           `json:"structures_kind"`
	TypeName           string           `json:"type_name"`
	UpToDate           bool             `json:"up_to_date"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

type Patronage struct {
	CreatedAt   time.Time       `json:"created_at"`
	Godfather   TransactionUser `json:"godfather"`
	GodfatherID int64           `json:"godfather_id"`
	ID          int64           `json:"id"`
	Ongoing     bool            `json:"ongoing"`
	UpdatedAt   time.Time       `json:"updated_at"`
	User        TransactionUser `json:"user"`
	UserID      int64           `json:"user_id"`
}

type GroupUser struct {
	CreatedAt time.Time `json:"created_at"`
	GroupID   int64     `json:"group_id"`
	ID        int64     `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int64     `json:"user_id"`
}

type Question struct {
	CreatedAt  time.Time `json:"created_at"`
	Guidelines string    `json:"guidelines"`
	ID         int64     `json:"id"`
	Kind       string    `json:"kind"`
	Name       string    `json:"name"`
	Rating     string    `json:"rating"`
}

type LanguageUser struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int64     `json:"id"`
	LanguageID int64     `json:"language_id"`
	Position   int64     `json:"position"`
	UserID     int64     `json:"user_id"`
}

type TeamUser struct {
	CreatedAt  time.Time       `json:"created_at"`
	ID         int64           `json:"id"`
	Leader     bool            `json:"leader"`
	Occurrence int64           `json:"occurrence"`
	Team       Team            `json:"team"`
	TeamID     int64           `json:"team_id"`
	User       TransactionUser `json:"user"`
	UserID     int64           `json:"user_id"`
	Validated  bool            `json:"validated"`
}

type AchievementUser struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	URL       string    `json:"url"`
	UserID    int64     `json:"user_id"`
}

type Cursus struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
}

type Theme struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Skill struct {
	CreatedAt time.Time `json:"created_at"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
}

type ProjectSessionSkill struct {
	CreatedAt        time.Time `json:"created_at"`
	ID               int64     `json:"id"`
	ProjectSessionID int64     `json:"project_session_id"`
	SkillID          int64     `json:"skill_id"`
	UpdatedAt        time.Time `json:"updated_at"`
	Value            int64     `json:"value"`
}

type Waitlist struct {
	ID               int       `json:"id"`
	WaitlistableID   int       `json:"waitlistable_id"`
	WaitlistableType string    `json:"waitlistable_type"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Flag struct {
	CreatedAt time.Time `json:"created_at"`
	Icon      string    `json:"icon"`
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Positive  bool      `json:"positive"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectUser struct {
	CurrentTeamID int64           `json:"current_team_id"`
	CursusIds     []int64         `json:"cursus_ids"`
	FinalMark     any             `json:"final_mark"`
	ID            int64           `json:"id"`
	Occurrence    int64           `json:"occurrence"`
	Project       Project         `json:"project"`
	Status        string          `json:"status"`
	Teams         []Team          `json:"teams"`
	User          TransactionUser `json:"user"`
	Validated     any             `json:"validated?"`
}

type Accreditation struct {
	CursusID  int64  `json:"cursus_id"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UserID    int64  `json:"user_id"`
	Validated bool   `json:"validated"`
}

type DasheUser struct {
	DashID    int64 `json:"dash_id"`
	FinalMark any   `json:"final_mark"`
	ID        int64 `json:"id"`
	RepoURL   any   `json:"repo_url"`
	RepoUUID  any   `json:"repo_uuid"`
	UserID    int64 `json:"user_id"`
}

type Meta struct {
	Date      string `json:"date"`
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	OtherUser string `json:"other_user"`
	User      string `json:"user"`
}

type Resources_sub100 struct {
	Description any    `json:"description"`
	Name        string `json:"name"`
}

type Section struct {
	Description string     `json:"description"`
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Questions   []Question `json:"questions"`
}

type Role struct {
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
}

type Resources_sub101 struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type FlashUser struct {
	EndAt   time.Time       `json:"end_at"`
	FlashID int64           `json:"flash_id"`
	ID      int64           `json:"id"`
	Seen    bool            `json:"seen"`
	User    TransactionUser `json:"user"`
}

type PartnershipUser struct {
	FinalMark     any             `json:"final_mark"`
	ID            int64           `json:"id"`
	PartnershipID int64           `json:"partnership_id"`
	User          TransactionUser `json:"user"`
}

type Language struct {
	ID         int64  `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

type Tag struct {
	ID         int64  `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Subnotions []any  `json:"subnotions"`
	Users      []any  `json:"users"`
}

type Evaluation struct {
	ID   int64  `json:"id"`
	Kind string `json:"kind"`
}

type UserCampus struct {
	ID          int64    `json:"id"`
	Language    Language `json:"language"`
	Name        string   `json:"name"`
	TimeZone    string   `json:"time_zone"`
	UsersCount  int64    `json:"users_count"`
	VogsphereID int64    `json:"vogsphere_id"`
}

type TransactionUser struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

type Partnership struct {
	ID                   int64  `json:"id"`
	Name                 string `json:"name"`
	PartnershipsUsersURL string `json:"partnerships_users_url"`
	Slug                 string `json:"slug"`
	Tier                 int64  `json:"tier"`
	URL                  string `json:"url"`
}

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	ID               int64           `json:"id"`
	Reason           string          `json:"reason"`
	TransactableID   int64           `json:"transactable_id"`
	TransactableType string          `json:"transactable_type"`
	User             TransactionUser `json:"user"`
	UserID           int64           `json:"user_id"`
	Value            int64           `json:"value"`
}

type ImageVersions struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Micro  string `json:"micro"`
	Small  string `json:"small"`
}

type Image struct {
	Link     string        `json:"link"`
	Versions ImageVersions `json:"versions"`
}

type AttachmentPdf struct {
	Pdf Pdf `json:"pdf"`
}

type Pdf struct {
	Thumb Thumb `json:"thumb"`
	URL   any   `json:"url"`
}

type ProductImage struct {
	Thumb Thumb  `json:"thumb"`
	URL   string `json:"url"`
}

type Thumb struct {
	URL string `json:"url"`
}
