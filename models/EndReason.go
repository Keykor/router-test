package models

type EndReason string

const (
	ReasonSuccess   EndReason = "Success"
	ReasonFusion    EndReason = "Fusion"
	ReasonCancelled EndReason = "Cancelled"
)
