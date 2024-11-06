variable "repository" {
  type = string
}

variable "live_image_tag" {
  type = string
}

variable "canary_image_tag" {
    type = string
    default = ""
}

variable "canary_enabled" {
  description = "Enable the canary"
  type = bool
  default = false
}

variable "canary_percent" {
  description = "Percentage of traffic to send to the canary"
  type = number
  default = 10
}

variable "force_new_revision" {
  type = bool
  default = false
}