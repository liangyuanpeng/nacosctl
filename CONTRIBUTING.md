How to contribute
If you would like to contribute code to this project, fork the repository and send a pull request.

Prerequisite
In this project, use go mod as the package management tool and make sure your Go version is higher then Go 1.11.

Fork
Before contributing, you need to fork pulsarctl to your GitHub repository.

Contribution flow
$ git remote add streamnative https://github.com/streamnative/pulsarctl.git
# sync with the remote master
$ git checkout master
$ git fetch streamnative
$ git rebase streamnative/master
$ git push origin master
# create a PR branch
$ git checkout -b your_branch   
# do something
$ git add [your change files]
$ git commit -sm "xxx"
$ git push origin your_branch
Configure GoLand
The pulsarctl uses go mod to manage dependencies, so make sure your IDE enables Go Modules(vgo).

To configure annotation processing in GoLand, follow the steps below.

To open the Go Modules Settings window, in GoLand, click Preferences > Go > Go Modules(vgo).

Select the Enable Go Modules(vgo) integration checkbox.

Click Apply and OK.

Code style
The code style suggested by the Golang community is used in pulsarctl. For more information, see Go Code Review Comments.

To make your pull request easy to review, maintain and develop, follow this style.

Create a new file
The project uses the open source protocol of Apache License 2.0. If you need to create a new file when developing new features, add the license at the beginning of each file. The location of the header file: header file.

Update dependencies
The pulsarctl uses Go 1.11 module to manage dependencies. To add or update a dependency, use the go mod edit command to change the dependency.