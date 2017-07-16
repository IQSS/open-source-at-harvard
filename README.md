# Open Source at Harvard

## What is this open-source-at-harvard project about? How do I add my project to the list?

How much open source software is written at Harvard? How popular is that software? When was it last updated? Which groups are writing the software? Which languages are used? These are the initial questions this **open-source-at-harvard** project attempts to answer, but there are many more questions we can tackle in the future. How many production services is Harvard hosting that are based on this open source software? How many contributors does each project have? Who are the people who are contributing?

Our starting point is a simple list of Harvard open source projects on GitHub at [github.tsv](github.tsv). To add your project to the list, please make a pull request. (Thank you to everyone who has already made a pull request!) If your project isn't on GitHub, please open an issue.

(If you work at Harvard and are interested in open sourcing your project, you might find this "Harvard Library Open Source Project Considerations" page helpful: https://wiki.harvard.edu/confluence/display/LibraryStaffDoc/Harvard+Library+Open+Source+Project+Considerations .)

## How do I run the code in this open-source-at-harvard repo?

Code in this repo operates the list of projects at [github.tsv](github.tsv) and creates a file called `data.tsv` which should look something like [data/2017-06-30.tsv](data/2017-06-30.tsv) or the data published at http://dx.doi.org/10.7910/DVN/TJCLKP .

There are enough Harvard open source projects on GitHub listed that to download metadata about them from the GitHub API. Before running any code it this repo, you must create a personal access token for yourself at https://github.com/settings/tokens and copy and paste it into a file called `.github-token`.

In order to reduce the number of calls to the GitHub API, we first download metadata about all the projects as JSON files and then operate on those JSON files to gather the information into a tab-separated values (TSV) file, again called `data.tsv`.

To run the code in this repo, Go 1.4 or higher is required. You can download a snapshot of the code as a zip file or tarball, but if you have Git installed you can simply clone the repo:

    git clone https://github.com/IQSS/open-source-at-harvard.git

This will create a directory called `open-source-at-harvard`. Navigate to that directory with `cd`:

    cd open-source-at-harvard

Finally, issue the following commands to download JSON files about each project and parse them:

    go run download.go
    go run parse.go

A file called `data.tsv` should result, representing the data you just downloaded. It should look something like [data/2017-06-30.tsv](data/2017-06-30.tsv) or the data published at http://dx.doi.org/10.7910/DVN/TJCLKP .

If you have any trouble with the commands above or have any ideas for this project, please open a GitHub issue. Pull request are also very welcome!

## opensource.harvard.edu

The code in this repo was hacked together by Philip Durbin who thinks that some day we should launch a site at opensource.harvard.edu and makes the argument in a [Google doc][] where public comments are enabled and very much encouraged! You are also welcome to start a discussion about the opensource.harvard.edu idea by opening an issue in this repo.

[Google doc]: https://docs.google.com/document/d/1CSWV9VxHfJj_ahArNYTsCAG0D8OtSfZhrCwpNiIKWQw/edit?usp=sharing
