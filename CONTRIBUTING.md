# How to contribute #

We'd love to accept your patches and contributions to this project. There are
a just a few small guidelines you need to follow.


## Contributor License Agreement ##

TODO

## Reporting issues ##

Bugs, feature requests, and development-related questions should be directed to
our [GitHub issue tracker](https://github.com/becoded/billit/issues).  If
reporting a bug, please try and provide as much context as possible such as
your operating system, Go version, and anything else that might be relevant to
the bug.  For feature requests, please explain what you're trying to do, and
how the requested feature would help you do that.

Security related bugs can either be reported in the issue tracker, or if they
are more sensitive, emailed to <bert@becoded.be>.

## Submitting a patch ##

1. It's generally best to start by opening a new issue describing the bug or
   feature you're intending to fix. Even if you think it's relatively minor,
   it's helpful to know what people are working on. Mention in the initial
   issue that you are planning to work on that bug or feature so that it can
   be assigned to you.

1. Follow the normal process of [forking][] the project, and setup a new
   branch to work in. It's important that each group of changes be done in
   separate branches in order to ensure that a pull request only includes the
   commits related to that bug or feature.

1. Go makes it very simple to ensure properly formatted code, so always run
   `go fmt` on your code before committing it. You should also run
   [golint][] over your code. As noted in the [golint readme][], it's not
   strictly necessary that your code be completely "lint-free", but this will
   help you find common style issues.

1. Any significant changes should almost always be accompanied by tests. The
   project already has good test coverage, so look at some of the existing
   tests if you're unsure how to go about it. [gocov][] and [gocov-html][]
   are invaluable tools for seeing which parts of your code aren't being
   exercised by your tests.

1. Please run:
    * `go test github.com/becoded/billit/...`
    * `go vet github.com/becoded/billit/...`

1. Do your best to have [well-formed commit messages][] for each change.
   This provides consistency throughout the project, and ensures that commit
   messages are able to be formatted properly by various git tools.

1. Finally, push the commits to your fork and submit a [pull request][].

[forking]: https://help.github.com/articles/fork-a-repo
[golangci]: https://github.com/golangci/golangci-lint
[gocov]: https://github.com/axw/gocov
[gocov-html]: https://github.com/matm/gocov-html
[well-formed commit messages]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html
[squash]: http://git-scm.com/book/en/Git-Tools-Rewriting-History#Squashing-Commits
[pull request]: https://help.github.com/articles/creating-a-pull-request

