#!/usr/bin/env python3
import csv
date = '2023-01-03'
tsv_file = 'data/' + date + '.tsv'
with open('language.tsv', 'w', newline='') as tsvfile:
    output = csv.writer(tsvfile, delimiter='\t')
    header = [
        'language',
        'repo',
    ]
    output.writerow(header)
    with open(tsv_file) as tsvfile:
        reader = csv.DictReader(tsvfile, delimiter="\t")
        rows = [row for row in reader]
        for row in rows:
            lang = row['language']
            repo = row['repo']
            output.writerow([
                lang,
                repo,
            ])
