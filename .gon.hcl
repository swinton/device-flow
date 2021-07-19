source = ["./device-flow"]
bundle_id = "com.github.swinton.device-flow"

apple_id {
  username = "@env:AC_USERNAME"
  password = "@env:AC_PASSWORD"
}

sign {
  application_identity = "25AF0257E423C5F1020177ABFE73B71EFC08B71D"
}

zip {
  output_path = "device-flow.zip"
}
