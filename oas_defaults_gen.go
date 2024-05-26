// Code generated by ogen, DO NOT EDIT.

package telnyx

// setDefaults set default value of fields.
func (s *AnswerRequest) setDefaults() {
	{
		val := AnswerRequestStreamTrack("inbound_track")
		s.StreamTrack.SetTo(val)
	}
	{
		val := bool(false)
		s.SendSilenceWhenIdle.SetTo(val)
	}
	{
		val := AnswerRequestWebhookURLMethod("POST")
		s.WebhookURLMethod.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *AudioTranscriptionRequestMultipart) setDefaults() {
	{
		val := AudioTranscriptionRequestMultipartResponseFormat("json")
		s.ResponseFormat.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *BridgeRequest) setDefaults() {
	{
		val := bool(false)
		s.PlayRingtone.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CallControlApplication) setDefaults() {
	{
		val := bool(true)
		s.Active.SetTo(val)
	}
	{
		val := AnchorsiteOverride("Latency")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := CallControlApplicationDtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := bool(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := int(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := CallControlApplicationRecordType("call_control_application")
		s.RecordType.SetTo(val)
	}
	{
		val := CallControlApplicationWebhookAPIVersion("1")
		s.WebhookAPIVersion.SetTo(val)
	}
	{
		val := string("")
		s.WebhookEventFailoverURL.SetTo(val)
	}
	{
		s.WebhookTimeoutSecs.Null = true
	}
}

// setDefaults set default value of fields.
func (s *CallControlApplicationInbound) setDefaults() {
	{
		val := bool(false)
		s.ShakenStirEnabled.SetTo(val)
	}
	{
		val := CallControlApplicationInboundSipSubdomainReceiveSettings("from_anyone")
		s.SipSubdomainReceiveSettings.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CallForwarding) setDefaults() {
	{
		val := bool(true)
		s.CallForwardingEnabled.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CallRecording) setDefaults() {
	{
		val := bool(false)
		s.InboundCallRecordingEnabled.SetTo(val)
	}
	{
		val := CallRecordingInboundCallRecordingFormat("wav")
		s.InboundCallRecordingFormat.SetTo(val)
	}
	{
		val := CallRecordingInboundCallRecordingChannels("single")
		s.InboundCallRecordingChannels.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CallRequest) setDefaults() {
	{
		val := int32(30)
		s.TimeoutSecs.SetTo(val)
	}
	{
		val := int32(14400)
		s.TimeLimitSecs.SetTo(val)
	}
	{
		val := CallRequestAnsweringMachineDetection("disabled")
		s.AnsweringMachineDetection.SetTo(val)
	}
	{
		val := CallRequestMediaEncryption("disabled")
		s.MediaEncryption.SetTo(val)
	}
	{
		val := CallRequestSipTransportProtocol("UDP")
		s.SipTransportProtocol.SetTo(val)
	}
	{
		val := CallRequestStreamTrack("inbound_track")
		s.StreamTrack.SetTo(val)
	}
	{
		val := bool(false)
		s.SendSilenceWhenIdle.SetTo(val)
	}
	{
		val := CallRequestWebhookURLMethod("POST")
		s.WebhookURLMethod.SetTo(val)
	}
	{
		val := CallRequestRecordChannels("dual")
		s.RecordChannels.SetTo(val)
	}
	{
		val := CallRequestRecordFormat("mp3")
		s.RecordFormat.SetTo(val)
	}
	{
		val := int32(0)
		s.RecordMaxLength.SetTo(val)
	}
	{
		val := int32(0)
		s.RecordTimeoutSecs.SetTo(val)
	}
	{
		val := bool(false)
		s.EnableDialogflow.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CallRequestAnsweringMachineDetectionConfig) setDefaults() {
	{
		val := int32(3500)
		s.TotalAnalysisTimeMillis.SetTo(val)
	}
	{
		val := int32(800)
		s.AfterGreetingSilenceMillis.SetTo(val)
	}
	{
		val := int32(50)
		s.BetweenWordsSilenceMillis.SetTo(val)
	}
	{
		val := int32(3500)
		s.GreetingDurationMillis.SetTo(val)
	}
	{
		val := int32(3500)
		s.InitialSilenceMillis.SetTo(val)
	}
	{
		val := int32(5)
		s.MaximumNumberOfWords.SetTo(val)
	}
	{
		val := int32(3500)
		s.MaximumWordLengthMillis.SetTo(val)
	}
	{
		val := int32(256)
		s.SilenceThreshold.SetTo(val)
	}
	{
		val := int32(5000)
		s.GreetingTotalAnalysisTimeMillis.SetTo(val)
	}
	{
		val := int32(1500)
		s.GreetingSilenceDurationMillis.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CnamListing) setDefaults() {
	{
		val := bool(false)
		s.CnamListingEnabled.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateCallControlApplicationRequest) setDefaults() {
	{
		val := bool(true)
		s.Active.SetTo(val)
	}
	{
		val := CreateCallControlApplicationRequestAnchorsiteOverride("\"Latency\"")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := CreateCallControlApplicationRequestDtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := bool(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := int(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := CreateCallControlApplicationRequestWebhookAPIVersion("1")
		s.WebhookAPIVersion.SetTo(val)
	}
	{
		val := string("")
		s.WebhookEventFailoverURL.SetTo(val)
	}
	{
		s.WebhookTimeoutSecs.Null = true
	}
}

// setDefaults set default value of fields.
func (s *CreateGroupMMSMessageRequest) setDefaults() {
	{
		val := bool(true)
		s.UseProfileWebhooks.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateLongCodeMessageRequest) setDefaults() {
	{
		val := bool(true)
		s.UseProfileWebhooks.SetTo(val)
	}
	{
		val := bool(false)
		s.AutoDetect.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateMessageRequest) setDefaults() {
	{
		val := bool(true)
		s.UseProfileWebhooks.SetTo(val)
	}
	{
		val := bool(false)
		s.AutoDetect.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateNumberPoolMessageRequest) setDefaults() {
	{
		val := bool(true)
		s.UseProfileWebhooks.SetTo(val)
	}
	{
		val := bool(false)
		s.AutoDetect.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateShortCodeMessageRequest) setDefaults() {
	{
		val := bool(true)
		s.UseProfileWebhooks.SetTo(val)
	}
	{
		val := bool(false)
		s.AutoDetect.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateTexmlApplicationRequest) setDefaults() {
	{
		val := ConnectionActive(true)
		s.Active.SetTo(val)
	}
	{
		val := AnchorsiteOverride("Latency")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := DtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := FirstCommandTimeout(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := FirstCommandTimeoutSecs(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := CreateTexmlApplicationRequestVoiceMethod("post")
		s.VoiceMethod.SetTo(val)
	}
	{
		val := CreateTexmlApplicationRequestStatusCallbackMethod("post")
		s.StatusCallbackMethod.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CreateTexmlApplicationRequestInbound) setDefaults() {
	{
		val := bool(false)
		s.ShakenStirEnabled.SetTo(val)
	}
	{
		val := CreateTexmlApplicationRequestInboundSipSubdomainReceiveSettings("from_anyone")
		s.SipSubdomainReceiveSettings.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *CsvDownload) setDefaults() {
	{
		val := CsvDownloadStatus("pending")
		s.Status.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *DialConferenceParticipantRequest) setDefaults() {
	{
		val := bool(false)
		s.EarlyMedia.SetTo(val)
	}
	{
		val := int(3500)
		s.MachineDetectionSpeechThreshold.SetTo(val)
	}
	{
		val := int(800)
		s.MachineDetectionSpeechEndThreshold.SetTo(val)
	}
	{
		val := int(3500)
		s.MachineDetectionSilenceTimeout.SetTo(val)
	}
	{
		val := bool(true)
		s.CancelPlaybackOnMachineDetection.SetTo(val)
	}
	{
		val := bool(true)
		s.CancelPlaybackOnDetectMessageEnd.SetTo(val)
	}
	{
		val := int(0)
		s.ConferenceRecordingTimeout.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *DialogflowConfig) setDefaults() {
	{
		val := bool(false)
		s.AnalyzeSentiment.SetTo(val)
	}
	{
		val := bool(false)
		s.PartialAutomatedAgentReply.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *EmergencySettings) setDefaults() {
	{
		val := bool(false)
		s.EmergencyEnabled.SetTo(val)
	}
	{
		val := EmergencySettingsEmergencyStatus("disabled")
		s.EmergencyStatus.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *EnqueueRequest) setDefaults() {
	{
		val := int(100)
		s.MaxSize.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GatherRequest) setDefaults() {
	{
		val := int32(1)
		s.MinimumDigits.SetTo(val)
	}
	{
		val := int32(128)
		s.MaximumDigits.SetTo(val)
	}
	{
		val := int32(60000)
		s.TimeoutMillis.SetTo(val)
	}
	{
		val := int32(5000)
		s.InterDigitTimeoutMillis.SetTo(val)
	}
	{
		val := int32(5000)
		s.InitialTimeoutMillis.SetTo(val)
	}
	{
		val := string("#")
		s.TerminatingDigit.SetTo(val)
	}
	{
		val := string("0123456789#*")
		s.ValidDigits.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GatherUsingAudioRequest) setDefaults() {
	{
		val := int32(1)
		s.MinimumDigits.SetTo(val)
	}
	{
		val := int32(128)
		s.MaximumDigits.SetTo(val)
	}
	{
		val := int32(3)
		s.MaximumTries.SetTo(val)
	}
	{
		val := int32(60000)
		s.TimeoutMillis.SetTo(val)
	}
	{
		val := string("#")
		s.TerminatingDigit.SetTo(val)
	}
	{
		val := string("0123456789#*")
		s.ValidDigits.SetTo(val)
	}
	{
		val := int32(5000)
		s.InterDigitTimeoutMillis.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *GatherUsingSpeakRequest) setDefaults() {
	{
		val := GatherUsingSpeakRequestPayloadType("text")
		s.PayloadType.SetTo(val)
	}
	{
		val := GatherUsingSpeakRequestServiceLevel("premium")
		s.ServiceLevel.SetTo(val)
	}
	{
		val := int32(1)
		s.MinimumDigits.SetTo(val)
	}
	{
		val := int32(128)
		s.MaximumDigits.SetTo(val)
	}
	{
		val := int32(3)
		s.MaximumTries.SetTo(val)
	}
	{
		val := int32(60000)
		s.TimeoutMillis.SetTo(val)
	}
	{
		val := string("#")
		s.TerminatingDigit.SetTo(val)
	}
	{
		val := string("0123456789#*")
		s.ValidDigits.SetTo(val)
	}
	{
		val := int32(5000)
		s.InterDigitTimeoutMillis.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *InitiateCallRequest) setDefaults() {
	{
		val := InitiateCallRequestUrlMethod("POST")
		s.UrlMethod.SetTo(val)
	}
	{
		val := InitiateCallRequestStatusCallbackMethod("POST")
		s.StatusCallbackMethod.SetTo(val)
	}
	{
		val := InitiateCallRequestStatusCallbackEvent("completed")
		s.StatusCallbackEvent.SetTo(val)
	}
	{
		val := InitiateCallRequestMachineDetection("Disable")
		s.MachineDetection.SetTo(val)
	}
	{
		val := InitiateCallRequestDetectionMode("Regular")
		s.DetectionMode.SetTo(val)
	}
	{
		val := bool(false)
		s.AsyncAmd.SetTo(val)
	}
	{
		val := InitiateCallRequestAsyncAmdStatusCallbackMethod("POST")
		s.AsyncAmdStatusCallbackMethod.SetTo(val)
	}
	{
		val := int(30000)
		s.MachineDetectionTimeout.SetTo(val)
	}
	{
		val := int(3500)
		s.MachineDetectionSpeechThreshold.SetTo(val)
	}
	{
		val := int(800)
		s.MachineDetectionSpeechEndThreshold.SetTo(val)
	}
	{
		val := int(3500)
		s.MachineDetectionSilenceTimeout.SetTo(val)
	}
	{
		val := bool(true)
		s.CancelPlaybackOnMachineDetection.SetTo(val)
	}
	{
		val := bool(true)
		s.CancelPlaybackOnDetectMessageEnd.SetTo(val)
	}
	{
		val := int(0)
		s.RecordingTimeout.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *MediaFeatures) setDefaults() {
	{
		val := bool(true)
		s.RtpAutoAdjustEnabled.SetTo(val)
	}
	{
		val := bool(false)
		s.AcceptAnyRtpPacketsEnabled.SetTo(val)
	}
	{
		val := bool(false)
		s.T38FaxGatewayEnabled.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *NoiseSuppressionStart) setDefaults() {
	{
		val := NoiseSuppressionDirection("inbound")
		s.Direction.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PhoneNumberDeletedDetailed) setDefaults() {
	{
		val := bool(true)
		s.CallForwardingEnabled.SetTo(val)
	}
	{
		val := PhoneNumberDeletedDetailedNumberLevelRouting("disabled")
		s.NumberLevelRouting.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PhoneNumberDetailed) setDefaults() {
	{
		val := bool(true)
		s.CallForwardingEnabled.SetTo(val)
	}
	{
		val := PhoneNumberDetailedNumberLevelRouting("disabled")
		s.NumberLevelRouting.SetTo(val)
	}
	{
		val := PhoneNumberDetailedInboundCallScreening("disabled")
		s.InboundCallScreening.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PhoneNumberWithVoiceSettings) setDefaults() {
	{
		val := bool(false)
		s.TechPrefixEnabled.SetTo(val)
	}
	{
		val := string("")
		s.TranslatedNumber.SetTo(val)
	}
	{
		val := PhoneNumberWithVoiceSettingsUsagePaymentMethod("pay-per-minute")
		s.UsagePaymentMethod.SetTo(val)
	}
	{
		val := PhoneNumberWithVoiceSettingsInboundCallScreening("disabled")
		s.InboundCallScreening.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PhoneNumbersJob) setDefaults() {
	{
		val := PhoneNumbersJobStatus("pending")
		s.Status.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PlayAudioUrlRequest) setDefaults() {
	{
		val := bool(false)
		s.Overlay.SetTo(val)
	}
	{
		val := string("self")
		s.TargetLegs.SetTo(val)
	}
	{
		val := bool(true)
		s.CacheAudio.SetTo(val)
	}
	{
		val := PlayAudioUrlRequestAudioType("mp3")
		s.AudioType.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PlaybackStopRequest) setDefaults() {
	{
		val := bool(false)
		s.Overlay.SetTo(val)
	}
	{
		val := string("all")
		s.Stop.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *PortoutComment) setDefaults() {
	{
		s.PortoutID.Null = true
	}
}

// setDefaults set default value of fields.
func (s *SendDTMFRequest) setDefaults() {
	{
		val := int32(250)
		s.DurationMillis.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *SlimPhoneNumberDetailed) setDefaults() {
	{
		val := bool(true)
		s.CallForwardingEnabled.SetTo(val)
	}
	{
		val := SlimPhoneNumberDetailedNumberLevelRouting("disabled")
		s.NumberLevelRouting.SetTo(val)
	}
	{
		val := SlimPhoneNumberDetailedInboundCallScreening("disabled")
		s.InboundCallScreening.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *SoundModifications) setDefaults() {
	{
		val := string("outbound")
		s.Track.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *SpeakRequest) setDefaults() {
	{
		val := SpeakRequestPayloadType("text")
		s.PayloadType.SetTo(val)
	}
	{
		val := SpeakRequestServiceLevel("premium")
		s.ServiceLevel.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *StartForkingRequest) setDefaults() {
	{
		val := StartForkingRequestStreamType("raw")
		s.StreamType.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *StartRecordingRequest) setDefaults() {
	{
		val := int32(0)
		s.MaxLength.SetTo(val)
	}
	{
		val := int32(0)
		s.TimeoutSecs.SetTo(val)
	}
	{
		val := StartRecordingRequestRecordingTrack("both")
		s.RecordingTrack.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *StartStreamingRequest) setDefaults() {
	{
		val := StartStreamingRequestStreamTrack("inbound_track")
		s.StreamTrack.SetTo(val)
	}
	{
		val := bool(false)
		s.EnableDialogflow.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *StopForkingRequest) setDefaults() {
	{
		val := StopForkingRequestStreamType("raw")
		s.StreamType.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TexmlApplication) setDefaults() {
	{
		val := ConnectionActive(true)
		s.Active.SetTo(val)
	}
	{
		val := AnchorsiteOverride("Latency")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := DtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := FirstCommandTimeout(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := FirstCommandTimeoutSecs(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := TexmlApplicationVoiceMethod("post")
		s.VoiceMethod.SetTo(val)
	}
	{
		val := TexmlApplicationStatusCallbackMethod("post")
		s.StatusCallbackMethod.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TexmlApplicationInbound) setDefaults() {
	{
		val := bool(false)
		s.ShakenStirEnabled.SetTo(val)
	}
	{
		val := TexmlApplicationInboundSipSubdomainReceiveSettings("from_anyone")
		s.SipSubdomainReceiveSettings.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TexmlCreateCallRecordingRequestBody) setDefaults() {
	{
		val := PlayBeep(true)
		s.PlayBeep.SetTo(val)
	}
	{
		val := TexmlStatusCallbackMethod("POST")
		s.RecordingStatusCallbackMethod.SetTo(val)
	}
	{
		val := TexmlRecordingChannels("dual")
		s.RecordingChannels.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TexmlCreateCallRecordingResponseBody) setDefaults() {
	{
		val := TwimlRecordingChannels(2)
		s.Channels.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TexmlGetCallRecordingResponseBody) setDefaults() {
	{
		val := TwimlRecordingChannels(2)
		s.Channels.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TranscriptionStartRequest) setDefaults() {
	{
		val := TranscriptionStartRequestTranscriptionEngine("A")
		s.TranscriptionEngine.SetTo(val)
	}
	{
		val := TranscriptionStartRequestLanguage("en")
		s.Language.SetTo(val)
	}
	{
		val := bool(false)
		s.InterimResults.SetTo(val)
	}
	{
		val := string("inbound")
		s.TranscriptionTracks.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TransferCallRequest) setDefaults() {
	{
		val := int32(30)
		s.TimeoutSecs.SetTo(val)
	}
	{
		val := int32(14400)
		s.TimeLimitSecs.SetTo(val)
	}
	{
		val := TransferCallRequestAnsweringMachineDetection("disabled")
		s.AnsweringMachineDetection.SetTo(val)
	}
	{
		val := TransferCallRequestMediaEncryption("disabled")
		s.MediaEncryption.SetTo(val)
	}
	{
		val := TransferCallRequestSipTransportProtocol("UDP")
		s.SipTransportProtocol.SetTo(val)
	}
	{
		val := TransferCallRequestWebhookURLMethod("POST")
		s.WebhookURLMethod.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *TransferCallRequestAnsweringMachineDetectionConfig) setDefaults() {
	{
		val := int32(3500)
		s.TotalAnalysisTimeMillis.SetTo(val)
	}
	{
		val := int32(800)
		s.AfterGreetingSilenceMillis.SetTo(val)
	}
	{
		val := int32(50)
		s.BetweenWordsSilenceMillis.SetTo(val)
	}
	{
		val := int32(3500)
		s.GreetingDurationMillis.SetTo(val)
	}
	{
		val := int32(3500)
		s.InitialSilenceMillis.SetTo(val)
	}
	{
		val := int32(5)
		s.MaximumNumberOfWords.SetTo(val)
	}
	{
		val := int32(3500)
		s.MaximumWordLengthMillis.SetTo(val)
	}
	{
		val := int32(256)
		s.SilenceThreshold.SetTo(val)
	}
	{
		val := int32(5000)
		s.GreetingTotalAnalysisTimeMillis.SetTo(val)
	}
	{
		val := int32(1500)
		s.GreetingSilenceDurationMillis.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *UpdateCallControlApplicationRequest) setDefaults() {
	{
		val := bool(true)
		s.Active.SetTo(val)
	}
	{
		val := AnchorsiteOverride("Latency")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := UpdateCallControlApplicationRequestDtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := bool(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := int(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := UpdateCallControlApplicationRequestWebhookAPIVersion("1")
		s.WebhookAPIVersion.SetTo(val)
	}
	{
		val := string("")
		s.WebhookEventFailoverURL.SetTo(val)
	}
	{
		s.WebhookTimeoutSecs.Null = true
	}
}

// setDefaults set default value of fields.
func (s *UpdatePhoneNumberRequest) setDefaults() {
	{
		val := UpdatePhoneNumberRequestNumberLevelRouting("enabled")
		s.NumberLevelRouting.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *UpdatePhoneNumberVoiceSettingsRequest) setDefaults() {
	{
		val := bool(false)
		s.TechPrefixEnabled.SetTo(val)
	}
	{
		val := bool(false)
		s.CallerIDNameEnabled.SetTo(val)
	}
	{
		val := UpdatePhoneNumberVoiceSettingsRequestUsagePaymentMethod("pay-per-minute")
		s.UsagePaymentMethod.SetTo(val)
	}
	{
		val := UpdatePhoneNumberVoiceSettingsRequestInboundCallScreening("disabled")
		s.InboundCallScreening.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *UpdateTexmlApplicationRequest) setDefaults() {
	{
		val := ConnectionActive(true)
		s.Active.SetTo(val)
	}
	{
		val := AnchorsiteOverride("Latency")
		s.AnchorsiteOverride.SetTo(val)
	}
	{
		val := DtmfType("RFC 2833")
		s.DtmfType.SetTo(val)
	}
	{
		val := FirstCommandTimeout(false)
		s.FirstCommandTimeout.SetTo(val)
	}
	{
		val := FirstCommandTimeoutSecs(30)
		s.FirstCommandTimeoutSecs.SetTo(val)
	}
	{
		val := UpdateTexmlApplicationRequestVoiceMethod("post")
		s.VoiceMethod.SetTo(val)
	}
	{
		val := UpdateTexmlApplicationRequestStatusCallbackMethod("post")
		s.StatusCallbackMethod.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *UpdateTexmlApplicationRequestInbound) setDefaults() {
	{
		val := bool(false)
		s.ShakenStirEnabled.SetTo(val)
	}
	{
		val := UpdateTexmlApplicationRequestInboundSipSubdomainReceiveSettings("from_anyone")
		s.SipSubdomainReceiveSettings.SetTo(val)
	}
}
