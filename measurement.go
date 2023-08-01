package saia

import "time"

type MeasurementStatus string

const (
	MeasurementStatusPending MeasurementStatus = "pending"
	MeasurementStatusSuccess MeasurementStatus = "success"
	MeasurementStatusFailed  MeasurementStatus = "failed"
)

type Measurement struct {
	ID                 int               `json:"id"`
	UUID               string            `json:"uuid"`
	DeliveryStatus     string            `json:"delivery_status"`
	NotificationMethod any               `json:"notification_method"`
	Status             MeasurementStatus `json:"status"`
	CalculationsCount  int               `json:"calculations_count"`
	Created            time.Time         `json:"created"`
	Updated            time.Time         `json:"updated"`
	State              struct {
		Email    string `json:"email"`
		Units    string `json:"units"`
		Gender   Gender `json:"gender"`
		Height   int    `json:"height"`
		Status   string `json:"status"`
		Weight   int    `json:"weight"`
		PersonID int    `json:"personId"`
		Settings struct {
			FinalPage string `json:"final_page"`
		} `json:"settings"`
		MtmClientID  int `json:"mtmClientId"`
		Measurements struct {
			ID      int       `json:"id"`
			URL     string    `json:"url"`
			Gender  Gender    `json:"gender"`
			Height  int       `json:"height"`
			Weight  int       `json:"weight"`
			Created time.Time `json:"created"`
			TaskSet struct {
				IsReady  bool `json:"is_ready"`
				SubTasks []struct {
					Name    string     `json:"name"`
					Status  TaskStatus `json:"status"`
					Message string     `json:"message"`
					TaskID  string     `json:"task_id"`
				} `json:"sub_tasks"`
				IsSuccessful bool `json:"is_successful"`
			} `json:"task_set"`
			IsViewed   bool   `json:"is_viewed"`
			IPAddress  string `json:"ip_address"`
			PhotoFlow  string `json:"photo_flow"`
			IsArchived bool   `json:"is_archived"`
			SideParams struct {
				ClothesType struct {
					Types []any `json:"types"`
				} `json:"clothes_type"`
				NeckToChest    float64 `json:"neck_to_chest"`
				ChestToWaist   float64 `json:"chest_to_waist"`
				WaistToAnkle   float64 `json:"waist_to_ankle"`
				SoftValidation struct {
					Messages []any `json:"messages"`
				} `json:"soft_validation"`
				ShouldersToKnees        float64 `json:"shoulders_to_knees"`
				SideNeckPointToUpperHip float64 `json:"side_neck_point_to_upper_hip"`
				SideUpperHipLevelToKnee float64 `json:"side_upper_hip_level_to_knee"`
			} `json:"side_params"`
			CountryCode string `json:"country_code"`
			CountryName string `json:"country_name"`
			FrontParams struct {
				Neck        float64 `json:"neck"`
				Rise        float64 `json:"rise"`
				Waist       float64 `json:"waist"`
				Inseam      float64 `json:"inseam"`
				Outseam     float64 `json:"outseam"`
				ChestTop    float64 `json:"chest_top"`
				HighHips    float64 `json:"high_hips"`
				Shoulders   float64 `json:"shoulders"`
				HipHeight   float64 `json:"hip_height"`
				BodyHeight  float64 `json:"body_height"`
				BustHeight  float64 `json:"bust_height"`
				KneeHeight  float64 `json:"knee_height"`
				NeckLength  float64 `json:"neck_length"`
				ClothesType struct {
					Types struct {
						Top struct {
							Code   string `json:"code"`
							Detail string `json:"detail"`
						} `json:"top"`
						Bottom struct {
							Code   string `json:"code"`
							Detail string `json:"detail"`
						} `json:"bottom"`
					} `json:"types"`
				} `json:"clothes_type"`
				TorsoHeight    float64 `json:"torso_height"`
				WaistHeight    float64 `json:"waist_height"`
				CrotchLength   float64 `json:"crotch_length"`
				JacketLength   float64 `json:"jacket_length"`
				SleeveLength   float64 `json:"sleeve_length"`
				ShoulderSlope  float64 `json:"shoulder_slope"`
				WaistToKnees   float64 `json:"waist_to_knees"`
				ShoulderLength float64 `json:"shoulder_length"`
				SoftValidation struct {
					Messages []any `json:"messages"`
				} `json:"soft_validation"`
				UnderarmLength                         float64 `json:"underarm_length"`
				BackNeckHeight                         float64 `json:"back_neck_height"`
				LowerArmLength                         float64 `json:"lower_arm_length"`
				UpperArmLength                         float64 `json:"upper_arm_length"`
				UpperHipHeight                         float64 `json:"upper_hip_height"`
				AcrossBackWidth                        float64 `json:"across_back_width"`
				InsideLegHeight                        float64 `json:"inside_leg_height"`
				ShoulderToWaist                        float64 `json:"shoulder_to_waist"`
				WaistToLowHips                         float64 `json:"waist_to_low_hips"`
				BackCrotchLength                       float64 `json:"back_crotch_length"`
				OuterAnkleHeight                       float64 `json:"outer_ankle_height"`
				BackShoulderWidth                      float64 `json:"back_shoulder_width"`
				FrontCrotchLength                      float64 `json:"front_crotch_length"`
				TotalCrotchLength                      float64 `json:"total_crotch_length"`
				UpperKneeToAnkle                       float64 `json:"upper_knee_to_ankle"`
				BackNeckToHipLength                    float64 `json:"back_neck_to_hip_length"`
				UpperHipToHipLength                    float64 `json:"upper_hip_to_hip_length"`
				NapeToWaistCentreBack                  float64 `json:"nape_to_waist_centre_back"`
				SideNeckPointToArmpit                  float64 `json:"side_neck_point_to_armpit"`
				AcrossBackShoulderWidth                float64 `json:"across_back_shoulder_width"`
				AbdomenToUpperKneeLength               float64 `json:"abdomen_to_upper_knee_length"`
				InsideCrotchLengthToCalf               float64 `json:"inside_crotch_length_to_calf"`
				InsideCrotchLengthToKnee               float64 `json:"inside_crotch_length_to_knee"`
				OutseamFromUpperHipLevel               float64 `json:"outseam_from_upper_hip_level"`
				BackNeckPointToWristLength             float64 `json:"back_neck_point_to_wrist_length"`
				InsideCrotchLengthToMidThigh           float64 `json:"inside_crotch_length_to_mid_thigh"`
				BackNeckPointToWristLength15Inch       float64 `json:"back_neck_point_to_wrist_length_1_5_inch"`
				InsideLegLengthToThe1InchAboveTheFloor float64 `json:"inside_leg_length_to_the_1_inch_above_the_floor"`
			} `json:"front_params"`
			VolumeParams struct {
				Calf                  float64 `json:"calf"`
				Knee                  int     `json:"knee"`
				Neck                  float64 `json:"neck"`
				Ankle                 float64 `json:"ankle"`
				Bicep                 float64 `json:"bicep"`
				Chest                 float64 `json:"chest"`
				Thigh                 float64 `json:"thigh"`
				Waist                 float64 `json:"waist"`
				Wrist                 float64 `json:"wrist"`
				Abdomen               float64 `json:"abdomen"`
				Forearm               float64 `json:"forearm"`
				LowHips               float64 `json:"low_hips"`
				HighHips              float64 `json:"high_hips"`
				NeckGirth             float64 `json:"neck_girth"`
				WaistGray             float64 `json:"waist_gray"`
				ElbowGirth            float64 `json:"elbow_girth"`
				WaistGreen            float64 `json:"waist_green"`
				ArmscyeGirth          float64 `json:"armscye_girth"`
				OverarmGirth          float64 `json:"overarm_girth"`
				MidThighGirth         float64 `json:"mid_thigh_girth"`
				UnderBustGirth        float64 `json:"under_bust_girth"`
				UpperBicepGirth       float64 `json:"upper_bicep_girth"`
				UpperChestGirth       float64 `json:"upper_chest_girth"`
				NeckGirthRelaxed      float64 `json:"neck_girth_relaxed"`
				Thigh1InchBelowCrotch float64 `json:"thigh_1_inch_below_crotch"`
			} `json:"volume_params"`
			PhonePosition *PhonePosition `json:"phone_position"`
		} `json:"measurements"`
		ProcessStatus  string `json:"processStatus"`
		SoftValidation struct {
			LooseTop          bool `json:"looseTop"`
			WideLegs          bool `json:"wideLegs"`
			SmallLegs         bool `json:"smallLegs"`
			LooseBottom       bool `json:"looseBottom"`
			BodyPercentage    bool `json:"bodyPercentage"`
			LooseTopAndBottom bool `json:"looseTopAndBottom"`
		} `json:"softValidation"`
	} `json:"state"`
	WidgetFlowStatus string `json:"widget_flow_status"`
	IsActive         bool   `json:"is_active"`
	ShortLink        string `json:"short_link"`
	Phone            any    `json:"phone"`
	Email            string `json:"email"`
	EmailMessageData struct {
	} `json:"email_message_data"`
	SmsMessageData struct {
	} `json:"sms_message_data"`
	Source     string `json:"source"`
	Unit       string `json:"unit"`
	Notes      string `json:"notes"`
	IsViewed   bool   `json:"is_viewed"`
	IsArchived bool   `json:"is_archived"`
	Person     struct {
		ID            int       `json:"id"`
		URL           string    `json:"url"`
		Gender        string    `json:"gender"`
		Height        int       `json:"height"`
		Created       time.Time `json:"created"`
		Weight        int       `json:"weight"`
		PhonePosition struct {
			SidePhoto struct {
				BetaX  float64 `json:"betaX"`
				AlphaZ float64 `json:"alphaZ"`
				GammaY float64 `json:"gammaY"`
			} `json:"sidePhoto"`
			FrontPhoto struct {
				BetaX  float64 `json:"betaX"`
				AlphaZ float64 `json:"alphaZ"`
				GammaY float64 `json:"gammaY"`
			} `json:"frontPhoto"`
		} `json:"phone_position"`
		PhotoFlow   string `json:"photo_flow"`
		IPAddress   string `json:"ip_address"`
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
		IsViewed    bool   `json:"is_viewed"`
		IsArchived  bool   `json:"is_archived"`
		TaskSets    []struct {
			ID           int           `json:"id"`
			FrontParams  *FrontParams  `json:"front_params"`
			SideParams   *SideParams   `json:"side_params"`
			VolumeParams *VolumeParams `json:"volume_params"`
			IsSuccessful bool          `json:"is_successful"`
			IsReady      bool          `json:"is_ready"`
			IsPrimary    bool          `json:"is_primary"`
			SubTasks     []struct {
				Name    string `json:"name"`
				Status  string `json:"status"`
				TaskID  string `json:"task_id"`
				Message string `json:"message"`
			} `json:"sub_tasks"`
			Created time.Time `json:"created"`
		} `json:"task_sets"`
	} `json:"person"`
	MtmClient struct {
		ID        int       `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Phone     string    `json:"phone"`
		Notes     string    `json:"notes"`
		Created   time.Time `json:"created"`
	} `json:"mtm_client"`
	IsDemoTry bool `json:"is_demo_try"`
}

type PhonePosition struct {
	SidePhoto struct {
		BetaX  float64 `json:"betaX"`
		AlphaZ float64 `json:"alphaZ"`
		GammaY float64 `json:"gammaY"`
	} `json:"sidePhoto"`
	FrontPhoto struct {
		BetaX  float64 `json:"betaX"`
		AlphaZ float64 `json:"alphaZ"`
		GammaY float64 `json:"gammaY"`
	} `json:"frontPhoto"`
}

type FrontParams struct {
	SoftValidation struct {
		Messages []any `json:"messages"`
	} `json:"soft_validation"`
	ClothesType struct {
		Types struct {
			Top struct {
				Code   string `json:"code"`
				Detail string `json:"detail"`
			} `json:"top"`
			Bottom struct {
				Code   string `json:"code"`
				Detail string `json:"detail"`
			} `json:"bottom"`
		} `json:"types"`
	} `json:"clothes_type"`
	BodyAreaPercentage                     float64 `json:"body_area_percentage"`
	BodyHeight                             float64 `json:"body_height"`
	Outseam                                float64 `json:"outseam"`
	OutseamFromUpperHipLevel               float64 `json:"outseam_from_upper_hip_level"`
	Inseam                                 float64 `json:"inseam"`
	InsideLegLengthToThe1InchAboveTheFloor float64 `json:"inside_leg_length_to_the_1_inch_above_the_floor"`
	InsideCrotchLengthToMidThigh           float64 `json:"inside_crotch_length_to_mid_thigh"`
	InsideCrotchLengthToKnee               float64 `json:"inside_crotch_length_to_knee"`
	InsideCrotchLengthToCalf               float64 `json:"inside_crotch_length_to_calf"`
	CrotchLength                           float64 `json:"crotch_length"`
	SleeveLength                           float64 `json:"sleeve_length"`
	UnderarmLength                         float64 `json:"underarm_length"`
	BackNeckPointToWristLength             float64 `json:"back_neck_point_to_wrist_length"`
	BackNeckPointToWristLength15Inch       float64 `json:"back_neck_point_to_wrist_length_1_5_inch"`
	HighHips                               float64 `json:"high_hips"`
	Shoulders                              float64 `json:"shoulders"`
	ChestTop                               float64 `json:"chest_top"`
	JacketLength                           float64 `json:"jacket_length"`
	ShoulderLength                         float64 `json:"shoulder_length"`
	ShoulderSlope                          float64 `json:"shoulder_slope"`
	Neck                                   float64 `json:"neck"`
	WaistToLowHips                         float64 `json:"waist_to_low_hips"`
	WaistToUpperKneeLength                 float64 `json:"waist_to_upper_knee_length"`
	WaistToKnees                           float64 `json:"waist_to_knees"`
	AbdomenToUpperKneeLength               float64 `json:"abdomen_to_upper_knee_length"`
	UpperKneeToAnkle                       float64 `json:"upper_knee_to_ankle"`
	NapeToWaistCentreBack                  float64 `json:"nape_to_waist_centre_back"`
	ShoulderToWaist                        float64 `json:"shoulder_to_waist"`
	SideNeckPointToArmpit                  float64 `json:"side_neck_point_to_armpit"`
	BackNeckHeight                         float64 `json:"back_neck_height"`
	BustHeight                             float64 `json:"bust_height"`
	HipHeight                              float64 `json:"hip_height"`
	UpperHipHeight                         float64 `json:"upper_hip_height"`
	KneeHeight                             float64 `json:"knee_height"`
	OuterAnkleHeight                       float64 `json:"outer_ankle_height"`
	WaistHeight                            float64 `json:"waist_height"`
	InsideLegHeight                        float64 `json:"inside_leg_height"`
	AcrossBackShoulderWidth                float64 `json:"across_back_shoulder_width"`
	AcrossBackWidth                        float64 `json:"across_back_width"`
	TotalCrotchLength                      float64 `json:"total_crotch_length"`
	Waist                                  float64 `json:"waist"`
	NeckLength                             float64 `json:"neck_length"`
	UpperArmLength                         float64 `json:"upper_arm_length"`
	LowerArmLength                         float64 `json:"lower_arm_length"`
	UpperHipToHipLength                    float64 `json:"upper_hip_to_hip_length"`
	BackShoulderWidth                      float64 `json:"back_shoulder_width"`
	Rise                                   float64 `json:"rise"`
	BackNeckToHipLength                    float64 `json:"back_neck_to_hip_length"`
	TorsoHeight                            float64 `json:"torso_height"`
	FrontTorsoHeight                       float64 `json:"front_torso_height"`
	FrontCrotchLength                      float64 `json:"front_crotch_length"`
	BackCrotchLength                       float64 `json:"back_crotch_length"`
	LegsDistance                           float64 `json:"legs_distance"`
}

type SideParams struct {
	SoftValidation struct {
		Messages []any `json:"messages"`
	} `json:"soft_validation"`
	ClothesType struct {
		Types []any `json:"types"`
	} `json:"clothes_type"`
	BodyAreaPercentage      float64 `json:"body_area_percentage"`
	SideUpperHipLevelToKnee float64 `json:"side_upper_hip_level_to_knee"`
	SideNeckPointToUpperHip float64 `json:"side_neck_point_to_upper_hip"`
	NeckToChest             float64 `json:"neck_to_chest"`
	ChestToWaist            float64 `json:"chest_to_waist"`
	WaistToAnkle            float64 `json:"waist_to_ankle"`
	ShouldersToKnees        float64 `json:"shoulders_to_knees"`
	WaistDepth              any     `json:"waist_depth"`
}

type VolumeParams struct {
	Chest                 float64              `json:"chest"`
	UnderBustGirth        float64              `json:"under_bust_girth"`
	UpperChestGirth       float64              `json:"upper_chest_girth"`
	OverarmGirth          float64              `json:"overarm_girth"`
	Waist                 float64              `json:"waist"`
	AlternativeWaistGirth float64              `json:"alternative_waist_girth"`
	HighHips              float64              `json:"high_hips"`
	LowHips               float64              `json:"low_hips"`
	WaistGreen            float64              `json:"waist_green"`
	WaistGray             float64              `json:"waist_gray"`
	PantWaist             float64              `json:"pant_waist"`
	Bicep                 float64              `json:"bicep"`
	UpperBicepGirth       float64              `json:"upper_bicep_girth"`
	UpperKneeGirth        float64              `json:"upper_knee_girth"`
	Knee                  float64              `json:"knee"`
	Ankle                 float64              `json:"ankle"`
	Wrist                 float64              `json:"wrist"`
	Calf                  float64              `json:"calf"`
	Thigh                 float64              `json:"thigh"`
	Thigh1InchBelowCrotch float64              `json:"thigh_1_inch_below_crotch"`
	MidThighGirth         float64              `json:"mid_thigh_girth"`
	Neck                  float64              `json:"neck"`
	Abdomen               float64              `json:"abdomen"`
	ArmscyeGirth          float64              `json:"armscye_girth"`
	NeckGirth             float64              `json:"neck_girth"`
	NeckGirthRelaxed      float64              `json:"neck_girth_relaxed"`
	Forearm               float64              `json:"forearm"`
	ElbowGirth            float64              `json:"elbow_girth"`
	BodyType              string               `json:"body_type"`
	BodyModel             string               `json:"body_model"`
	Textures              any                  `json:"textures"`
	CoatSleeveInseam      any                  `json:"coat_sleeve_inseam"`
	FrontDebugInfo        map[string][]float64 `json:"front_debug_info"`
	VolumeDebugInfo       map[string][]float64 `json:"volume_debug_info"`
}
