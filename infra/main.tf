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
  name        = "MoviesApp"
  description = "Watch movies anywhere"
}

### Hosted Feature Flags ###

resource "aws_appconfig_configuration_profile" "feature_flags_hosted" {
  application_id = aws_appconfig_application.main.id
  description    = "Feature flags that hosted by AppConfig"
  name           = "Hosted Feature Flags"
  location_uri   = "hosted"
  type           = "AWS.AppConfig.FeatureFlags"
}

resource "aws_appconfig_hosted_configuration_version" "feature_flags_hosted_v1" {
  application_id           = aws_appconfig_application.main.id
  configuration_profile_id = aws_appconfig_configuration_profile.feature_flags_hosted.configuration_profile_id
  description              = "Example Feature Flag Configuration Version"
  content_type             = "application/json"

  content = jsonencode({
    flags : {
      blackfriday : {
        name : "Black Friday Promotion",
        description : "Discount feature for the Black Friday sales.",
        _deprecation : {
          "status" : "planned"
        }
      },
      recommendations : {
        name : "Smart Recommendations",
        description : "Movie recommendations feature.",
        attributes : {
          itemsQuantity : {
            constraints : {
              type : "number",
              required : true
            }
          },
          someAttribute : {
            constraints : {
              type : "string",
              required : true
            }
          }
        }
      }
    },
    values : {
      blackfriday : {
        enabled : "true",
      },
      recommendations : {
        enabled : "true",
        itemsQuantity : 3,
        someAttribute : "Hello World"
      }
    },
    version : "1"
  })
}
