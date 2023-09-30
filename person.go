package saia

import (
	"strings"
	"time"
)

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

func (t *TaskSet) IsFailed() bool {
	return !t.IsSuccessful && t.IsReady
}

type SubTaskName string

const (
	SubTaskNameFrontSkeletonProcessing    SubTaskName = "front_skeleton_processing"
	SubTaskNameSideSkeletonProcessing     SubTaskName = "side_skeleton_processing"
	SubTaskNameFrontProcessing            SubTaskName = "front_processing"
	SubTaskNameSideProcessing             SubTaskName = "side_processing"
	SubTaskNameMeasurementModelProcessing SubTaskName = "measurement_model_processing"
)

type SubTaskErrorCode int

// We use simple iota instead of using the application code provided by 3dlook (https://saia.3dlook.me/docs/#task-errors)
// because they are incomplete (e.g. "Front photo in the side" is not defined)
const (
	SubTaskErrorCodeUnknown SubTaskErrorCode = iota
	SubTaskErrorCodeWrongPose
	SubTaskErrorCodeHumanBodyNotDetected
	SubTaskErrorCodeObjectIsNotHuman
	SubTaskErrorCodeSidePhotoInTheFront
	SubTaskErrorCodeFrontPhotoInTheSide
	SubTaskErrorCodeDetermineLookingSideFailed
	SubTaskErrorCodeBodyIsNotFull
)

type SubTask struct {
	Name    SubTaskName `json:"name"`
	Status  TaskStatus  `json:"status"`
	TaskID  string      `json:"task_id"`
	Message string      `json:"message"`
}

// ErrorCode is a workaround method to extract the error code
// It returns unknown when the sub task is not failed
// TODO: Ask 3dlook to return application error code which is more valid than parsing error message
func (s *SubTask) ErrorCode() SubTaskErrorCode {
	if !s.IsFailed() {
		return SubTaskErrorCodeUnknown
	}

	switch {
	case strings.HasPrefix(s.Message, "The pose is wrong"):
		return SubTaskErrorCodeWrongPose
	case s.Message == "Can't detect the human body":
		return SubTaskErrorCodeHumanBodyNotDetected
	case s.Message == "The detected object is not human":
		return SubTaskErrorCodeObjectIsNotHuman
	case s.Message == "Side photo in the front":
		return SubTaskErrorCodeSidePhotoInTheFront
	case s.Message == "Front photo in the side":
		return SubTaskErrorCodeFrontPhotoInTheSide
	case s.Message == "The body is not full":
		return SubTaskErrorCodeBodyIsNotFull
	case s.Message == "Failed to determine looking side":
		return SubTaskErrorCodeDetermineLookingSideFailed
	default:
		return SubTaskErrorCodeUnknown
	}
}

func (s *SubTask) IsFailed() bool {
	return s.Status == TaskStatusFailure
}
