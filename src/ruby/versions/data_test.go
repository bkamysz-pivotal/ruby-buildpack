package versions_test

const windowsGemfileLockFixture = `GEM
  remote: https://rubygems.org/
  specs:
    rack (1.5.2)
    rack-protection (1.5.2)
      rack
    sinatra (1.4.4)
      rack (~> 1.4)
      rack-protection (~> 1.4)
      tilt (~> 1.3, >= 1.3.4)
    tilt (1.4.1)

PLATFORMS
  x64-mingw32

DEPENDENCIES
  sinatra
`

const linuxGemfileLockFixture = `GEM
  remote: https://rubygems.org/
  specs:
    rack (1.5.2)
    rack-protection (1.5.2)
      rack
    sinatra (1.4.4)
      rack (~> 1.4)
      rack-protection (~> 1.4)
      tilt (~> 1.3, >= 1.3.4)
    tilt (1.4.1)

PLATFORMS
  ruby

DEPENDENCIES
  sinatra
`
