# Open Source at Harvard

[![doi:10.7910/DVN/TJCLKP](https://img.shields.io/badge/DOI-10.7910%2FDVN%2FTJCLKP-orange.svg)](https://doi.org/10.7910/DVN/TJCLKP)

## What is this open-source-at-harvard project about? How do I add my project to the list?

How much open source software is written at Harvard? How popular is that software? When was it last updated? Which groups are writing the software? Which languages are used? These are the initial questions this **open-source-at-harvard** project attempts to answer, but there are many more questions we can tackle in the future. How many production services is [Harvard hosting][] that are based on this open source software? How many contributors does each project have? Who are the people who are contributing?

[Harvard hosting]: https://github.com/IQSS/open-source-at-harvard/issues/3

Our starting point is a simple list of Harvard open source projects on GitHub at [github.tsv](github.tsv). To add your project to the list, please make a pull request. (Thank you to everyone who has already made a pull request!) If your project isn't on GitHub, please open an issue.

(If you work at Harvard and are interested in open sourcing your project, you might find this "Harvard Library Open Source Project Considerations" page helpful: https://wiki.harvard.edu/confluence/display/LibraryStaffDoc/Harvard+Library+Open+Source+Project+Considerations .)

## How do I run the code in this open-source-at-harvard repo?

Code in this repo operates in two stages, download and parse. In the parse stage, the code uses the list of projects at [github.tsv](github.tsv) as input and downloads metadata about each project from the GitHub API as a series of JSON files, one per project. In the second stage, those JSON files are used as input to create a file called `data.tsv` which should look something like [data/2017-07-31.tsv](data/2017-07-31.tsv) or the data published at http://dx.doi.org/10.7910/DVN/TJCLKP . The reason for the two stages is that we are attempting to reduce the number requests to the GitHub API. Once you have downloaded the JSON files, you can hack away at the parsing code without burdening the GitHub API or spending any time or bandwidth downloading the JSON files over and over.

To run the code in this repo, [Go](https://golang.org) 1.4 or higher is required. You can download a snapshot of the code as a zip file or tarball, but if you have Git installed you can simply clone the repo:

    git clone https://github.com/IQSS/open-source-at-harvard.git

This will create a directory called `open-source-at-harvard`. Navigate to that directory with `cd`:

    cd open-source-at-harvard

There are enough Harvard open source projects on GitHub listed at [github.tsv](github.tsv) that you must create a personal access token for yourself at https://github.com/settings/tokens and place it into a file called `.github-token`. Otherwise, GitHub's rate limiting will prevent you from downloading metadata about all of the projects. (JSON files will always be downloaded but they'll contain errors about rate limiting.) In the example below, we're using vi to create the `.github-token` file:

    vi .github-token

Finally, issue the following commands to download JSON files about each project and parse them:

    go run download.go
    go run parse.go

A file called `data.tsv` should result, representing the data you just downloaded. Again, it should look something like the tsv file described above or the one at http://dx.doi.org/10.7910/DVN/TJCLKP .

If you have any trouble with the commands above or have any ideas for this project, please open a GitHub issue. Pull request are also very welcome!

## opensource.harvard.edu

The code in this repo was hacked together by Philip Durbin who thinks that some day we should launch a site at opensource.harvard.edu and makes the argument in a [Google doc][] where public comments are enabled and very much encouraged! You are also welcome to start a discussion about the opensource.harvard.edu idea by opening an issue in this repo. Harvard will never write as much open source software as Google, but we can use their showcase at https://opensource.google.com/projects for inspiration!

[Google doc]: https://docs.google.com/document/d/1CSWV9VxHfJj_ahArNYTsCAG0D8OtSfZhrCwpNiIKWQw/edit?usp=sharing
