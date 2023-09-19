terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.17.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_appconfig_application" "main" {
  name        = "appconfig-golang-api"
  description = "AppConfig Application for a Golang REST API"
}

resource "aws_appconfig_configuration_profile" "feature_flags" {
  application_id = aws_appconfig_application.main.id
  description    = "Example Configuration Profile"
  name           = "Hosted Feature Flags"
  location_uri   = "hosted"
  type           = "AWS.AppConfig.FeatureFlags"
}

resource "aws_appconfig_hosted_configuration_version" "example" {
  application_id           = aws_appconfig_application.main.id
  configuration_profile_id = aws_appconfig_configuration_profile.feature_flags.configuration_profile_id
  description              = "Example Feature Flag Configuration Version"
  content_type             = "application/json"

  content = jsonencode({
    flags : {
      foo : {
        name : "foo",
        _deprecation : {
          "status" : "planned"
        }
      },
      bar : {
        name : "bar",
        attributes : {
          someAttribute : {
            constraints : {
              type : "string",
              required : true
            }
          },
          someOtherAttribute : {
            constraints : {
              type : "number",
              required : true
            }
          }
        }
      }
    },
    values : {
      foo : {
        enabled : "true",
      },
      bar : {
        enabled : "true",
        someAttribute : "Hello World",
        someOtherAttribute : 123
      }
    },
    version : "1"
  })
}
