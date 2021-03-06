variable "name" {}

variable "project" {}

variable "contact" {}

variable "key_name" {}

variable "region" {}

variable "stack" {
  default = ""
}

variable "state_bucket" {
  default = ""
}

variable "stack_name_prefix" {
  default = ""
}

data "template_file" "stack_name" {
  template = "${var.stack_name_prefix}${var.environment}-${var.name}"
}

variable "allowed_account_ids" {
  type    = "list"
  default = []
}

variable "environment" {
  default = "nonprod"
}

variable "vault_init_token_master" {
  default = ""
}

variable "vault_init_token_worker" {
  default = ""
}

variable "vault_init_token_etcd" {
  default = ""
}

variable "state_cluster_name" {
  default = "hub"
}

variable "vault_cluster_name" {
  default = "hub"
}

variable "tools_cluster_name" {
  default = "hub"
}
