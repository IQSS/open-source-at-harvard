# Contributing

## Running the code in this repo

Code in this repo operates in two stages, download and parse. In the parse stage, the code uses the list of projects at [github.tsv](github.tsv) as input and downloads metadata about each project from the GitHub API as a series of JSON files, one per project. In the second stage, those JSON files are used as input to create a file called `data.tsv` which should look something like [data/2023-01-03.tsv](data/2023-01-03.tsv) or the data published at http://dx.doi.org/10.7910/DVN/TJCLKP . The reason for the two stages is that we are attempting to reduce the number requests to the GitHub API. Once you have downloaded the JSON files, you can hack away at the parsing code without burdening the GitHub API or spending any time or bandwidth downloading the JSON files over and over.

To run the code in this repo, [Go](https://golang.org) 1.4 or higher is required.

Start by cloning the repo:

    git clone https://github.com/IQSS/open-source-at-harvard.git

Then navigate into the directory that was created:

    cd open-source-at-harvard

There are enough Harvard open source projects on GitHub listed at [github.tsv](github.tsv) that you must create a personal access token for yourself at https://github.com/settings/tokens and place it into a file called `.github-token`. Otherwise, GitHub's rate limiting will prevent you from downloading metadata about all of the projects. (JSON files will always be downloaded but they'll contain errors about rate limiting.) In the example below, we're using vi to create the `.github-token` file:

    vi .github-token

Finally, issue the following commands to download JSON files about each project and parse them:

    go run download.go
    go run parse.go

A file called `data.tsv` should result, representing the data you just downloaded. Again, it should look something like the tsv file described above or the one at http://dx.doi.org/10.7910/DVN/TJCLKP .

If you have any trouble with the commands above or have any ideas for this project, please open a GitHub issue. Pull request are also very welcome!
