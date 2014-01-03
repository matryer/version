# Version

Command line tool (and Go package) for keeping track of the versions of projects or directories.  Version creates and maintains a `.version` file in the directory containing the current version number, and provides a command line tool to easily get, and update the version number.

Perfect for:

  * Automated build/release scripts
  * Integration with GitHub tags

### Command line

The `version` command line has the following syntax:

    version path [option]

  * `path` - Path to set the version for.  Use `./` for current directory.
  * `option`
    * No option will just read and return the current value and will not change it
    * `+` Increase the build number (`1.0.0` -> `1.0.1`) and return the new value
    * `++` Increase the minor number (`1.0.0` -> `1.1.0`) and return the new value
    * `+++` Increase the major number (`1.0.0` -> `2.0.0`) and return the new value

### Tips for writing scripts

Use backticks to get the current version and use it in another command:

    echo `version ./`
    
... or increase the build version at the same time as getting it:

    echo `version ./ +`

To use the version multiple times, use variables:    

    VERSION=`./version ./ +`; echo $VERSION
    
#### Tagging in GitHub

    VERSION=`version ./ +`; git tag -a `echo $VERSION` -m "Release $VERSION"; git push --tags
    
The above script will increase the build number, tag the repo and push it to the server.
