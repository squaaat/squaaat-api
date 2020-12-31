locals {

}

module "alpha" {
  source = "./module"

  env = "alpha"
}

module "production" {
  source = "./module"

  env = "production"
}