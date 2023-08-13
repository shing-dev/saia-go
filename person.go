package saia

import "time"

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type Person struct {
	ID            int            `json:"id"`
	URL           string         `json:"url"`
	Gender        Gender         `json:"gender"`
	Height        int            `json:"height"`
	Created       time.Time      `json:"created"`
	Weight        float64        `json:"weight"`
	PhonePosition *PhonePosition `json:"phone_position"`
	PhotoFlow     string         `json:"photo_flow"`
	IPAddress     string         `json:"ip_address"`
	CountryName   string         `json:"country_name"`
	CountryCode   string         `json:"country_code"`
	TaskSet       struct {
		IsSuccessful bool `json:"is_successful"`
		IsReady      bool `json:"is_ready"`
		SubTasks     []struct {
			Name    string     `json:"name"`
			Status  TaskStatus `json:"status"`
			TaskID  string     `json:"task_id"`
			Message string     `json:"message"`
		} `json:"sub_tasks"`
	} `json:"task_set"`
	FrontParams  *FrontParams  `json:"front_params"`
	SideParams   *SideParams   `json:"side_params"`
	VolumeParams *VolumeParams `json:"volume_params"`
	IsViewed     bool          `json:"is_viewed"`
	IsArchived   bool          `json:"is_archived"`
}

type TaskSet struct {
	IsSuccessful bool       `json:"is_successful"`
	IsReady      bool       `json:"is_ready"`
	SubTasks     []*SubTask `json:"sub_tasks"`
}

type SubTask struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	TaskID  string `json:"task_id"`
	Message string `json:"message"`
}
