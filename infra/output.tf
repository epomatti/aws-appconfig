output "application_id" {
  value = aws_appconfig_application.main.id
}

output "configuration_profile_id" {
  value = aws_appconfig_configuration_profile.feature_flags_hosted.configuration_profile_id
}

output "environment_id" {
  value = aws_appconfig_environment.production.environment_id
}
